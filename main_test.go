package main

import (
	"reflect"
	"testing"
)

func TestCreatePhoneBook(t *testing.T) {
	tests := []struct {
		name           string
		names, numbers []string
		want           map[string]string
	}{
		{
			name:    "empty",
			names:   []string{},
			numbers: []string{},
			want:    map[string]string{},
		},
		{
			name:    "letter",
			names:   []string{"d"},
			numbers: []string{"0"},
			want:    map[string]string{"a": "0"},
		},
		{
			name:    "only number",
			names:   []string{""},
			numbers: []string{"+3711231241241"},
			want:    map[string]string{"": "+3711231241241"},
		},
		{
			name:    "only names",
			names:   []string{"asd", "asdasd"},
			numbers: []string{"", ""},
			want:    map[string]string{"asd": "", "asdasd": ""},
		},
		{
			name:    "one contact",
			names:   []string{"asd"},
			numbers: []string{"+3711231241241"},
			want:    map[string]string{"asd": "+3711231241241"},
		},
		{
			name:    "two unique contacts",
			names:   []string{"asd", "asdasd"},
			numbers: []string{"+3711231241241", "+371123124124"},
			want:    map[string]string{"asd": "+3711231241241", "asdasd": "+371123124124"},
		},
		{
			name:    "two repeating contacts",
			names:   []string{"asd", "asd"},
			numbers: []string{"+3711231241241", "+3711231241241"},
			want:    map[string]string{"asd": "+3711231241241"},
		},
		{
			name:    "names with the same number",
			names:   []string{"asd", "asdasd", "as", "ad"},
			numbers: []string{"+3711231241241", "+3711231241241", "+3711231241241", "+3711231241241"},
			want:    map[string]string{"asd": "+3711231241241", "asdasd": "+3711231241241", "as": "+3711231241241", "ad": "+3711231241241"},
		},
		{
			name:    "example from the task",
			names:   []string{"Alina", "Deniss B", "Antons", "Alina", "Antons"},
			numbers: []string{"+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154"},
			want:    map[string]string{"Alina": "+37126017505", "Antons": "+37123622588", "Deniss B": "+37127785804"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createPhoneBook(tt.names, tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPhoneBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
