package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sendRandNum := make(chan int)
	sendDoubleNum := make(chan int)
	func() {
		wg.Add(2)
		go createTenRandomNum(sendRandNum, &wg)
		go doubleRandomNum(sendRandNum, sendDoubleNum, &wg)
	}()
	go func() {
		wg.Wait()
		close(sendRandNum)
		close(sendDoubleNum)
	}()
	for num := range sendDoubleNum {
		if num == 10201 {
			break
		}
		fmt.Println(num)
	}
}

func createTenRandomNum(sendNum chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var randNum int
	for i := 0; i < 10; i++ {
		randNum = rand.Intn(101)
		sendNum <- int(randNum)
	}
	sendNum <- int(101)
}

func doubleRandomNum(getNum chan int, sendDoubleNum chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range getNum {
		num = int(math.Pow(float64(num), 2))
		sendDoubleNum <- int(num)
	}
}
