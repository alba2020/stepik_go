package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)


func main() {
    file, err := os.Open("task.data") 
    if err != nil { 
         fmt.Println("Unable to open file:", err) 
         return
    } 
    defer file.Close() 
  
	pos := 0

    reader := bufio.NewReader(file) 
    for { 
        line, err := reader.ReadString(';')
		pos++
        if err != nil { 
            if err == io.EOF { 
                break
            } else { 
                fmt.Println(err) 
                return
            } 
        }

		if line == "0;" {
			fmt.Print(line)
			fmt.Println(pos)
		}
    }
}