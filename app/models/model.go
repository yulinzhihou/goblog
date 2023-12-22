// models
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/22

package models

import (
	"time"

	"myblog/pkg/types"
)

// BaseModel 模型基类
type BaseModel struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetStringID 获取 ID 的字符串格式
func (b BaseModel) GetStringID() string {
	return types.Uint64ToString(b.ID)
}
