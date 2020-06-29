package worker

import "database/sql"

type Review struct {
	DB *sql.DB
}
