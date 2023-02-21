package utility

import (
	"math"
	"strings"
)

func ComputeHash(url string) int64 {
	var hash int64 = 0
	tokens := strings.Split(url, ".")
	pow := 0
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		for _, c := range token {
			p := math.Pow(2, float64(pow))
			hash += int64(int64(c) * int64(p))
			pow++
		}
	}
	return hash
}
