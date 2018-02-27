package decorator

import (
	"fmt"
)

type decoratorA struct {
	addState string
	c conponent
}

func (this *decoratorA)setConponent(c conponent) {
	this.c = c
}
func (this *decoratorA)operation() {

}
