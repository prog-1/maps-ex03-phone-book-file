# Phone Book From File (Exercise)

In this exercise, you need to write a program that creates a phone book storing names
and phone numbers of your contacts from a file.

Write a program that asks a user to enter names and phone number of your contacts from `phone_numbers.txt`,
stores them in a map, and prints the map.

Write a function `func createPhoneBook(names []string, numbers []string) map[string]string`,
where `names` is a slice of names entered by the user and `numbers` is a slice of phone numbers read from the file.

A name may be entered multiple times. In this case, the map must store the **first** phone number occurrence.

Create tests for the `createPhoneBook` function.
