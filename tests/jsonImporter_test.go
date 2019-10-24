package test

import (
	"github.com/Ahrimal/teamCMP/fileReaders"
	"io/ioutil"
	"log"
	"testing"
)

func TestJsonImport(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	var jsonImporter fileReaders.JsonImporter

	// Test from an empty name file
	videos, err := jsonImporter.GetVideosData("")

	if videos != nil {
		t.Errorf("GetVideosData(\"\") failed. Expected %v, got %v", "nil", videos)
	} else {
		t.Logf("GetVideosData(\"\") success. Expected <%v>, got %v", "nil", videos)
	}

	if err == nil {
		t.Errorf("GetVideosData(\"\") failed. Expected %v, got %v", "error", "nil")
	} else {
		t.Logf("GetVideosData(\"\") success. Expected <%v>, got <%v>", "open : The system cannot find the file specified.", err)
	}

	// Test for un-existent file
	videos, err = jsonImporter.GetVideosData("doesNotExist.json")

	if videos != nil {
		t.Errorf("GetVideosData(\"doesNotExist.json\") failed. Expected %v, got %v", "nil", videos)
	} else {
		t.Logf("GetVideosData(\"doesNotExist.json\") success. Expected <%v>, got %v", "nil", videos)
	}

	if err == nil {
		t.Errorf("GetVideosData(\"doesNotExist.json\") failed. Expected %v, got %v", "error", "nil")
	} else {
		t.Logf("GetVideosData(\"doesNotExist.json\") success. Expected <%v>, got <%v>",
			"open : The system cannot find the file specified.", err)
	}

	// Test for invalid file
	videos, err = jsonImporter.GetVideosData("/feed-exports/flub.yaml")

	if videos != nil {
		t.Errorf("GetVideosData(\"feed-exports/flub.yaml\") failed. Expected %v, got %v", "nil", videos)
	} else {
		t.Logf("GetVideosData(\"feed-exports/flub.yaml\") success. Expected <%v>, got %v", "nil", videos)
	}

	if err == nil {
		t.Errorf("GetVideosData(\"feed-exports/flub.yaml\") failed. Expected %v, got %v", "error", "nil")
	} else {
		t.Logf("GetVideosData(\"feed-exports/flub.yaml\") success. Expected <%v>, got <%v>",
			"yaml: unmarshal errors: line 1: cannot unmarshal !!map into []fileReaders.YamlVideoData.", err)
	}

	// Test for valid file - Note that this test should test that all the info is loaded correctly, which is not doing
	videos, err = jsonImporter.GetVideosData("feed-exports/glorf.json")

	if err != nil {
		t.Errorf("GetVideosData(\"feed-exports/glorf.json\") failed. Expected %v, got %v", "error", "nil")
	} else {
		t.Logf("GetVideosData(\"feed-exports/glorf.json\") success. Expected <%v>, got %v",
			"nil", err)
	}

	if videos == nil {
		t.Errorf("GetVideosData(\"feed-exports/glorf.json\") failed. Expected a length of %v, got %v", 2, videos)
	} else {
		t.Logf("GetVideosData(\"feed-exports/glorf.json\") success. Expected a length of %v, got %v", 2, len(videos.Videos))
		for _, video := range videos.Videos {
			if video.Title == "" {
				t.Errorf("GetVideosData(\"feed-exports/glorf.json\") failed. Expected title value different than empty, got %v", video.Title)
			} else {
				t.Logf("GetVideosData(\"feed-exports/glorf.json\") success. Expected a value of title with a name, got %v", video.Title)
			}

			if video.Url == "" {
				t.Errorf("GetVideosData(\"feed-exports/glorf.json\") failed. Expected url value different than empty, got %v", video.Title)
			} else {
				t.Logf("GetVideosData(\"feed-exports/glorf.json\") success. Expected a value of url, got %v", video.Title)
			}

			// TODO: Pending tags test
		}
	}
}
