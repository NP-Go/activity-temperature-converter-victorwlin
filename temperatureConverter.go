package main

import "fmt"

func convertKelvin(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := (temperatureInput / (5.0 / 9.0)) - 459.67
	resultCelsius := (5.0 / 9.0) * (resultFahrenheit - 32.0)
	//Do not remove this line
	return resultFahrenheit, resultCelsius
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := (temperatureInput / (5.0 / 9.0)) + 32.0
	resultKelvin := (5.0 / 9.0) * (resultFahrenheit + 459.67)
	//Do not remove this line
	return resultFahrenheit, resultKelvin
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultKelvin := (5.0 / 9.0) * (temperatureInput + 459.67)
	resultCelsius := (5.0 / 9.0) * (temperatureInput - 32.0)
	//Do not remove this line
	return resultKelvin, resultCelsius
}

func main() {
	var temperatureChoice int
	var temperatureInput float64
	fmt.Printf("Enter your temperature format (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit: ")
	fmt.Scanln(&temperatureChoice)
	fmt.Printf("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)

	if temperatureChoice == 1 {
		//Insert Code here
		resultFahrenheit, resultCelsius := convertKelvin(temperatureInput)
		//DO not remove this
		fmt.Printf("Fahrenheit: %.2f  Celsius: %.2f", resultFahrenheit, resultCelsius)
	} else if temperatureChoice == 2 {
		//Insert Code here
		resultFahrenheit, resultKelvin := convertCelsius(temperatureInput)
		//DO not remove this
		fmt.Printf("Fahrenheit: %.2f  Kelvin: %.2f", resultFahrenheit, resultKelvin)
	} else if temperatureChoice == 3 {
		//Insert Code here
		resultKelvin, resultCelsius := convertFahrenheit(temperatureInput)
		//DO not remove this
		fmt.Printf("Kelvin: %.2f  Celsius: %.2f", resultKelvin, resultCelsius)
	} else {
		fmt.Println("No Conversion")
	}

}
