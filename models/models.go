package models

import "time"

// @Summary 通用响应结构体
// @Description 通用响应结构体
// @param code int "状态码"
// @param message string "消息"
// @param data object "数据，类型根据接口变化"
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// @Summary 用户信息
// @Description 用户结构体
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	Group     string    `gorm:"default:'user'" json:"group"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

// @Summary 书籍信息
// @Description 书籍信息
type Book struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title  string `json:"title"`
	Author string
}

type BorrowRecord struct {
}
