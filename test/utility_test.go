package test

import (
	"utility"
)
import "testing"

func TestStringUtils(t *testing.T) {
	t.Run("string array contains", func(t *testing.T) {
		strArr := []string{"a", "b", "c"}

		if !utility.StringArrContains(strArr, "a") {
			t.Errorf("Expected the true got false")
		}

		if utility.StringArrContains(strArr, "d") {
			t.Errorf("Expected the false got true")
		}

		if !utility.StringArrContains(strArr, "b") {
			t.Errorf("Expected the true got false")
		}

	})
}
