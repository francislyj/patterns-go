package command

func TestCommand() {
	tv := &tv{}

	onCommand := onCommand{
		device: tv,
	}

	offCommand := offCommand{
		device: tv,
	}

	onButton := button{
		command: &onCommand,
	}

	offButton := button{
		command: &offCommand,
	}

	onButton.onPress()
	offButton.onPress()
}
