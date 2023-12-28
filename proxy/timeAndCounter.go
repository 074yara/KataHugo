package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

var indexFilePath = "/app/static/tasks/_index.md"

func TimeAndCounterWorker() {
	var timeField, counterField []byte
	var counter int
	data, err := os.ReadFile(indexFilePath)
	checkError(err)
	newData := make([]byte, len(data))
	fields := bytes.Split(data, []byte("\n"))
	for _, field := range fields {
		if bytes.Contains(field, []byte("Текущее время:")) {
			timeField = field
		}
		if bytes.Contains(field, []byte("Счетчик:")) {
			counterField = field
		}
	}
	for {
		time.Sleep(time.Second * 5)
		counter++
		currTime := time.Now().Format("2006-01-02 15-04-05")
		newTimeField := []byte(fmt.Sprintf("Текущее время: %v", currTime))
		newCounterField := []byte(fmt.Sprintf("Счетчик: %v", counter))
		newData = bytes.Replace(data, timeField, newTimeField, 1)
		newData = bytes.Replace(newData, counterField, newCounterField, 1)
		err = os.WriteFile(indexFilePath, newData, 0644)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
