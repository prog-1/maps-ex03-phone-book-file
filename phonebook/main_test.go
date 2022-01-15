package main

func TestCreatePhonebook(t *testing.T) {
	tests := [] struct {
		name string
		names, numbers [string]
		want map[string]string
	}{
		name : "empty"
		names : []string{},
		numbers : []string {}
		want : map[string]string{},
	}
	for _, a :=range resrts {
		t.Run(a.names, func(t *testing.T)) {
			if got := createPhoneBook(a.number, a.names); !equal(got, a.want){
				t.Errorf("createPhoneBook() = %v, want %v", got, a.want)
}