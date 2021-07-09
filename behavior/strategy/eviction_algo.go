package strategy

type evictionAlgo interface {
	evict(*cache)
}
