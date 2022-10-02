package goperation

import (
	"fmt"
	"testing"
)

func TestRandomIntRange(t *testing.T) {
	result := randomIntRange(1, 100)
	fmt.Println(result)
}
