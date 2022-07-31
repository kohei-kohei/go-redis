package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kohei-kohei/go-redis/domain"
)

type bread struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func GetBread(id int) (*domain.Bread, error) {
	db, err := sqlx.Open("mysql", "go_test:password@tcp(db:3306)/go_database?parseTime=true")
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var b bread
	err = db.Get(&b, "SELECT * FROM bread WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &domain.Bread{
		ID:   b.ID,
		Name: b.Name,
	}, nil
}
