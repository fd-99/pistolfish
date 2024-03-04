package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func WriteToCSV(data []map[string]string, outputFile string, webURL string, Company string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Name", "Title", "OpenToWork", "Email1", "Email2", "Email3", "Email4"}
	writer.Write(header)

	for _, person := range data {
		name := strings.ReplaceAll(strings.ReplaceAll(person["Name"], ".", ""), " ", ".")

		var trimmedWebURL string
		if webURL != "" {
			trimmedWebURL = strings.TrimPrefix(webURL, "http://")
			trimmedWebURL = strings.TrimPrefix(trimmedWebURL, "https://")
			trimmedWebURL = strings.TrimPrefix(trimmedWebURL, "www.")
			trimmedWebURL = strings.TrimSuffix(trimmedWebURL, "/")
		}
		if webURL == "" {
			trimmedWebURL = Company + ".com"
		}

		// Remove special characters from the name
		name = strings.Map(func(r rune) rune {
			if r == '\'' || r < 32 || r > 126 {
				return -1 // Remove non-printable and non-ASCII characters
			}
			return r
		}, name)

		nameParts := strings.FieldsFunc(name, func(r rune) bool {
			return r == '.' || r == '(' || r == ')' || r == ' '
		})

		var emailNoDot, emailFirstNameLastName, emailFirstInitialLastName, emailAddress string
		if len(nameParts) == 1 {
			emailNoDot = fmt.Sprintf("%s@%s", strings.ToLower(nameParts[0]), trimmedWebURL)
			emailFirstNameLastName = emailNoDot
			emailFirstInitialLastName = emailNoDot
			emailAddress = emailNoDot
		} else if len(nameParts) >= 2 {
			// Use first name, last initial
			emailFirstNameLastName = fmt.Sprintf("%s@%s", strings.ToLower(nameParts[0]), trimmedWebURL)

			// Use first initial, last name
			emailFirstInitialLastName = fmt.Sprintf("%s.%s@%s", strings.ToLower(nameParts[0][:1]), strings.ToLower(nameParts[len(nameParts)-1]), trimmedWebURL)

			// Use full name without dot
			emailNoDot = fmt.Sprintf("%s@%s", strings.ToLower(strings.Join(nameParts, "")), trimmedWebURL)

			// Use first name, last name
			emailAddress = fmt.Sprintf("%s.%s@%s", strings.ToLower(nameParts[0]), strings.ToLower(nameParts[len(nameParts)-1]), trimmedWebURL)
		}

		row := []string{person["Name"], person["Title"], person["OpenToWork"], emailAddress, emailFirstNameLastName, emailFirstInitialLastName, emailNoDot}
		writer.Write(row)
	}

	return nil
}
