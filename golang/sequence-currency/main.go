package main

import (
	"fmt"
	"sync"
	"time"
)

// trick use chan chan struct{}

func main() {

	nS := time.Now()
	NormalCurrency()
	nE := time.Since(nS)
	fmt.Println("normal cost -----------------------", nE)
	sS := time.Now()
	SequenceCurrency()
	sE := time.Since(sS)
	fmt.Println("sequence cost -----------------------", sE)
}

// no order
func NormalCurrency() {
	printChan := make(chan string)
	wg := sync.WaitGroup{}
	go PrintNormal(printChan)
	Normal2(printChan, "1", &wg)
	Normal2(printChan, "2", &wg)
	Normal1(printChan, "3", &wg)
	Normal1(printChan, "4", &wg)
	wg.Wait()
}

func Normal1(printChan chan string, str string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(str string, wg *sync.WaitGroup) {
		time.Sleep(time.Second * 1)
		printChan <- "Normal1:" + str
		wg.Done()
	}(str, wg)
}

func Normal2(printChan chan string, str string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(str string, wg *sync.WaitGroup) {
		time.Sleep(time.Second * 2)
		printChan <- "Normal2:" + str
		wg.Done()
	}(str, wg)
}

func PrintNormal(msg chan string) {
	for {
		fmt.Println(<-msg)
	}
}

// in order
func SequenceCurrency() {
	printChan := make(chan chan string, 1024)
	wg := sync.WaitGroup{}
	go PrintSequence(printChan)
	Sequence2(printChan, "1", &wg)
	Sequence2(printChan, "2", &wg)
	Sequence1(printChan, "3", &wg)
	Sequence1(printChan, "4", &wg)
	wg.Wait()
}

func Sequence1(ch chan chan string, str string, wg *sync.WaitGroup) {
	c := make(chan string)
	ch <- c
	wg.Add(1)
	go func(str string) {
		time.Sleep(time.Second * 1)
		c <- "Sequence1:" + str
		wg.Done()
	}(str)
}

func Sequence2(ch chan chan string, str string, wg *sync.WaitGroup) {
	c := make(chan string)
	ch <- c
	wg.Add(1)
	go func(str string) {
		time.Sleep(time.Second * 2)
		c <- "Sequence2:" + str
		wg.Done()
	}(str)
}

func PrintSequence(msg chan chan string) {
	for {
		// ensure order
		c := <-msg
		r := <-c
		fmt.Println(r)
	}
}
