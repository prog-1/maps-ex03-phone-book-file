package main

import "testing"

func equal(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestPhonebook(t *testing.T) {
	tests := []struct {
		name    string
		names   []string
		numbers []string
		want    map[string]string
	}{
		{
			name:    "example from exercise",
			names:   []string{"Alina", "Deniss B", "Antons", "Alina", "Antons"},
			numbers: []string{"+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154"},
			want:    map[string]string{"Alina": "+37126017505", "Antons": "+37123622588", "Deniss B": "+37127785804"},
		},
		{
			name:    "repeated contact",
			names:   []string{"Alina", "Alina", "Alina", "Alina", "Alina"},
			numbers: []string{"+37126017505", "+37126017505", "+37126017505", "+37126017505", "+37126017505"},
			want:    map[string]string{"Alina": "+37126017505"},
		},
		{
			name:    "empty",
			names:   nil,
			numbers: nil,
			want:    map[string]string{},
		},
		{
			name:    "3 uniqe contacts",
			names:   []string{"Alina", "Deniss B", "Antons"},
			numbers: []string{"+37126017505", "+37127785804", "+37123622588"},
			want:    map[string]string{"Alina": "+37126017505", "Deniss B": "+37127785804", "Antons": "+37123622588"},
		},
	}

	for _, a := range tests {
		t.Run(a.name, func(t *testing.T) {
			if got := createPhoneBook(a.names, a.numbers); !equal(got, a.want) {
				t.Errorf("createPhoneBook() = %v, want %v", got, a.want)
			}
		})
	}
}
