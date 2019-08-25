package queue

var x []string

// Push добавит новый элемент в очередь
func Push(str string) {
	x = append(x, str)
}

// Pop вернет последний добавленный в очередь элемент
func Pop() string {

	numOfElements := len(x)
	// Когда очередь будет пустая, она вернет пустую строку
	if x[0] == "null" {
		return ""
	}

	popElem := x[0]
	for i := 0; i < (numOfElements - 1); i++ {
		x[i] = x[i+1]
		x[i+1] = "null"
	}
	return popElem
}
