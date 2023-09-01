package dao

import (
	"database/sql/driver"
	"encoding/json"
)

type JSON []map[string]string

func (c JSON) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JSON) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
