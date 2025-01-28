package mongosession

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type SessionOptions []SessionOption

type SessionOption func(*options.SessionOptions)

func (so SessionOptions) Apply(o *options.SessionOptions) {
	for _, opt := range so {
		opt(o)
	}
}

func NewSessionOptions(opts ...SessionOption) SessionOptions {
	return opts
}

func WithCausalConsistency(enabled bool) SessionOption {
	return func(opts *options.SessionOptions) {
		opts.SetCausalConsistency(enabled)
	}
}

func WithDefaultReadConcern(rc *readconcern.ReadConcern) SessionOption {
	return func(opts *options.SessionOptions) {
		opts.SetDefaultReadConcern(rc)
	}
}

func WithDefaultReadPreference(rp *readpref.ReadPref) SessionOption {
	return func(opts *options.SessionOptions) {
		opts.SetDefaultReadPreference(rp)
	}
}

func WithDefaultWriteConcern(wc *writeconcern.WriteConcern) SessionOption {
	return func(opts *options.SessionOptions) {
		opts.SetDefaultWriteConcern(wc)
	}
}

func WithDefaultMaxCommitTime(mct time.Duration) SessionOption {
	return func(opts *options.SessionOptions) {
		opts.SetDefaultMaxCommitTime(&mct)
	}
}

func WithSnapshot(enabled bool) SessionOption {
	return func(opts *options.SessionOptions) {
		opts.SetSnapshot(enabled)
	}
}

type TxOptions []TxOption

type TxOption func(o *options.TransactionOptions)

func (to TxOptions) Apply(o *options.TransactionOptions) {
	for _, opt := range to {
		opt(o)
	}
}

func NewTxOptions(opts ...SessionOption) SessionOptions {
	return opts
}

func WithWriteConcern(wc *writeconcern.WriteConcern) TxOption {
	return func(o *options.TransactionOptions) {
		o.SetWriteConcern(wc)
	}
}

func WithReadConcern(rc *readconcern.ReadConcern) TxOption {
	return func(o *options.TransactionOptions) {
		o.SetReadConcern(rc)
	}
}

func WithReadPreference(rp *readpref.ReadPref) TxOption {
	return func(o *options.TransactionOptions) {
		o.SetReadPreference(rp)
	}
}

func WithMaxCommitTime(d time.Duration) TxOption {
	return func(o *options.TransactionOptions) {
		o.SetMaxCommitTime(&d)
	}
}
