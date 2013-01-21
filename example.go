package main

import (
	"fmt"
	"github.com/melvinmt/gt"
)

var g = &gt.Build{
	Index: gt.Strings{
		"homepage-title": {
			"en":    "Hello World!",
			"es":    "¡Hola mundo!",
			"zh-CN": "你好世界!",
		},
		"homepage-welcome": {
			"en":    "Welcome to %s, %s.",
			"es-LA": "Bienvenidos a %s, %s.",
			"tr-TR": "%s, %s'ya hoşgeldiniz.",
			"nl":    "Welkom bij %s, %s",
		},
		"dashboard-notice": {
			"en":    "Hello %s#name, you have a new message from %s#friend.",
			"tr-TR": "%s#name merhaba, %s#friend'den yeni bir mesaj var.",
		},
		"invoice": {
			"en":    "You need to pay %5.2f#amount dollars in %d#days days.",
			"pt-BR": "Você precisa pagar %5.2f#amount dólares em %d#days dias.",
		},
	},
	Origin: "en", // the language you'll be translating from
}

func main() {
	// Key based:
	g.SetTarget("es")
	s1 := g.T("homepage-title")
	fmt.Println(s1) // outputs: ¡Hola mundo!

	// String based:
	g.SetTarget("zh-CN")
	s2 := g.T("Hello World!")
	fmt.Println(s2) // outputs: 你好世界!

	// Parse arguments:
	g.SetTarget("nl")
	s3 := g.T("Welcome to %s, %s.", "Github", "Melvin")
	fmt.Println(s3) // outputs: Welkom bij Github, Melvin

	// As you can see in the previous example, you can use regular printf verbs
	// to parse arguments. The problem that you're facing here is that the order
	// of words is different in some languages:

	g.SetOrigin("es-LA")
	g.SetTarget("tr-TR")
	s4, err := g.Translate("Bienvenidos a %s, %s.", "Github", "Melvin")
	if err != nil {
		panic(err)
	}
	fmt.Println(s4)
	// This outputs: Github, Melvin'ya hoşgeldiniz. This is roughly translated as:
	// Welcome to Melvin, Github.  Which is NOT what you want. You can solve this with
	// tag notation.

	// Tag notation:
	g.SetOrigin("en")
	s5 := g.T("Hello %s#name, you have a new message from %s#friend.")
	fmt.Println(g.T(s5, "Melvin", "Alex"))
	// Outputs: Melvin merhaba, Alex'den yeni bir mesaj var. 
	// Which is in a different order, but correctly translated.

	// You can use any legal printf verb in combination with tags:
	g.SetOrigin("pt-BR")
	g.SetTarget("en")
	fmt.Printf(g.T("Você precisa pagar %5.2f#amount dólares em %d#days dias."), 499.95, 5)
	// Outputs: You need to pay 499.95 dollars in 5 days.
}
