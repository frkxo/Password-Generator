package main

import (
	"fmt"
	"github.com/frkxo/Password-Generator/generator"
	"github.com/spf13/cobra"
	"os"
)

var (
	mainCmd                    *cobra.Command
	length                     uint
	characterSet               string
	includeSymbol              bool
	includeNumbers             bool
	includeUppercaseLetters    bool
	includeLowercaseLetters    bool
	excludeSimilarCharacters   bool
	excludeAmbiguousCharacters bool
	times                      uint
)

func generate(_ *cobra.Command, args []string) {
	config := generator.Config{
		Length:                     length,
		CharacterSet:               characterSet,
		IncludeSymbols:             includeSymbol,
		IncludeNumbers:             includeNumbers,
		IncludeUppercaseLetters:    includeUppercaseLetters,
		IncludeLowercaseLetters:    includeLowercaseLetters,
		ExcludeSimilarCharacters:   excludeSimilarCharacters,
		ExcludeAmbiguousCharacters: excludeAmbiguousCharacters,
	}
	g, err := generator.NewGenerator(&config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	passwd, err := g.MultipleGenerator(times)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, p := range passwd {
		fmt.Println(p)
	}
}

func main() {
	mainCmd = &cobra.Command{
		Run:   generate,
		Use:   "generate-password",
		Short: "generate-password is a password generator.",
		Long:  "generate-password is a password generator written in Golang.",
	}

	mainCmd.PersistentFlags().UintVarP(&length, "length", "l", generator.DefaultConfig.Length, "length of the password.")
	mainCmd.PersistentFlags().StringVarP(&characterSet, "character-set", "c", generator.DefaultConfig.CharacterSet, "character set of the password.")
	mainCmd.PersistentFlags().BoolVarP(&includeNumbers, "include-numbers", "n", generator.DefaultConfig.IncludeNumbers, "include numbers in the password.")
	mainCmd.PersistentFlags().BoolVarP(&includeSymbol, "include-symbols", "s", generator.DefaultConfig.IncludeSymbols, "include symbols in the password.")
	mainCmd.PersistentFlags().BoolVarP(&includeUppercaseLetters, "include-uppercase-letters", "u", generator.DefaultConfig.IncludeUppercaseLetters, "include uppercase letters in the password.")
	mainCmd.PersistentFlags().BoolVarP(&includeLowercaseLetters, "include-lowercase-letters", "f", generator.DefaultConfig.IncludeLowercaseLetters, "include lowercase letters in the password.")
	mainCmd.PersistentFlags().BoolVarP(&excludeSimilarCharacters, "exclude-similar-characters", "e", generator.DefaultConfig.ExcludeSimilarCharacters, "exclude characters that look the same in the characters.")
	mainCmd.PersistentFlags().BoolVarP(&excludeAmbiguousCharacters, "exclude-ambiguous-characters", "a", generator.DefaultConfig.ExcludeAmbiguousCharacters, "exclude characters that hard to remember.")
	mainCmd.PersistentFlags().UintVarP(&times, "times", "t", 1, "How many passwords you want to generate.")

	err := mainCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
