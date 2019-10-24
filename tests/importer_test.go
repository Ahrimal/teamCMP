package test

import (
	"github.com/Ahrimal/teamCMP/fileReaders"
	"testing"
)

func TestImporter(t *testing.T) {

	// Test url.. this should give an error right now
	var feedImporter fileReaders.FeedImporter
	err := feedImporter.GetVideosData("https://urltest.com", "gogogo", "", "string", false)

	if err == nil {
		t.Errorf("GetVideosData(\"https://urltest.com\", ...) failed. Expected not implemented, got %v",
			err)
	} else {
		t.Logf("GetVideosData(\"https://urltest.com\", ...) success. Expected not implemented, got %v",
			err)
	}

	// Test no url wrong source
	err = feedImporter.GetVideosData("", "gogogo", "", "string", false)
	if err == nil {
		t.Errorf("GetVideosData(\"https://urltest.com\", ...) failed. Expected got no source files found, got %v",
			err)
	} else {
		t.Logf("GetVideosData(\"https://urltest.com\", ...) success. Expected got no source files found, got %v",
			err)
	}

	// Test no url correct source
	err = feedImporter.GetVideosData("", "glorf", "", "string", false)
	if err != nil {
		t.Errorf("GetVideosData(\"https://urltest.com\", ...) failed. nil, got %v",
			err)
	} else {
		t.Logf("GetVideosData(\"https://urltest.com\", ...) success. nil, got %v",
			err)
	}

	// Test no url correct source yaml
	err = feedImporter.GetVideosData("", "flub", "", "string", false)
	if err != nil {
		t.Errorf("GetVideosData(\"https://urltest.com\", ...) failed. nil, got %v",
			err)
	} else {
		t.Logf("GetVideosData(\"https://urltest.com\", ...) success. nil, got %v",
			err)
	}

	// Test no url correct all
	err = feedImporter.GetVideosData("", "", "", "string", true)
	if err != nil {
		t.Errorf("GetVideosData(..., true,) failed. nil, got %v",
			err)
	} else {
		t.Logf("GetVideosData(..., true,) success. nil, got %v",
			err)
	}
}
