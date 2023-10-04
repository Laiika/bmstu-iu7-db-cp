package integrational

import (
	"context"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_User(t *testing.T) {
	ctx := context.Background()
	err := TruncateTables(client, ctx)
	if err != nil {
		return
	}

	name := "jfdjkfd"
	user := &entity.CreateUser{
		Name:     name,
		Password: "jfdjkdf",
		Role:     "user",
	}
	id, err := srvc.CreateUser(client, ctx, user)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	t1, err := srvc.GetUserById(client, ctx, id)
	assert.NoError(t, err)

	t2, err := srvc.GetUserByName(client, ctx, name)
	assert.NoError(t, err)
	assert.Equal(t, t1, t2)

	ts1, err := srvc.GetAllUsers(client, ctx)
	assert.NoError(t, err)
	assert.Equal(t, t1, ts1[0])

	err = srvc.UpdateUserRole(client, ctx, id, "employee")
	assert.NoError(t, err)

	t2.Role = "employee"
	t1, err = srvc.GetUserById(client, ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, t2, t1)

	err = srvc.DeleteUser(client, ctx, id)
	assert.NoError(t, err)

	err = srvc.DeleteUser(client, ctx, id)
	assert.Error(t, err)

	_, err = srvc.GetUserById(client, ctx, id)
	assert.Error(t, err)

	_, err = srvc.GetUserByName(client, ctx, name)
	assert.Error(t, err)

	ts2, err := srvc.GetAllUsers(client, ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts2)

	err = srvc.UpdateUserRole(client, ctx, id, "admin")
	assert.Error(t, err)
}
