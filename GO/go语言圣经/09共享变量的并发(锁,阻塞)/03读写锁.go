package tmp
import "sync"

var mu sync.RWMutex

var balance int

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}