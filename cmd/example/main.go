package main

import (
	"flag"
	"fmt"
	lab2 "github.com/maxnetyaga/software-architecture-lab1"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	postfix := "3 4 + 5 *" // (3 + 4) * 5
	infix, err := lab2.PostfixToInfix(postfix)
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}
	fmt.Println("Інфіксний вираз:", infix)
}