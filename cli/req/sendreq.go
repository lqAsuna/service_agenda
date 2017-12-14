package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service_agenda/entity"
)

// mock server
//var host = "https://private-ea20c-agenda16.apiary-mock.com/#"

// local server
var host = "http://localhost:8080"

// UserPost .
func UserPost(ur entity.User) int {
	b, err := json.Marshal(ur)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))

	res, err := http.Post(host+"/v1/users", "application/json;charset=utf-8", body)
	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	if res.StatusCode == http.StatusCreated {
		fmt.Println("created successfully")
		fmt.Println(string(result))
	} else {
		fmt.Println("created failed")
	}

	return res.StatusCode
}

// UsersGet .
func UsersGet() int {
	res, err := http.Get(host + "/v1/users")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if res.StatusCode == http.StatusOK {
		fmt.Println(string(body))
	} else {
		fmt.Println("Get failed")
	}

	return res.StatusCode
}

// UserPatch .
func UserPatch(ur entity.User) int {
	client := &http.Client{}

	b, err := json.Marshal(ur)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))

	req, err := http.NewRequest(http.MethodPatch, host+"/v1/users/"+ur.Name, body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("updated successfully")
	} else {
		fmt.Println("updated failed")
	}

	return resp.StatusCode
}

// UserDelete .
func UserDelete(ur entity.User) int {
	client := &http.Client{}

	b, err := json.Marshal(ur)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))

	req, err := http.NewRequest(http.MethodDelete, host+"/v1/users/"+ur.Name, body)

	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	if http.StatusNoContent == resp.StatusCode {
		fmt.Println("deleted successfully")
	}

	return resp.StatusCode
}

// MeetingPost .
func MeetingPost(mt entity.Meeting) int {
	b, err := json.Marshal(mt)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))

	res, err := http.Post(host+"/v1/meetings", "application/json;charset=utf-8", body)
	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	if res.StatusCode == http.StatusCreated {
		fmt.Println("created successfully")
		fmt.Println(string(result))
	} else {
		fmt.Println("created failed")
	}

	return res.StatusCode
}

// MeetingGet .
func MeetingGet(title string, name string) int {
	res, err := http.Get(host + "/v1/meetings/" + title + "?name=" + name)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if res.StatusCode == http.StatusOK {
		fmt.Println(string(body))
	} else {
		fmt.Println("Get failed")
	}
	return res.StatusCode
}

// MeetingDelete .
func MeetingDelete(mt entity.Meeting) int {
	b, err := json.Marshal(mt)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, host+"/v1/meetings/"+mt.Title, body)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if http.StatusNoContent == resp.StatusCode {
		fmt.Println("Deleted successfully")
	} else {
		fmt.Println("Deleted failed")
	}

	return resp.StatusCode
}
