package lock

import (
	log "github.com/corgi-kx/logcustom"
	"sync"
)

func ByMutex() {
	var syncLock sync.Mutex

	syncLock.Lock()
	go func() {
		log.Info("先打印，再解锁")
		syncLock.Unlock()
	}()
	syncLock.Lock()
}

func ByChannel() {
	var channel = make(chan string)

	go func() {
		log.Info("通过管道实现加锁")
		channel <- "黄棋"
	}()
	log.Info("取出通道内内容:", <-channel)

}
