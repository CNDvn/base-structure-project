package utils

import (
	"reflect"
	"time"
)

func SetDefaultInsert(obj interface{}) {
	createdAtField := reflect.ValueOf(obj).Elem().FieldByName("CreatedAt")
	if createdAtField.CanSet() && createdAtField.Type() == reflect.TypeOf(time.Now()) {
		createdAtField.Set(reflect.ValueOf(time.Now()))
	}

	updatedAtField := reflect.ValueOf(obj).Elem().FieldByName("UpdatedAt")
	if updatedAtField.CanSet() && updatedAtField.Type() == reflect.TypeOf(time.Now()) {
		updatedAtField.Set(reflect.ValueOf(time.Now()))
	}
}

func SetDefaultUpdate(obj interface{}) {
	updatedAtField := reflect.ValueOf(obj).Elem().FieldByName("UpdatedAt")
	if updatedAtField.CanSet() && updatedAtField.Type() == reflect.TypeOf(time.Now()) {
		updatedAtField.Set(reflect.ValueOf(time.Now()))
	}
}
