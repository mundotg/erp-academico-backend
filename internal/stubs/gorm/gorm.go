package gorm

import "time"

type Model struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt DeletedAt `json:"-"`
}
type DeletedAt struct {
	Time  time.Time
	Valid bool
}
type Config struct{}
type Dialector interface{}
type DB struct{ Error error }

func Open(Dialector, *Config) (*DB, error)           { return &DB{}, nil }
func (db *DB) AutoMigrate(...interface{}) error      { return nil }
func (db *DB) Preload(string, ...interface{}) *DB    { return db }
func (db *DB) Find(interface{}, ...interface{}) *DB  { return db }
func (db *DB) Where(interface{}, ...interface{}) *DB { return db }
func (db *DB) Create(interface{}) *DB                { return db }
