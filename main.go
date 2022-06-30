package main

import (
	"fmt"
	"github.com/fatih/color"
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
	//Help page for the command line
	mainCmd = &cobra.Command{
		Run:   generate,
		Use:   "Password-Generator",
		Short: "Password-Generator is a password generator.",
		Long:  "Password-Generator is a password generator written in Golang.",
	}

	mainCmd.Flags().UintVarP(&length, "length", "l", generator.DefaultConfig.Length, "length of the password.")
	mainCmd.Flags().StringVarP(&characterSet, "character-set", "c", generator.DefaultConfig.CharacterSet, "character set of the password.")
	mainCmd.Flags().BoolVarP(&includeNumbers, "include-numbers", "n", generator.DefaultConfig.IncludeNumbers, "include numbers in the password.")
	mainCmd.Flags().BoolVarP(&includeSymbol, "include-symbols", "s", generator.DefaultConfig.IncludeSymbols, "include symbols in the password.")
	mainCmd.Flags().BoolVarP(&includeUppercaseLetters, "include-uppercase-letters", "u", generator.DefaultConfig.IncludeUppercaseLetters, "include uppercase letters in the password.")
	mainCmd.Flags().BoolVarP(&includeLowercaseLetters, "include-lowercase-letters", "o", generator.DefaultConfig.IncludeLowercaseLetters, "include lowercase letters in the password.")
	mainCmd.Flags().BoolVarP(&excludeSimilarCharacters, "exclude-similar-characters", "e", generator.DefaultConfig.ExcludeSimilarCharacters, "exclude characters that look the same in the characters.")
	mainCmd.Flags().BoolVarP(&excludeAmbiguousCharacters, "exclude-ambiguous-characters", "a", generator.DefaultConfig.ExcludeAmbiguousCharacters, "exclude characters that hard to remember.")
	mainCmd.Flags().UintVarP(&times, "times", "t", 1, "How many passwords you want to generate.")
	color.Magenta("Your Password/s is/are:")
	err := mainCmd.Execute()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
