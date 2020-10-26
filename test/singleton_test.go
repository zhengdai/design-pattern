package test

import (
	sync_mutex "design-pattern/signleton/sync-mutex"
	sync_once "design-pattern/signleton/sync-once"
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 30; i++ {
		go sync_mutex.GetInstance()
	}
	fmt.Scanln()
}

func TestSingletonOnce(t *testing.T) {
	for i := 0; i < 30; i++ {
		go sync_once.GetInstance()
	}
	fmt.Scanln()
}
