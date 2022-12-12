package basics

import (
	log "github.com/corgi-kx/logcustom"
	"sync"
)

func ExecRunQ() {

	wg := sync.WaitGroup{}
	wg.Add(3)

	go printN(1, &wg)
	go printN(2, &wg)
	go printN(3, &wg)
	wg.Wait()
}

func printN(i int, wg *sync.WaitGroup) {
	log.Info(i)
	wg.Done()
}
