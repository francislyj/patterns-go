package adapter

func TestAdapter() {
	mac := &mac{}

	client := client{}
	client.insertIntoLightningPortToComputer(mac)


	windows := &windows{}
	windowAdapter := &windowAdapter{
		windows: windows,
	}

	client.insertIntoLightningPortToComputer(windowAdapter)


}
