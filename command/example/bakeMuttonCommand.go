package command

type bakeMuttonCommand struct {
	r receiver
}

func (this *bakeMuttonCommand)ExcuteCommand() {
	this.r.action()
}
