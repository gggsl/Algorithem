package main

import (
	"math/bits"
	"strconv"
)

func readBinaryWatch(turnedOn int) []string {
	ans := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if bits.OnesCount(uint(i))+bits.OnesCount(uint(j)) == turnedOn {
				ans = append(ans, strconv.Itoa(i)+":"+If(j < 10, "0"+strconv.Itoa(j), strconv.Itoa(j)))
			}
		}
	}
	return ans
}

func If(isTrue bool, a, b string) string {
	if isTrue {
		return a
	}
	return b
}
