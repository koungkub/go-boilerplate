package model

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// Conf struct contain configuration
type Conf struct {
	DB  *sql.DB
	Log *logrus.Entry
}

// NewConf get conf
func NewConf(db *sql.DB, log *logrus.Entry) *Conf {
	return &Conf{
		DB:  db,
		Log: log,
	}
}
