package main

import "fmt"

func main(){

	fmt.Printf("Hello World\n\n")
var firstName string = "Pankaj"
var lastName string = "Kumar"
fmt.Printf(firstName+" " +lastName+"\n\n")
firstName, lastName = lastName, firstName
fmt.Printf(firstName+" "+ lastName+"\n")

  var x [5]float64
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83

var total float64 = 0
for i, value := range x {
  total += value
  
  fmt.Printf("\n")
  fmt.Printf("For i  = %d total is %f", i,total)
  fmt.Printf("\n")
  //fmt.Printf("value=%f",value)
}

fmt.Println("Average :", total / float64(len(x)))
fmt.Printf("\n")
fmt.Printf("\n\n\n\n\n\n\n\n\n\n")

runGoRoutines()
goChannels()
}
