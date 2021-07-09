package observer

type subject interface {
	register(observer)
	deregister(observer)
	notifyAll()
}
