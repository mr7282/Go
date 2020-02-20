package figures

import (
	"fmt"
)

//Triangle - ...
type Triangle struct {
	A float64
	B float64
	H float64
}

//Area - ...
func (t Triangle) Area() float64 {
	return (t.A * t.H) / 2
}

//SetCatet - ...
func (t *Triangle) SetCatet() {
	fmt.Println("Введите значение Стороны А треугольника")
	fmt.Scanln(&t.A)
	fmt.Println("Введите значение Стороны В треугольника")
	fmt.Scanln(&t.B)
}

//SetAltitude - ...
func (t *Triangle) SetAltitude() {
	fmt.Println("Введите значение Высоты с основанием стороной А треугольника")
	fmt.Scanln(&t.H)
}

//GetCatet - ...
func (t Triangle) GetCatet() {
	fmt.Println("Сторона А =", t.A)
	fmt.Println("Сторона В =", t.B)

}

//Trapeze - ...
type Trapeze struct {
	A float64
	B float64
	H float64
}

//Area - ...
func (tz Trapeze) Area() float64 {
	return ((tz.A + tz.B) / 2) * tz.H
}

//SetBasis - ...
func (tz *Trapeze) SetBasis() {
	fmt.Println("Введите значение Основания А трапеции")
	fmt.Scanln(&tz.A)
	fmt.Println("Введите значение Основания В трапеции")
	fmt.Scanln(&tz.B)
}

//SetAltitude - ...
func (tz *Trapeze) SetAltitude() {
	fmt.Println("Введите значение Высоты трапеции")
	fmt.Scanln(&tz.H)
}

//GetBasis - ...
func (tz Trapeze) GetBasis() {
	fmt.Println("Основание А =", tz.A)
	fmt.Println("Основание В =", tz.B)

}
