package style

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/sebnyberg/mapboxcli/pkg/mapbox"
)

func formatStyleList(styles []mapbox.ListStyle, outputFormat string) (string, error) {
	switch outputFormat {
	case "table":
		return StyleListToTable(styles)
	case "json":
		return StyleListToJson(styles)
	}
	panic("Unsupported output format..")
}

func formatStyle(style *mapbox.Style, outputFormat string) (string, error) {
	switch outputFormat {
	case "table":
		return StyleToTable(style)
	case "json":
		return StyleToJson(style)
	}
	panic("Unsupported output format..")
}

func StyleListToTable(styles []mapbox.ListStyle) (string, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"id", "name", "owner"})

	data := make([][]string, len(styles))

	for _, style := range styles {
		rowData := []string{
			style.Id,
			style.Name,
			style.Owner,
		}
		data = append(data, rowData)
	}
	table.AppendBulk(data)
	table.Render()

	return buf.String(), nil
}

func StyleListToJson(styles []mapbox.ListStyle) (string, error) {
	jsonStr, err := json.Marshal(styles)

	if err != nil {
		return "", fmt.Errorf("failed parse response: %v", err)
	}

	return fmt.Sprintln(string(jsonStr)), nil
}

func StyleToTable(style *mapbox.Style) (string, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"id", "name", "owner"})

	data := make([][]string, 0)

	rowData := []string{
		style.Id,
		style.Name,
		style.Owner,
	}
	data = append(data, rowData)

	table.AppendBulk(data)
	table.Render()

	return buf.String(), nil
}

func StyleToJson(style *mapbox.Style) (string, error) {
	jsonStr, err := json.Marshal(style)

	if err != nil {
		return "", fmt.Errorf("failed parse response: %v", err)
	}

	return fmt.Sprintln(string(jsonStr)), nil
}

func GetAll(accessToken string, username string, outputFormat string) (string, error) {
	if err := checkFormatAvailable(outputFormat); err != nil {
		return "", err
	}

	styles, err := mapbox.GetStyles(accessToken, username)
	if err != nil {
		return "", err
	}

	s, err := formatStyleList(styles, outputFormat)
	if err != nil {
		return "", err
	}

	return s, nil
}

func Get(accessToken string, username string, styleId string, outputFormat string) (string, error) {
	if err := checkFormatAvailable(outputFormat); err != nil {
		return "", err
	}

	style, err := mapbox.GetStyle(accessToken, username, styleId)
	if err != nil {
		return "", err
	}

	s, err := formatStyle(style, outputFormat)

	return s, err
}
