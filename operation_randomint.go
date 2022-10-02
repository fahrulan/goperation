package goperation

import "math/rand"

func randomIntRange(min, max int) int {
	result := rand.Intn(max-min) + min
	return result
}
