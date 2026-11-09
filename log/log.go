package log

import (
	"fmt"

	"github.com/MarcosGabrielPF/GoRandomNick/vars"
)

func Help() {
	fmt.Println(
		vars.Green+"\t'-c 12': "+vars.Reset+"Number of characters.",
		vars.Green+"\n\t'-n false'(or true): "+vars.Reset+"If the nickname will have numbers.",
	)
}

func HelpError() {
	fmt.Println(
		vars.Red + "- " + vars.Reset + "Use 'grn -h' for help.",
	)
}

func ValueEmpty() {
	fmt.Println(
		vars.Red + "- " + vars.Reset + "The value is empty",
	)
}

func ValueErrorC() {
	fmt.Println(
		vars.Red + "- " + vars.Reset + "Enter the number of characters after '-c'.",
	)
}

func ValueErrorN() {
	fmt.Println(
		vars.Red + "- " + vars.Reset + "Use '-n' to specify whether the nickname will include numbers.",
	)
}

func ZeroErrorC() {
	fmt.Println(
		vars.Red + "- " + vars.Reset + "The value of '-c' cannot be less than or equal to 0.",
	)
}

func Version() {
	fmt.Println(
		"The program version is: " + vars.Green + vars.Version + vars.Reset,
	)
}

func Log(message string, err ...bool) {
	if len(err) == 0 {
		err = []bool{false}
	}
	if err[0] == true {
		fmt.Println(
			vars.Red + "- " + vars.Reset + message,
		)
	} else {
		fmt.Println(
			vars.Green + "+ " + vars.Reset + message,
		)
	}
}
