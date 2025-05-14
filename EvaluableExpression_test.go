package govaluate

import "testing"

func TestA(t *testing.T) {
	t.Parallel()

	_, err := NewEvaluableExpression("foo == 5")

	if err != nil {
		t.Fatalf("Error creating expression: %v", err)
	}
}

func TestB(t *testing.T) {
	t.Parallel()

	_, err := NewEvaluableExpression("foo == 1")

	if err != nil {
		t.Fatalf("Error creating expression: %v", err)
	}
}
