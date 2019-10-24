package fileReaders

import (
	"github.com/Ahrimal/teamCMP/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

/// Yaml Importer
type YamlImporter struct {
}

/// Yaml format importer
type YamlVideoData struct {
	Labels string `json:"labels"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

/// Gets yaml video data
func (yamlImporter YamlImporter) GetVideosData(file string) (*models.Videos, error) {

	// Open our yamlFile
	yamlFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}

	defer yamlFile.Close()

	println("Successfully Opened " + file)

	// Read and Parse file
	byteValue, _ := ioutil.ReadAll(yamlFile)

	var yamlVideos []YamlVideoData

	err = yaml.Unmarshal(byteValue, &yamlVideos)
	if err != nil {
		return nil, err
	}

	// Convert to models.Videos data format
	return yamlImporter.convertToVideosFormat(yamlVideos)
}

/// Converts info from []YamlVideoData to Videos format
func (YamlImporter) convertToVideosFormat(yamlVideos []YamlVideoData) (*models.Videos, error) {
	var videos models.Videos
	var video models.Video

	// Convert YamlVideos to Videos format
	for _, yamlVideo := range yamlVideos {
		video.Title = yamlVideo.Name
		video.Url = yamlVideo.Url
		yamlTags := strings.Split(yamlVideo.Labels, ",")
		for k, yamlTag := range yamlTags {
			yamlTag = strings.Trim(yamlTag, " ")
			if yamlTag != "" {
				yamlTags[k] = yamlTag
			}
		}

		if len(yamlTags) == 1 && yamlTags[0] == "" {
			yamlTags = nil
		}

		video.Tags = yamlTags

		videos.Videos = append(videos.Videos, video)
	}

	return &videos, nil
}
