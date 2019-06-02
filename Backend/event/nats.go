package event

import (
	"bytes"
	"encoding/gob"

	"github.com/nats-io/go-nats"
	"github.com/krijnrien/Kwetter/schema"
)

type NatsEventStore struct {
	nc                      *nats.Conn
	KweetCreatedSubscription *nats.Subscription
	KweetCreatedChan         chan KweetCreatedMessage
}

func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

func (e *NatsEventStore) SubscribeKweetCreated() (<-chan KweetCreatedMessage, error) {
	m := KweetCreatedMessage{}
	e.KweetCreatedChan = make(chan KweetCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	e.KweetCreatedSubscription, err = e.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	// Decode message
	go func() {
		for {
			select {
			case msg := <-ch:
				e.readMessage(msg.Data, &m)
				e.KweetCreatedChan <- m
			}
		}
	}()
	return (<-chan KweetCreatedMessage)(e.KweetCreatedChan), nil
}

func (e *NatsEventStore) OnKweetCreated(f func(KweetCreatedMessage)) (err error) {
	m := KweetCreatedMessage{}
	e.KweetCreatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		e.readMessage(msg.Data, &m)
		f(m)
	})
	return
}

func (e *NatsEventStore) Close() {
	if e.nc != nil {
		e.nc.Close()
	}
	if e.KweetCreatedSubscription != nil {
		e.KweetCreatedSubscription.Unsubscribe()
	}
	close(e.KweetCreatedChan)
}

func (e *NatsEventStore) PublishKweetCreated(Kweet schema.Kweet) error {
	m := KweetCreatedMessage{Kweet.ID, Kweet.Body, Kweet.CreatedAt}
	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (mq *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (mq *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
