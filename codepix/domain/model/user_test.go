package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/thauska/imersao-fullstack-fullcycle/codepix/domain/model"
)

func TestModel_NewUser(t *testing.T) {
	name := "Thayana"
	email := "email@email.com"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser("", "")
	require.NotNil(t, err)
}
