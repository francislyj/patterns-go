package observer

func TestObserver() {

	shirt := newItem("shirt")

	xiaoming := &customer{
		id: "xiaoming",
	}
	zhanghua := &customer{
		id: "zhanghua",
	}

	shirt.register(xiaoming)
	shirt.register(zhanghua)

	shirt.updateAvailability()


}
