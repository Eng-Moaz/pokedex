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
        if scanner.Scan(){
            text := scanner.Text()
            cleaned := strings.ToLower(strings.TrimSpace(text))

            if len(cleaned) == 0{
            continue
            }
            command := strings.Fields(cleaned)[0]
            fmt.Println("Your command was:", command)
        }
    }
}



