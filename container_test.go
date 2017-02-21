package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBroker(t *testing.T) {
	require := require.New(t)
	b := Broker()
	require.NotNil(b)

	b2 := Broker()
	require.Exactly(b, b2)

	q, err := b.Queue("foo")
	require.NotNil(q)
	require.NoError(err)
	require.NoError(b.Close())
}

func TestDatabase(t *testing.T) {
	require := require.New(t)
	db := Database()
	require.NotNil(db)

	db2 := Database()
	require.Exactly(db, db2)
}

func TestModelRepositoryStore(t *testing.T) {
	require := require.New(t)
	s := ModelRepositoryStore()
	require.NotNil(s)

	s2 := ModelRepositoryStore()
	require.Exactly(s, s2)
}

func TestModelMentionStore(t *testing.T) {
	require := require.New(t)
	s := ModelMentionStore()
	require.NotNil(s)

	s2 := ModelMentionStore()
	require.Exactly(s, s2)
}

func TestTemporaryFilesystem(t *testing.T) {
	require := require.New(t)

	fs := TemporaryFilesystem()
	require.NotNil(fs)

	fs2 := TemporaryFilesystem()
	require.Exactly(fs, fs2)

	f, err := fs.TempFile("", "test")
	require.NoError(err)
	fPath := f.Filename()
	defer func() { require.NoError(fs.Remove(fPath)) }()
	require.NoError(f.Close())
}
