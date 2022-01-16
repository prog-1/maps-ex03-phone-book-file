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
			name:    "just a single letter",
			names:   []string{"a"},
			numbers: []string{"0"},
			want:    map[string]string{"a": "0"},
		},
		{
			name:    "no name, only phone number",
			names:   []string{""},
			numbers: []string{"+37167812439"},
			want:    map[string]string{"": "+37167812439"},
		},
		{
			name:    "no phone numbers, only names",
			names:   []string{"Iļja", "Daniils"},
			numbers: []string{"", ""},
			want:    map[string]string{"Daniils": "", "Iļja": ""},
		},
		{
			name:    "one contact",
			names:   []string{"Sebastjans"},
			numbers: []string{"+37127604728"},
			want:    map[string]string{"Sebastjans": "+37127604728"},
		},
		{
			name:    "two unique contacts",
			names:   []string{"Andrejs", "Kirils"},
			numbers: []string{"+37125134696", "+37123207377"},
			want:    map[string]string{"Andrejs": "+37125134696", "Kirils": "+37123207377"},
		},
		{
			name:    "two repeating contacts",
			names:   []string{"Ričards", "Ričards"},
			numbers: []string{"+37126144762", "+37132547991"},
			want:    map[string]string{"Ričards": "+37126144762"},
		},
		{
			name:    "three unique contacts",
			names:   []string{"Aleksandrs", "Валерия", "Alīna"},
			numbers: []string{"+37177345910", "+3265807134", "+26734167784"},
			want:    map[string]string{"Aleksandrs": "+37177345910", "Alīna": "+26734167784", "Валерия": "+3265807134"},
		},
		{
			name:    "three repeating contacts",
			names:   []string{"Vladimirs", "Vladimirs", "Vladimirs"},
			numbers: []string{"+4112562870", "+37141390776", "+5423607491"},
			want:    map[string]string{"Vladimirs": "+4112562870"},
		},
		{
			name:    "unique names with the same phone number",
			names:   []string{"Aleksejs", "Anastasija", "Vladislavs", "Gļebs"},
			numbers: []string{"+37128167023", "+37128167023", "+37128167023", "+37128167023"},
			want:    map[string]string{"Aleksejs": "+37128167023", "Anastasija": "+37128167023", "Gļebs": "+37128167023", "Vladislavs": "+37128167023"},
		},
		{
			name:    "repeating names with the same phone number",
			names:   []string{"Sergejs", "Sergejs", "Sergejs", "Sergejs"},
			numbers: []string{"+37124971390", "+37124971390", "+37124971390", "+37124971390"},
			want:    map[string]string{"Sergejs": "+37124971390"},
		},
		{
			name:    "example from the task (and in the file 'phone_numbers.txt')",
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
