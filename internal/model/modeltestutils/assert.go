package modeltestutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/internal/model"
)

func AssertUsersEqual(t *testing.T, a model.User, b model.User, msg string) {
	assert.Equal(t, a.ID(), b.ID(), "%s (IDs are different)", msg)
	assert.Equal(t, a.Username(), b.Username(), "%s: (Usernames are different)", msg)
	assert.Equal(t, a.Email(), b.Email(), "%s (Emails are different)", msg)
	assert.Equal(t, a.Aliases(), b.Aliases(), "%s (Aliases are different)", msg)
}

func AssertUsersEqualIgnoreID(t *testing.T, a model.User, b model.User, msg string) {
	assert.Equal(t, a.Username(), b.Username(), "%s: (Usernames are different)", msg)
	assert.Equal(t, a.Email(), b.Email(), "%s (Emails are different)", msg)
	assert.Equal(t, a.Aliases(), b.Aliases(), "%s (Aliases are different)", msg)
}
