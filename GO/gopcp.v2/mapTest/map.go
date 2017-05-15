package mapTest

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type StringMap struct {
	data map[string]string
	dataLock map[string]sync.Mutex
	lock sync.RWMutex
}

func NewStringMap(szie int) *StringMap {
	return &StringMap{
		data: make(map[string]string, szie),
		dataLock: make(map[string]sync.Mutex, szie),
		lock: sync.RWMutex{},
	}
}

func (s *StringMap)Set(key string, str string) {
	s.lock.Lock()


	//datalock.Lock()
	s.data[key] = str
	//datalock.Unlock()
	s.lock.Unlock()

}

func (s *StringMap)Get(key string) (reData string) {
	s.lock.RLock()
	reData = s.data[key]

	s.lock.RUnlock()
	return
}

//-----------------------------------------------------


func hash(str string) uint64 {
	str = str
	seed := uint64(13131)
	var hash uint64
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint64(str[i])
	}
	return (hash & 0x7FFFFFFFFFFFFFFF)
}

type Item struct {
	Key string
	Data string
	Next unsafe.Pointer
}

type Linkedlist struct {
	firstValue atomic.Value
	lock sync.Mutex
	len *int64
}

type Piece struct {
	links []Linkedlist
	oldLinks []Linkedlist
	linkLen uint64
	len *int64
	lock sync.Mutex
	reSet *uint32
	maxsize int64
}

func NewPiece(size uint64) *Piece {
	var i int64
	var reset uint32
	piece := &Piece{
		links: make([]Linkedlist, size),
		linkLen: size,
		len: &i,
		lock: sync.Mutex{},
		reSet: &reset,
		maxsize: 30,
	}

	for i:=0;i<len(piece.links);i++ {
		piece.links[i] = *NewLinkedlist()
	}

	return piece
}

func (p *Piece)Set(key, value string) {

	p.lock.Lock()

	i := hash(key) % p.linkLen
	if size, nex := p.links[i].Set(key, value); nex {
		atomic.AddInt64(p.len, 1)
		if size >= p.maxsize {
			p.oldLinks = make([]Linkedlist, p.linkLen)
			copy(p.oldLinks, p.links)
			p.linkLen = p.linkLen * 2
			newlinks := make([]Linkedlist, p.linkLen)
			var i uint64
			for i=0;i<p.linkLen;i++ {
				newlinks[i] = *NewLinkedlist()
			}
			for _, links := range p.oldLinks {
				v := links.firstValue.Load()
				if v == nil {
					panic("error not firstValue")
				}
				item := v.(*Item)
				for {
					pointer := atomic.LoadPointer(&item.Next)
					if pointer == nil {
						break
					}
					item = (*Item)(pointer)
					i := hash(item.Key) % p.linkLen
					newlinks[i].Set(item.Key, item.Data)
				}
			}

			p.oldLinks = nil
			p.links = newlinks
		}
	}

	p.lock.Unlock()
}

func (p *Piece)Get(key string) (reValue string) {
	i := hash(key) % p.linkLen
	reValue = p.links[i].Get(key)

	return
}

func (p *Piece)Len() int64 {
	return atomic.LoadInt64(p.len)
}

func NewLinkedlist() *Linkedlist {
	var len int64

	link := &Linkedlist{
		firstValue: atomic.Value{},
		lock: sync.Mutex{},
		len: &len,
	}

	link.firstValue.Store(&Item{})

	return link
}

func (l *Linkedlist) Set(key string, value string) (int64, bool) {
	l.lock.Lock()

	v := l.firstValue.Load()
	if v == nil {
		panic("error not firstValue")
	}
	item := v.(*Item)
	for {
		pointer := atomic.LoadPointer(&item.Next)
		if pointer == nil {
			break
		}
		item = (*Item)(pointer)
		if item.Key == key {
			item.Data = value
			l.lock.Unlock()
			return 0, false
		}
	}

	atomic.StorePointer(&item.Next, unsafe.Pointer(&Item{
		Key: key,
		Data: value,
	}))
	l.lock.Unlock()

	return atomic.AddInt64(l.len, 1), true
}

func (l *Linkedlist) Get(key string) string {

	v := l.firstValue.Load()
	if v == nil {
		panic("error not firstValue")
	}
	item := v.(*Item)
	for {
		pointer := atomic.LoadPointer(&item.Next)
		if pointer == nil {
			break
		}
		item = (*Item)(pointer)
		if item.Key == key {
			return item.Data
		}
	}

	return ""
}

func (l *Linkedlist) Len() int64 {
	return atomic.LoadInt64(l.len)
}