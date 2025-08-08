package ports

type QueueConsumer interface {
	Consume() error
}
