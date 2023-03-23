package alarm

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	for true {
		time.Sleep(delay * time.Second)
		fmt.Printf("The time is %d.%d.%d: %s\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second(), text)
	}
}

func main() {
	go Remind("Time to eat", 10)
	go Remind("Time to work", 30)
	go Remind("Time to sleep", 60)
	select {}
}
