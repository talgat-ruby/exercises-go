package problem2

import "github.com/talgat-ruby/exercises-go/pkg/util"

func capitalize(slice []string) []string {
	res := make([]string, len(slice))
	for i := 0; i < len(res); i++ {
		slice[i] = util.CapitalizeString(slice[i])
	}
	return slice
}
