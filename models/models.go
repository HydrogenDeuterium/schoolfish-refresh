// Code generated by sql2gorm. DO NOT EDIT.
package models

import (
	"time"
)

type Comments struct {
	CreatedAt   time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Cid         uint       `gorm:"column:cid;primary_key;AUTO_INCREMENT" json:"cid"`
	Product     uint       `gorm:"column:product;NOT NULL" json:"product"`
	Commentator uint       `gorm:"column:commentator;NOT NULL" json:"commentator"`
	ResponseTo  uint       `gorm:"column:response_to" json:"response_to"`
	Text        string     `gorm:"column:text" json:"text"`
}

type Images struct {
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Iid       uint       `gorm:"column:iid;AUTO_INCREMENT;NOT NULL" json:"iid"`
	Address   string     `gorm:"column:address;NOT NULL" json:"address"`
	Products  uint       `gorm:"column:products;NOT NULL" json:"products"`
}

type Messages struct {
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Mid       uint       `gorm:"column:mid;primary_key;AUTO_INCREMENT" json:"mid"`
	From      uint       `gorm:"column:from;NOT NULL" json:"from"`
	To        uint       `gorm:"column:to;NOT NULL" json:"to"`
}

type Products struct {
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Pid       uint       `gorm:"column:pid;primary_key;AUTO_INCREMENT" json:"pid"`
	Title     string     `gorm:"column:title;NOT NULL" json:"title"`
	Info      string     `gorm:"column:info" json:"info"`
	Price     string     `gorm:"column:price;NOT NULL" json:"price"`
	Owner     uint       `gorm:"column:owner;NOT NULL" json:"owner"`
	Location  string     `gorm:"column:location" json:"location"`
}

type Users struct {
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Uid       uint       `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"uid"`
	Username  string     `gorm:"column:username;NOT NULL" json:"username"`
	Email     string     `gorm:"column:email;NOT NULL" json:"email"`
	Hashed    string     `gorm:"column:hashed;NOT NULL" json:"hashed"` // 密码哈希
	Avatar    string     `gorm:"column:avatar" json:"avatar"`          // 头像
	Info      string     `gorm:"column:info" json:"info"`              // 简历
	Profile   string     `gorm:"column:profile" json:"profile"`        // 详历
	Location  string     `gorm:"column:location" json:"location"`      // 默认地址
}
