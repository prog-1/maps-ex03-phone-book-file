package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
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
	var str []string
	for scanner.Scan() {
		a := scanner.Text()
		str = append(str, a)

	}
	defer f.Close()
	b := len(str)
	for i := 0; i < b; i++ {
		if i%2 == 0 {
			names = append(names, str[i])
		} else {
			numbers = append(numbers, str[i])
		}
	}
	fmt.Println(createPhoneBook(names, numbers))
}
