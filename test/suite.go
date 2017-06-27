package test

import (
	"database/sql"
	"errors"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/src-d/framework.v0/database"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	dbName string
	DB     *sql.DB
}

func (s *Suite) Setup() {
	s.dbName = fmt.Sprintf("db_%d", time.Now().UnixNano())
	db, err := database.Default()
	s.NoError(err)

	_, err = db.Exec("CREATE DATABASE " + s.dbName)
	s.NoError(err)
	s.NoError(db.Close())

	s.DB, err = database.Default(database.WithName(s.dbName))
	s.NoError(err)

	bytes, err := ioutil.ReadFile(schemaPath)
	s.NoError(err)

	_, err = s.DB.Exec(string(bytes))
	s.NoError(err)
}

func (s *Suite) TearDown() {
	s.NoError(s.DB.Close())

	db, err := database.Default()
	s.NoError(err)

	_, err = db.Exec("DROP DATABASE " + s.dbName)
	s.NoError(err)
	s.NoError(db.Close())
}

var (
	rootDir    string
	schemaPath string
)

func init() {
	// First look at possible vendor directories
	srcs := vendorDirectories()

	// And then GOPATH
	srcs = append(srcs, build.Default.SrcDirs()...)

	for _, src := range srcs {
		rf := filepath.Join(src, "srcd.works", "core.v0")

		if _, err := os.Stat(rf); err == nil {
			rootDir = rf
			schemaPath = filepath.Join(rootDir, "schema", "schema.sql")
			return
		}
	}

	panic(errors.New("core.v0 directory not found"))
}

func vendorDirectories() []string {
	dir, err := os.Getwd()
	if err != nil {
		return nil
	}

	var dirs []string

	for {
		if dir == "." || dir == "/" {
			break
		}

		dirs = append(dirs, filepath.Join(dir, "vendor"))
		dir = filepath.Dir(dir)
	}

	return dirs
}
