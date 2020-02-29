package test

import . "../src"
import "testing"

func TestStringUtils(t *testing.T) {
	t.Run("string array contains", func(t *testing.T) {
		strArr := []string{"a", "b", "c"}

		if !StringArrContains(strArr, "a") {
			t.Errorf("Expected the true got false")
		}

		if StringArrContains(strArr, "d") {
			t.Errorf("Expected the false got true")
		}

		if !StringArrContains(strArr, "b") {
			t.Errorf("Expected the true got false")
		}

	})
}
