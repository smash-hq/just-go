package digital

import (
	"container/ring"
	log "github.com/corgi-kx/logcustom"
)

const WindowSize = 10

func GoRing() {
	r := ring.New(WindowSize)
	for i := 0; i < 100; i++ {
		r.Value = i
		r = r.Next()
	}
	var sum int
	r.Do(func(i interface{}) {
		sum += i.(int)
	})
	log.Info(sum)
}
