package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/brittonhayes/dnd"
)

// monsterList returns a list of monster
// names
func monsterList() []string {
	c := dnd.NewClient()
	monsters, err := c.Monsters.List()
	if err != nil {
		log.Fatal(err)
	}

	return monsters.ResultsNames()
}

// spellList returns a list of spell
// names
func spellList() []string {
	c := dnd.NewClient()
	spells, err := c.Spells.List()
	if err != nil {
		log.Fatal(err)
	}

	return spells.ResultsNames()
}

// q is the list of questions
// to ask the user
var q = []*survey.Question{
	{
		Name: "monster",
		Prompt: &survey.Select{
			Message: "Select a monster",
			Options: monsterList(),
			Help:    "Choose a monster from the list",
		},
		Validate: survey.Required,
	},
	{
		Name: "spell",
		Prompt: &survey.Select{
			Message: "Select a spell",
			Options: spellList(),
			Help:    "Choose a spell from the list",
		},
		Validate: survey.Required,
	},
}

func main() {
	answers := struct {
		Monster string
		Spell   string
	}{}

	// Ask the questions
	if err := survey.Ask(q, &answers); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the selected result to the screen
	fmt.Printf("Monster chosen:\t%s\nSpell chosen:\t%s\n", answers.Monster, answers.Spell)
}
