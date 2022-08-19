package dummy

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	IsActive() (ok bool, err error)
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"kolo_test_db"`
}

// PGRepo is postgres implementation
type PGRepo struct {
	log *logrus.Logger
	db  *pg.DB
}

// NewDBRepository returns a new persistence layer object which can be used for
// CRUD on db
func NewDBRepository(i NewRepositoryIn) (repo Repository, err error) {

	repo = &PGRepo{
		log: i.Log,
		db:  i.DB,
	}

	return repo, nil
}

// IsActive checks if DB is connected
func (r *PGRepo) IsActive() (ok bool, err error) {
	ctx := context.Background()
	err = r.db.Ping(ctx)
	if err == nil {
		ok = true
	}
	return
}
