// Package mongoapi provides a auto-generated package which contains a mongo base pkg for db operations.
//
//
package mdb

import (
	"errors"

	"context"

	"runtime"

	"time"

	"sync"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/influx6/faux/metrics"
)

// errors ...
var (
	ErrNotFound       = errors.New("record not found")
	ErrExpiredContext = errors.New("context has expired")
)

//**********************************************************
// MongoDB Config and Setup
//**********************************************************

// Config embodies the data used to connect to user's mongo connection.
type Config struct {
	DB       string `toml:"db" json:"db"`
	AuthDB   string `toml:"authdb" json:"authdb"`
	User     string `toml:"user" json:"user"`
	Password string `toml:"password" json:"password"`
	Host     string `toml:"host" json:"host"`
}

// Empty returns true/false if all Config values are at default/empty/non-set
// state.
func (mgc Config) Empty() bool {
	return mgc.AuthDB == "" &&
		mgc.DB == "" &&
		mgc.User == "" &&
		mgc.Password == "" &&
		mgc.Host == ""
}

// Validate returns an error if the config is invalid.
func (mgc Config) Validate() error {
	if mgc.User == "" {
		return errors.New("Config.User is required")
	}
	if mgc.Password == "" {
		return errors.New("Config.Password is required")
	}
	if mgc.AuthDB == "" {
		return errors.New("Config.AuthDB is required")
	}
	if mgc.Host == "" {
		return errors.New("Config.Host is required")
	}
	if mgc.DB == "" {
		return errors.New("Config.DB is required")
	}
	return nil
}

// MongoDB defines a interface which exposes a method for retrieving a
// mongo.Database and mongo.Session.
type MongoDB interface {
	New(isread bool) (*mgo.Database, *mgo.Session, error)
}

// NewMongoDB returns a new instance of a MongoDB.
func NewMongoDB(conf Config) MongoDB {
	mg := &mongoDB{
		Config: conf,
	}

	// Add finalizer to ensure closure of master session.
	runtime.SetFinalizer(mg, func() {
		mg.ml.Lock()
		defer mg.ml.Unlock()
		if mg.master != nil {
			mg.master.Close()
			mg.master = nil
		}
	})

	return mg
}

// mongoDB defines a mongo connection manager that builds
// allows usage of a giving configuration to generate new mongo
// sessions and database instances.
type mongoDB struct {
	Config
	ml     sync.Mutex
	master *mgo.Session
}

// New returns a new session and database from the giving configuration.
//
// Argument:
//  isread: bool
//
// 1. If `isread` is false, then the mgo.Session is cloned so that we re-use the existing
// sessiby not closing, so others get use ofn connection, in such case, it lets you optimize writes, so try not
// the session instance connection for other writes.
//
// 2. If `isread` is true, then session is copied which creates a new unique session which you
// should close after use, this lets you handle large reads that may contain complicated queries.
//
func (m *mongoDB) New(isread bool) (*mgo.Database, *mgo.Session, error) {
	m.ml.Lock()
	defer m.ml.Unlock()

	// if m.master is alive then continue else, reset as empty.
	if err := m.master.Ping(); err != nil {
		m.master = nil
	}

	ses, err := getSession(m.Config)
	if err != nil {
		return nil, nil, err
	}

	m.master = ses

	if isread {
		copy := m.master.Copy()
		db := copy.DB(m.Config.DB)
		return db, copy, nil
	}

	clone := m.master.Clone()
	db := clone.DB(m.Config.DB)
	return db, clone, nil
}

// getSession attempts to retrieve the giving session for the given config.
func getSession(config Config) (*mgo.Session, error) {
	info := mgo.DialInfo{
		Addrs:    []string{config.Host},
		Timeout:  60 * time.Second,
		Database: config.AuthDB,
		Username: config.User,
		Password: config.Password,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	ses, err := mgo.DialWithInfo(&info)
	if err != nil {
		return nil, err
	}

	ses.SetMode(mgo.Monotonic, true)

	return ses, nil
}

//**********************************************************
// DB Functions
//**********************************************************

// AddIndex adds provided index if any to giving collection within database exposed by the provided
// MongoDB instance.
func AddIndex(db MongoDB, m metrics.Metrics, col string, indexes ...mgo.Index) error {
	defer m.CollectMetrics("DB.AddIndex")

	if len(indexes) == 0 {
		return nil
	}

	database, session, err := db.New(false)
	if err != nil {
		m.Emit(metrics.Errorf("Failed to create session for index"), metrics.With("collection", col), metrics.With("error", err.Error()))
		return err
	}

	defer session.Close()

	collection := database.C(col)

	for _, index := range indexes {
		if err := collection.EnsureIndex(index); err != nil {
			m.Emit(metrics.Errorf("Failed to ensure session index"), metrics.With("collection", col), metrics.With("index", index), metrics.With("error", err.Error()))
			return err
		}

		m.Emit(metrics.Info("Succeeded in ensuring collection index"), metrics.With("collection", col), metrics.With("index", index))
	}

	m.Emit(metrics.Info("Finished adding index"), metrics.With("collection", col))
	return nil
}

// Count attempts to return the total number of record from the db.
func Count(ctx context.Context, db MongoDB, m metrics.Metrics, col string) (int, error) {
	defer m.CollectMetrics("DB.Count")

	if isContextExpired(ctx) {
		err := ErrExpiredContext

		m.Emit(metrics.Errorf("Failed to get record count"), metrics.With("collection", col), metrics.With("error", err.Error()))
		return -1, err
	}

	database, session, err := db.New(true)
	if err != nil {
		m.Emit(metrics.Errorf("Failed to get record count"), metrics.With("collection", col), metrics.With("error", err.Error()))

		return -1, err
	}

	defer session.Close()

	query := bson.M{}
	total, err := database.C(col).Find(query).Count()
	if err != nil {
		m.Emit(metrics.Errorf("Failed to get record count"), metrics.With("collection", col), metrics.With("query", query), metrics.With("error", err.Error()))

		return -1, err
	}

	m.Emit(metrics.Info("Deleted record"), metrics.With("collection", col), metrics.With("query", query))

	return total, err
}

// Exec provides a function which allows the execution of a custom function against the collection.
func Exec(ctx context.Context, db MongoDB, m metrics.Metrics, col string, isread bool, fx func(col *mgo.Collection) error) error {
	defer m.CollectMetrics("DB.Exec")

	if isContextExpired(ctx) {
		err := ErrExpiredContext
		m.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("collection", col), metrics.With("error", err.Error()))
		return err
	}

	database, session, err := db.New(isread)
	if err != nil {
		m.Emit(metrics.Errorf("Failed to create session"), metrics.With("collection", col), metrics.With("error", err.Error()))
		return err
	}

	defer session.Close()

	if err := fx(database.C(col)); err != nil {
		m.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("collection", col), metrics.With("error", err.Error()))
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}

	m.Emit(metrics.Info("Operation executed"), metrics.With("collection", col))

	return nil
}

func isContextExpired(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
