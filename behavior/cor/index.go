package cor

func TestCor() {
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)


	patient := &patient{name: "hello"}

	reception.execute(patient)

}
