package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Look for the file
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No file specified")
	}

	//Read the file
	dat, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	//Create the memory
	memory := make([]uint, 30000, 30000)
	//Create the pointer to the memory
	pos := 0
	jump := 0
	skip := false
	//The reader
	reader := bufio.NewReader(os.Stdin)

	//Process the file
	for i := 0; i < len(dat); i++ {
		if string(dat[i]) == ">" {
			pos++
		} else if string(dat[i]) == "<" && !skip {
			pos--
		} else if string(dat[i]) == "+" && !skip {
			memory[pos]++
		} else if string(dat[i]) == "-" && !skip {
			memory[pos]--
		} else if string(dat[i]) == "." && !skip {
			fmt.Printf("%c", rune(memory[pos]))
		} else if string(dat[i]) == "," && !skip {
			input, _ := reader.ReadString('\n')
			memory[pos] = uint(rune(input[0]))
		} else if string(dat[i]) == "[" && !skip {
			jump = i
			if memory[pos] == 0 {
				skip = true
			} else {
				skip = false
			}
		} else if string(dat[i]) == "]" {
			if memory[pos] != 0 {
				i = jump
			} else {
				skip = false
			}
		}
	}
}
