package services

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Person struct {
	ID        int32     `db:"id"`
	UID       string    `db:"uid"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"createdAt"`
	NickName  string    `db:"nickName"`
}

func (person *Person) GetAll() (*sqlx.Rows, error) {
	rows, err := db.Queryx(`SELECT id, phone, "createdAt", "nickName" FROM primary_user_user ORDER BY id`)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
