package problem5

import "sort"

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func products(vals map[string]int, threshold int) []string {

	result := make(PairList, len(vals))
	i := 0
	for s, v := range vals {
		if v >= threshold {
			result[i] = Pair{s, v}
			i++
		}
	}
	sort.Sort(result)

	return result
}
