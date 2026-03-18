package model

import "time"

// Points 积分表
type Points struct {
	ID          int64     `json:"id"`          // ID
	MemberID    int64     `json:"member_id"`   // 会员ID
	Type        int       `json:"type"`        // 类型: 1-收入, 2-支出
	Points      int       `json:"points"`      // 积分数量
	Source      string    `json:"source"`      // 来源
	SourceID    string    `json:"source_id"`   // 来源ID
	Description string    `json:"description"` // 描述
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
}

// PointsRule 积分规则表
type PointsRule struct {
	ID         int64     `json:"id"`          // 规则ID
	Name       string    `json:"name"`        // 规则名称
	Code       string    `json:"code"`        // 规则代码
	Type       int       `json:"type"`        // 类型: 1-收入, 2-支出
	Points     int       `json:"points"`      // 积分数量
	PointsRate float64   `json:"points_rate"` // 抵扣比例
	Status     int       `json:"status"`      // 状态
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
}
