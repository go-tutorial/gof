package decorator



type decoratorB struct {
	c conponent
}

func (this *decoratorB)setConponent(c conponent) {
	this.c = c
}
func (this *decoratorB)operation() {

}
func (this *decoratorB)AddBehavior() {

}
