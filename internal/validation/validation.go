package validation

func ListUserOutputValidation(output string) bool {
	switch output {
	case
		"active",
		"activeroles",
		"name",
		"uuid":
		return true
	}
	return false
}

func UserOutputValidation(output string) bool {
	switch output {
	case
		"active",
		"activeroles",
		"name",
		"orguuid",
		"username",
		"uuid":
		return true
	}
	return false
}
