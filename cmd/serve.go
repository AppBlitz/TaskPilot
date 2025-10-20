package cmd

func VerificationNumber(number string) bool {
	for _, value := range number {
		if value < 47 || value > 57 {
			return false
		}
	}
	return true
}
