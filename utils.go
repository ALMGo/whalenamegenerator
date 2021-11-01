package whalenamegenerator

import (
	"math/rand"
)

func RandomFromArray(arr *[]string) string {
	randomIndex := rand.Intn(len(*arr))
	return (*arr)[randomIndex]
}

func RandomColor() string {
	return RandomFromArray(&Colors)
}

func RandomAdjective() string {
	return RandomFromArray(&Adjectives)
}

func RandomWhale() string {
	return RandomFromArray(&Whales)
}
