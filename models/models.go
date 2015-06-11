package models

import (
	"encoding/json"
	"fmt"
	"github.com/containerops/wrench/db"
	"github.com/containerops/wrench/utils"
	"reflect"
	"strings"
)

func GetUUID(ObjectType, Object string) (UUID string, err error) {

	index := ""

	switch strings.TrimSpace(ObjectType) {

	case "user":
		index = db.GLOBAL_USER_INDEX
	case "repository":
		index = db.GLOBAL_REPOSITORY_INDEX
	case "organization":
		index = db.GLOBAL_ORGANIZATION_INDEX
	case "team":
		index = db.GLOBAL_TEAM_INDEX
	case "image":
		index = db.GLOBAL_IMAGE_INDEX
	case "tarsum":
		index = db.GLOBAL_TARSUM_INDEX
	case "tag":
		index = db.GLOBAL_TAG_INDEX
	case "compose":
		index = db.GLOBAL_COMPOSE_INDEX
	case "admin":
		index = db.GLOBAL_ADMIN_INDEX
	case "log":
		index = db.GLOBAL_LOG_INDEX
	default:

	}

	if UUID, err = db.Client.HGet(index, Object).Result(); err != nil {
		return "", err
	} else {
		return UUID, nil
	}
}

func Save(obj interface{}, key string) (err error) {
	s := reflect.TypeOf(obj).Elem()

	for i := 0; i < s.NumField(); i++ {

		value := reflect.ValueOf(obj).Elem().Field(s.Field(i).Index[0])

		switch value.Kind() {

		case reflect.String:
			if _, err := db.Client.HSet(key, s.Field(i).Name, value.String()).Result(); err != nil {
				return err
			}

		case reflect.Bool:
			if _, err := db.Client.HSet(key, s.Field(i).Name, utils.BoolToString(value.Bool())).Result(); err != nil {
				return err
			}

		case reflect.Int64:
			if _, err := db.Client.HSet(key, s.Field(i).Name, utils.Int64ToString(value.Int())).Result(); err != nil {
				return err
			}

		case reflect.Slice:
			if "[]string" == value.Type().String() && !value.IsNil() {
				strJson := "["

				for i := 0; i < value.Len(); i++ {
					nowUUID := value.Index(i).String()
					if i != 0 {
						strJson += fmt.Sprintf(`,"%s"`, nowUUID)
					} else {
						strJson += fmt.Sprintf(`"%s"`, nowUUID)
					}
				}

				strJson += "]"

				if db.Client.HSet(key, s.Field(i).Name, strJson).Result(); err != nil {
					return err
				}
			}

		default:
		}

	}
	return nil
}

func Get(obj interface{}, UUID string) (err error) {

	nowTypeElem := reflect.ValueOf(obj).Elem()
	types := nowTypeElem.Type()

	for i := 0; i < nowTypeElem.NumField(); i++ {
		nowField := nowTypeElem.Field(i)
		nowFieldName := types.Field(i).Name

		switch nowField.Kind() {

		case reflect.String:
			nowValue, err := db.Client.HGet(UUID, nowFieldName).Result()
			nowField.SetString(nowValue)
			if err != nil {
				return err
			}

		case reflect.Bool:
			nowValue, err := db.Client.HGet(UUID, nowFieldName).Result()
			nowField.SetBool(utils.StringToBool(nowValue))
			if err != nil {
				return err
			}

		case reflect.Int64:
			nowValue, err := db.Client.HGet(UUID, nowFieldName).Result()
			nowField.SetInt(utils.StringToInt64(nowValue))
			if err != nil {
				return err
			}

		case reflect.Slice:
			if "[]string" == nowField.Type().String() {
				nowValue, err := db.Client.HGet(UUID, nowFieldName).Result()

				var stringSlice []string
				err = json.Unmarshal([]byte(nowValue), &stringSlice)

				//TBD : code as below just for testing,it will be updated later
				if err != nil && (len(nowValue) > 0) {
					//return err
				} else {
					sliceValue := reflect.ValueOf(stringSlice)
					nowField.Set(sliceValue)
				}
			}

		default:
		}
	}

	return nil
}
