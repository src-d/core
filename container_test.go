package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
