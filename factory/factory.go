package factory

func operationFactory(numberA, numberB float64, op string) operation {
	switch op {
	case "+":
		return &Add{numberA:numberA, numberB:numberB}
	case "-":
		return &Sub{numberA:numberA, numberB:numberB}
		}
		return nil
}
