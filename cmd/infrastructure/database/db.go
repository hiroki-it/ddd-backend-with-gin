package database

import (
	"fmt"
	"os"

	"github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	user "github.com/hiroki-it/ddd-backend-with-gin/cmd/infrastructure/user/dtos"
)

// DB ORMのメソッドをラッピングしやすいように，ORMを構造体に保持するようにします．
type DB struct {
	conn *gorm.DB
}

// NewDB コンストラクタ
func NewDB() (*DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
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
	err := d.conn.AutoMigrate(&user.UserDTO{})

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

// Create 集約を作成します．
func (d *DB) Create(DTO infrastructure.DTO) error {
	return d.conn.Create(DTO).Error
}

// Find 集約を一つ取得します．
func (d *DB) Find(DTO infrastructure.DTO, id int) error {
	return d.conn.First(DTO, id).Error
}

// FindAll 集約を全て取得します．
func (d *DB) FindAll(DTO infrastructure.DTO) error {
	return d.conn.First(DTO).Error
}

// Updates 集約の複数値を更新します．
func (d *DB) Updates(DTO infrastructure.DTO) error {
	return d.conn.Select("*").Updates(DTO).Error
}

// Delete 集約を削除します．
func (d *DB) Delete(DTO infrastructure.DTO) error {
	return d.conn.Delete(DTO).Error
}
