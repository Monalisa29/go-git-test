package main

import "testing"

// TestGreet tests the Greet function
func TestGreet(t *testing.T) {
    result := Greet("Git")
    expected := "Hello, Git!"

    // If the result is not equal to the expected value, the test fails
    if result != expected {
        t.Errorf("Greet() = %v; want %v", result, expected)

    }
}
