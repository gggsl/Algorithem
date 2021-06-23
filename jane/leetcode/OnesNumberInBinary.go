package main

func hammingWeight(num uint32) int {
	var res uint32
	for num != 0 {
		res += num & 1
		num = num >> 1
	}
	return int(res)
}
