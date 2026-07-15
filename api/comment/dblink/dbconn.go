package dblink

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"xorm.io/xorm"
)

func Database(connString string) *xorm.Engine {
	var err error
	engine, err := xorm.NewEngine("mysql", connString)
	if connString == "" || err != nil {
		fmt.Printf("Fail to create xorm system logger: %v\n", err)
	}
	engine.SetConnMaxLifetime(10000 * time.Hour)
	engine.SetMaxIdleConns(20)
	engine.ShowSQL(true)
	return engine
}

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format("2006-01-02 15:04:05") + `"`), nil
}
