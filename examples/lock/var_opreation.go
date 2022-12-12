package lock

import (
	log "github.com/corgi-kx/logcustom"
	"runtime"
)

var i = 0

var numCh = make(chan interface{}, 1)

func add() {
	i++
}

func Handler() {
	numCh <- "aa"
	// 协程goroutine排队执行，提交给processor执行
	runtime.GOMAXPROCS(100)
	for i := 0; i < 1000000; i++ {
		go add()
	}
	<-numCh
	log.Info("计算结果：", i)
}
