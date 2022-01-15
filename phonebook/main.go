package main

import (
	"bufio"
	"strings"
)

func  createPhoneBook(names []string, numbers []string) map[string]string{
	m := make(map[string]string)
	for i,key := range names {
		m[key] = numbers[i]
	}
	return m 
}
	
func main() {
	name := []string 
	numbers := []string
	// я не понимаю в чем моя проблема в 17 строке 
	f, err := os.Open("phone_numbers.txt")
	bufio.NewScanner(a)
	is name = true
	for f.Scan() {
		if is name {
			names = append(names, s.Text())
			is name = false
		} else {
			numbers = append(numbers, s.Text())
			isname = true
		}
		}
	}
	fmt.Println(create.Phone Book(names, numbers))
}
