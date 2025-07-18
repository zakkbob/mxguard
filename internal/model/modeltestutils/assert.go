package modeltestutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/internal/model"
)

func AssertAliasesEqual(t *testing.T, a model.Alias, b model.Alias, msg string) {
	assert.Equal(t, a.Name(), b.Name(), "%s (Names are different)", msg)
	assert.Equal(t, a.Description(), b.Description(), "%s: (Descriptions are different)", msg)
	assert.Equal(t, a.Enabled(), b.Enabled(), "%s (Enabled isn't equal)", msg)
}

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
