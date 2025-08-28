package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang,
		"lang",
		"en",
		"The required language to greet. e.g en, ur...")
	flag.Parse()

	greeting := greetingsWorld(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

func greetingsWorld(lang language) string {
	greeting, ok := phrasebook[lang]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", lang)
	}

	return greeting
}
