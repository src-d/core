package core

import (
	"database/sql"
	"io/ioutil"
	"os"

	"srcd.works/core.v0/model"

	"srcd.works/framework.v0/configurable"
	"srcd.works/framework.v0/database"
	"srcd.works/framework.v0/queue"
	"srcd.works/go-billy.v1"
	"srcd.works/go-billy.v1/osfs"
)

type containerConfig struct {
	configurable.BasicConfiguration
	TempDir string `default:"/tmp/sourced"`
	Broker  string `default:"amqp://localhost:5672"`
}

var config = &containerConfig{}

func init() {
	configurable.InitConfig(config)
}

var container struct {
	Broker               queue.Broker
	Database             *sql.DB
	ModelRepositoryStore *model.RepositoryStore
	TempDirFilesystem    billy.Filesystem
}

// Broker returns a queue.Broker for the default queue system.
func Broker() queue.Broker {
	if container.Broker == nil {
		b, err := queue.NewBroker(config.Broker)
		if err != nil {
			panic(err)
		}

		container.Broker = b
	}

	return container.Broker
}

// Database returns a sql.DB for the default database. If it is not possible to
// connect to the database, this function will panic. Multiple calls will always
// return the same instance.
func Database() *sql.DB {
	if container.Database == nil {
		container.Database = database.Must(database.Default())
	}

	return container.Database
}

// ModelRepositoryStore returns the default *model.RepositoryStore, using the
// default database. If it is not possible to connect to the database, this
// function will panic. Multiple calls will always return the same instance.
func ModelRepositoryStore() *model.RepositoryStore {
	if container.ModelRepositoryStore == nil {
		container.ModelRepositoryStore = model.NewRepositoryStore(Database())
	}

	return container.ModelRepositoryStore
}

// TemporaryFilesystem returns a billy.Filesystem that can be used to store
// temporary files. This directory is dedicated to the running application.
func TemporaryFilesystem() billy.Filesystem {
	if container.TempDirFilesystem == nil {
		if err := os.MkdirAll(config.TempDir, os.FileMode(0755)); err != nil {
			panic(err)
		}

		dir, err := ioutil.TempDir(config.TempDir, "")
		if err != nil {
			panic(err)
		}

		container.TempDirFilesystem = osfs.New(dir)
	}

	return container.TempDirFilesystem
}
