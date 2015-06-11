package model

import (
	"fmt"
	"time"
)

type Component struct {
	Version string `json:"version"`
	CSS string `json:"css"`
	JS string `json:"js"`
}

func ReadComponentVersions(componentName string) []Component {
	rows, err := conn.Query(
		`SELECT version FROM component
		WHERE name=?
		ORDER BY date_created`, componentName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	history := make([]Component, 0)
	for rows.Next() {
		var version string
		var dateCreated time.Time
		rows.Scan(&version, &dateCreated)
		history = append(history, Component{version, dateCreated})
	}
	return history
}
