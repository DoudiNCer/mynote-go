package dao

import "github.com/DoudiNCer/mynote-go/dao/sqlite"

// 调用初始化数据库函数
func InitDB() error {
	return sqlite.Init()
}
