package calculation

// Add performs addition
func Add(num1, num2 float64) float64 {
	return num1 + num2
}

// Sub performs subtraction
func Sub(num1, num2 float64) float64 {
	return num1 - num2
}

// Mul performs multiplication
func Mul(num1, num2 float64) float64 {
	return num1 * num2
}

// Div performs division
// FIXME: return an error if the second number is 0
func Div(num1, num2 float64) float64 {
	return num1 / num2
}
