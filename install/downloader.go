package install

import (
	"net/http"
	"io"
	"os"
	"fmt"
)

type GoDownloader struct {
	version string
	url string
	downloadPath string
}

func NewGoDownloader(version, url, downloadPath string) GoDownloader {
	return GoDownloader{
		version: version,
		url: url,
		downloadPath: downloadPath,
	}
}

func (gd GoDownloader) DownloadGolangPackage() {
	fmt.Println("Downloading...")

	head, err := http.Head(gd.url)

	fmt.Println(head)

	out, err := os.Create("test.tar.gz")
	defer out.Close()

	resp, err := http.Get(gd.url)

	if err != nil {
		fmt.Println("Download Error!")
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		fmt.Println("Cannot save downloaded file")
	}

}
