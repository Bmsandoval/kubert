package hello_test

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

var tz, _ = time.LoadLocation("UTC")

var presentTime time.Time
var presentUnixTime time.Time
var presentTimeStamp timestamp.Timestamp

var pastTime time.Time
var pastUnixTime time.Time
var pastTimeStamp timestamp.Timestamp

var futureTime time.Time
var futureUnixTime time.Time
var futureTimeStamp timestamp.Timestamp

func setTimeStampsForTests() {
	presentTime = time.Now().Round(time.Second)
	presentUnixTime = time.Unix(presentTime.Unix(), 0).In(tz)
	presentTimeStamp = timestamp.Timestamp{Seconds: presentTime.Unix()}

	pastTime = presentTime.Add(-(time.Minute * 10))
	pastUnixTime = time.Unix(pastTime.Unix(), 0).In(tz)
	pastTimeStamp = timestamp.Timestamp{Seconds: pastTime.Unix()}

	futureTime = presentTime.Add(time.Minute * 10)
	futureUnixTime = time.Unix(futureTime.Unix(), 0).In(tz)
	futureTimeStamp = timestamp.Timestamp{Seconds: futureTime.Unix()}
}

func stringPointer(sample string) *string {
	return &sample
}

type BaseTestData struct {
	Description string
	Request interface{}
	Response interface{}
}