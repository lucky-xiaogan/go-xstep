package main

import (
	"fmt"
	"sync"
)

type XXData struct {
	work    int         //工作协程数量
	message chan string //消费 消息chan
	wg      sync.WaitGroup
	close   chan struct{}
	once    sync.Once
}

func NewXXData(work, buffer int) *XXData {
	return &XXData{
		work:    work,
		message: make(chan string, buffer),
		close:   make(chan struct{}),
	}
}

func (d *XXData) Run(stop chan struct{}) {
	go func() {
		<-stop
		d.Shutdown()
	}()

	for i := 0; i < d.work; i++ {
		d.wg.Add(1)
		go func() {
			defer d.wg.Done()
			for {
				select {
				case <-d.close:
					fmt.Println("close")
					return
				//消费 chan 数据
				case msg := <-d.message:
					fmt.Println(msg)
				}
			}
		}()
	}
	d.wg.Wait()
}

func (d *XXData) Shutdown() {
	d.once.Do(func() {
		close(d.close)
	})
}

func (d *XXData) SendData(name string) {
	select {
	case <-d.close:
		fmt.Println("send data close")
	default:
	}

	select {
	case <-d.close:
		fmt.Println("send data close")
	case d.message <- name:
	}
}
