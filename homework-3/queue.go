package main

var x []string

// Push добавит новый элемент в стек
func Push(str string) {
	x = append(x, str)
}

// Pop вернет последний добавленный в стек элемент
func Pop() string {

	numOfElements := len(x)
	// Когда стек будет пустым, он вернет пустую строку
	if numOfElements == 0 {
		return ""
	}
	popElem := x[0]
	for i := 0; i < (numOfElements - 1); i++ {
		x[i] = x[i+1]
	}
	x = x[:numOfElements-1]
	return popElem
}
