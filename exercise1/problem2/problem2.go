package problem2

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func capitalize(names []string) []string {
	var res []string
	for _, name := range names {
		res = append(res, cases.Title(language.English, cases.Compact).String(name))
	}

	return res
}
