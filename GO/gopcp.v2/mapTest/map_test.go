package mapTest

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sync"
	"strconv"
)

func TestStringMap_Set(t *testing.T) {
	stringMap := NewStringMap(1)
	stringMap.Set("test", "test_value")
	assert.Equal(t, "test_value", stringMap.Get("test"))
}

func TestLinkedlist_Set(t *testing.T) {
	stringMap := NewPiece(1)
	stringMap.Set("test", "test_value")
	assert.Equal(t, "test_value", stringMap.Get("test"))

	stringMap.Set("test1", "test_value1")
	assert.Equal(t, "test_value1", stringMap.Get("test1"))
}

func TestHash(t *testing.T) {
	t.Log(hash("1"))
	t.Log(hash("100"))
	t.Log(hash("au980ujiubcd"))
}

func BenchmarkStringMap_SetGet(b *testing.B) {
	stringMap := NewStringMap(10)
	work := sync.WaitGroup{}

	b.StopTimer()
	for i:=0;i<b.N;i++ {
		stringMap.Set(strconv.Itoa(i), "test_value")
	}
	b.StartTimer()

	for i:=0;i<b.N;i++ {
		work.Add(1)
		go func() {
			stringMap.Set(strconv.Itoa(i), "test_value")
			stringMap.Get(strconv.Itoa(i))
			work.Done()
		}()
	}

	work.Wait()
	b.Log(b.N, len(stringMap.data))
}


func BenchmarkLinkedlist_SetGet(b *testing.B) {
	stringMap := NewLinkedlist()
	work := sync.WaitGroup{}

	b.StopTimer()
	for i:=0;i<b.N;i++ {
		stringMap.Set(strconv.Itoa(i), "test_value")
	}
	b.StartTimer()

	for i:=0;i<b.N;i++ {
		work.Add(1)
		go func() {
			stringMap.Set(strconv.Itoa(i), "test_value")
			stringMap.Get(strconv.Itoa(i))
			work.Done()
		}()
	}

	work.Wait()
	b.Log(b.N, *stringMap.len)
}


func BenchmarkPieceLinked_SetGet(b *testing.B) {
	stringMap := NewPiece(10)
	work := sync.WaitGroup{}

	b.StopTimer()
	for i:=1;i<b.N;i++ {
		stringMap.Set(strconv.Itoa(i), "test_value")
	}
	b.StartTimer()

	for i:=0;i<b.N;i++ {
		work.Add(1)
		go func() {
			stringMap.Set(strconv.Itoa(i), "test_value")
			stringMap.Get(strconv.Itoa(i))
			work.Done()
		}()
	}

	work.Wait()
	b.Log(b.N, *stringMap.len, stringMap.linkLen)
}