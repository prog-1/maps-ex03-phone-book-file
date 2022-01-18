package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	a := len(names)
	for i := a - 1; i >= 0; i-- {
		m[names[i]] = numbers[i]
	}
	return m
}
func main() {
	var names, numbers []string
	f, err := os.Open("phone_numbers.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		a := scanner.Text()
		if i%2 == 0 {
			names = append(names, a)
		} else {
			numbers = append(numbers, a)
		}
		i++
	}
	defer f.Close()
	fmt.Println(createPhoneBook(names, numbers))
}
