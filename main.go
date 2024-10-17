package main

import (
	"example.com/m/v2/services"
	"fmt"
	"math/rand"
)

func main() {
	runQueue()
}

func runQueue() {
	loop := 500
	qService := services.NewQueue(int64(loop))
	list := make([]int, 0)
	for i := 0; i <= loop; i++ {
		val := rand.Intn(loop)
		fmt.Printf("rand : %d\n", val)
		go func() {
			qService.Lock.Lock()

			list = append(list, val)
			qService.InputQueue(val)
			qService.Lock.Unlock()
		}()

	}

	for i := 0; i <= loop; i++ {
		go func() {
			qService.PrintQueue()
		}()
	}
	qService.WaitQueue()

}
