package utils

// credit here: https://stackoverflow.com/questions/29002724/implement-ruby-style-cartesian-product-in-go

func CartesianProduct(items [][]string) [][]string {
	out := [][]string{}

	lens := func(i int) int { return len(items[i]) }

	for ix := make([]int, len(items)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []string
		for j, k := range ix {
			r = append(r, items[j][k])
		}

		out = append(out, r)
	}

	return out
}

// func arrange(set map[string][]string) {
// NextIndex sets ix to the lexicographically next value,
// such that for each i>0, 0 <= ix[i] < lens(i).
func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}
