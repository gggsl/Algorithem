package leetcode

import "math/bits"

func maxLength(arr []string) int {
	bitArr := []int{}
outer:
	for _, s := range arr {
		mask := 0
		for _, ch := range s {
			ch -= 'a'
			if mask&(1<<ch) > 0 {
				continue outer
			}
			mask |= 1 << ch
		}
		bitArr = append(bitArr, mask)
	}
	return dfs(0, 0, bitArr)
}

func dfs(position, mask int, bitArr []int) int {
	if position == len(bitArr) {
		return bits.OnesCount(uint(mask))
	}
	res := 0
	if bitArr[position]&mask == 0 {
		res = dfs(position+1, mask|bitArr[position], bitArr)
	}
	res = max(dfs(position+1, mask, bitArr), res)
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}