package integrational

import (
	"context"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Sensor(t *testing.T) {
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

	anId := "20081UC-061"
	analyzer := &entity.Analyzer{
		Id:              anId,
		Type:            anType,
		PartNumber:      "VTS-L0001100709",
		JobNumber:       "20091U",
		SoftwareVersion: "10.11.01",
	}
	err = srvc.CreateAnalyzer(client, ctx, analyzer)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	anId2 := "20081UC-023"
	analyzer2 := &entity.Analyzer{
		Id:              anId2,
		Type:            anType,
		PartNumber:      "VTS-L0001100709",
		JobNumber:       "20091U",
		SoftwareVersion: "10.11.01",
	}
	err = srvc.CreateAnalyzer(client, ctx, analyzer2)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	//

	id := "20081UC-023"
	sensor := &entity.Sensor{
		Id:              id,
		Type:            "электрохимический",
		AnalyzerId:      anId,
		Gas:             gName,
		LowLimitAlarm:   "10 % НКПВ",
		UpperLimitAlarm: "50 % НКПВ",
	}
	err = srvc.CreateSensor(client, ctx, sensor)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	ts1, err := srvc.GetAllSensors(client, ctx)
	assert.NoError(t, err)
	assert.Equal(t, sensor, ts1[0])

	ts2, err := srvc.GetAnalyzerSensors(client, ctx, anId)
	assert.NoError(t, err)
	assert.Equal(t, sensor, ts2[0])

	err = srvc.UpdateSensorAnalyzer(client, ctx, id, anId2)

	sensor.AnalyzerId = anId2
	ts2, err = srvc.GetAllSensors(client, ctx)
	assert.NoError(t, err)
	assert.Equal(t, sensor, ts2[0])

	ts2, err = srvc.GetAnalyzerSensors(client, ctx, anId)
	assert.NoError(t, err)
	assert.Empty(t, ts2)

	err = srvc.DeleteSensor(client, ctx, id)
	assert.NoError(t, err)

	err = srvc.DeleteSensor(client, ctx, id)
	assert.Error(t, err)

	ts3, err := srvc.GetAllAnalyzers(client, ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts3)

	err = srvc.UpdateSensorAnalyzer(client, ctx, id, anId)
	assert.Error(t, err)
}
