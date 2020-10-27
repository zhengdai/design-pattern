package strategy

import "fmt"

type Lru struct {
}

func (l *Lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}
