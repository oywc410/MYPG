package tmp


//file:///Users/ouyou/gopath/src/github.com/golang-china/gopl-zh/_book/ch9/ch9-06.html
隻要在go build，go run或者go test命令後面加上-race的flag，就會使編譯器創建一個你的應用的“脩改”版或者一個附帶了能夠記録所有運行期對共享變量訪問工具的test，併且會記録下每一個讀或者寫共享變量的goroutine的身份信息。另外，脩改版的程序會記録下所有的同步事件，比如go語句，channel操作，以及對(*sync.Mutex).Lock，(*sync.WaitGroup).Wait等等的調用。(完整的同步事件集合是在The Go Memory Model文檔中有説明，該文檔是和語言文檔放在一起的。譯註：https://golang.org/ref/mem)