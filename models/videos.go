package models

import (
	"errors"
	"math/rand"
	"sync"
	"teamCMP/connector"
	"time"
)

type Video struct {
	connector.Connector
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
	Url   string   `json:"url"`
}

type Videos struct {
	connector.Connector
	Videos []Video `json:"videos"`
}

/// Imports all video files and stores the info into the DB
func (videos Videos) ImportAll() []error {

	var errs []error
	var videoNameList []string
	var wg sync.WaitGroup
	err := videos.Connect()
	if err != nil {
		return append(errs, err)
	}

	wg.Add(len(videos.Videos))
	for _, video := range videos.Videos {
		go func(video Video) {
			defer wg.Done()
			tags := ""
			for i, tag := range video.Tags {
				tags += tag
				if (i + 1) < len(video.Tags) {
					tags += ", "
				}
			}

			println("Importing: \"" + video.Title + "\"; Url: " + video.Url + "; Tags: " + tags)
			name, err := video.Import()
			if err != nil {
				errs = append(errs, err)
				println("error importing: \"" + video.Title + "\"; Url: " + video.Url + "; Tags: " + tags)
			} else {
				println("Successfully imported: " + video.Title)
				_, err = video.saveToDB(name)
				if err != nil {
					println("error saving in database: \"" + video.Title + "\"; Url: " + video.Url + "; Tags: " + tags)
				}
				videoNameList = append(videoNameList, name)
			}
		}(video)
	}

	wg.Wait()
	err = videos.Close()
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}

// Imports video from stream
func (video Video) Import() (string, error) {

	var err error
	// TODO: look if video exists and if not show:
	// TODO: save video in server

	if video.Title == "" {
		println("video imported without title")
	}

	if video.Url == "" {
		return "", errors.New("no url to import video")
	}

	// Sleep 1-10 seconds as if was downloading and saving (in this case its always successful
	seconds := rand.Intn(9) + 1
	time.Sleep(time.Second * time.Duration(seconds))

	// TODO: save video with name generated randomly or via an id, etc.
	return "savedVideoName", err
}

// Saves video info into our BD
func (video Video) saveToDB(name string) (interface{}, error) {

	query := "my query to insert"
	result, err := video.Execute(query)

	// TODO: save video to database (Connect + Save + Defer Close
	println("Saving to DB: " + name + " - " + video.Title)
	// id, and error if any...
	return result, err
}
