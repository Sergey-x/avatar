package utils

import (
	"strconv"
	"strings"
)

func StrToArrayInt(str string) []int {
	chunks := strings.Split(str, ",")

	var res []int
	for _, c := range chunks {
		i, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		res = append(res, i)
	}

	return res
}
