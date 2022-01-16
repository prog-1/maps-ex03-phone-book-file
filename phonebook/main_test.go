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