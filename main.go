package main

import (
	"fmt"
	"time"
	"udemy/go/model"
)

func fibbonaci(total int, c chan int, finish chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("masuk")
			x, y = y, x+y
		case <-finish:
			fmt.Println("finish")
			close(c)
			return
		}
	}
}
func main() {
	// ball := make(chan *int)
	// go player("hanny", ball)
	// go player("fransisco", ball)

	// init := 0
	// ball <- &init

	// time.Sleep(1 * time.Second)
	// <-ball

	c := make(chan int)
	finish := make(chan bool)

	go func(c chan int) {
		index := 0
		for {
			select {
			case c <- index:
				index++
			case <-finish:
				close(c)
				return
			}
		}
	}(c)

	go func(c chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		finish <- true
	}(c)

	<-finish

	person := model.Human{Person: model.Asian{Height: 10, Name: "him"}}
	person.DoSpeak()
	person = model.Human{Person: model.Western{Height: 10, Name: "him"}}
	person.DoSpeak()

}

func player(name string, ball chan *int) {
	for {

		getBall := <-ball
		*getBall++
		fmt.Println(name, "hit ball: ", *getBall)
		time.Sleep(50 * time.Millisecond)
		ball <- getBall
	}
}
