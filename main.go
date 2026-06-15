package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    )

func main(){
	scanner := bufio.NewScanner(os.Stdin)

		for {
			fmt.Print("Pokedex > ")

			if !scanner.Scan() {
				break 
			}

			input := scanner.Text()
			cleaned := strings.ToLower(strings.TrimSpace(input))
			fields := strings.Fields(cleaned)

			if len(fields) == 0 {
				continue
			}

			command := fields[0]
			
			cmd, ok := AllCommands[command]
			if !ok{
				fmt.Println("Invalid command")
				continue
			}
			cmd.callback()
		}
	if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		}
}


