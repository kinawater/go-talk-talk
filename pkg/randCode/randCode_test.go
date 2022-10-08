package randCode

import (
	"fmt"
	"testing"
)

func TestNotRepeatingRandCode(t *testing.T) {
	// 相同的code
	similarNuum := make(map[string]int)

	for i := 0; i < 10000000; i++ {
		code := NotRepeatingRandCode(string(rune(i)), 8)
		if similarNuum[string(code)] > 0 {
			similarNuum[string(code)] += 1
		} else {
			similarNuum[string(code)] = 1
		}
	}
	count := 0
	for k, v := range similarNuum {
		if v > 1 {
			fmt.Println(k, v)
			count++
		}
	}
	fmt.Println("count is", count)

}
