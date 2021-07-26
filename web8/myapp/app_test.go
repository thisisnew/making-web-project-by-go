package myapp

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No Users")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:89")

	res, err = http.Get(ts.URL + "/users/56")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:56")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"thisisnew", "last_name":"kim", "email":"thisisnew@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := &User{}
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	user2 := &User{}
	err = json.NewDecoder(res.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/users/1", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID:1")
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/users", strings.NewReader(`{"id":1, "first_name":"updated", "last_name":"updated", "email":"updated@updated.com"}`))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID:1")

	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"thisisnew", "last_name":"kim", "email":"thisisnew@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := &User{}
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	updateStr := fmt.Sprintf(`{"id":%d, "first_name":"jason"}`, user.ID)

	req, _ = http.NewRequest(http.MethodPut, ts.URL+"/users", strings.NewReader(updateStr))
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	updateUser := &User{}
	err = json.NewDecoder(res.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(updateUser.ID, user.ID)
	assert.Equal("jason", updateUser.FirstName)
	//assert.Equal("", updateUser.LastName)
	assert.Equal(updateUser.Email, user.Email)
}

func TestUsers_WithUsersData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"thisisnew", "last_name":"kim", "email":"thisisnew@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"json", "last_name":"lee", "email":"jason@ttt.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	res, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	users := []*User{}
	err = json.NewDecoder(res.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))
}
