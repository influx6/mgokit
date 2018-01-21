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

	mdb "github.com/gokit/mgokit/example/api/usermgo"

	fixtures "github.com/gokit/mgokit/example/api/usermgo/fixtures"
)

var (
	config = mdb.Config{
		Mode:     mgo.Monotonic,
		DB:       os.Getenv("API_MONGO_TEST_DB"),
		Host:     os.Getenv("API_MONGO_TEST_HOST"),
		User:     os.Getenv("API_MONGO_TEST_USER"),
		AuthDB:   os.Getenv("API_MONGO_TEST_AUTHDB"),
		Password: os.Getenv("API_MONGO_TEST_PASSWORD"),
	}

	testCol = "user_test_collection"
)

// TestGetUser validates the retrieval of a User
// record from a mongodb.
func TestGetUser(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	_, err = api.Get(ctx, elem.PublicID)
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

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	records, _, err := api.GetAll(ctx, "asc", "public_id", -1, -1)
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
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
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
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
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

	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	elem2, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	elem2.PublicID = elem.PublicID

	if err := api.Update(ctx, elem2.PublicID, elem2); err != nil {
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
	api := mdb.New(testCol, events, mdb.NewMongoDB(config))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for User into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for User into db.")

	if _, err = api.Get(ctx, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for User into db.")
	}
	tests.Passed("Successfully failed to get deleted record for User into db.")
}
