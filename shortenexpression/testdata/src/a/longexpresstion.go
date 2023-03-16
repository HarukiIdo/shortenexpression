package a

import "fmt"

func longExpression() {
	year := 2023
	leap := (year%4 == 0) && (!(year%100 == 0) || (year%400 == 0)) // want `expression is too long \(.*\)`
	fmt.Println(leap)
}
