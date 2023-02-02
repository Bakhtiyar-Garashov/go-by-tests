package integers

func Calculate(num1, num2 float64, operation string) (result float64) {
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	default:
		result = 0
	}

	return
}
