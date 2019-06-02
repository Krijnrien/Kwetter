package event

import "github.com/krijnrien/Kwetter/schema"

type EventStore interface {
	Close()
	PublishKweetCreated(Kweet schema.Kweet) error
	SubscribeKweetCreated() (<-chan KweetCreatedMessage, error)
	OnKweetCreated(f func(KweetCreatedMessage)) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishKweetCreated(Kweet schema.Kweet) error {
	return impl.PublishKweetCreated(Kweet)
}

func SubscribeKweetCreated() (<-chan KweetCreatedMessage, error) {
	return impl.SubscribeKweetCreated()
}

func OnKweetCreated(f func(KweetCreatedMessage)) error {
	return impl.OnKweetCreated(f)
}
