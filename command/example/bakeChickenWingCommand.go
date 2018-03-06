package command

type bakeChildrenWingCommand struct {
	r receiver
	}

func (this *bakeChildrenWingCommand)()  {
	this.r.action()
}
