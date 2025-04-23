package main

import (
	"fmt"
	// "go_tuts/readinputs"
	// "os"
	"strconv"
	"strings"
	"time"
)

func oddNumbers(c chan []string) {
	for i := 1; i < 100; i = i + 2 {
		c <- []string{strconv.Itoa(i), "odd"}
		time.Sleep(800 * time.Millisecond)
	}
	close(c)
}

func evenNumbers(c chan []string) {
	for i := 2; i < 100; i = i + 2 {
		time.Sleep(400 * time.Millisecond)
		c <- []string{strconv.Itoa(i), "even"}
		time.Sleep(400 * time.Millisecond)
	}
	close(c)
}

func main() {
	str := fmt.Sprintf("not \n")
	fmt.Printf("::::::::::This is %s ::::::::::\n", str)
	str = strings.TrimSpace(str)
	fmt.Printf("::::::::::This is %s ::::::::::", str)
	// ch := make(chan []string)
	// go oddNumbers(ch)
	// go evenNumbers(ch)
	// fmt.Println(readinputs.ReadTxt("APIKEY"))
	// go writeAPI()

	// for rec := range ch {
	// 	num, str := rec[0], rec[1]
	// 	fmt.Printf("this is %v:{%v}\n", str, num)
	// }

	// for {
	// 	select {
	// 	case v := <-ch:
	// 		fmt.Printf("This is V: %v \n", v)
	// 	case <-time.After(300 * time.Millisecond):
	// 		fmt.Println("Waited too long for response")
	// 		os.Exit(0)
	// 	}
	// }

}
