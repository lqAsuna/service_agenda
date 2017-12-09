package entity

import (
	"encoding/json"

	"github.com/go-xorm/xorm"
	// .
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	db, err := xorm.NewEngine("sqlite3", "./test.db")
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

// func conv(m map[string]string) map[string]interface{} {
// 	result := make(map[string]interface{})
// 	for k, v := range m {
// 		result[k] = v
// 	}
// 	return result
// }

// //SetField .
// func SetField(obj interface{}, name string, value interface{}) error {
// 	name = firstToUpper(name, 1, len(name))
// 	structValue := reflect.ValueOf(obj).Elem()
// 	structFieldValue := structValue.FieldByName(name)

// 	if !structFieldValue.IsValid() {
// 		return fmt.Errorf("No such field: %s in obj", name)
// 	}

// 	if !structFieldValue.CanSet() {
// 		return fmt.Errorf("Cannot set %s field value", name)
// 	}

// 	structFieldType := structFieldValue.Type()
// 	val := reflect.ValueOf(value)
// 	if structFieldType != val.Type() {
// 		fmt.Println(val)
// 		return errors.New("Provided value type didn't match obj field type")
// 	}

// 	structFieldValue.Set(val)
// 	return nil
// }

// func firstToUpper(str string, begin, length int) (substr string) {
// 	// 将字符串的转换成[]rune
// 	rs := []rune(str)
// 	lth := len(rs)

// 	// 简单的越界判断
// 	if begin < 0 {
// 		begin = 0
// 	}
// 	if begin >= lth {
// 		begin = lth
// 	}
// 	end := begin + length
// 	if end > lth {
// 		end = lth
// 	}

// 	// 返回子串
// 	return strings.ToUpper(string(rs[0:begin])) + string(rs[begin:end])
// }
