package entity

import (
	"container/list"
	"fmt"
	"sync"
)

//AgendaService .
type AgendaService struct {
	storage *Storage
}

var agendaInstance *AgendaService
var agendamu sync.Mutex

//GetAgendaService .
func GetAgendaService() *AgendaService {
	if agendaInstance == nil {
		agendamu.Lock()
		defer agendamu.Unlock()
		if agendaInstance == nil {
			sto := GetStorage()
			agendaInstance = &AgendaService{}
			agendaInstance.storage = sto
			GetStorage().ReadFromDb()
		}
	}
	return agendaInstance
}

//GetAgendaServiceStorage .
func (agendaS *AgendaService) GetAgendaServiceStorage() *Storage {
	return agendaS.storage
}

//UserLogIn .
func (agendaS *AgendaService) UserLogIn(user User) bool {
	filter := func(ur User) bool {
		return ur.GetName() == user.GetName()
	}

	qr := agendaS.storage.QueryUser(filter)

	if qr.Len() <= 0 {
		return false
	}

	usr := qr.Front().Value.(User)

	agendaS.storage.Current = &usr
	return true
}

//UserRegister .
func (agendaS *AgendaService) UserRegister(user User) bool {
	nameFilter := func(ur User) bool {
		return ur.GetName() == user.GetName()
	}
	if agendaS.storage.QueryUser(nameFilter).Len() > 0 {
		return false
	}

	err := userService.Save(&user)
	if err != nil {
		return false
	}

	agendaS.storage.CreateUser(user)
	return true
}

// DeleteUser .
func (agendaS *AgendaService) DeleteUser(N string, P string) bool {
	filter := func(ur User) bool {
		return N == ur.GetName() && P == ur.GetPassword()
	}
	u := &User{}
	u.SetName(N)
	u.SetPassword(P)

	err := userService.Delete(u)
	if err != nil {
		return false
	}

	return agendaS.storage.DeleteUser(filter) > 0
}

//QueryAllUsers .
func (agendaS *AgendaService) QueryAllUsers() *list.List {
	filter := func(ur User) bool {
		return true
	}
	return agendaS.storage.QueryUser(filter)
}

//QueryUserByName .
func (agendaS *AgendaService) QueryUserByName(name string) User {
	filter := func(ur User) bool {
		return name == ur.Name
	}

	ulist := agendaS.storage.QueryUser(filter)
	if ulist.Len() > 0 {
		return ulist.Front().Value.(User)
	}

	return User{}

}

//UpdateUser .
func (agendaS *AgendaService) UpdateUser(u User) bool {
	tmp := agendaS.QueryUserByName(u.Name)
	if u.Password == "" {
		u.Password = tmp.Password
	}
	if u.Email == "" {
		u.Email = tmp.Email
	}
	if u.Phone == "" {
		u.Phone = tmp.Phone
	}

	tmp = u

	err := userService.Update(&tmp)
	if err != nil {
		return false
	}

	return agendaS.storage.UpdateUser(u.Name, u) > 0
}

// CreateMeeting .
func (agendaS *AgendaService) CreateMeeting(meeting Meeting) bool {
	titleFilter := func(mt Meeting) bool {
		return mt.GetTitle() == meeting.GetTitle()
	}

	if agendaS.storage.QueryMeeting(titleFilter).Len() > 0 {
		return false
	}

	overlapFilter := func(mt Meeting) bool {
		pas := meeting.GetAllPAS()
		for _, p := range pas {
			if mt.ContainParticipator(p) {
				return !(mt.GetStartDate().After(meeting.GetEndDate()) || mt.GetStartDate().Equal(meeting.GetEndDate()) || mt.GetEndDate().Before(meeting.GetStartDate()) || mt.GetEndDate().Equal(meeting.GetStartDate()))
			}
		}
		return false
	}

	if agendaS.storage.QueryMeeting(overlapFilter).Len() > 0 {
		return false
	}

	if !agendaS.storage.ContainUser(meeting.GetSponsor()) {
		return false
	}
	for _, p := range meeting.GetParticipators() {
		if !agendaS.storage.ContainUser(p) {
			return false
		}
	}

	err := meetingService.Save(&meeting)
	if err != nil {
		return false
	}

	fmt.Println("2")

	agendaS.storage.CreateMeeting(meeting)
	return true
}

//AddMeetingParticipators .
func (agendaS *AgendaService) AddMeetingParticipators(user, title string, ptors []string) {
	mt := agendaS.QueryMyMeetingByTitle(user, title)
	for _, p := range ptors {
		mt.AddParticipators(p)
	}

	err := meetingService.Update(&mt)
	if err != nil {
		return
	}

	agendaS.storage.UpdateMeetingByTitle(title, mt)
}

//DeelteMeetingParticipators .
func (agendaS *AgendaService) DeelteMeetingParticipators(user, title string, ptors []string) {
	mt := agendaS.QueryMyMeetingByTitle(user, title)
	for _, p := range ptors {
		mt.DeleteParticipator(p)
	}

	err := meetingService.Update(&mt)
	if err != nil {
		return
	}

	agendaS.storage.UpdateMeetingByTitle(title, mt)

}

//Quit .
func (agendaS *AgendaService) Quit(user, title string) {
	mt := agendaS.QueryAllMeetingByTitle(user, title)
	if mt.Sponsor == user {
		mt.Sponsor = ""
	} else {
		mt.DeleteParticipator(user)
	}

	err := meetingService.Update(&mt)
	if err != nil {
		return
	}

	agendaS.storage.UpdateMeetingByTitle(title, mt)

}

//QueryMeetingByUserAndTime .
func (agendaS *AgendaService) QueryMeetingByUserAndTime(meeting Meeting) *list.List {
	timeAndUserFilter := func(mt Meeting) bool {
		return (mt.ContainParticipator(meeting.Sponsor) && mt.StartDate.After(meeting.StartDate) && mt.EndDate.Before(meeting.EndDate))
	}

	return agendaS.storage.QueryMeeting(timeAndUserFilter)
}

//QueryMyMeetingByTitle .
func (agendaS *AgendaService) QueryMyMeetingByTitle(user, title string) Meeting {
	titleFilter := func(mt Meeting) bool {
		return user == mt.GetSponsor() && title == mt.GetTitle()
	}
	//可能nil
	return agendaS.storage.QueryMeeting(titleFilter).Front().Value.(Meeting)
}

//QueryAllMeetingByTitle .
func (agendaS *AgendaService) QueryAllMeetingByTitle(user, title string) Meeting {
	titleFilter := func(mt Meeting) bool {
		return mt.ContainParticipator(user) && title == mt.GetTitle()
	}
	//可能nil

	mtlist := agendaS.storage.QueryMeeting(titleFilter)
	if mtlist.Len() > 0 {
		return mtlist.Front().Value.(Meeting)
	}
	return Meeting{}
}

//DeleteMeetingByTitle .
func (agendaS *AgendaService) DeleteMeetingByTitle(user, title string) bool {
	titleFilter := func(mt Meeting) bool {
		return user == mt.GetSponsor() && title == mt.GetTitle()
	}

	mt := &Meeting{}
	mt.SetTitle(title)
	mt.SetSponsor(user)
	err := meetingService.Delete(mt)
	if err != nil {
		return false
	}

	return agendaS.storage.DeleteMeeting(titleFilter) > 0
}

//DeleteMeetingAll .
// func (agendaS *AgendaService) DeleteMeetingAll(user string) bool {
// 	allFilter := func(mt Meeting) bool {
// 		return user == mt.GetSponsor()
// 	}

// 	mt := &Meeting{}
// 	mt.SetSponsor(user)
// 	err := meetingService.Delete(user, mt)
// 	if err != nil {
// 		return false
// 	}

// 	return agendaS.storage.DeleteMeeting(allFilter) > 0
// }

//ReadFromDb .
func (agendaS *AgendaService) ReadFromDb() {
	agendaS.storage.ReadFromDb()
}
