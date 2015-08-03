package modules

import (
	"fmt"

	"github.com/containerops/crew/models"
	"github.com/containerops/wrench/db"
)

func SaveUser(username, passwd, email string) (string, error) {
	return "", nil
}

func GetUser(username, passwd string) (models.User, error) {
	user := new(models.User)
	if exist, UUID, err := user.Has(username); err != nil {
		return models.User{}, err
	} else if exist == false && err == nil {
		return models.User{}, fmt.Errorf("User is not exist: %s", username)
	} else if exist == true && err == nil {
		if err := db.Get(user, UUID); err != nil {
			return models.User{}, err
		} else {
			if user.Password != passwd {
				return models.User{}, fmt.Errorf("User password error.")
			} else {
				return models.User{}, nil
			}
		}
	}

	return models.User{}, nil
}
