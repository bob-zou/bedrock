package dao

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(New, NewDB, NewRedis)

//go:generate wire
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
}

// dao dao.
type dao struct {
	db    *gorm.DB
	redis *redis.Client
}

// New new a dao and return.
func New(r *redis.Client, db *gorm.DB) (d Dao, cf func(), err error) {
	return newDao(r, db)
}

func newDao(r *redis.Client, db *gorm.DB) (d *dao, cf func(), err error) {
	d = &dao{
		db:    db,
		redis: r,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	_ = d.redis.Close()
	if db, _ := d.db.DB(); db != nil {
		_ = db.Close()
	}
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
