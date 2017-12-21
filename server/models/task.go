package models

import (
	"github.com/AlkBur/web1c/server/db1c"
)

type Task struct {
	ID  int    `json:"id"`
	Cmd string `json:"cmd"`
}

func GetTasks(db *db1c.DB) interface{} {
	return ""
}

func PostTasks(db *db1c.DB) interface{} {
	return ""
}
