package dbo

import "gorm.io/gorm"

type dbMigration struct {
	models []any
}

func (m *dbMigration) Add(models ...any) {
	m.models = append(m.models, models...)
}

func (m *dbMigration) Run(db *gorm.DB, seed ...bool) (err error) {
	for _, model := range m.models {
		err = db.AutoMigrate(model)

		if err != nil {
			return
		}
	}

	if len(seed) > 0 && seed[0] {
		err = Seeder.Run(db)
	}

	return
}

// ======================

var Migration dbMigration

func initMigration(m *dbMigration) {
	m.models = make([]any, 0)
}

func init() {
	initMigration(&Migration)
}
