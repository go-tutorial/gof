package factory

type Sub struct {
	numberA float64
	numberB float64
}


func (this *Sub)getResult() float64 {
	return this.numberA - this.numberB
}
