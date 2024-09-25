package problem2

import (
	"slices"
	"testing"
)

func TestCapitalize(t *testing.T) {
	table := []struct {
		names []string
		exp   []string
	}{
		{[]string{"A", "B", ""}, []string{"A", "B", ""}},
		{[]string{"mavis", "senaida", "letty"}, []string{"Mavis", "Senaida", "Letty"}},
		{[]string{"samuel", "MABELLE", "letitia", "meridith"}, []string{"Samuel", "Mabelle", "Letitia", "Meridith"}},
		{[]string{"Slyvia", "Kristal", "Sharilyn", "Calista"}, []string{"Slyvia", "Kristal", "Sharilyn", "Calista"}},
		{[]string{"krisTopher", "olIva", "herminiA"}, []string{"Kristopher", "Oliva", "Herminia"}},
		{[]string{"luke", "marsha", "stanford"}, []string{"Luke", "Marsha", "Stanford"}},
		{[]string{"kara"}, []string{"Kara"}},
		{[]string{"mARIANN", "jOI", "gEORGEANN"}, []string{"Mariann", "Joi", "Georgeann"}},
	}

	for _, r := range table {
		capitalized := capitalize(r.names)
		if !slices.Equal(r.exp, capitalized) {
			t.Errorf("capitalize(%v) was incorrect, got: %v, expected: %v.", r.names, capitalized, r.exp)
		}
	}
}
