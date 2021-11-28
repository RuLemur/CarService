package repo

//
//import (
//	"context"
//	"database/sql"
//	"fmt"
//	"github.com/jmoiron/sqlx"
//	"github.com/sirupsen/logrus"
//)
//
//type QueryLogger struct {
//	Queryer sqlx.Queryer
//	Logger  *logrus.Logger
//}
//
//func (p *QueryLogger) Init(ctx context.Context) error {
//	db := ctx.Value("db").(*sqlx.DB)
//	dbLogger := &QueryLogger{Queryer: db, Logger: &logrus.Logger{}}
//	ctx = context.WithValue(ctx, "dbLogger", dbLogger)
//	return nil
//}
//
//func (p *QueryLogger) Ping(ctx context.Context) error {
//	fmt.Println("ping")
//	return nil
//}
//
//func (p *QueryLogger) Close() error {
//	return nil
//}
//
//func (p *QueryLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
//	p.Logger.Infof(query, args...)
//	return p.Queryer.Query(query, args...)
//}
//
//func (p *QueryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
//	p.Logger.Infof(query, args...)
//	return p.Queryer.Queryx(query, args...)
//}
//
//func (p *QueryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
//	p.Logger.Infof(query, args...)
//	return p.Queryer.QueryRowx(query, args...)
//}
