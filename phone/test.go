package main

import (
	"fmt"
	"testing"
)

func TestCreatePhoneBook(t *testing.T) {
	for _, tc := range []struct {
		name    string
		names   []string
		numbers []string
		want    map[string]string
	}{
		{"simple", []string{"Alina"}, []string{"+37126017505"}, map[string]string{"Alina": "+37126017505"}},
		{"isn't in phone book", []string{"Anna"}, []string{"+37125896457"}, map[string]string{}},
		{"empty", []string{}, []string{}, map[string]string{}},
		{"is and iasn't", []string{"Alina", "Anna", "Antons"}, []string{"+37126017505", "+37125896457", "+37123622588"}, map[string]string{"Alina": "+37126017505", "Antons": "+37123622588"}},
		{"random symbols", []string{"fafmflh"}, []string{"+356461466496"}, map[string]string{}},
	} {
		if got := createPhoneBook(tc.names, tc.numbers); got != tc.want {
			fmt.Printf("FAIL %s: got %v, want %v", tc.names, tc.numbers, got, tc.want)

		}
	}
}
