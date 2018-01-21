package types

import (
	"context"

	"github.com/gokit/mgokit/example/api"
)

// UserDBBackend defines a backend which represents the giving
// methods exposed by the DB implementation for the giving type User.
// @implement_mock
type UserDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem api.User) error
	Get(ctx context.Context, publicID string) (api.User, error)
	Update(ctx context.Context, publicID string, elem api.User) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]api.User, error)
	GetByField(ctx context.Context, key string, value interface{}) (api.User, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]api.User, int, error)
}
