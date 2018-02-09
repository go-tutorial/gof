package factory

type Add struct {
	numberA float64
	numberB float64
}

func (this *Add)getResult() float64 {
	return this.numberA + this.numberB
}
