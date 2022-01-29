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
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var names, numbers []string
	i := 0
	for s.Scan() {
		if i == 0 {
			names = append(names, s.Text())
			i = 1
		} else {
			numbers = append(numbers, s.Text())
			i = 0
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
