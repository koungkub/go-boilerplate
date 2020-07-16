package connection

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewLog get log instance
func NewLog(svc string) *logrus.Entry {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	l := log.WithFields(logrus.Fields{
		"service-name": svc,
	})

	return l
}

// NewDB get db instance
func NewDB(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "[NewDB] prepare dsn")
	}

	// db.SetMaxOpenConns()

	return db, nil
}
