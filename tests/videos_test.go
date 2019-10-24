package test

import (
	"github.com/Ahrimal/teamCMP/models"
	"testing"
)

func TestVideos(t *testing.T) {
	var videos models.Videos

	// Empty videos test
	err := videos.ImportAll()
	if err != nil {
		t.Errorf("Videos().ImportAll() failed. Expected %v, got %v", "error", "nil")
	} else {
		t.Logf("Videos().ImportAll() success. Expected <[]>, got <%v>", err)
	}

	//One video without url test
	videos.Videos = append(videos.Videos, models.Video{Title: "test", Url: "", Tags: nil})
	err = videos.ImportAll()
	if err == nil {
		t.Errorf("Videos() - ImportAll() failed. Expected <no url to import video>, got <%v>", err)
	} else {
		t.Logf("Videos() - ImportAll() success. Expected <no url to import video>, got <%v>", err[0])
	}

	//One video without title test
	var videos1 models.Videos
	videos1.Videos = append(videos1.Videos, models.Video{Title: "", Url: "https://this.video.com", Tags: nil})
	err = videos1.ImportAll()
	if err != nil {
		t.Errorf("Videos(Video(``, url, nil) - ImportAll() failed. Expected <[]>, got <%v>", err)
	} else {
		t.Logf("Videos(``, url,  nil) - ImportAll() success. Expected <[]>, got <%v>", err)
	}

	//One video url test
	var videos2 models.Videos
	videos2.Videos = append(videos2.Videos, models.Video{Title: "test", Url: "https://this.video.com", Tags: nil})
	err = videos2.ImportAll()
	if err != nil {
		t.Errorf("Videos(test, url, nil) - ImportAll() failed. Expected <[]>, got <%v>", err[0])
	} else {
		t.Logf("Videos(test, url, nil) - ImportAll() success. Expected <[]>, got <%v>", err)
	}
}
