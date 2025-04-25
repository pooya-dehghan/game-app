package migrator

import (
	"database/sql"
	"fmt"

	"github.com/pooya-dehghan/repository/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect  string
	dbConfig mysql.Config
	migrate  *migrate.FileMigrationSource
}

func New(dbConfig mysql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	return Migrator{dialect: "mysql", dbConfig: dbConfig, migrate: migrations}
}

func (m *Migrator) Up() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s",
		m.dbConfig.Username,
		m.dbConfig.Password,
		m.dbConfig.Host,
		m.dbConfig.Port,
		m.dbConfig.DBName))

	if err != nil {
		panic(fmt.Errorf("cant open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrate, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("cant migrate mysql db: %v", err))
	}

	fmt.Printf("Applied %d migrations!\n", n)
}

func (m *Migrator) Down() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s",
		m.dbConfig.Username,
		m.dbConfig.Password,
		m.dbConfig.Host,
		m.dbConfig.Port,
		m.dbConfig.DBName))

	if err != nil {
		panic(fmt.Errorf("cant open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrate, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("cant rollback migrations: %v", err))
	}

	fmt.Printf("Rollbacked %d migrations!\n", n)
}

func (m *Migrator) Status() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s",
		m.dbConfig.Username,
		m.dbConfig.Password,
		m.dbConfig.Host,
		m.dbConfig.Port,
		m.dbConfig.DBName))

	if err != nil {
		panic(fmt.Errorf("cant open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrate, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("cant migrate mysql db: %v", err))
	}

	fmt.Printf("Applied %d migrations!\n", n)

}
