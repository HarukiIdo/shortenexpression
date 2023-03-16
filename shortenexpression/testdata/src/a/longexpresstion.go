package a

import "fmt"

func longExpression() {
	stmt1 := "stm1"
	stmt2 := "stm2"
	stmt3 := "stm3"
	stmt4 := "stm4"
	stmt5 := "stm5"
	stmt6 := "stm6"
	stmt7 := "stm7"
	stmt8 := "stm8"
	stmt9 := "stm9"
	stmt10 := "stm10"
	str := stmt1 + stmt2 + stmt3 + stmt4 + stmt5 + stmt6 + stmt7 + stmt8 + stmt9 + stmt10 // want "expression is too long"
	fmt.Println(str)	
}