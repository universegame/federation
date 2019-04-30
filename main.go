package main

import (
	"fmt"
	"github.com/universegame/universes"
)



func main() {
	go func() {
		for {
			go universes.Universe_producer{}.Product()
		}
	}()

	for {
		fmt.Println(<-universes.Wormhole)
	}

	//select {
	//case truck := <- universes.Wormhole:
	//	fmt.Println(truck)
	//}
	fmt.Println("finished")
}
