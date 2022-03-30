package helpers

func Contains(slice []string, element string) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}

	return false
}
