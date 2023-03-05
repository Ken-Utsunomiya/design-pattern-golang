package main

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Creation/Singleton/singleton"
)

func main() {

	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
	}

	fmt.Scanln()
}
