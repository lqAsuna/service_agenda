package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service_agenda/entity"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func userRegisterHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()

		u := entity.User{}
		err := json.Unmarshal(result, &u)
		entity.CheckErr(err)
		if !entity.GetAgendaService().UserRegister(u) {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Register failed!!!"})
			return
		}
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserByNameHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		name := vars["name"]

		u := entity.GetAgendaService().QueryUserByName(name)

		formatter.JSON(w, http.StatusOK, u)
	}
}

func getAllUsersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ulist := entity.GetAgendaService().QueryAllUsers()

		var userSlice []entity.User

		for e := ulist.Front(); e != nil; e = e.Next() {
			userSlice = append(userSlice, e.Value.(entity.User))
		}

		formatter.JSON(w, http.StatusOK, userSlice)
	}
}

func deleteUserByNameHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()

		u := entity.User{}
		err := json.Unmarshal(result, &u)
		entity.CheckErr(err)
		if !entity.GetAgendaService().DeleteUser(u.Name, u.Password) {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"delete failed!!!"})
			return
		}
		formatter.JSON(w, http.StatusNoContent, struct{ info string }{"Deleted"})
	}
}

func updateUserByNameHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()

		u := entity.User{}
		err := json.Unmarshal(result, &u)
		entity.CheckErr(err)

		if !entity.GetAgendaService().UpdateUser(u) {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"update failed!!!"})
			return
		}
		formatter.JSON(w, http.StatusCreated, struct{ info string }{"Updated"})
	}
}

func createMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()

		mt := entity.Meeting{}
		err := json.Unmarshal(result, &mt)
		entity.CheckErr(err)

		if !entity.GetAgendaService().CreateMeeting(mt) {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"creating a meeting failed!!!"})
			return
		}
		formatter.JSON(w, http.StatusOK, mt)
	}
}

func getMeetingByTitleHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		title := vars["title"]

		req.ParseForm()
		//name := req.Form["name"][0]
		name := "zhangsan"

		fmt.Println(title, name)

		mt := entity.GetAgendaService().QueryAllMeetingByTitle(name, title)

		formatter.JSON(w, http.StatusOK, mt)
	}
}

func deleteMeetingByTitleHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()

		mt := entity.Meeting{}
		err := json.Unmarshal(result, &mt)
		entity.CheckErr(err)
		if !entity.GetAgendaService().DeleteMeetingByTitle(mt.Sponsor, mt.Title) {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"delete failed!!!"})
			return
		}
		formatter.JSON(w, http.StatusNoContent, struct{ info string }{"Deleted"})
	}
}
