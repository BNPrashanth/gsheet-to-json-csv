package main

import (
	s "gsheet-to-json-csv/src/services"
	u "gsheet-to-json-csv/src/utils"
	"os"
)

func main() {
	// Strting the Application
	u.GeneralLogger.Println("Starting Extracting Language Files from GoogleSheet - downloading csv approach..")

	fileName := "../outputs/gsheet.csv"
	var errorResponse *u.ErrorResponse
	errorResponse = s.Download(
		"https://docs.google.com/spreadsheets/d/e/2PACX-1vQIhLNNfUKVjxMkMwdtTFnvuV8oN1H_OmgOWRCwHBkSfOo1fzA08LXDfcK4EA86fx18M4FeAIwOoBBR/pub?output=csv",
		fileName,
		5000,
	)
	if errorResponse.Err != nil {
		u.ErrorLogger.Println(errorResponse.Message, errorResponse.Err)
		os.Exit(1)
	}
	s.WriteLanguageFiles(fileName)
	u.GeneralLogger.Println("Completed Execution..")
}
