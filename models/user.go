package models

import (
	"fmt"
	"github.com/containerops/wrench/db"
	"regexp"
)

type User struct {
	Id                string   `json:"id"`                //
	Username          string   `json:"username"`          //
	Password          string   `json:"password"`          //
	Email             string   `json:"email"`             //
	Fullname          string   `json:"fullname"`          //
	Company           string   `json:"company"`           //
	Location          string   `json:"location"`          //
	Mobile            string   `json:"mobile"`            //
	URL               string   `json:"url"`               //
	Gravatar          string   `json:"gravatar"`          //
	Created           int64    `json:"created"`           //
	Updated           int64    `json:"updated"`           //
	Repositories      []string `json:"repositories"`      //
	Organizations     []string `json:"organizations"`     //
	Teams             []string `json:"teams"`             //
	JoinOrganizations []string `json:"joinorganizations"` //
	JoinTeams         []string `json:"jointeams"`         //
	Starts            []string `json:"starts"`            //
	Comments          []string `json:"comments"`          //
	Memo              []string `json:"memo"`              //
}

func (user *User) Has(username string) (bool, string, error) {
	UUID, err := db.GetUUID("user", username)
	if err != nil {
		return false, "", err
	}

	if len(UUID) <= 0 {
		return false, "", nil
	}

	err = db.Get(user, UUID)
	return true, UUID, err
}

func (user *User) Save() error {
	validNamespace := regexp.MustCompile(`^([a-z0-9_]{4,30})$`)
	if !validNamespace.MatchString(user.Username) {
		return fmt.Errorf("Username must be 4 - 30, include a-z, 0-9 and '_'")
	}

	if len(user.Password) < 5 {
		return fmt.Errorf("Password length should be more than 5")
	}

	validEmail := regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")
	if !validEmail.MatchString(user.Email) {
		return fmt.Errorf("Email illegal")
	}

	if err := db.Save(user, user.Id); err != nil {
		return err
	}

	if _, err := db.Client.HSet(db.GLOBAL_USER_INDEX, user.Username, user.Id).Result(); err != nil {
		return err
	}

	return nil
}
