package main

import (
	"flag"
	"log"
	"os"

	"github.com/Binject/bintriage/bt"
)

func main() {

	var srcFile = flag.String("inFile", "a.out", "Input file to triage")
	flag.StringVar(srcFile, "i", "a.out", "Input file to triage")
	flag.StringVar(srcFile, "input", "a.out", "Input file to triage")

	var logFile = flag.String("logFile", "", "send stdout to a log file")
	flag.StringVar(logFile, "l", "", "send stdout to a log file")

	flag.Parse()

	if *logFile != "" {
		logTown, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer logTown.Close()
		log.SetOutput(logTown)
		log.Println("Log file started!")
	}

	err := bt.BinTriage(*srcFile)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Done!")
	}
}
