package test

import (
	"net/http"
	"service_agenda/cli/req"
	"service_agenda/entity"
	"testing"
)

//TestUserPost .
func TestUserPost(t *testing.T) {

	in1 := entity.GetUser("zhangsan", "12345", "zhangsan.mail2.sysu.edu", "13512510211")

	cases := []struct {
		in   entity.User
		want int
	}{
		{in1, http.StatusCreated},
	}
	for _, c := range cases {
		got := req.UserPost(c.in)
		if got != c.want {
			t.Errorf("getResult(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}

// TestUserGet .
func TestUserGet(t *testing.T) {

	got := req.UsersGet()
	if got != http.StatusOK {
		t.Errorf("got %q, want %q", got, http.StatusOK)
	}

}

// TestUserPatch
func TestUserPatch(t *testing.T) {
	in1 := entity.GetUser("zhangsan", "12345", "zhang3.mail2.sysu.edu", "13512510211")

	cases := []struct {
		in   entity.User
		want int
	}{
		{in1, http.StatusCreated},
	}
	for _, c := range cases {
		got := req.UserPatch(c.in)
		if got != c.want {
			t.Errorf("getResult(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// TestMeetingPost .
func TestMeetingPost(t *testing.T) {
	in1 := entity.GetMeeting("zhangsan", []string{"zhangsan", "zhangsan"}, "2017-01-22 12:00", "2017-02-10 10:00", "computer")

	cases := []struct {
		in   entity.Meeting
		want int
	}{
		{in1, http.StatusCreated},
	}
	for _, c := range cases {
		got := req.MeetingPost(c.in)
		if got != c.want {
			t.Errorf("getResult(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// TestMeetingGet .
func TestMeetingGet(t *testing.T) {
	in1 := entity.GetMeeting("zhangsan", []string{"zhangsan", "zhangsan"}, "2017-01-22 12:00", "2017-02-10 10:00", "computer")

	cases := []struct {
		in   entity.Meeting
		want int
	}{
		{in1, http.StatusOK},
	}
	for _, c := range cases {
		got := req.MeetingGet(c.in.Title, c.in.Sponsor)
		if got != c.want {
			t.Errorf("getResult(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// TestMeetingDelete .
func TestMeetingDelete(t *testing.T) {
	in1 := entity.GetMeeting("zhangsan", []string{"zhangsan", "zhangsan"}, "2017-01-22 12:00", "2017-02-10 10:00", "computer")

	cases := []struct {
		in   entity.Meeting
		want int
	}{
		{in1, http.StatusNoContent},
	}
	for _, c := range cases {
		got := req.MeetingDelete(c.in)
		if got != c.want {
			t.Errorf("getResult(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// TestUserDelete
func TestUserDelete(t *testing.T) {
	in1 := entity.GetUser("zhangsan", "12345", "zhang3.mail2.sysu.edu", "13512510211")

	cases := []struct {
		in   entity.User
		want int
	}{
		{in1, http.StatusNoContent},
	}
	for _, c := range cases {
		got := req.UserDelete(c.in)
		if got != c.want {
			t.Errorf("getResult(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
