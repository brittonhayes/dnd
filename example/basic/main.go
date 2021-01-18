package main

import (
	"fmt"
	"github.com/brittonhayes/dnd"
	"os"
	"text/template"
)

// Print out a spell with go text/template
func main() {
	// In your terminal, type:
	// go run main.go acid-arrow
	args := os.Args[1]
	renderSpellTemplate(args)
}

// renderSpellTemplate will search the REST API
// for the provided spell name
func renderSpellTemplate(name string) {
	// Create a new dnd client
	c := dnd.NewClient()

	// Search for a spell
	spell, err := c.FindSpell(name)
	if err != nil {
		panic(err)
	}

	// If the spell wasn't found, let the user know!
	if spell.Name == "" {
		fmt.Println("No spell by that name was found")
		return
	}

	// Render the spell with a go text/template
	// And you're done! 🎉
	tpl := template.Must(template.ParseFiles("spell.tmpl"))
	if err := tpl.ExecuteTemplate(os.Stdout, "spell.tmpl", spell); err != nil {
		panic(err)
	}
}
