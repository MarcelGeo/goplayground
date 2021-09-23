package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func PrintQuote() {
	fmt.Println(quote.Go())
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

	return formats[rand.Intn(len(formats))]
}