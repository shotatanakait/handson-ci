package main

import "testing"

// Define test func
func TestEvenOrOdd(t *testing.T) {
    result := EvenOrOdd(10)
    if result != "even" {
        t.Errorf("expected: even, actual: %s", result)
    }
}

