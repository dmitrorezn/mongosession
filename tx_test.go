package mongosession

import (
	"context"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test(t *testing.T) {
	uri := "mongodb://127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/?replicaSet=rs0"
	err := mtest.Setup(mtest.NewSetupOptions().SetURI(uri))
	require.NoError(t, err)
	var (
		ctx     = context.Background()
		mongoT  = mtest.New(t)
		connOpt = options.Client().ApplyURI(uri)
		opts    = mtest.NewOptions().
			CreateClient(true).
			CreateCollection(true).
			ShareClient(true).
			ClientOptions(connOpt)
	)
	mongoT.RunOpts("Session", opts, func(mt *mtest.T) {
		testSession(ctx, mt)
	})
}

func testSession(ctx context.Context, mt *mtest.T) {
	var (
		required = require.New(mt.T)
		client   = mt.Client
		docs     = client.Database("db").Collection("docs")
		users    = client.Database("db").Collection("users")
		session  = NewTransactor(client)
	)
	doc := map[string]any{
		"name":   "doc1",
		"userId": 1,
	}
	user := map[string]any{
		"_id":  1,
		"docs": []string{"doc1"},
	}
	_, err := session.Tx(ctx, func(ctx context.Context) (_ any, err error) {
		if _, err = docs.InsertOne(ctx, doc); err != nil {
			return
		}
		if _, err = users.InsertOne(ctx, user); err != nil {
			return
		}

		return
	})
	required.NoError(err)

	err = docs.FindOne(ctx, bson.M{"name": "doc1"}).Err()
	required.NoError(err)
	err = users.FindOne(ctx, bson.M{"_id": 1}).Err()
	required.NoError(err)

	doc = map[string]any{
		"name":   "doc2",
		"userId": 1,
	}
	id := math.MaxUint32 + 1
	user = map[string]any{
		"_id":  id,
		"docs": []string{"doc2"},
	}
	_, err = session.Tx(ctx, func(ctx context.Context) (_ any, err error) {
		if _, err = docs.InsertOne(ctx, doc); err != nil {
			return
		}
		if _, err = users.InsertOne(ctx, user); err != nil {
			return
		}

		return
	})
	required.Error(err)

	err = docs.FindOne(ctx, bson.M{"name": "doc2"}).Err()
	required.ErrorIs(err, mongo.ErrNoDocuments)
	err = users.FindOne(ctx, bson.M{"_id": id}).Err()
	required.ErrorIs(err, mongo.ErrNoDocuments)
}
