package main

import (
	"os"
	"strconv"

	"github.com/MarcosGabrielPF/GoRandomNick/generator"
	"github.com/MarcosGabrielPF/GoRandomNick/log"
	"github.com/MarcosGabrielPF/GoRandomNick/vars"
)

func main() {
	if len(os.Args) < 2 {
		log.HelpError()
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h", "--help":
			log.Help()
			os.Exit(0)

		case "-v", "--version":
			log.Version()
			os.Exit(0)

		case "-c":
			if i+1 >= len(os.Args) {
				log.ValueEmpty()
				os.Exit(1)
			}

			n, err := strconv.Atoi(os.Args[i+1])

			if err != nil {
				log.ValueErrorC()
				os.Exit(1)
			}

			vars.Characters = n
			i++

		case "-n":
			if i+1 >= len(os.Args) {
				log.ValueEmpty()
				os.Exit(1)
			}

			switch os.Args[i+1] {
			case "true", "t":
				vars.Numbers = true

			case "false", "f":
				vars.Numbers = false

			default:
				log.ValueErrorN()
				os.Exit(1)
			}

			i++

		default:
			log.HelpError()
			os.Exit(1)
		}
	}

	if vars.Characters <= 0 {
		log.ZeroErrorC()
		os.Exit(1)
	}

	generator.Generate()
}
