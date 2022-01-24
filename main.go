package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	a := len(names)
	for i := a - 1; i >= 0; i-- {
		m[names[i]] = numbers[i]
	}
	return m
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("This program creates a map from names and numbers from a file.")
	var names, numbers []string
	f, err := os.Open("phone_numbers.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	isname := true
	for scanner.Scan() {
		if isname {
			names = append(names, scanner.Text())
			isname = false
		} else {
			numbers = append(numbers, scanner.Text())
			isname = true
		}
	}
	fmt.Println(createPhoneBook(names, numbers))

}
