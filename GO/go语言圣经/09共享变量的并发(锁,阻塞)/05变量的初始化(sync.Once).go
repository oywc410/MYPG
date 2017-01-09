package tmp
import (
	"image"
	"sync"
)

var icons map[string]image.Image


func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// NOTE: not concurrency-safe!()多进程时有可能反复初始化
func Icon1(name string) image.Image {
	if icons == nil {
		loadIcons() // one-time initialization
	}
	return icons[name]
}

var mu sync.RWMutex

func Icon(name string) image.Image {

	//读取共享锁
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()

	//互斥锁
	mu.Lock()
	if icons == nil {
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}