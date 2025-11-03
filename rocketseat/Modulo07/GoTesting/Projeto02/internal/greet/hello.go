package greet

import "fmt"

var Greetings = map[string]string{
	"pt": "Oi",
	"es": "Hola",
	"fr": "Bonjour",
	"en": "Hello",
}

const (
	defaultGreet = "??"
	defaultName  = "Anonymous"
)

func helloName(lang, name string) string {
	if name == "" {
		name = defaultName
	}

	greet, ok := Greetings[lang]
	if !ok || lang == "" {
		greet = defaultGreet
	}

	return fmt.Sprintf("%s, %s", greet, name)
}
