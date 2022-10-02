package goperation

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	result := RandomString(20)
	fmt.Println(result)
}
