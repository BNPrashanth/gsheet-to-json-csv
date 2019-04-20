package main

import (
	s "gsheet-to-json-csv/src/services"
	u "gsheet-to-json-csv/src/utils"
	"os"
)

func main() {
	// Strting the Application
	u.GeneralLogger.Println("Starting Extracting Language Files from GoogleSheet - downloading csv approach..")

	csvFilePath := "../outputs/gsheet.csv"
	errorResponse := s.Download(
		"https://docs.google.com/spreadsheets/d/e/2PACX-1vQIhLNNfUKVjxMkMwdtTFnvuV8oN1H_OmgOWRCwHBkSfOo1fzA08LXDfcK4EA86fx18M4FeAIwOoBBR/pub?output=csv",
		csvFilePath,
		5000,
	)
	if errorResponse.Err != nil {
		u.ErrorLogger.Println(errorResponse.Message, errorResponse.Err)
		os.Exit(1)
	}
	s.WriteLanguageFiles(csvFilePath)
	u.GeneralLogger.Println("Completed Execution..")
}
