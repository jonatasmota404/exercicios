// Package weather: package that contains the Forecast Weather.
package weather

//CurrentCondition: variable that contains the current condition of weather.
var CurrentCondition string
//CurrentLocation: variable that contains the current location of forecast weather.
var CurrentLocation string

//Forecast: function that returns the current weather condition and his location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
