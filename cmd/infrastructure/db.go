package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	persist *gorm.DB
}

// NewDB コンストラクタ
func NewDB(dbUser string, dbPassword string, dbHost string, dbPort string, dbDatabase string) (*DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbPassword, dbUser, dbHost, dbPort, dbDatabase)

	persist, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DB{
		persist: persist,
	}, nil
}

// AutoMigrate マイグレーションを実行します．
func (d *DB) AutoMigrate() error {
	return d.persist.AutoMigrate()
}

// Close DBとの接続を終了します．
func (d *DB) Close() error {

	sqlDb, err := d.persist.DB()

	if err != nil {
		return err
	}

	return sqlDb.Close()
}

// Persist persistを返却します．
func (d *DB) Persist() *gorm.DB {
	return d.persist
}
