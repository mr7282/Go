package main

import (
	"fmt"

	"./queue"
)

//Car of struct
/*Описать несколько структур — любой легковой автомобиль и грузовик.
Структуры должны содержать марку авто, год выпуска, объем багажника/кузова,
 запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.*/
type Car struct {
	Brand       string
	Model       string
	Vin         string
	Year        int
	TrankVolume int
	Engine      string
	Windows     string
	FuelTank    int
}

//Truck of struct
type Truck struct {
	general       Car
	carrying      int
	cruisingRange int
}

//VAZ2106 of struct
var VAZ2106 = Car{
	Brand:       "Ваз",
	Model:       "2106",
	FuelTank:    40,
	TrankVolume: 150,
	Vin:         "XTA21063T4R678543",
	Year:        1993,
	Engine:      "Запущен",
	Windows:     "Закрыты",
}

//MAZ of struct
/*Инициализировать несколько экземпляров структур. Применить к ним различные действия. Вывести значения свойств экземпляров в консоль.*/
var MAZ = Truck{
	general: Car{
		Brand:       "Маз",
		Model:       "48053",
		FuelTank:    150,
		TrankVolume: 600,
		Vin:         "XTA48053T4R665341",
		Year:        2003,
		Engine:      "Заглушен",
		Windows:     "Закрыты",
	},
	carrying:      28000,
	cruisingRange: 890,
}

func main() {

	//* Реализовать очередь. Это структура данных, работающая по принципу FIFO (First Input — first output, или «первым зашел — первым вышел»).
	queue.Push("1")
	queue.Push("2")
	queue.Push("3")
	queue.Push("4")
	queue.Push("5")
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	fmt.Println("Марка:", VAZ2106.Brand, "\nМодель:", VAZ2106.Model, "\nVIN:", VAZ2106.Vin, "\nГод:", VAZ2106.Year, "\nОбъем багажника:", VAZ2106.TrankVolume, "л", "\nСостояние двигателя:", VAZ2106.Engine, "\nОкна:", VAZ2106.Windows, "\nОбъем топливного бака:", VAZ2106.FuelTank, "л")
	fmt.Println("\n\nМарка:", MAZ.general.Brand, "\nМодель:", MAZ.general.Model, "\nVIN:", MAZ.general.Vin, "\nГод:", MAZ.general.Year, "\nОбъем багажника:", MAZ.general.TrankVolume, "л", "\nСостояние двигателя:", MAZ.general.Engine, "\nОкна:", MAZ.general.Windows, "\nОбъем топливного бака:", MAZ.general.FuelTank, "л", "\nГрузоподъемность:", MAZ.carrying, "кг", "\nЗапас хода:", MAZ.cruisingRange, "км")
}
