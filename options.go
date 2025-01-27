package mongosession

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func WithCausalConsistency(enabled bool) func(opts *options.SessionOptions) {
	return func(opts *options.SessionOptions) {
		opts.SetCausalConsistency(enabled)
	}
}

func WithDefaultReadConcern(rc *readconcern.ReadConcern) func(opts *options.SessionOptions) {
	return func(opts *options.SessionOptions) {
		opts.SetDefaultReadConcern(rc)
	}
}

func WithDefaultReadPreference(rp *readpref.ReadPref) func(opts *options.SessionOptions) {
	return func(opts *options.SessionOptions) {
		opts.SetDefaultReadPreference(rp)
	}
}

func WithDefaultWriteConcern(wc *writeconcern.WriteConcern) func(opts *options.SessionOptions) {
	return func(opts *options.SessionOptions) {
		opts.SetDefaultWriteConcern(wc)
	}
}

func WithDefaultMaxCommitTime(mct time.Duration) func(opts *options.SessionOptions) {
	return func(opts *options.SessionOptions) {
		opts.SetDefaultMaxCommitTime(&mct)
	}
}

func WithSnapshot(enabled bool) func(opts *options.SessionOptions) {
	return func(opts *options.SessionOptions) {
		opts.SetSnapshot(enabled)
	}
}
