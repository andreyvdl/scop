package src

func ValidateTokenAmount(len, min, max int, msg string) {
	if len < min || len > max {
		panic(msg)
	}
}
