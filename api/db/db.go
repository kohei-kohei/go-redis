package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Bread struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func GetBread(id int) (*Bread, error) {
	db, err := sqlx.Open("mysql", "go_test:password@tcp(db:3306)/go_database?parseTime=true")
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var bread Bread
	err = db.Get(&bread, "SELECT * FROM bread WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &bread, nil
}
