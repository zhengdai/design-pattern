package strategy

import "fmt"

type Lfu struct {
}

func (l *Lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}
