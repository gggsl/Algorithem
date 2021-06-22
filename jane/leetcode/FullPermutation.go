package main

var ans map[string]bool

func permutation(s string) []string {
	ans = map[string]bool{}
	indexSet := map[int]bool{}
	recursion(indexSet, s, "")
	res := []string{}
	for s := range ans {
		res = append(res, s)
	}
	return res
}

func recursion(indexSet map[int]bool, s, con string) {
	if len(s) == len(indexSet) {
		ans[con] = true
		return
	}
	for index, char := range s {
		exist := indexSet[index]
		if !exist {
			indexSet[index] = true
			recursion(indexSet, s, con+string(char))
			delete(indexSet, index)
		}
	}
}
