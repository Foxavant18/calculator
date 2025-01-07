package calculator

var logMessage = "[LOG]" //Privado

// Version de la calculadora
var Version = "1.0" // Publico

func internalSum(number int) int { //Privado
	return number - 1
}

// Suma de dos numeros enteros
func Sum(number1, number2 int) int { //Publico
	return number1 + number2
}

// Resta de dos numeros enteros
func Rest(number1, number2 int) int {
	return number1 - number2
}

// Multiplicacion de dos numeros enteros
func Plus(number1, number2 int) int {
	return number1 * number2
}
