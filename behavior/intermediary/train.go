package intermediary

type train interface {
	arrive()
	depart()
	permitArrival()
}
