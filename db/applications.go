package db

import (
	"encoding/json"
	"time"
)

// Application holds all the information about an application
type Application struct {
	RowID          string
	AppData        map[string]json.RawMessage
	EmploymentData map[string]json.RawMessage
	CreatedTime    time.Time
	UpdatedTime    time.Time
}

func AddApplication(
	rowID string,
	appData map[string]json.RawMessage,
	employmentData map[string]json.RawMessage,
) error {
	appDataJSON, err := json.Marshal(appData)
	if err != nil {
		return err
	}

	employmentDataJSON, err := json.Marshal(employmentData)
	if err != nil {
		return err
	}

	queryStr := `insert into applications (row_id, app_data, employment_data) values (?, ?, ?)` +
		` on duplicate key update app_data = ?, employment_data = ?`
	_, err = zestDB.Exec(queryStr, rowID, appDataJSON, employmentDataJSON, appDataJSON, employmentDataJSON)

	return err
}
