package repo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type QueryLogger struct {
	Queryer sqlx.Queryer
	Logger  *logrus.Logger
}

func (p *QueryLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	p.Logger.Infof(query, args...)
	return p.Queryer.Query(query, args...)
}

func (p *QueryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	p.Logger.Infof(query, args...)
	return p.Queryer.Queryx(query, args...)
}

func (p *QueryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	p.Logger.Infof(query, args...)
	return p.Queryer.QueryRowx(query, args...)
}
