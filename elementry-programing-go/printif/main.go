package main

func PrintIf(word string) string {
	if len(word) == 0 {
		return "G\n"
	}

	if len(word) >= 3 {
		return "G\n"
	}
	return "invalid input"
}
