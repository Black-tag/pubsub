package main

import (
	"fmt"
	"sync"

	
)



func publisher (wg *sync.WaitGroup, msgChan chan string) {
	defer wg.Done()
	for i:= range 10 {
		msgChan <- fmt.Sprintf("message %d", i)
	}
	close(msgChan)
	
}



func subscriber (id int, wg *sync.WaitGroup, subChan chan string) {
	defer wg.Done()
	for message := range subChan {
		fmt.Printf("subscriber %d recieved message: %s\n", id, message)
	}
	

}


func broadCaster(wg *sync.WaitGroup, msgChan chan string, subscribers []chan string) {
	defer wg.Done()
	for msg := range msgChan {
		for _, subChan := range subscribers {
			subChan <- msg
		}
	}
	for _, subChan := range subscribers {
		close(subChan)
		
	}
}




func main() {

	msgChan := make(chan string)
	subscribers := make([]chan string, 3)

	for i:= range 3 {
		subscribers[i] = make(chan string)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go publisher(wg, msgChan)

	wg.Add(1)
	go broadCaster(wg, msgChan, subscribers)
	for i:= range 3 {

		wg.Add(1)
		go subscriber(i, wg, subscribers[i])
	}
	wg.Wait()

}