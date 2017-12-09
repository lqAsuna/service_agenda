package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service_agenda/entity"
	"testing"
)

func getBody(u interface{}) *bytes.Buffer {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))

	return body
}

func getResult(u entity.User) entity.User {
	body := getBody(u)
	res, err := http.Post("http://localhost:8080/v1/users", "application/json;charset=utf-8", body)
	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	retUser := entity.User{}

	if err := json.Unmarshal(result, &retUser); err != nil {
		panic(err)
	}

	return retUser
}

// TestUserRegister .
func TestUserRegister(t *testing.T) {

	in1 := entity.GetUser("zhangsan", "12345", "zhangsan.mail2.sysu.edu", "13512510211")
	want1 := in1
	cases := []struct {
		in, want entity.User
	}{
		{in1, want1},
	}
	for _, c := range cases {
		got := getResult(c.in)
		if got != c.want {
			t.Errorf("getResult(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}

func getUserByName(name string) entity.User {
	res, err := http.Get("http://localhost:8080/v1/users/" + name)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	u := entity.User{}

	if err := json.Unmarshal(body, &u); err != nil {
		panic(err)
	}

	return u
}

//TestGetUserByNameHandler .
func TestGetUserByNameHandler(t *testing.T) {

	cases := []struct {
		in   string
		want entity.User
	}{
		{"zhangsan", entity.GetUser("zhangsan", "12345", "zhangsan.mail2.sysu.edu", "13512510211")},
	}
	for _, c := range cases {
		got := getUserByName(c.in)
		if got != c.want {
			t.Errorf("getUserByName(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func getMtResult(mt entity.Meeting) entity.Meeting {
	body := getBody(mt)
	res, err := http.Post("http://localhost:8080/v1/meetings", "application/json;charset=utf-8", body)
	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	retMeeting := entity.Meeting{}

	if err := json.Unmarshal(result, &retMeeting); err != nil {
		panic(err)
	}

	return retMeeting
}

// TestCreateMeetingHandler .
func TestCreateMeetingHandler(t *testing.T) {
	in1 := entity.GetMeeting("zhangsan", []string{"zhangsan", "zhangsan"}, "2017-01-22 12:00", "2017-02-10 10:00", "computer")
	want1 := in1
	cases := []struct {
		in, want entity.Meeting
	}{
		{in1, want1},
	}
	for _, c := range cases {
		getMtResult(c.in)
		// if got != c.want {
		// 	t.Errorf("getResult(%q) == %q, want %q", c.in, got, c.want)
		// }
	}

}

// TestUpdateUserByNameHandler .
func TestUpdateUserByNameHandler(t *testing.T) {
	client := &http.Client{}

	u := entity.GetUser("zhangsan", "12345", "z3.mail2.sysu.edu", "13512510211")
	req, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/v1/users/zhangsan", getBody(u))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	if http.StatusCreated != resp.StatusCode {
		t.Errorf("ret %v, want %v", resp.StatusCode, http.StatusCreated)
	}
}

// TestDeleteUserByNameHandler .
func TestDeleteUserByNameHandler(t *testing.T) {
	client := &http.Client{}

	u := entity.GetUser("zhangsan", "12345", "z3.mail2.sysu.edu", "13512510211")
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/v1/users/zhangsan", getBody(u))

	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	if http.StatusNoContent != resp.StatusCode {
		t.Errorf("ret %v, want %v", resp.StatusCode, http.StatusNoContent)
	}
}

// TestdeleteMeetingByTitleHandler .
func TestDeleteMeetingByTitleHandler(t *testing.T) {
	client := &http.Client{}

	mt := entity.GetMeeting("zhangsan", []string{"zhangsan", "zhangsan"}, "2017-01-22 12:00", "2017-02-10 10:00", "computer")
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/v1/meetings/title", getBody(mt))

	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	if http.StatusNoContent != resp.StatusCode {
		t.Errorf("ret %v, want %v", resp.StatusCode, http.StatusNoContent)
	}
}
