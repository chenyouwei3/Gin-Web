package model

import (
	"time"
)

type Moment struct {
	Id         int64     `gorm:"column:id;type:bigint;primarykey;not null"`
	CreateTime time.Time `gorm:"type:datetime;comment:'createTime'"`
	UserId     int64     `gorm:"not null;index"`
	User       User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content    string    `gorm:"size:2048"`
}

type Comment struct {
	Id         int64     `gorm:"column:id;type:bigint;primarykey;not null"`
	CreateTime time.Time `gorm:"type:datetime(3);comment:'createTime'"`
	MomentId   int64     `gorm:"not null;index"`
	Moment     Moment    `gorm:"foreignKey:MomentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId     int64     `gorm:"not null;index"`
	User       User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ParentId   *int      `gorm:"index;default:NULL"`
	Parent     *Comment  `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content    string    `gorm:"not null;size:1024"`
}

type CommentTree struct {
	CommentId int                    `json:"cid"`
	Content   string                 `json:"content"`
	Author    map[string]interface{} `json:"author"`
	CreatedAt string                 `json:"createdAt"`
	Children  []*CommentTree         `json:"children"`
}

type AddMomentForm struct {
	UserId  int64  `form:"uid" json:"uid" binding:"required"`
	Content string `form:"content" json:"content" binding:"required,max=2048"`
}

type AddCommentForm struct {
	UserId   int64  `json:"uid"`
	MomentId int64  `json:"mid"`
	ParentId int    `json:"pid"`
	Content  string `json:"content"`
}
