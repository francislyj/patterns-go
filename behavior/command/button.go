package command

type button struct {
	command command
}

func (b *button) onPress() {
	b.command.execute()
}
