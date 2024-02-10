package mock

import (
	"fmt"
	"testing"
)

var fmtMock SampleInterface = &SampleInterfaceImpl{}

func Sprintf(format string, args ...interface{}) string {
	return fmtMock.Sprintf(format, args...)
}

//////////////////////////////////////////////////////

type SampleInterface interface {
	Sprintf(format string, a ...any) string
}

// ////////////////////////////////////////////////////
type SampleInterfaceImpl struct{}

func (m *SampleInterfaceImpl) Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

type SampleInterfaceImplMock struct{}

func (m *SampleInterfaceImplMock) Sprintf(format string, a ...any) string {
	return "mocked Sprintf"
}

//////////////////////////////////////////////////////

func TestSprintf(t *testing.T) {
	// setup function mocks
	fmtMock = &SampleInterfaceImpl{}

	// inside functionThatUsesMultileGlobalFunctions: fmt.Sprintf is mocked
	got := Sprintf("-%d-", 50)
	want := "-50-"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
func TestSprintfMock(t *testing.T) {
	// setup function mocks
	fmtMock = &SampleInterfaceImplMock{}

	// inside functionThatUsesMultileGlobalFunctions: fmt.Sprintf is mocked
	got := Sprintf("-%d-", 50)
	want := "mocked Sprintf"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

//////////////////////////////////////////////////////

var mockf = fmt.Sprintf

func SprintfFunc(format string, args ...interface{}) string {
	return mockf(format, args...)
}

func TestSprintfFunc(t *testing.T) {
	// inside functionThatUsesMultileGlobalFunctions: fmt.Sprintf is mocked
	got := SprintfFunc("-%d-", 50)
	want := "-50-"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSprintfFuncMock(t *testing.T) {

	mockf = func(format string, a ...any) string {
		return "mocked Sprintf"
	}

	// inside functionThatUsesMultileGlobalFunctions: fmt.Sprintf is mocked
	got := SprintfFunc("-%d-", 50)
	want := "mocked Sprintf"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
