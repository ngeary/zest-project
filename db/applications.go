package db

import (
	"encoding/json"
)

// AddApplication adds an application to the database
func AddApplication(
	rowID string,
	id string,
	memberID string,
	firstName string,
	lastName string,
	dob string,
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

	queryStr := `insert into applications (row_id, id, member_id, first_name, last_name, dob, app_data, employment_data)` +
		` values (?, ?, ?, ?, ?, ?, ?, ?)` +
		` on duplicate key update id = ?, member_id = ?, first_name = ?, last_name = ?, dob = ?, app_data = ?, employment_data = ?`
	_, err = zestDB.Exec(queryStr, rowID, id, memberID, firstName, lastName, dob, appDataJSON, employmentDataJSON,
		id, memberID, firstName, lastName, dob, appDataJSON, employmentDataJSON)

	return err
}
