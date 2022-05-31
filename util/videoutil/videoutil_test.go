package videoutil

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGetSnapshot(t *testing.T) {
	GetSnapshot("../../tmp/test.flv", "../../tmp/test", 24)
}

func TestUploadData(t *testing.T) {
	file, err := os.Open("../../tmp/test.flv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileinfo.Size()
	buffer := make([]byte, fileSize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("bytes read:", bytesread)
	names := strings.Split(file.Name(), "/")
	UploadData(names[len(names)-1], buffer)
}

func TestGetDownloadUrl(t *testing.T) {
	url := GetDownloadUrl("test2.flv")
	fmt.Println("url:", url)
}
