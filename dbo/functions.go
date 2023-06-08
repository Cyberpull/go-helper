package dbo

import (
	"cyberpull.com/gotk/errors"

	"gorm.io/gorm"
)

type DBFunction func(tx ...*gorm.DB) (value *gorm.DB, err error)

func DB(db *gorm.DB) DBFunction {
	// Database Function
	return func(tx ...*gorm.DB) (value *gorm.DB, err error) {
		if len(tx) > 0 && tx[0] != nil {
			value = tx[0]
			return
		}

		if db != nil {
			value = NewSession(db)
			return
		}

		err = errors.New("Database connection not found")

		return
	}
}

func NewSession(tx *gorm.DB) *gorm.DB {
	return tx.Session(&gorm.Session{
		NewDB: true,
	})
}
