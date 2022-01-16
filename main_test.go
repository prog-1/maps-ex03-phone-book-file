package main

import (
	"reflect"
	"testing"
)

func TestCreatePhoneBook(t *testing.T) {
	tests := []struct {
		test_name string
		names     []string
		numbers   []string
		want      map[string]string
	}{
		{"empty",
			[]string{},
			[]string{},
			map[string]string{}},

		{"example from README with repeat",
			[]string{
				"Alina",
				"Deniss B",
				"Antons",
				"Alina",
				"Antons",
			},
			[]string{
				"+37126017505",
				"+37127785804",
				"+37123622588",
				"+37126505719",
				"+37128852154",
			},
			map[string]string{
				"Alina":    "+37126017505",
				"Antons":   "+37123622588",
				"Deniss B": "+37127785804",
			}},

		{"example from README without repeat",
			[]string{
				"Alina",
				"Antons",
				"Deniss B",
			},
			[]string{
				"+37126505719",
				"+37128852154",
				"+37127785804",
			},
			map[string]string{
				"Alina":    "+37126505719",
				"Antons":   "+37128852154",
				"Deniss B": "+37127785804",
			}},
		{"one name - one number",
			[]string{
				"Alesja",
			},
			[]string{
				"+37122331231",
			},
			map[string]string{
				"Alesja": "+37122331231",
			}},
		{"duplicate",
			[]string{
				"Alesja",
				"Alesja",
				"Alesja",
			},
			[]string{
				"+37122331231",
				"+37122331231",
				"+37122331231",
			},
			map[string]string{
				"Alesja": "+37122331231",
			}},
		{"1 name - 3 numbers",
			[]string{
				"Alesja",
				"Alesja",
				"Alesja",
			},
			[]string{
				"+37122331231",
				"+37122331280",
				"+37122331299",
			},
			map[string]string{
				"Alesja": "+37122331231",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			if got := createPhoneBook(tt.names, tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPhoneBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
