package usermgo_test

import (
	"os"

	"time"

	"context"

	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	mdb "github.com/gokit/mgokit/example/methods/usermgo"

	fixtures "github.com/gokit/mgokit/example/methods/usermgo/fixtures"
)

var (
	config = mdb.Config{
		Mode:     mgo.Monotonic,
		DB:       os.Getenv("METHODS_MONGO_TEST_DB"),
		Host:     os.Getenv("METHODS_MONGO_TEST_HOST"),
		User:     os.Getenv("METHODS_MONGO_TEST_USER"),
		AuthDB:   os.Getenv("METHODS_MONGO_TEST_AUTHDB"),
		Password: os.Getenv("METHODS_MONGO_TEST_PASSWORD"),
	}

	db      = mdb.NewMongoDB(config)
	testCol = "user_test_collection"
)

// TestGetUser validates the retrieval of a User
// record from a mongodb.
func TestGetUser(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer mdb.Delete(ctx, db, testCol, events, elem.PublicID)

	if err := mdb.Create(ctx, db, testCol, events, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	_, err = mdb.Get(ctx, db, testCol, events, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for User from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for User from db.")
}

// TestGetAllUser validates the retrieval of all User
// record from a mongodb.
func TestGetAllUser(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer mdb.Delete(ctx, db, testCol, events, elem.PublicID)

	if err := mdb.Create(ctx, db, testCol, events, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	records, _, err := mdb.GetAll(ctx, db, testCol, events, "asc", "public_id", -1, -1)
	if err != nil {
		tests.Failed("Successfully retrieved all records for User from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for User from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for User from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for User from db.")
}

// TestGetAllUserOrderBy validates the retrieval of all User
// record from a mongodb.
func TestGetAllUserByOrder(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer mdb.Delete(ctx, db, testCol, events, elem.PublicID)

	if err := mdb.Create(ctx, db, testCol, events, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	records, err := mdb.GetAllByOrder(ctx, db, testCol, events, "asc", "public_id")
	if err != nil {
		tests.Failed("Successfully retrieved all records for User from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for User from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for User from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for User from db.")
}

// TestUserCreate validates the creation of a User
// record with a mongodb.
func TestUserCreate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer mdb.Delete(ctx, db, testCol, events, elem.PublicID)

	if err := mdb.Create(ctx, db, testCol, events, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")
}

// TestUserUpdate validates the update of a User
// record with a mongodb.
func TestUserUpdate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer mdb.Delete(ctx, db, testCol, events, elem.PublicID)

	if err := mdb.Create(ctx, db, testCol, events, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	elem2, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	elem2.PublicID = elem.PublicID

	if err := mdb.Update(ctx, db, testCol, events, elem2.PublicID, elem2); err != nil {
		tests.Failed("Successfully updated record for User into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for User into db.")
}

// TestUserDelete validates the removal of a User
// record from a mongodb.
func TestUserDelete(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	if err := mdb.Create(ctx, db, testCol, events, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	if err := mdb.Delete(ctx, db, testCol, events, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for User into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for User into db.")

	if _, err = mdb.Get(ctx, db, testCol, events, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for User into db.")
	}
	tests.Passed("Successfully failed to get deleted record for User into db.")
}
