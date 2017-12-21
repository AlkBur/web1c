package server

import (
	"errors"
	"fmt"
	"github.com/AlkBur/web1c/server/db1c"
	"github.com/AlkBur/web1c/server/handlers"
)

var dbs map[string]*db1c.DB

func init() {
	dbs = make(map[string]*db1c.DB)
}

func (s *Server) AddDB(name, path, url string) error {
	var err error
	_, ok := dbs[name]
	if !ok {
		dbs[name], err = NewDB(path)
		if err != nil {
			return err
		}
		dbs[name].Url = url
		dbs[name].Name = name

		htmlVal := make(map[string]interface{})
		htmlVal["name"] = name

		s.router.GET(url, handlers.BaseIndex(htmlVal))
		s.router.GET(url+"/config", handlers.ConfigIndex(htmlVal))
		s.router.GET(url+"/api", handlers.GetTasks(dbs[name]))
		s.router.POST(url+"/api", handlers.PostTasks(dbs[name]))

		return nil
	}
	return errors.New("Base already exists")
}

func NewDB(path string) (*db1c.DB, error) {
	return db1c.New(path)
}

func CloseAllDB() error {
	e := make([]error, 0, len(dbs))
	for _, db := range dbs {
		err := db.Close()
		if err != nil {
			e = append(e, err)
		}
	}
	if len(e) > 0 {
		str := ""
		for _, v := range e {
			if str != "" {
				str += "\r\n"
			}
			str += v.Error()
		}
		return fmt.Errorf("Errors: %v", str)
	}
	return nil
}
