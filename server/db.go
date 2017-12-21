package server

import (
	"path/filepath"
	"github.com/AlkBur/web1c/server/md"
	"errors"
)

const fileMD  = "1cv7.md"

var db map[string]*DB

type DB struct {
	dir  string
	md *md.MD
}

func init() {
	db = make(map[string]*DB)
}

func (s *Server) AddDB(name, path, url string) error {
	var err error
	_, ok := db[name]
	if !ok {
		db[name], err = NewDB(path)
		if err != nil {
			return err
		}
	}
	return errors.New("Base already exists")
}

func NewDB(path string) (*DB, error) {
	md, err := md.New(filepath.Join(path, fileMD))
	if err != nil {
		return nil, err
	}
	db := &DB{
		dir: path,
		md: md,
	}
	return db, nil
}

