package integrational

import (
	"context"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Type(t *testing.T) {
	ctx := context.Background()
	err := TruncateTables(client, ctx)
	if err != nil {
		return
	}

	gName := "кислород"
	gas := &entity.Gas{
		Name:    gName,
		Formula: "H2O",
		Type:    "горючие газы",
	}
	err = srvc.CreateGas(client, ctx, gas)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	//

	name := "VENTIS MX4"
	gases := []string{gName}
	anType := &entity.CreateType{
		Name:       name,
		MaxSensors: 3,
		Gases:      gases,
	}
	err = srvc.CreateType(client, ctx, anType)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	crType := &entity.Type{
		Name:       name,
		MaxSensors: 3,
	}
	ts1, err := srvc.GetAllTypes(client, ctx)
	assert.NoError(t, err)
	assert.Equal(t, crType, ts1[0])

	err = srvc.DeleteType(client, ctx, name)
	assert.NoError(t, err)

	err = srvc.DeleteType(client, ctx, name)
	assert.Error(t, err)

	ts3, err := srvc.GetAllTypes(client, ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts3)
}
