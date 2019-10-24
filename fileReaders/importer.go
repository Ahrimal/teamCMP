package fileReaders

import (
	"errors"
	"os"
	"strings"
	"sync"
	"teamCMP/models"
)

type FileImporter interface {
	GetVideosData(file string) (*models.Videos, error)
}

/// Feed Importer
type FeedImporter struct {
}

/// Get videos data of a given source
func (feedImporter FeedImporter) GetVideosData(url, source, user, password string, all bool) error {
	println("Reading feed-exports...")

	if url == "" {
		url = "feed-exports"
	} else {
		return errors.New("not implemented")
	}

	if all {
		source = ""
	}

	var videos *models.Videos
	var err error

	// Gets all the filenames in the folder
	fileNames, err := feedImporter.getFileNames(url, source)
	if err != nil {
		println(err.Error())
	}

	if len(fileNames) == 0 {
		println("No source files found... nothing to import")
		return err
	}

	// Range between all fileNames and import the videos form each file
	var wg sync.WaitGroup
	wg.Add(len(fileNames))
	for _, fileName := range fileNames {
		// Read Files
		go func(fileName string) {
			defer wg.Done()
			fileSuffix := strings.Split(fileName, ".")
			switch fileSuffix[len(fileSuffix)-1] {
			case "json":
				videos, err = FileImporter.GetVideosData(JsonImporter{}, url+"/"+fileName)
				break
			case "yaml":
				videos, err = FileImporter.GetVideosData(YamlImporter{}, url+"/"+fileName)
				break
			default:
				videos = nil
				err = errors.New("wrong file format found: " + url + "/" + fileName)
				println(err.Error())
				return
			}
			// Import Videos
			videos.ImportAll()
		}(fileName)
	}

	wg.Wait()
	return err
}

/// Reads all filenames in the given url, if no url then it searches in folder feed-exports
func (FeedImporter) getFileNames(url, source string) ([]string, error) {

	var fileNames []string

	file, err := os.Open(url)
	if err != nil {
		return fileNames, err
	}
	defer file.Close()

	list, err := file.Readdirnames(0) // 0 to read all files and folders

	for _, name := range list {
		if strings.Contains(name, source) {
			fileNames = append(fileNames, name)
			println(name)
		}
	}

	if len(fileNames) == 0 {
		err = errors.New("no source files found")
	}

	return fileNames, err
}
