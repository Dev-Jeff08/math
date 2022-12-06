package main

import (
	"log"
)

var x = 1

var y = 1

var z = 1

func cal() {
	result := false

	calResult := x * x * x + y * y * y

	zResult := z * z * z

	if (calResult == zResult) {
		result = true

		log.Println(result)
	} else {
		result = false

		log.Println("x:", x, "y:", y, "z:", z,"result:", result)
	}
}

func addX(value ... int) {
	x = value[0]
}

func addY(value ... int) {
	y = value[0]
}

func addZ(value ... int) {
	z = value[0]
}

func main() {
	length := 1

	for {

		length++

		for i := 1; i < length; i++ {
			addX(i)

			cal()

			for i := 1; i < length; i++ {
				addY(i)

				cal()

				for i:= 1; i < length; i++ {
					addZ(i)

					cal()
				}
			}
		}
	}
}