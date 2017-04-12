package test

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"srcd.works/framework.v0/database"

	"github.com/stretchr/testify/suite"
)

var schemaFile = fmt.Sprintf("%s/src/srcd.works/core.v0/schema/schema.sql", os.Getenv("GOPATH"))

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

	bytes, err := ioutil.ReadFile(schemaFile)
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
