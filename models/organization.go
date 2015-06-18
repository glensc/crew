package models

import (
	"github.com/containerops/wrench/db"
)

type Organization struct {
	Id           string   `json:"id"`           //
	Name         string   `json:"name"`         //
	Username     string   `json:"username"`     //
	Description  string   `json:"description"`  //
	Repositories []string `json:"repositories"` //
	Teams        []string `json:"teams"`        //
	Created      int64    `json:"created"`      //
	Updated      int64    `json:"updated"`      //
	Memo         []string `json:"memo"`         //
}

func (organization *Organization) Has(organizationName string) (bool, string, error) {
	UUID, err := db.GetUUID("organization", organizationName)
	if err != nil {
		return false, "", err
	}
	if len(UUID) <= 0 {
		return false, "", nil
	}

	err = db.Get(organization, UUID)

	return true, UUID, err
}

func (organization *Organization) Save() error {
	if err := db.Save(organization, organization.Id); err != nil {
		return err
	}

	if _, err := db.Client.HSet(db.GLOBAL_ORGANIZATION_INDEX, organization.Name, organization.Id).Result(); err != nil {
		return err
	}

	return nil
}
