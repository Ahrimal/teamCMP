package fileReaders

import (
	"encoding/json"
	"github.com/Ahrimal/teamCMP/models"
	"io/ioutil"
	"os"
)

/// Json Importer
type JsonImporter struct {
	Videos []models.Video
}

/// Gets json video data
func (JsonImporter) GetVideosData(file string) (*models.Videos, error) {

	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	println("Successfully Opened " + file)

	// Read and Parse file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var videos models.Videos

	err = json.Unmarshal(byteValue, &videos)

	return &videos, err
}
