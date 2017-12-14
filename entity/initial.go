package entity

import (
	"encoding/json"

	"github.com/go-xorm/xorm"
	// .
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

// Connectdb .
func Connectdb() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	db, err := xorm.NewEngine("sqlite3", "./data/test.db")
	if err != nil {
		panic(err)
	}
	engine = db

	err = CreateMtTable()
	err = CreateUserTable()

}

// CheckErr .
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// ToJSONString .
func ToJSONString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}
