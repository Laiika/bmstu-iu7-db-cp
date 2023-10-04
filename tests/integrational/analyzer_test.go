package integrational

import (
	"context"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Analyzer(t *testing.T) {
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

	anType2 := "VENTIS MX6"
	tp2 := &entity.CreateType{
		Name:       anType2,
		MaxSensors: 5,
		Gases:      gases,
	}
	err = srvc.CreateType(client, ctx, tp2)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	//

	id := "20081UC-061"
	analyzer := &entity.Analyzer{
		Id:              id,
		Type:            anType2,
		PartNumber:      "VTS-L0001100709",
		JobNumber:       "20091U",
		SoftwareVersion: "10.11.01",
	}
	err = srvc.CreateAnalyzer(client, ctx, analyzer)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	t1, err := srvc.GetAnalyzerById(client, ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, analyzer, t1)

	ts1, err := srvc.GetAllAnalyzers(client, ctx)
	assert.NoError(t, err)
	assert.Equal(t, t1, ts1[0])

	ts2, err := srvc.GetTypeAnalyzers(client, ctx, anType)
	assert.NoError(t, err)
	assert.Empty(t, ts2)

	id2 := "20081UC-023"
	analyzer2 := &entity.Analyzer{
		Id:              id2,
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

	ts2, err = srvc.GetTypeAnalyzers(client, ctx, anType)
	assert.NoError(t, err)
	assert.Equal(t, analyzer2, ts2[0])

	err = srvc.DeleteAnalyzer(client, ctx, id)
	assert.NoError(t, err)

	err = srvc.DeleteAnalyzer(client, ctx, id)
	assert.Error(t, err)

	err = srvc.DeleteAnalyzer(client, ctx, id2)
	assert.NoError(t, err)

	_, err = srvc.GetAnalyzerById(client, ctx, id)
	assert.Error(t, err)

	ts3, err := srvc.GetAllAnalyzers(client, ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts3)
}
