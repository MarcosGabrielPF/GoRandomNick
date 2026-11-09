package generator

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/MarcosGabrielPF/GoRandomNick/log"
	"github.com/MarcosGabrielPF/GoRandomNick/vars"
)

func subGenerate() string {
	var nickname string
	var lastCharacter string
	charactersMissingForNumber := rand.IntN((-3)-(-4)+1) + (-4) // (maiorNumero - menorNumero + 1) + menorNumero

	for i := 0; i < vars.Characters; i++ {
		if vars.Numbers == true && i-vars.Characters >= charactersMissingForNumber {
			index := rand.IntN(len(vars.PossibleCharacters["Numbers"]))
			lastCharacter = string(vars.PossibleCharacters["Numbers"][index])
			nickname += lastCharacter
		} else if i == 0 {
			index := rand.IntN(len(vars.PossibleCharacters["Letters"]))
			lastCharacter = string(vars.PossibleCharacters["Letters"][index])
			nickname += lastCharacter
			nickname = strings.ToUpper(nickname)
		} else {
			switch {
			case strings.Contains(vars.PossibleCharacters["Consonants"], lastCharacter):
				index := rand.IntN(len(vars.PossibleCharacters["Vowels"]))
				lastCharacter = string(vars.PossibleCharacters["Vowels"][index])
				nickname += lastCharacter

			case strings.Contains(vars.PossibleCharacters["Vowels"], lastCharacter):
				index := rand.IntN(len(vars.PossibleCharacters["Letters"]))
				lastCharacter = string(vars.PossibleCharacters["Letters"][index])
				nickname += lastCharacter

			default:
				index := rand.IntN(len(vars.PossibleCharacters["Letters"]))
				lastCharacter = string(vars.PossibleCharacters["Letters"][index])
				nickname += lastCharacter
			}
		}
	}

	return nickname
}

func Generate() {
	log.Log("Creating a nickname.")
	nickname := subGenerate()
	if len(nickname) == vars.Characters {
		log.Log(fmt.Sprintf("Nick: %s", nickname))
	} else {
		i := 0
		for len(nickname) != vars.Characters {
			i++
			nickname = subGenerate()
			fmt.Printf("\rAttempt: %d", i)
			if i == 1000000 {
				fmt.Println()
				log.Log("Too many attempts; please try again.", true)
				os.Exit(1)
			}
		}
		log.Log(fmt.Sprintf("Nick: %s", nickname))
	}
}
