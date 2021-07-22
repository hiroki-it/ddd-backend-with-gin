package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

// NewDB コンストラクタ
func NewDB(dbUser string, dbPassword string, dbHost string, dbPort string, dbDatabase string) (*DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbPassword, dbUser, dbHost, dbPort, dbDatabase)

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
	return d.conn.AutoMigrate()
}

// Close DBとの接続を終了します．
func (d *DB) Close() error {

	sqlDb, err := d.conn.DB()

	if err != nil {
		return err
	}

	return sqlDb.Close()
}

// Conn connを返却します．
func (d *DB) Conn() *gorm.DB {
	return d.conn
}
