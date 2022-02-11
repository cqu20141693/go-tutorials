package mysql

import "time"

type Command struct {
	ID        uint      `gorm:"primaryKey"`
	GmtCreate time.Time `gorm:"gmt_create"`
	HourTime  uint      `gorm:"hour_time"`
}

func (c *Command) TableName() string {
	return "tb_command"
}
