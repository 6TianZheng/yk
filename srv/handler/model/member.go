package model

import "time"

// Member 会员表
type Member struct {
	ID        int64     `json:"id"`         // 会员ID
	Username  string    `json:"username"`   // 用户名
	Password  string    `json:"-"`          // 密码
	Mobile    string    `json:"mobile"`     // 手机号
	Nickname  string    `json:"nickname"`   // 昵称
	Avatar    string    `json:"avatar"`     // 头像
	LevelID   int64     `json:"level_id"`   // 会员等级ID
	Points    int       `json:"points"`     // 积分
	Balance   float64   `json:"balance"`    // 余额
	Status    int       `json:"status"`     // 状态: 0-正常, 1-禁用
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
