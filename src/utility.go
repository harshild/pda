package src

func StringArrContains(strArr []string, lookupItem string) bool {
	for _, elem := range strArr {
		if elem == lookupItem {
			return true
		}
	}
	return false

}
