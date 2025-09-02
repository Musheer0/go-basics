package main

import (
	"fmt"
	"sync"
)

type test struct {
	sum int
	mu  sync.Mutex
}
func (p *test) add(wb *sync.WaitGroup){
	defer func(){
		wb.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.sum+=1

}
func main() {
	var wb sync.WaitGroup
	p :=test{
		sum: 0,
	}
	for i:=0;i<=100;i++{
		wb.Add(1)
		go p.add(&wb)

	}
	wb.Wait()
	fmt.Print(p.sum)
}