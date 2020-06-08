package main

import (
	"fmt"
	"strings"

	"code.mantid.org/wordcase"
)

//var exampleStrings = []string{
//	"",
//	"a",
//	"A",
//	"a---------b",
//	"dooker",
//	"dookerSpam99_rawr",
//	"IDOne_XMLHttp_ON",
//	"maxID",
//	"maxId",
//	"ENV_VAR",
//	"snake_case",
//	"kebab-case",
//	"IDOne",
//	"99two",
//	"99Two",
//	"*a.B",
//	"interface{}",
//	"something$",
//	"something$$$",
//}

func main() {
	examples := []func(){
		exampleStandAlone,
		exampleTitleCase,
	}

	for _, fn := range examples {
		fn()
		fmt.Print("\n")
	}
}

func exampleStandAlone() {
	standAlone := "One Two three-four."
	methods := map[string]func(string) string{
		"SnakeCase":          wordcase.SnakeCase,
		"KebabCase":          wordcase.KebabCase,
		"DotCase":            wordcase.DotCase,
		"ScreamingSnakeCase": wordcase.ScreamingSnakeCase,
		"CamelCase":          wordcase.CamelCase,
		"PascalCase":         wordcase.PascalCase,
		"Words":              wordcase.Words,
	}
	for k, fn := range methods {
		fmt.Printf("%-20s -> %s\n", k, fn(standAlone))
	}
}

func exampleTitleCase() {

	titleCase := wordcase.
		NewPipeline().
		TokenizeUsing(wordcase.LookAroundCategorizer, wordcase.NotLetterOrDigit, true).
		TokenizeUsing(wordcase.LookAroundCategorizer, wordcase.NotLowerOrDigit, false).
		WithFormatter(strings.ToUpper, wordcase.KeyWords).
		WithAllFormatter(wordcase.UppercaseFirst).
		JoinWith(" ")

	ex := []string{
		"fooBarBaz",
		"one two three",
		"EIGHT html rawr",
	}

	fmt.Println("title case")
	for _, e := range ex {
		fmt.Printf("%-20s -> %-20s\n", e, titleCase(e))
	}

}
