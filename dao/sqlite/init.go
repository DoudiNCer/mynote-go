package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

const (
	// 本地配置文件夹
	cfgdir = "~/.config/mynote-go"
	// SQLite DSN
	dsn = "~/.config/mynote-go/note.db?_auth_user=zml&_auth_pass=19930111&mode=rwc"
	// 初始化表使用的 SQL 语句
	mkTable = `
		CREATE TABLE IF NOT EXIST "note" (
  			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  			"date" TIMESTAMP NOT NULL,
  			"title" VARCHAR(20),
  			"context" TEXT
		);
	`
)

func Init() error {
	// 检查工作目录是否存在
	var err error
	err = os.Chdir(cfgdir)
	if err != nil {
		err := os.MkdirAll(cfgdir, 0755)
		if err != nil {
			return err
		}
		err = nil
	}
	// 连接数据库
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// 初始化数据表
	err = DB.Exec(mkTable).Error
	if err != nil {
		return err
	}
	return nil
}
