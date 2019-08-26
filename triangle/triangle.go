package triangle

//Triangle - ...
type Triangle struct {
	A float64
	B float64
	C float64
	H float64
}

//Area - ...
func (t Triangle) Area() float64 {
	return (t.A * t.H) / 2
}

//SetCatet - ...
func (t *Triangle) SetCatet(lenghtA, lenghtB float64) {
	t.A = lenghtA
	t.B = lenghtB
}

//SetAltitude - ...
func (t *Triangle) SetAltitude(lenghtH float64) {
	t.H = lenghtH

}

//GetCatet - ...
func (t Triangle) GetCatet() (a, b float64) {
	a = t.A
	b = t.B
	return a, b
}
