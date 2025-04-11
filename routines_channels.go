package main

import (
	"context"
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 100; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d, ", i)
	}
}

func printNumbers_(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(c)
}

func printNumbers__(c chan int) {
	for i := 100; i > 0; i-- {
		c <- i
		time.Sleep(50 * time.Millisecond)
	}
	close(c)
}

func main() {
	// go printNumbers() // Use go keyword to specify a go routine
	// printNumbers() // Processed on the main thread

	// c := make(chan int)
	// go printNumbers__(c)

	// // // Reading from the channel
	// // for num := range c {
	// // 	fmt.Printf("%v, ", num)
	// // }

	// // For/Switch to Retrieve values from Go channel
	// for value := range c {
	// 	switch value % 2 {
	// 	case 0:
	// 		fmt.Printf("%v is even\n", value)
	// 	case 1:
	// 		fmt.Printf("%v is odd\n", value)
	// 	default:
	// 		fmt.Println("%v is weird\n")
	// 	}
	// }

	// ch := make(chan string)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- fmt.Sprintf("hello ")
	// }()

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- fmt.Sprintf("World")
	// }()

	// for {
	// 	select {
	// 	case v := <-ch:
	// 		fmt.Printf("%s\n", v)
	// 	case <-time.After(3 * time.Second):
	// 		fmt.Println("Waited for 3 seconds")
	// 		os.Exit(0) // exiting the program after 3 seconds wait
	// 	}
	// }

	// Using Go Contexts
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task finished")
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context Done")
		err := ctx.Err()
		if err != nil {
			fmt.Printf("err: %s", err)
		}
	}
}
