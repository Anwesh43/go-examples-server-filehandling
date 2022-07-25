package main 

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type MyFile interface {
	WriteString(string) (int, error)
}

func WriteFileAfterSomeTime(n int, file MyFile, chan1 chan bool) {
	for i := 0; i< n; i++ {
		file.WriteString(fmt.Sprintf("Hello World %d \n", i))
		for j := 0; j < i + 1; j++ {
			fmt.Print(".")
		}
		fmt.Println("");
		time.Sleep(time.Second)
	}
	chan1 <- true 
}
func CreateMore(filename string) {
	file, err := os.Create(filename)
	chan1 := make(chan bool)
	if err != nil  {
		log.Fatalf("error creating file")
	}
	
	go WriteFileAfterSomeTime(10, file, chan1)
	<- chan1 
	file.Close()
	data, err1 := ioutil.ReadFile(filename)
	if err1 != nil {
		log.Fatalf("error reading file")
	}
	fmt.Println(string(data))
	
}

func main() {
	CreateMore("hello_world.txt")
}