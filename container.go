package core

import (
	"database/sql"

	"srcd.works/core.v0/model"

	"srcd.works/framework.v0/configurable"
	"srcd.works/framework.v0/database"
	"srcd.works/go-billy.v1"
	"srcd.works/go-billy.v1/osfs"
)

type configType struct {
	configurable.BasicConfiguration
	TempDir string `default:"/tmp/sourced"`
}

var config = &configType{}

func init() {
	configurable.InitConfig(config)
}

var container struct {
	Database             *sql.DB
	ModelRepositoryStore *model.RepositoryStore
	ModelMentionStore    *model.MentionStore
	TempDirFilesystem    billy.Filesystem
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

// ModelMentionStore returns the default *model.ModelMentionStore, using the
// default database. If it is not possible to connect to the database, this
// function will panic. Multiple calls will always return the same instance.
func ModelMentionStore() *model.MentionStore {
	if container.ModelMentionStore == nil {
		container.ModelMentionStore = model.NewMentionStore(Database())
	}

	return container.ModelMentionStore
}

// TemporaryFilesystem returns a billy.Filesystem that can be used to store
// temporary files. This directory might or might not be shared with other
// applications, so subdirectories should be used.
func TemporaryFilesystem() billy.Filesystem {
	if container.TempDirFilesystem == nil {
		container.TempDirFilesystem = osfs.New(config.TempDir)
	}

	return container.TempDirFilesystem
}
