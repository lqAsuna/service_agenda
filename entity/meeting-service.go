package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

//MeetingAtomicService .
type MeetingAtomicService struct{}

var meetingService = MeetingAtomicService{}

var meetingQueryAll = "SELECT * FROM meeting"
var meetingInsertSmt = "INSERT INTO meeting VALUES (?,?,?,?,?)"
var meetingUpdateSmt = "UPDATE meeting SET Sponsor=?,Participators=?,StartDate=?,EndDate=?,Title=? WHERE Title = ?"
var meetingDeleteSmt = "DELETE FROM meeting WHERE Sponsor=? AND Title=?"

var createMeetingTableSmt = "CREATE TABLE `meeting` (`Sponsor` TEXT NULL, `Participators` TEXT NULL, `StartDate` TEXT NULL, `EndDate` TEXT NULL, `Title` TEXT NULL)"

// CreateMtTable .
func CreateMtTable() error {
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(createMeetingTableSmt)

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
func (*MeetingAtomicService) Save(mt *Meeting) error {

	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(meetingInsertSmt, mt.GetSponsor(),
		ToJSONString(mt.GetParticipators()), ToJSONString(mt.GetStartDate()),
		ToJSONString(mt.GetEndDate()), mt.GetTitle())

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

// FindAll .
func (*MeetingAtomicService) FindAll() []Meeting {

	results, err := engine.QueryString(meetingQueryAll)
	CheckErr(err)

	mtSlice := make([]Meeting, 0, 0)

	for _, result := range results {
		var part []string
		if err := json.Unmarshal([]byte(result["Participators"]), &part); err != nil {
			panic(err)

		}

		var st time.Time
		if err := json.Unmarshal([]byte(result["StartDate"]), &st); err != nil {
			panic(err)
		}

		var et time.Time
		if err := json.Unmarshal([]byte(result["EndDate"]), &et); err != nil {
			panic(err)
		}

		mt := Meeting{}
		mt.Sponsor = result["Sponsor"]
		mt.Participators = part
		mt.StartDate = st
		mt.EndDate = et
		mt.Title = result["Title"]

		mtSlice = append(mtSlice, mt)
	}

	return mtSlice
}

// Update .
func (*MeetingAtomicService) Update(mt *Meeting) error {
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(meetingUpdateSmt, mt.GetSponsor(),
		ToJSONString(mt.GetParticipators()), ToJSONString(mt.GetStartDate()),
		ToJSONString(mt.GetEndDate()), mt.GetTitle(), mt.GetTitle())

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

// Delete .
func (*MeetingAtomicService) Delete(mt *Meeting) error {
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(meetingDeleteSmt, mt.Sponsor, mt.Title)
	fmt.Println(mt.Sponsor, mt.Title)

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
