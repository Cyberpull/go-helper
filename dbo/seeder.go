package dbo

import "gorm.io/gorm"

type SeederHandler func(db *gorm.DB) (err error)

type dbSeeder struct {
	handlers []SeederHandler
}

func (s *dbSeeder) Add(handlers ...SeederHandler) {
	s.handlers = append(s.handlers, handlers...)
}

func (s *dbSeeder) Run(db *gorm.DB) (err error) {
	for _, handler := range s.handlers {
		tx := NewSession(db)

		if err = handler(tx); err != nil {
			return
		}
	}

	return
}

// ===================

var Seeder dbSeeder

func init() {
	Seeder.handlers = make([]SeederHandler, 0)
}
