package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

// NewDB コンストラクタ
func NewDB() (*DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DB{
		conn: conn,
	}, nil
}

// AutoMigrate マイグレーションを実行します．
func (d *DB) AutoMigrate() error {
	err := d.conn.AutoMigrate()

	if err != nil {
		return err
	}

	return nil
}

// Close DBとの接続を終了します．
func (d *DB) Close() error {
	sqlDb, err := d.conn.DB()

	if err != nil {
		return err
	}

	err = sqlDb.Close()

	if err != nil {
		return err
	}

	return nil
}

// Conn connを返却します．
func (d *DB) Conn() *gorm.DB {
	return d.conn
}
