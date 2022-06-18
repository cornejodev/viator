package log

import (
	"io"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	InfoLogger = log.New(logFile, "[INFO]: ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(logFile, "[WARNING]: ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "[ERROR]: ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLogger.Println("Starting the application...")
	InfoLogger.Println("Something noteworthy happened")
	WarningLogger.Println("There is something you should know about")
	ErrorLogger.Println("Something went wrong")
}
