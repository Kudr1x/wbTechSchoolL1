// file: main.go
package main

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
)

// Остановка по условию
func demoAtomic() {
	fmt.Println("== atomic flag ==")
	var stop atomic.Bool

	go func() {
		for {
			if stop.Load() {
				fmt.Println("worker[atomic]: stop")
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(350 * time.Millisecond)
	stop.Store(true)
	time.Sleep(100 * time.Millisecond)
}

// Через done-канал
func demoDoneChannel() {
	fmt.Println("== done channel ==")
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("worker[done]: stop")
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(350 * time.Millisecond)
	close(done)
	time.Sleep(100 * time.Millisecond)
}

// Закрытие рабочего канала (range завершается)
func demoCloseJobs() {
	fmt.Println("== close jobs channel ==")
	jobs := make(chan int)
	results := make(chan string, 8)

	go func() {
		for j := range jobs {
			results <- fmt.Sprintf("worker[jobs] handled %d", j)
		}
		results <- "worker[jobs]: exit after close"
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 4; i++ {
		fmt.Println(<-results)
	}
}

// Context, ручная отмена WithCancel
func demoContextCancel() {
	fmt.Println("== context.WithCancel ==")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker[ctx-cancel]:", ctx.Err())
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(350 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

// Context, таймаут
func demoContextTimeout() {
	fmt.Println("== context.WithTimeout ==")
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		t := time.NewTicker(100 * time.Millisecond)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker[ctx-timeout]:", ctx.Err())
				return
			case <-t.C:
				// periodic work
			}
		}
	}()

	<-done
}

// Остановка по сигналам ОС
func demoSignalNotifyContext() {
	fmt.Println("== signal.NotifyContext ==")
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		fmt.Println("worker[signal]: stopping on OS signal (simulated)")
	}()

	time.AfterFunc(200*time.Millisecond, stop)
	time.Sleep(400 * time.Millisecond)
}

// runtime.Goexit
func demoGoexit() {
	fmt.Println("== runtime.Goexit ==")
	done := make(chan struct{})
	go func() {
		defer fmt.Println("worker[Goexit]: deferred runs")
		time.Sleep(150 * time.Millisecond)
		runtime.Goexit()
	}()
	go func() {
		time.Sleep(300 * time.Millisecond)
		close(done)
	}()
	<-done
}

// panic
func demoPanicRecover() {
	fmt.Println("== panic + recover ==")
	done := make(chan struct{})
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("worker[panic]: recovered and exit")
			}
			close(done)
		}()
		time.Sleep(150 * time.Millisecond)
		panic("force stop (demo)")
	}()
	<-done
}

func main() {
	demoAtomic()
	demoDoneChannel()
	demoCloseJobs()
	demoContextCancel()
	demoContextTimeout()
	demoSignalNotifyContext()
	demoGoexit()
	demoPanicRecover()
	fmt.Println("finish")
}
