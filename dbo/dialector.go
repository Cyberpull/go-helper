package dbo

import (
	"fmt"

	"cyberpull.com/gotk/v2/errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func dialector(opts *Options) (conn gorm.Dialector, err error) {
	switch dbDriver(opts) {
	case DRIVER_MYSQL:
		conn = func() gorm.Dialector {
			if opts.DSN != "" {
				return mysql.Open(opts.DSN)
			}

			return mysql.Open(fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				opts.Username,
				opts.Password,
				opts.Host,
				opts.Port,
				opts.DBName,
			))
		}()

	case DRIVER_PGSQL:
		conn = func() gorm.Dialector {
			if opts.DSN != "" {
				return postgres.Open(opts.DSN)
			}

			return postgres.Open(fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s",
				opts.Username,
				opts.Password,
				opts.Host,
				opts.Port,
				opts.DBName,
			))
		}()

	case DRIVER_SQLITE:
		conn = sqlite.Open(opts.DSN)

	default:
		err = errors.New("DB Driver not available")
	}

	return
}

func dbDriver(opts *Options) DRIVER {
	if opts.Driver == "" {
		opts.Driver = DRIVER_PGSQL
	}

	return opts.Driver
}
