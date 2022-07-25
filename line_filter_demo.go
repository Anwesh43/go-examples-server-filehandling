package main 

import (
	"fmt"
	"os"
	"bufio"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0 
	message := ""
	file_name := ""
	defer func() {
		file, _ := os.Create(file_name)
		file.WriteString(message)
	}()
	for scanner.Scan() {
		data := scanner.Text()
		if (i == 0) {
			file_name = data 
		} else {
			message = message + data + "\n"
		}
		if (i == 9) {
			break
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error", err)
	}
}
