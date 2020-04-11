package utility

import (
	"os"
	"reflect"
)

func StringArrContains(strArr []string, lookupItem string) bool {
	for _, elem := range strArr {
		if elem == lookupItem {
			return true
		}
	}
	return false

}

func Crash(err error) {
	_, _ = os.Stderr.WriteString("Error type = " + reflect.TypeOf(err).String() + "; Error Message = " + err.Error() + " ;\n")
	os.Exit(-1)
}
