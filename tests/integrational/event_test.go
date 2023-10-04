package integrational

import (
	"context"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_Event(t *testing.T) {
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

	sId := "20081UC-023"
	sensor := &entity.Sensor{
		Id:              sId,
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

	//

	ts1, err := srvc.GetSensorEvents(client, ctx, sId)
	assert.NoError(t, err)
	assert.Empty(t, ts1)

	left := time.Now()
	tm := left.Add(5 * time.Minute)
	right := left.Add(10 * time.Minute)
	ts1, err = srvc.GetEventsBySignalTime(client, ctx, left, right)
	assert.NoError(t, err)
	assert.Empty(t, ts1)

	event := &entity.CreateEvent{
		SignalTime:   tm.Format("2006-01-02 15:04:05"),
		SensorId:     "22101BL013",
		PeakReadings: 21.1,
	}
	id, err := srvc.CreateEvent(client, ctx, event)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	crEvent := &entity.Event{
		Id:           id,
		SignalTime:   tm,
		SensorId:     event.SensorId,
		PeakReadings: event.PeakReadings,
	}
	ts2, err := srvc.GetSensorEvents(client, ctx, sId)
	assert.NoError(t, err)
	assert.Equal(t, crEvent, ts2[0])

	ts3, err := srvc.GetEventsBySignalTime(client, ctx, left, right)
	assert.NoError(t, err)
	assert.Equal(t, crEvent, ts3[0])

	err = srvc.DeleteEvent(client, ctx, id)
	assert.NoError(t, err)

	err = srvc.DeleteEvent(client, ctx, id)
	assert.Error(t, err)

	ts4, err := srvc.GetSensorEvents(client, ctx, sId)
	assert.NoError(t, err)
	assert.Empty(t, ts4)

	ts5, err := srvc.GetEventsBySignalTime(client, ctx, left, right)
	assert.NoError(t, err)
	assert.Empty(t, ts5)
}
