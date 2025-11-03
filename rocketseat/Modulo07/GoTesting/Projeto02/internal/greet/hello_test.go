package greet

import "testing"

func TestHelloName(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "English Greet",
			got:  helloName("en", "Jhon"),
			want: "Hello, Jhon",
		},
		{
			name: "Portugese Greet",
			got:  helloName("pt", "Jhon"),
			want: "Oi, Jhon",
		},
		{
			name: "Spanish Greet",
			got:  helloName("es", "Jhon"),
			want: "Hola, Jhon",
		},
		{
			name: "French Greet",
			got:  helloName("fr", "Jhon"),
			want: "Bonjour, Jhon",
		},
		{
			name: "Empty string name",
			got:  helloName("pt", ""),
			want: "Oi, Anonymous",
		},
		{
			name: "Empty Language",
			got:  helloName("", "Gabriel"),
			want: "??, Gabriel",
		},
		{
			name: "Invalid Language",
			got:  helloName("rs", "Gabriel"),
			want: "??, Gabriel",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("Got: %q | want: %q", tc.got, tc.want)
			}
		})
	}
}
