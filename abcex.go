package abcex

import "fmt"
import "math"

var decToAbcex = map[int]string{
	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7",
	8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e",
	15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l",
	22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s",
	29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z" }

var abcexToDec = flipMap()

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func flipMap() map[string]int {
    flipped := make(map[string]int)
    for i, val := range decToAbcex {
        flipped[val] = i
    }

    return flipped
}

func Encode(number int64) string{
    result := ""

    for number > 0 {
        result = fmt.Sprintf("%s%s", decToAbcex[int(number % 36)], result)
        number /= 36
	}

    return result
}

func Decode(str string) int64 {
    var result int64
	result = 0
    str = reverse(str)

    for i, c := range str {
        result = int64(math.Pow(36, float64(i))) * int64(abcexToDec[string(c)]) + result
    }

    return result
}
