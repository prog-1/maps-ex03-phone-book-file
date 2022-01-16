package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	pb := make(map[string]string)
	for i := len(names) - 1; i >= 0; i-- {
		pb[names[i]] = numbers[i]
	}
	return pb
}

func main() {
	file, err := os.Open("phone_numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var names, numbers []string
	isName := true
	for scanner.Scan() {
		if isName {
			names = append(names, scanner.Text())
			isName = false
		} else {
			numbers = append(numbers, scanner.Text())
			isName = true
		}
	}
	fmt.Print(createPhoneBook(names, numbers))
}
