package go_tutorials

import (
	"fmt"
	"github.com/cqu20141693/go-service-common/v2/boot"
	"github.com/cqu20141693/go-service-common/v2/mysql"
	mysql2 "github.com/cqu20141693/go-tutorials/mysql"
	"testing"
	"time"
)

func Create() uint {
	command := mysql2.Command{GmtCreate: time.Now()}
	result := mysql.MysqlDB.Create(&command)
	if result.Error == nil {
		fmt.Sprintf("id=%d,rows=%d", command.ID, result.RowsAffected)
	}
	return command.ID
}

func init() {
	boot.Task()
}

func TestFirst(t *testing.T) {
	command := mysql2.Command{}
	// Get the first record ordered by primary key
	first := mysql.MysqlDB.First(&command)
	if first.Error == nil {
		fmt.Println(command)
	}
}

func TestOrder(t *testing.T) {
	limit := 1
	id := 0
	var commands []mysql2.Command
	find := mysql.MysqlDB.Limit(limit).Order("id asc").Select("id", "gmt_create").Where("id >?", id).Find(&commands)
	if find.Error == nil {
		fmt.Println(commands)
	}
}

func TestUpdate(t *testing.T) {
	command := mysql2.Command{ID: 1}
	hour := 10
	// 如果model中primaryKey 存在，则条件会增加id=1
	update := mysql.MysqlDB.Model(&command).Update("hour_time", hour)
	update = mysql.MysqlDB.Model(&mysql2.Command{}).Where("id=0").Update("hour_time", hour)
	if update.Error == nil {
		fmt.Println(update.RowsAffected)
	}
}
