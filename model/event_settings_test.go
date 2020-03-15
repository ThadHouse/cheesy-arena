// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventSettingsReadWrite(t *testing.T) {
	db := setupTestDb(t)

	eventSettings, err := db.GetEventSettings()
	assert.Nil(t, err)
	assert.Equal(t, EventSettings{Id: 0, Name: "Untitled Event", NumElimAlliances: 8, SelectionRound2Order: "L",
		SelectionRound3Order: "", TBADownloadEnabled: true, ApTeamChannel: 157, ApAdminChannel: 0,
		ApAdminWpaKey: "1234Five", WarmupDurationSec: 0, AutoDurationSec: 15, PauseDurationSec: 0,
		TeleopDurationSec: 135, Warning1RemainingDurationSec: 30, Warning2RemainingDurationSec: 20,
		HabDockingThreshold: 15}, *eventSettings)

	eventSettings.Name = "Chezy Champs"
	eventSettings.NumElimAlliances = 6
	eventSettings.SelectionRound2Order = "F"
	eventSettings.SelectionRound3Order = "L"
	err = db.SaveEventSettings(eventSettings)
	assert.Nil(t, err)
	eventSettings2, err := db.GetEventSettings()
	assert.Nil(t, err)
	assert.Equal(t, eventSettings, eventSettings2)
}
