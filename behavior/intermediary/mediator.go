package intermediary

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}
