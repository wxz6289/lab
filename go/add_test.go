package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var intResult int = Add(3, 5)
	var floatResult float64 = Add(3.5, 2.521)
	if intResult != 8 {
		t.Errorf("Expected 8, but got %d", intResult)
	}
	if floatResult != 6.021 {
		t.Errorf("Expected 6.021, but got %.3f", floatResult)
	}

	t.Log("intResult:", intResult)
	t.Log("floatResult:", floatResult)
}

func TestSum2(t *testing.T) {
	t.Run("case1", func (t *testing.T) {
		result := sum2(3, 5)
		if result != 8 {
			t.Errorf("Expected 8, but got %d", result)
		}
		t.Log("result:", result)
	})

	t.Run("case2", func (t *testing.T) {
		result := sum2(3.5, 2.521)
		if result != 6.021 {
			t.Errorf("Expected 6.021, but got %.3f", result)
		}
		t.Log("result:", result)
	})
}


func TestAddBatch(t *testing.T) {
	tests := []struct {
		name     string
		a        any
		b        any
		expected any
	}{
		{"int case", 3, 5, 8},
		// {"int float", 2, 2.3, 4.3},
		{"float case", 3.5, 2.521, 6.021},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result any
			switch a := tt.a.(type) {
			case int:
				result = Add(a, tt.b.(int))
			case float64:
				result = Add(a, tt.b.(float64))
			default:
				t.Fatalf("Unsupported type: %T", a)
			}

			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
			t.Log("result:", result)
		})
	}
}

