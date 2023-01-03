package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/a-khutiev/Widberries_internship/internal/cache"
	"github.com/a-khutiev/Widberries_internship/internal/model"
	"github.com/nats-io/stan.go"
)

const (
	natsURL   = stan.DefaultNatsURL
	clusterID = "test-cluster"
	clientID  = "orders-handler"
	subject   = "test-subject"
	qgroup    = "test-qgroup"
	durable   = "test-durable"
)

type MsgHandler struct {
	sc        stan.Conn
	cacheImpl *cache.Cache

	isStarted bool
	sub       stan.Subscription
}

func New(cacheImpl *cache.Cache) (*MsgHandler, error) {
	sc, err := stan.Connect(clusterID, clientID,
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", natsURL, clusterID, clientID)

	return &MsgHandler{
		sc:        sc,
		cacheImpl: cacheImpl,
	}, nil
}

func (handler *MsgHandler) Start() error {
	if !handler.isStarted {
		startOpt := stan.StartWithLastReceived()
		sub, err := handler.sc.QueueSubscribe(subject, qgroup, handler.HandleOrderMsg, startOpt, stan.DurableName(durable))
		if err != nil {
			return err
		}

		handler.sub = sub
	}

	return nil
}

func (handler *MsgHandler) Close() {
	if handler.isStarted {
		err := handler.sub.Unsubscribe()
		if err != nil {
			log.Printf("failed to unsubscribe from nats: %v", err)
		}
	}

	err := handler.sc.Close()
	if err != nil {
		log.Printf("failed to close nats connection: %v", err)
	}
}

func (handler *MsgHandler) HandleOrderMsg(msg *stan.Msg) {
	order := &model.Order{}
	err := json.Unmarshal(msg.Data, order)
	if err != nil {
		log.Println(err)
	}

	err = handler.cacheImpl.CreateOrder(context.Background(), order)
	if err != nil {
		log.Println(err)
	}

	log.Printf("create order with id %s\n", order.OrderUid)

	// todo: write cache
}
