package header

import (
	"net/http"
	"testing"
)

func TestCopyValueWithSingleKey(t *testing.T) {
	headers := &http.Header{}
	headers.Add("source", "key1=value1")

	copyHeaderValue(headers, &Config{
		From:              "source",
		PairSeparator:     ";",
		KeyValueSeparator: "=",
		Key:               "key1",
		To:                "target",
		Overwrite:         false,
	})

	if headers.Get("target") != "value1" {
		t.Errorf("Does not contain header 'target' with value 'value1'.")
	}
}

func TestCopyValueWithMultipleKeys(t *testing.T) {
	headers := &http.Header{}
	headers.Add("source", "key1=value1; key2=value2")

	copyHeaderValue(headers, &Config{
		From:              "source",
		PairSeparator:     ";",
		KeyValueSeparator: "=",
		Key:               "key2",
		To:                "target",
		Overwrite:         false,
	})

	if headers.Get("target") != "value2" {
		t.Errorf("Does not contain header 'target' with value 'value2'")
	}
}

func TestPrefix(t *testing.T) {
	headers := &http.Header{}
	headers.Add("source", "key1=value1; key2=value2")

	copyHeaderValue(headers, &Config{
		From:              "source",
		PairSeparator:     ";",
		KeyValueSeparator: "=",
		Key:               "key2",
		To:                "target",
		Prefix:            "prefix ",
	})

	if headers.Get("target") != "prefix value2" {
		t.Errorf("header 'target' does not include prefix 'prefix'")
	}
}

func TestDoesNotContainTargetWhenNoKeyInValue(t *testing.T) {
	headers := &http.Header{}
	headers.Add("source", "key1=value1; key2=value2")

	copyHeaderValue(headers, &Config{
		From:              "source",
		PairSeparator:     ";",
		KeyValueSeparator: "=",
		Key:               "key3",
		To:                "target",
		Overwrite:         false,
	})

	if headers.Get("target") != "" {
		t.Errorf("Contains 'target' header.")
	}
}

func TestNotOverwriteTargetIfItExists(t *testing.T) {
	headers := &http.Header{}
	headers.Add("source", "key1=value1; key2=value2")
	headers.Add("target", "value3")

	copyHeaderValue(headers, &Config{
		From:              "source",
		PairSeparator:     ";",
		KeyValueSeparator: "=",
		Key:               "key2",
		To:                "target",
		Overwrite:         false,
	})

	if headers.Get("target") != "value3" {
		t.Errorf("'Target' has been overwritten.")
	}
}

func TestOverwriteTargetIfItExists(t *testing.T) {
	headers := &http.Header{}
	headers.Add("source", "key1=value1; key2=value2")
	headers.Add("target", "value3")

	copyHeaderValue(headers, &Config{
		From:              "source",
		PairSeparator:     ";",
		KeyValueSeparator: "=",
		Key:               "key2",
		To:                "target",
		Overwrite:         true,
	})

	if headers.Get("target") != "value2" {
		t.Errorf("'Target' has not been overwritten.")
	}
}
