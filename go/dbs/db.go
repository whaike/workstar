package adapter

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ISqlXPlus interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type SqlXAdapter struct {
	DB     *sqlx.DB
	Logger ILogger
}

func (s SqlXAdapter) Get(dest interface{}, query string, args ...interface{}) error {
	start := time.Now()
	err := s.DB.Get(dest, query, args...)
	s.Logger.Info("ISqlXPlus Get >> ",
		zap.String("query", query),
		zap.Any("args", args),
		zap.Error(err),
		zap.Int64("cost", time.Since(start).Milliseconds()))
	return err
}

func (s SqlXAdapter) Select(dest interface{}, query string, args ...interface{}) error {
	start := time.Now()
	err := s.DB.Select(dest, query, args...)
	s.Logger.Info("ISqlXPlus Select >> ",
		zap.String("query", query),
		zap.Any("args", args),
		zap.Error(err),
		zap.Int64("cost", time.Since(start).Milliseconds()))
	return err
}

func (s SqlXAdapter) QueryRow(query string, args ...interface{}) *sql.Row {
	start := time.Now()
	row := s.DB.QueryRow(query, args...)
	s.Logger.Info("ISqlXPlus QueryRow >> ",
		zap.String("query", query),
		zap.Any("args", args),
		zap.Int64("cost", time.Since(start).Milliseconds()))
	return row
}

func (s SqlXAdapter) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	res, err := s.DB.Exec(query, args...)
	s.Logger.Info("ISqlXPlus Exec >> ",
		zap.String("query", query),
		zap.Any("args", args),
		zap.Error(err),
		zap.Int64("cost", time.Since(start).Milliseconds()))
	return res, err
}

// 初始化 sqlx adapter
func NewSqlXAdapter(db *sqlx.DB, logger ILogger) ISqlXPlus {
	return SqlXAdapter{
		DB:     db,
		Logger: logger,
	}
}
