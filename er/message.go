package er

var messages = map[string]string{
	"1": "Oops! Something went wrong. Please try later",
	"2": "Invalid Application ID or token",
	"3": "CSV file data format not proper",
	"4": "error not nil while inserting data to db",
}

var codes = map[Code]string{
	UncaughtException:          "1",
	InvalidAppToken:            "2",
	UserNotPresent:             "2",
	CSVFileDataFormatNotProper: "3",
	ErrorInsertingToDatabase:   "4",
}
