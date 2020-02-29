package main

import "testing"

func TestStringUtils(t *testing.T) {
	t.Run("string array contains", func(t *testing.T) {
		strArr := []string{"a", "b", "c"}

		if !stringArrContains(strArr, "a") {
			t.Errorf("Expected the true got false")
		}

		if stringArrContains(strArr, "d") {
			t.Errorf("Expected the false got true")
		}

		if !stringArrContains(strArr, "b") {
			t.Errorf("Expected the true got false")
		}

	})
}
