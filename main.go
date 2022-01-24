package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	for i := len(names) - 1; i >= 0; i-- {
		m[names[i]] = numbers[i]
	}
	return m
}

func main() {
	f, err := os.Open("phone_numbers.txt")
	if err != nil {
		log.Fatalf("open: %v", err) // or just log.Fatal(err)
	}
	defer f.Close()

	var names, numbers []string
	s := bufio.NewScanner(f)
	isname := true
	for s.Scan() {
		if isname {
			names = append(names, s.Text())
			isname = false
		} else {
			numbers = append(numbers, s.Text())
			isname = true
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("scanning: %v", err) // or just log.Fatal(err)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
