package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreatePhoneBook(names []string, numbers []string) map[string]string {
	x := len(names)
	phonebook := make(map[string]string)
	for ; x > 0; x-- {
		phonebook[names[x-1]] = numbers[x-1]
	}
	return phonebook
}

func main() {
	var names, numbers []string
	f, err := os.Open("phone_numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for tmp := 1; s.Scan(); tmp++ {
		input := s.Text()
		if tmp%2 != 0 {
			names = append(names, input)
		} else {
			numbers = append(numbers, input)
		}
	}
	fmt.Println(CreatePhoneBook(names, numbers))
}
