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

func (s *Session) Session(ctx context.Context, fn func(ctx context.Context) error) error {
	return s.client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		if err := sessionContext.StartTransaction(); err != nil {
			return err
		}
		if err := fn(sessionContext); err != nil {
			return err
		}

		return sessionContext.CommitTransaction(ctx)
	})
}

func (s *Session) SessionWithSessionOptions(ctx context.Context, sopts *options.SessionOptions, fn func(ctx context.Context) error) error {
	return s.client.UseSessionWithOptions(ctx, sopts, func(sessionContext mongo.SessionContext) error {
		if err := sessionContext.StartTransaction(); err != nil {
			return err
		}
		if err := fn(sessionContext); err != nil {
			return err
		}

		return sessionContext.CommitTransaction(ctx)
	})
}

func (s *Session) SessionWithOptions(ctx context.Context, fn func(ctx context.Context) error, opts ...func(sopts *options.SessionOptions)) error {
	sopts := options.Session()
	for _, opt := range opts {
		opt(sopts)
	}

	return s.SessionWithSessionOptions(ctx, sopts, fn)
}
