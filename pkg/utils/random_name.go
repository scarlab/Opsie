package utils

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var adjectives = []string{
	"red", "blue", "green", "silent", "frozen", "golden", "wild",
	"brave", "lucky", "mystic", "shadow", "rapid", "gentle", "bright",
}

var nouns = []string{
	"dragon", "falcon", "river", "wolf", "star", "lion", "storm",
	"hawk", "tiger", "phoenix", "comet", "mountain", "whale", "forest",
}

// GenerateTeamName returns a random two-word org name like “Blue Falcon”.
func GenerateTeamName() string {
	adj := adjectives[rand.IntN(len(adjectives))]
	noun := nouns[rand.IntN(len(nouns))]

	name := fmt.Sprintf("%s %s", adj, noun)
	title := cases.Title(language.English)
	return title.String(strings.ToLower(name))
}
