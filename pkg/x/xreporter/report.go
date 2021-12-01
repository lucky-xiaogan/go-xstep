package xreporter

import (
	"fmt"
	"sync"
)

type Reporter struct {
	worker   int
	messages chan string
	wg       sync.WaitGroup
	closed   chan struct{}
	once     sync.Once
}

// NewReporter 嘻嘻嘻
func NewReporter(worker, buffer int) *Reporter {
	return &Reporter{
		worker:   worker,
		messages: make(chan string, buffer),
		closed:   make(chan struct{}),
	}
}

func (r *Reporter) Run(stop <-chan struct{}) {
	go func() {
		<-stop
		fmt.Println("stop...")
		r.Shutdown()
	}()
	for i := 0; i < r.worker; i++ {
		r.wg.Add(1)
		go func() {
			defer r.wg.Done()
			for {
				select {
				case <-r.closed:
					fmt.Printf("report: %s\n", "close")
					return
				case msg := <-r.messages:
					fmt.Printf("report: %s\n", msg)
				}
			}
		}()
	}
	r.wg.Wait()
	fmt.Println("report workers exit...")
}

// 这里不必关闭 messages
// 因为 closed 关闭之后，发送端会直接丢弃数据不再发送
// Run 方法中的消费者也会退出
// Run 方法会随之退出
func (r *Reporter) Shutdown() {
	r.once.Do(func() { close(r.closed) })
}

//Report 模拟耗时
func (r *Reporter) Report(data string) {
	// 这个是为了及早退出
	// 并且为了避免我们消费者能力很强，发送者这边一直不阻塞，可能还会一直写数据
	select {
	case <-r.closed:
		fmt.Printf("reporter is closed, data will be discarded: %s \n", data)
	default:
	}

	select {
	case <-r.closed:
		fmt.Printf("reporter is closed, data will be discarded: %s \n", data)
	case r.messages <- data:
	}
}
