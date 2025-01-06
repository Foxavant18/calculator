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
