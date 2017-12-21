package db1c

import "path/filepath"

type DB struct {
	Url  string
	Name string
	dir  string
	md   *MD
}

func New(path string) (*DB, error) {
	md, err := newMD(filepath.Join(path, fileMD))
	if err != nil {
		return nil, err
	}
	db := &DB{
		dir: path,
		md:  md,
	}

	return db, nil
}

func (db *DB) Close() error {
	return nil
}
