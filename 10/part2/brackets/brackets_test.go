package brackets

import "testing"

func TestCalculateScore(t *testing.T) {
	stack := []string{}
	stack = append(stack, "{")
	score := CalculateScore(stack)
	expected := 10

	if score != expected {
		t.Errorf("expected %d but go %d", expected, score)
	}
}
