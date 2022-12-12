package basics

import (
	log "github.com/corgi-kx/logcustom"
	"sync"
	"time"
)

//chan1 创建一个管道
var chan1 = make(chan string, 5)
var chan2 = make(chan string, 5)

var wg sync.WaitGroup

func chan1send(name string) {
	for i := 0; i < 10; i++ {
		chan1 <- name
		log.Info("producer-->", name)
	}
}
func chan2send(name string) {
	for i := 0; i < 10; i++ {
		chan2 <- name
		log.Info("producer-->", name)
	}
}

func TestChannel() {
	go chan1send("张三")
	go chan2send("李四")
	go func() {
		for {
			select {
			case <-chan1:
				log.Info("consumer:", <-chan1)
			case <-chan2:
				log.Info("consumer:", <-chan2)
			}
		}

	}()
	time.Sleep(time.Second * 3)
}
