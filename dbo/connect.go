package dbo

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connect(opts *Options) (i Instance, err error) {
	var db *gorm.DB
	var conn gorm.Dialector

	config := &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{},
	}

	if conn, err = dialector(opts); err != nil {
		return
	}

	if db, err = gorm.Open(conn, config); err != nil {
		return
	}

	switch dbDriver(opts) {
	case DRIVER_PGSQL:
		err = db.Exec(`SET DEFAULT_TRANSACTION_ISOLATION TO SERIALIZABLE`).Error
		// SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
		// SET DEFAULT_TRANSACTION_ISOLATION TO SERIALIZABLE;
	}

	i = NewInstance(db, opts)

	return
}
