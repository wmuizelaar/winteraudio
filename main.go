package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/wmuizelaar/winteraudio/easteregg"
)

func main() {

	if rand.Float32() < 0.5 {
		go easteregg.PlayWinter()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
