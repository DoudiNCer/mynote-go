package sqlite

import (
	"time"
)

// Note 表
type Note struct {
	Id      int64     `gorm:"column:id"`       // 日记ID
	Date    time.Time `gorm:"column:date"`     // 创建时间
	Title   string    `gorm:"column:title"`    // 日记标题
	Context string    `gorm:"context:context"` // 日记内容
}

// 表名
func (note Note) TableName() string {
	return "note"
}

// 插入新日记
func InsertNote(title, context string) {
	var note Note
	note.Date = time.Now()
	note.Title = title
	note.Context = context
	DB.Create(note)
}

// 查询所有日记
func QueryAllNotes() []Note {
	var notes []Note
	DB.Find(&notes)
	return notes
}

// 根据 ID 删除指定日记
func DropNoteByID(id int64) {
	DB.Delete(&Note{}, id)
}

// 修改指定日记
func UpdateNoteById(note Note) {
	DB.Save(note)
}
