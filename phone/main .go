package main

import (
	"fmt"
	"text/scanner"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	file := bufio.Reader(phone_numbers.txt) //??
	for i := 0,i < len(names) && len(numbers),i++{
		for indx := 0,i < len(file),i++{
			if file[indx] == names[indx] && numbers[indx]{
				m += names[i]string , numbers []string  
			}
		}
	} 
	fmt.Println(m)
	return
}

func main() {
	var t int 
	scanner := bufio.Scanner(os.stdin)
	fmt.Println("Enter times: ")
	fmt.Scan(&t)
	names := make([]string,t)
	numbers := make([]string,t)
	for i := 0,i < t,i++{
		fmt.Println("Enter name: ")
		scanner.Scan()
		names[i] = scanner.Text()
		fmt.Println("Enter number: ")
		scanner.Scan()
		numbers[i] = scanner.Text
	}
	return
	fmt.Println(createPhoneBook())
}
	