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

func readFile2(id int, scanner *bufio.Scanner, counter *SafeCounter) {
	scanner.Scan()
	text := scanner.Text()

	counter.mu.Lock()
	defer counter.mu.Unlock()
	// fmt.Println(id, text)
	num, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("err - ", err, text)
		counter.v += 0
		// panic(err)
	}

	counter.v += num
	// ch <- parse(text)
}
func main() {
	file, err := os.Open("../files/file-num.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// var wg sync.WaitGroup

	var counter SafeCounter = SafeCounter{}

	scanner := bufio.NewScanner(file)

	for id := range 10 {
		// go readFile(id, scanner, ch)
		go readFile2(id, scanner, &counter)
	}

	time.Sleep(time.Second)
	fmt.Println(counter.v)

	fmt.Println("Terminado")
}
