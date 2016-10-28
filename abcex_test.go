package abcex_test

import "fmt"
import "abcex"
import "testing"

func TestEncode(t *testing.T) {
    fmt.Println(abcex.Encode(12345))
    // Output: 9ix
}

func TestDecode(t *testing.T) {
    fmt.Println(abcex.Decode("9ix"))
    // Output: 12345
}
