package utils

// Perm calls f with each permutation of a.
func Permutation(a []string, f func([]string)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// func main() {
// 	Permutation([]string{"one", "two", "three", "four", "five"}, func(a []string) {
// 		fmt.Println(a)
// 	})
// }
