package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  int
}

func readFile(id int, scanner *bufio.Scanner, ch chan string) {
	scanner.Scan()
	text := scanner.Text()
	// fmt.Println(id, text)

	ch <- text
}

func readFile2(id int, scanner *bufio.Scanner, counter *SafeCounter, wg *sync.WaitGroup) bool {
	// func readFile2(scanner *bufio.Scanner, counter *SafeCounter) {
	ok := scanner.Scan()
	if !ok {
		return false
	}
	text := scanner.Text()

	// fmt.Println(text)

	counter.mu.Lock()
	defer counter.mu.Unlock()
	defer wg.Done()
	// fmt.Println(id, text)
	num, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("err - ", err, text)
		counter.v += 0
		return true
		// panic(err)
	}

	counter.v += num
	return true
	// ch <- parse(text)
}

func main() {
	now := time.Now()
	file, err := os.Open("../files/file-num.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineChan := make(chan string)
	numbersToSumChann := make(chan int)
	scanner := bufio.NewScanner(file)

	const workersCounter = 10

	go func() {
		for scanner.Scan() {
			lineChan <- scanner.Text()
		}
		close(lineChan)
	}()

	var wg sync.WaitGroup

	for i := 0; i < workersCounter; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for line := range lineChan {
				num, err := strconv.Atoi(line)
				if err != nil {
					fmt.Println("err - ", err, line)
					numbersToSumChann <- 0
				}

				numbersToSumChann <- num
			}
		}(i)
	}

	var counter = 0
	var wg2 sync.WaitGroup
	wg2.Add(1)

	wg2.Go(func() {
		for num := range numbersToSumChann {
			counter += num
		}
	})

	wg.Wait()
	close(numbersToSumChann)
	// wg2.Wait()
	// var counter SafeCounter = SafeCounter{}

	fmt.Println("Terminado - Soma:", counter)
	fmt.Println("Tempo: ", time.Since(now))
}
