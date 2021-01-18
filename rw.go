package mutex

import (
	"sync"
	"time"
)

var (
	producerCount = 10
)

func society(count int, mutex, rwMutex sync.Locker) time.Duration {
	var wg sync.WaitGroup
	wg.Add(count + 1)
	bgn := time.Now()

	go producer(&wg, mutex)

	for i := count; i > 0; i-- {
		go observer(&wg, rwMutex)
	}

	wg.Wait()
	return time.Since(bgn)

}

func producer(wg *sync.WaitGroup, l sync.Locker) {
	defer wg.Done()
	for i := producerCount; i > 0; i-- {
		l.Lock()
		l.Unlock()
		time.Sleep(1)
	}
}

func observer(wg *sync.WaitGroup, l sync.Locker) {
	defer wg.Done()
	l.Lock()
	defer l.Unlock()
}
