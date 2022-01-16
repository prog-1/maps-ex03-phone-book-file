package main

import (
	"reflect"
	"testing"
)

func TestPhonebook(t *testing.T) {
	tests := []struct {
		name    string
		names   []string
		numbers []string
		want    map[string]string
	}{
		{
			name:    "empty",
			names:   nil,
			numbers: nil,
			want:    map[string]string{},
		},
		{
			name:    "example from exercise",
			names:   []string{"Alina", "Deniss B", "Antons", "Alina", "Antons"},
			numbers: []string{"+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154"},
			want:    map[string]string{"Alina": "+37126017505", "Antons": "+37123622588", "Deniss B": "+37127785804"},
		},
		{
			name:    "example with unique contacts",
			names:   []string{"Alex", "Vladimir", "Dmitry", "Daniel"},
			numbers: []string{"+37151713242", "+37191942672", "+37166980350", "+37185365063"},
			want:    map[string]string{"Alex": "+37151713242", "Vladimir": "+37191942672", "Dmitry": "+37166980350", "Daniel": "+37185365063"},
		},
		{
			name:    "example with only repeated contacts",
			names:   []string{"Vladimir", "Dmitry", "Daniel", "Dmitry", "Vladimir", "Daniel", "Dmitry"},
			numbers: []string{"+37183460228", "+37117548706", "+37171941446", "+37124314810", "+37138972685", "+37151246471", "+37172228266"},
			want:    map[string]string{"Vladimir": "+37183460228", "Dmitry": "+37117548706", "Daniel": "+37171941446"},
		},
		{
			name:    "example with multiple repeated contacts",
			names:   []string{"Alex", "Vladimir", "Dmitry", "Vladimir", "Daniel", "Alex", "Daniel", "Vladimir"},
			numbers: []string{"+37194566528", "+37198992079", "+37175819718", "+37168512597", "+37122345593", "+37180531738", "+37156100264", "+37147086024"},
			want:    map[string]string{"Alex": "+37194566528", "Vladimir": "+37198992079", "Dmitry": "+37175819718", "Daniel": "+37122345593"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createPhoneBook(tt.names, tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PhoneBook() = %v, want = %v", got, tt.want)
			}
		})
	}
}
