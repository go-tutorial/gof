package policy

func getPrice(value float64, count int, discount float64) float64 {
	return value * (float64(count)) * discount
}
