package integrational

import (
	"context"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Gas(t *testing.T) {
	ctx := context.Background()
	err := TruncateTables(client, ctx)
	if err != nil {
		return
	}

	ts0, err := srvc.GetAllGases(client, ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts0)

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

	anType := "VENTIS MX4"
	gases := []string{gName}
	tp := &entity.CreateType{
		Name:       anType,
		MaxSensors: 3,
		Gases:      gases,
	}
	err = srvc.CreateType(client, ctx, tp)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	ts1, err := srvc.GetTypeGases(client, ctx, anType)
	assert.NoError(t, err)
	assert.Equal(t, gas, ts1[0])

	ts2, err := srvc.GetAllGases(client, ctx)
	assert.NoError(t, err)
	assert.Equal(t, gas, ts2[0])
}
