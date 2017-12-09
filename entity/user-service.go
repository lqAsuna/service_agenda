package entity

import "fmt"

//UserAtomicService .
type UserAtomicService struct{}

//UserService .
var userService = UserAtomicService{}

var userQueryAll = "SELECT * FROM user"

//var userInsertSmt = "INSERT user SET name=?,password=?,email=?,phone=?"
var userInsertSmt = "INSERT INTO user VALUES (?,?,?,?)"
var userUpdateSmt = "UPDATE user SET Name=?,Password=?,Email=?,Phone=? WHERE Name = ?"
var userDeleteSmt = "DELETE FROM user WHERE name=? AND password=?"

var createUserTableSmt = "CREATE TABLE `user` (`Name` TEXT NULL, `Password` TEXT NULL, `Email` TEXT NULL, `Phone` TEXT NULL)"

// CreateUserTable .
func CreateUserTable() error {
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(createUserTableSmt)

	if err != nil {
		session.Rollback()
		fmt.Println(err)
		return err
	}

	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Save .
func (*UserAtomicService) Save(u *User) error {

	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(userInsertSmt, u.Name, u.Password, u.Email, u.Phone)

	if err != nil {
		fmt.Println(err)
		session.Rollback()
		return err
	}

	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

// FindAll .
func (*UserAtomicService) FindAll() []User {

	results, err := engine.QueryString(userQueryAll)
	CheckErr(err)

	var uSlice []User
	for _, result := range results {
		u := User{}
		// (&u).FillStruct(conv(result))
		u.Name = result["Name"]
		u.Password = result["Password"]
		u.Email = result["Email"]
		u.Phone = result["Phone"]

		uSlice = append(uSlice, u)
	}

	return uSlice
}

//Update .
func (*UserAtomicService) Update(u *User) error {
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(userUpdateSmt, u.Name, u.Password, u.Email, u.Phone, u.Name)

	if err != nil {
		session.Rollback()
		return err
	}

	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

//Delete .
func (*UserAtomicService) Delete(u *User) error {

	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(userDeleteSmt, u.Name, u.Password)

	if err != nil {
		session.Rollback()
		return err
	}

	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

// FindByID .
// func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
// 	var result = make(map[string]string)
// 	has, err := engine.Where("uid = ?", id).Get(&result)
// 	checkErr(err)
// 	if has {
// 		u := UserInfo{}
// 		FillStruct(&u, conv(result))
// 		return &u
// 	}
// 	return &UserInfo{}
// }
