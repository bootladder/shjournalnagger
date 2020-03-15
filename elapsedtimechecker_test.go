package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockCurrentTimeGetter struct {
	mockTime time.Time
}

func (m *MockCurrentTimeGetter) Now() time.Time {
	return m.mockTime
}

func Test_LessThanTimeLimit_ReturnsFalse(t *testing.T) {

	lastNaggingTimeReader := bytes.NewReader([]byte("0000003000"))

	mockCurrentTimeGetter := MockCurrentTimeGetter{}
	mockCurrentTimeGetter.mockTime = time.Unix(3010, 0) //seconds, nanoseconds

	elapsedTimeChecker := ElapsedTimeChecker{lastNaggingTimeReader, nil, &mockCurrentTimeGetter}
	yesno := elapsedTimeChecker.isNaggingIntervalExpired(300)
	assert.False(t, yesno)
}

func Test_GreaterThanTimeLimit_ReturnsTrue(t *testing.T) {

	lastNaggingTimeReader := bytes.NewReader([]byte("0000003000"))

	mockCurrentTimeGetter := MockCurrentTimeGetter{}
	mockCurrentTimeGetter.mockTime = time.Unix(10000, 0)

	elapsedTimeChecker := ElapsedTimeChecker{lastNaggingTimeReader, nil, &mockCurrentTimeGetter}
	yesno := elapsedTimeChecker.isNaggingIntervalExpired(300)
	assert.True(t, yesno)
}
