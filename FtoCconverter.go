package main

import ("fmt"
    "math" 
)

func main() {
    var fahr_value float64
    fmt.Print("Please enter Fahrenheit value: ")
    fmt.Scanf("%f", &fahr_value)
    fahr_value = math.Round(fahr_value*10)/10
    s := fmt.Sprintf("%.1f", (fahr_value-32)*5/9)
    fmt.Printf("\n%f Fahrenheit is %s Celsius\n", fahr_value, s)
}
