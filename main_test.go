package main

import (
	"reflect"
	"testing"
)

func TestCreatePhoneBook(t *testing.T) {
	tests := []struct {
		name            string
		names, phonenum []string
		want            map[string]string
	}{
		{
			name:     "empty",
			names:    []string{},
			phonenum: []string{},
			want:     map[string]string{},
		},
		{
			name:     "Only phone number",
			names:    []string{""},
			phonenum: []string{"+3712963883830"},
			want:     map[string]string{"": "+3712963883830"},
		},
		{
			name:     "Only name",
			names:    []string{"Vova"},
			phonenum: []string{""},
			want:     map[string]string{"Vova": ""},
		},
		{
			name:     "one contact",
			names:    []string{"tony stark"},
			phonenum: []string{"+79122773413"},
			want:     map[string]string{"tony stark": "+79122773413"},
		},
		{
			name:     "with capital letters",
			names:    []string{"kOt VasILIJ"},
			phonenum: []string{"+37124425678"},
			want:     map[string]string{"kOt VasILIJ": "+37124425678"},
		},
		{
			name:     "different contacts",
			names:    []string{"Neznakomec", "Znakomec"},
			phonenum: []string{"+37100000000", "+37123456789"},
			want:     map[string]string{"Neznakomec": "+37100000000", "Znakomec": "+37123456789"},
		},
		{
			name:     "repeating contacts",
			names:    []string{"Genij", "Genij"},
			phonenum: []string{"+37112345678", "+37187654321"},
			want:     map[string]string{"Genij": "+37187654321"},
		},
		{
			name:     "duplicate names",
			names:    []string{"Alina", "Alina", "Antons", "Antons", "Deniss", "Deniss"},
			phonenum: []string{"+37126017505", "+37126017505", "+37123622588", "+37123622588", "+37128852154", "+37128852154"},
			want:     map[string]string{"Alina": "+37126017505", "Antons": "+37123622588", "Deniss": "+37128852154"},
		},
		{
			name:     "special letters",
			names:    []string{"Polīna", "Bērze", "Alīna"},
			phonenum: []string{"+3713344789", "+92754327890", "+31466547890"},
			want:     map[string]string{"Polīna": "+3713344789", "Bērze": "+92754327890", "Alīna": "+31466547890"},
		},
		{
			name:     "example which was given",
			names:    []string{"Alina", "Deniss B", "Antons", "Alina", "Antons"},
			phonenum: []string{"+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154"},
			want:     map[string]string{"Alina": "+37126017505", "Antons": "+37123622588", "Deniss B": "+37127785804"},
		},
		{
			name:     "spaces before the name",
			names:    []string{"    Alina"},
			phonenum: []string{"+37126017505"},
			want:     map[string]string{"    Alina": "+37126017505"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createPhoneBook(tt.names, tt.phonenum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPhoneBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
