package test

import (
	"io/ioutil"
	"log"
	"teamCMP/fileReaders"
	"testing"
)

func TestYamlImport(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	var yamlImporter fileReaders.YamlImporter

	// Test from an empty name file
	videos, err := yamlImporter.GetVideosData("")

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
	videos, err = yamlImporter.GetVideosData("doesNotExist.yaml")

	if videos != nil {
		t.Errorf("GetVideosData(\"doesNotExist.yaml\") failed. Expected %v, got %v", "nil", videos)
	} else {
		t.Logf("GetVideosData(\"doesNotExist.yaml\") success. Expected <%v>, got %v", "nil", videos)
	}

	if err == nil {
		t.Errorf("GetVideosData(\"doesNotExist.yaml\") failed. Expected %v, got %v", "error", "nil")
	} else {
		t.Logf("GetVideosData(\"doesNotExist.yaml\") success. Expected <%v>, got <%v>",
			"open : The system cannot find the file specified.", err)
	}

	// Test for invalid file
	videos, err = yamlImporter.GetVideosData("feed-exports/glorf.json")

	if videos != nil {
		t.Errorf("GetVideosData(\"feed-exports/glorf.json\") failed. Expected %v, got %v", "nil", videos)
	} else {
		t.Logf("GetVideosData(\"feed-exports/glorf.json\") success. Expected <%v>, got %v", "nil", videos)
	}

	if err == nil {
		t.Errorf("GetVideosData(\"feed-exports/glorf.json\") failed. Expected %v, got %v", "error", "nil")
	} else {
		t.Logf("GetVideosData(\"feed-exports/glorf.json\") success. Expected <%v>, got <%v>",
			"yaml: unmarshal errors: line 1: cannot unmarshal !!map into []fileReaders.YamlVideoData.", err)
	}

	// Test for valid file - Note that this test should test that all the info is loaded correctly, which is not doing
	videos, err = yamlImporter.GetVideosData("feed-exports/flub.yaml")

	if err != nil {
		t.Errorf("GetVideosData(\"feed-exports/flub.yaml\") failed. Expected %v, got %v", "error", "nil")
	} else {
		t.Logf("GetVideosData(\"feed-exports/flub.yaml\") success. Expected <%v>, got %v",
			"nil", err)
	}

	if videos == nil {
		t.Errorf("GetVideosData(\"feed-exports/flub.yaml\") failed. Expected a length of %v, got %v", 2, videos)
	} else {
		t.Logf("GetVideosData(\"feed-exports/flub.yaml\") success. Expected a length of %v, got %v", 2, len(videos.Videos))
		for _, video := range videos.Videos {
			if video.Title == "" {
				t.Errorf("GetVideosData(\"feed-exports/flub.yaml\") failed. Expected title value different than empty, got %v", video.Title)
			} else {
				t.Logf("GetVideosData(\"feed-exports/flub.yaml\") success. Expected a value of title with a name, got %v", video.Title)
			}

			if video.Url == "" {
				t.Errorf("GetVideosData(\"feed-exports/flub.yaml\") failed. Expected url value different than empty, got %v", video.Title)
			} else {
				t.Logf("GetVideosData(\"feed-exports/flub.yaml\") success. Expected a value of url, got %v", video.Title)
			}

			// TODO: Pending tags test
		}
	}
}
