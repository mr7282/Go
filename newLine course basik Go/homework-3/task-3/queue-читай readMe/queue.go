package queue

var q []string

//Push - this function adds some value in array
func Push(str string) {
	q = append(q, str)
}

//Pop - this function extracting first value from array
func Pop() string {
	lenghtArr := len(q)

	if lenghtArr == 0 {
		return ""
	}

	queueElem := q[0]
	for i := 0; i < len(q); i++ {
		q[i] = q[i+1]
	}
	q = q[:lenghtArr-1]

	return queueElem
}
