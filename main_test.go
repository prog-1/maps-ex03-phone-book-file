package main

import "testing"

func TestCreatePhoneBook(t *testing.T) {
	for _, tc := range []struct {
		names   []string
		numbers []string
		want    map[string]string
	}{
		{
			[]string{""},
			[]string{""},
			map[string]string{"": ""},
		},
		{
			[]string{"Artur"},
			[]string{""},
			map[string]string{"Artur": ""},
		},
		{
			[]string{""},
			[]string{"+1234567890"},
			map[string]string{"": "+1234567890"},
		},
		{
			[]string{"Alina", "Deniss B", "Antons", "Alina", "Antons"},
			[]string{"+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154"},
			map[string]string{"Alina": "+37126017505", "Antons": "+37123622588", "Deniss B": "+37127785804"},
		},
		{
			[]string{},
			[]string{},
			map[string]string{},
		},
		{
			[]string{"Alina", "Alina", "Alina", "Alina"},
			[]string{"+37143780565", "+371705", "+371234234237505", "+37126017505"},
			map[string]string{"Alina": "+37143780565"},
		},
		{
			[]string{"Tihons", "Arina", "Egors"},
			[]string{"+37134988264", "+37192744934", "+3712281337"},
			map[string]string{"Tihons": "+37134988264", "Arina": "+37192744934", "Egors": "+3712281337"},
		},
	} {
		got := CreatePhoneBook(tc.names, tc.numbers)
		if !equal(got, tc.want) {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}
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
