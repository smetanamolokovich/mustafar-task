package kvstore

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

var (
	ErrConnectionError = errors.New("connection error")
	ErrNotFound        = errors.New("not found")
)

type KvStore struct {
	db        *sql.DB
	tableName string
}

type KvStoreRow struct {
	Key     string
	Value   []byte
	Expires time.Time
}

func New(db *sql.DB, tableName string) *KvStore {
	output := new(KvStore)
	output.db = db
	output.tableName = tableName

	err := output.createTable()
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func (s *KvStore) createTable() error {
	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS " + s.tableName + " (key text NOT NULL PRIMARY KEY, value text NOT NULL, expires timestamp(0) with time zone NOT NULL)")
	if err != nil {
		return err
	}

	s.db.Exec("CREATE INDEX key_index ON " + s.tableName + " (key)")

	return nil
}

func (s *KvStore) getRowByKey(key string) (*KvStoreRow, error) {
	row := new(KvStoreRow)
	query := "SELECT key, value, expires FROM " + s.tableName + " WHERE key = $1 AND expires < now()"
	err := s.db.QueryRow(query, key).Scan(&row.Key, &row.Value, &row.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return row, nil
}

func (s *KvStore) GetValue(key string) ([]byte, error) {
	row, err := s.getRowByKey(key)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return row.Value, nil
}

func (s *KvStore) SetValue(key string, value []byte, expires time.Time) error {

	row, err := s.getRowByKey(key)
	var query string

	if row == nil && err == nil {
		query = "INSERT INTO " + s.tableName + " (value, key, expires) VALUES($1, $2, $3)"
		_, err = s.db.Exec(query, value, key, expires)
		if err != nil {
			return err
		}
	} else {
		query = "UPDATE " + s.tableName + " SET value = $1 WHERE key = $2 AND expires < now()"
		_, err = s.db.Exec(query, value, key)
		if err != nil {
			return err
		}
	}

	return nil
}
