package dbo

import (
	"cyberpull.com/gotk/v2/errors"

	"gorm.io/gorm"
)

type Instance interface {
	New() *gorm.DB
	DB(db ...*gorm.DB) (value *gorm.DB, err error)
	Migrate(seed ...bool) (err error)
	Seed() (err error)
}

// ======================

type dbInstance struct {
	db *gorm.DB
}

func (s *dbInstance) New() *gorm.DB {
	return NewSession(s.db)
}

func (s *dbInstance) DB(db ...*gorm.DB) (value *gorm.DB, err error) {
	if len(db) > 0 && db[0] != nil {
		value = db[0]
		return
	}

	if s.db != nil {
		value = s.New()
		return
	}

	err = errors.New("Database connection not found")

	return
}

func (s *dbInstance) Migrate(seed ...bool) (err error) {
	return Migration.Run(s.db, seed...)
}

func (s *dbInstance) Seed() (err error) {
	return Seeder.Run(s.db)
}

// ======================

func NewInstance(db *gorm.DB) Instance {
	return &dbInstance{
		db: db,
	}
}
