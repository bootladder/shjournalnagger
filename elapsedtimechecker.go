package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//NaggingIntervalTracker bool
type NaggingIntervalTracker interface {
	isNaggingIntervalExpired(interval int) bool
	updateLastNaggingTime()
}

//CurrentTimeGetter gets the time for testing purposes
type CurrentTimeGetter interface {
	Now() time.Time
}

//ElapsedTimeChecker  checks if the elapsed time has passed
type ElapsedTimeChecker struct {
	lastNaggingTimeFileReader io.Reader
	lastNaggingTimeFileWriter io.Writer
	currentTimeGetter         CurrentTimeGetter
}

func (e *ElapsedTimeChecker) getElapsedSeconds() int {
	nowUnixSeconds := e.currentTimeGetter.Now().Unix()
	b, _ := ioutil.ReadAll(e.lastNaggingTimeFileReader)
	lastNaggingTimeUnixSeconds, _ := strconv.Atoi(string(b))

	return int(nowUnixSeconds) - lastNaggingTimeUnixSeconds
}

func (e *ElapsedTimeChecker) isNaggingIntervalExpired(interval int) bool {
	return e.getElapsedSeconds() > interval
}

func (e *ElapsedTimeChecker) updateLastNaggingTime() {
	timeStr := fmt.Sprintf("%d", e.currentTimeGetter.Now().Unix())
	e.lastNaggingTimeFileWriter.Write([]byte(timeStr))
}

//RealLastNaggingTimeFileReader blah terrible name
type RealLastNaggingTimeFileReader struct{}

var lastNaggingTimeFilename = "/home/steve/.shjournalnagger/lastnaggingtime.txt"

func (r *RealLastNaggingTimeFileReader) Read(b []byte) (int, error) {

	// if file not exist, create
	if _, err := os.Stat(lastNaggingTimeFilename); os.IsNotExist(err) {
		fmt.Printf("NO LAST NAGGING TIME FILE, CREATING: %v", err)
		err = os.MkdirAll("/home/steve/.shjournalnagger", 0777)
		if err != nil {
			fmt.Printf("WRITING ERROR: %v", err)
		}

		err = ioutil.WriteFile(lastNaggingTimeFilename, []byte("1"), 0777)
		if err != nil {
			fmt.Printf("WRITING ERROR: %v", err)
		}

	}

	f, err := os.Open(lastNaggingTimeFilename)
	if err != nil {
		return 0, err
	}

	reader := bufio.NewReader(f)
	buf, _ := ioutil.ReadAll(reader)
	copy(b, buf)
	return len(buf), io.EOF
}

//RealLastNaggingTimeFileWriter blah terrible name
type RealLastNaggingTimeFileWriter struct{}

func (w *RealLastNaggingTimeFileWriter) Write(b []byte) (int, error) {

	err := ioutil.WriteFile(lastNaggingTimeFilename, b, 0777)
	return len(b), err
}

//RealCurrentTimeGetter blah terrible name
type RealCurrentTimeGetter struct{}

//Now blah
func (r *RealCurrentTimeGetter) Now() time.Time {
	return time.Now()
}
