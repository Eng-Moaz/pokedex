package main

import "strings"

func cleanInput(text string) []string{
	clean := strings.Fields(text)
	return clean
}


