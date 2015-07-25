package models

import (
	"errors"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"

	"log"
	"reflect"
)

type User struct {
	Id      int64     `json:"id" datastore:"-"`
	Name    string    `json:"text" datastore:",noindex"`
	Pass    string    `json:"text" datastore:",noindex"`
	Email   string    `json:"text" datastore:",noindex"`
	Created time.Time `json:"created"`
}

func DefaultUserKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "User", "default", 0, nil)
}

// AddUser add a user to datastore
func AddUser(u *User, c appengine.Context) (*datastore.Key, error) {

	key := datastore.NewIncompleteKey(c, "User", DefaultUserKey(c))
	_, err := datastore.Put(c, key, u)
	log.Println("AddUser receibed an object of type", reflect.TypeOf(u))
	return key, err
}

// CheckExist verifies that the user has not been created earlier
func CheckExist(u *User, c appengine.Context) (bool, error) {

	q := datastore.NewQuery("User").Filter("Email =", u.Email)
	var users []User
	_, err := q.GetAll(c, &users)
	if err != nil {
		return false, errors.New("Error veryfing if user exist," + err.Error())
	}
	log.Println("Before check array", len(users))
	if len(users) < 1 {
		return false, nil
	} else {
		for _, us := range users {
			log.Println(us.Email, u.Email)
			if strings.EqualFold(us.Email, u.Email) {
				return true, nil
			}
		}

	}
	return false, nil
}
