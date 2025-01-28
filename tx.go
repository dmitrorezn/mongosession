package mongosession

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transactor struct {
	client *mongo.Client
	opts   []*options.SessionOptions
}

func NewTransactor(client *mongo.Client, opts ...*options.SessionOptions) *Transactor {
	return &Transactor{
		client: client,
		opts:   opts,
	}
}

func (s *Transactor) Tx(ctx context.Context, fn func(ctx context.Context) (any, error)) (any, error) {
	return s.TxWithOptions(ctx, fn)
}

func (s *Transactor) TxWithOptions(
	ctx context.Context,
	fn func(ctx context.Context) (any, error),
	opts ...*options.TransactionOptions,
) (any, error) {
	session, err := s.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	return session.WithTransaction(ctx, func(ctx mongo.SessionContext) (any, error) {
		return fn(ctx)
	}, opts...)
}
