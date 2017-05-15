package cmap

import (
	"sync"
	"testing"
)

type testg struct {

}

func BenchmarkLinkedlist_SetGet(b *testing.B) {
	stringMap := newBucket()
	work := sync.WaitGroup{}

	lock := &sync.Mutex{}

	for i:=0;i<b.N;i++ {
		work.Add(1)
		go func() {
			d, _ := newPair("test", "test_value")
			stringMap.Put(d, lock)
			stringMap.Get("test")
			work.Done()
		}()
	}

	work.Wait()
}
