package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmg  float64
}

// value receiver method
func (c car) kmh() float64 { // makes a copy of the struc
	c.top_speed_kmg = 5000 //invalid(function scoped) because not a poointer receiver method
	fmt.Println(c.top_speed_kmg, "teststaa")
	return float64(c.gas_pedal) * (c.top_speed_kmg / usixteenbitmax)
}

// pointer receiver method
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmg = newspeed
}

// value/pointer receiver with function
func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmg = speed
	return c
}

func main() {
	a_car := car{gas_pedal: 1, brake_pedal: 2, steering_wheel: 3, top_speed_kmg: 4}
	a_car.kmh()
	a_car.new_top_speed(21)
	fmt.Println(a_car.top_speed_kmg, "teststaa2")
	fmt.Println((a_car.gas_pedal))
	a_car = newer_top_speed(a_car, 10)
}
