package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	for i, key := range names { // or as I did before:         for i := len(names) - 1; i >= 0; i-- {
		if _, found := m[key]; !found { //                             m[names[i]] = numbers[i]
			m[key] = numbers[i] //                             }
		} //                                                   return m
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
	for s.Scan() {
		names = append(names, s.Text())
		s.Scan()
		numbers = append(numbers, s.Text())
	}

	// or as I did before:
	//
	// isname := true
	// for s.Scan() {
	//	   if isname {
	//		   names = append(names, s.Text())
	//		   isname = false
	//	   } else {
	//		   numbers = append(numbers, s.Text())
	//		   isname = true
	//	   }
	// }

	if err := s.Err(); err != nil {
		log.Fatalf("scanning: %v", err) // or just log.Fatal(err)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
