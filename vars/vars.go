package vars

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

var PossibleCharacters = map[string]string{
	"Letters":    "abcdefghijklmnopqrstuvwxyz",
	"Consonants": "bcdfghjklmnpqrstvwxyz",
	"Vowels":     "aeiou",
	"Numbers":    "0123456789",
}

var Version string = "1.0"

var (
	Characters int
	Numbers    bool
)
