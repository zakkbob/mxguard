package db_test

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zakkbob/mxguard/db"
	"github.com/zakkbob/mxguard/internal/service"
	"fmt"
)

func TestCreatingAValidUserReturnsNoError(t *testing.T) {
	userRepo := db.NewPostgresUserRepository(conn)

	username := "test"
	isAdmin := false
	
	user, err := userRepo.CreateUser(context.Background(), service.CreateUserParams{
		Username: username,
		IsAdmin: isAdmin,
	})

	assert.Equal(t, username, user.Username, "Created User's username should match requested username")
	assert.Equal(t, username, user.Username, "Created User's admin status should match requested admin status")
	assert.NoError(t, err, "CreateUser should not return and error")
	fmt.Print(user)

}
