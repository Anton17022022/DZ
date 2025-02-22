package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

const (
	numRandomNumbers = 10
	maxRandomNumber  = 101
	stopSignal       = 101
	stopCondition    = 10201
)

func main() {
	var wg sync.WaitGroup
	sendRandNum := make(chan int)
	sendDoubleNum := make(chan int)
	wg.Add(2)
	go createTenRandomNum(sendRandNum, &wg)
	go doubleRandomNum(sendRandNum, sendDoubleNum, &wg)
	go func() {
		wg.Wait()
		close(sendRandNum)
		close(sendDoubleNum)
	}()
	for num := range sendDoubleNum {
		if num == stopCondition {
			break
		}
		fmt.Println(num)
	}
}

func createTenRandomNum(sendNum chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var randNum int
	for i := 0; i < numRandomNumbers; i++ {
		randNum = rand.Intn(maxRandomNumber)
		sendNum <- int(randNum)
	}
	sendNum <- int(stopSignal)
}

func doubleRandomNum(getNum chan int, sendDoubleNum chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range getNum {
		num = int(math.Pow(float64(num), 2))
		sendDoubleNum <- int(num)
	}
}
