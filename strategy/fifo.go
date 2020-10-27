package strategy

import "fmt"

type Fifo struct {
}

func (l *Fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}
