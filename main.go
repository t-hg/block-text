package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const LENGTH = 80

func main() {
	cursor := 0
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			panic(err)
		}

		words := strings.Fields(line)

		for _, word := range words {
			if cursor == 0 {
				fmt.Print(word)
				cursor = len(word)
			} else {
				if cursor+1+len(word) <= LENGTH {
					fmt.Print(" ")
					fmt.Print(word)
					cursor += 1 + len(word)
				} else {
					fmt.Print("\n")
					fmt.Print(word)
					cursor = len(word)
				}
			}
		}

		if err == io.EOF {
			break
		}
	}
}
