package mongosession

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Session struct {
	client *mongo.Client
}

func New(client *mongo.Client) *Session {
	return &Session{
		client: client,
	}
}

func (r *Session) Session(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		if err := sessionContext.StartTransaction(); err != nil {
			return err
		}
		if err := fn(sessionContext); err != nil {
			return err
		}

		return sessionContext.CommitTransaction(ctx)
	})
}

func (r *Session) SessionWithSessionOptions(ctx context.Context, fn func(ctx context.Context) error, sopts *options.SessionOptions) error {
	return r.client.UseSessionWithOptions(ctx, sopts, func(sessionContext mongo.SessionContext) error {
		if err := sessionContext.StartTransaction(); err != nil {
			return err
		}
		if err := fn(sessionContext); err != nil {
			return err
		}

		return sessionContext.CommitTransaction(ctx)
	})
}

func (r *Session) SessionWithOptions(ctx context.Context, fn func(ctx context.Context) error, opts ...func(sopts *options.SessionOptions)) error {
	sopts := options.Session()
	for _, opt := range opts {
		opt(sopts)
	}

	return r.SessionWithSessionOptions(ctx, fn, sopts)
}
