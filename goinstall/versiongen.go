package goinstall

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

const GolangInstallPackageURL = "https://golang.org/dl/"

type GolangListParser struct {}

type GoVersion struct {
	fileName     string
	kind         string
	os           string
	arch         string
	url          string
	size         string
	checksumType string
	checksum     string
}

func NewGoVersion() *GolangListParser {
	return &GolangListParser{}
}

func (gv *GolangListParser) GetGoVersionListFromGolangOrg() []GoVersion {
	resp, err := http.Get(GolangInstallPackageURL)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var goVersionList []GoVersion

	z := html.NewTokenizer(resp.Body)
	defer resp.Body.Close()

	// TODO: Get go version string

	for {
		isEnd := false
		var goVersion GoVersion
		tokenType := z.Next()

		switch tokenType {
		case html.ErrorToken:
			isEnd = true
		case html.StartTagToken:
			token := z.Token()

			if token.Data == "td" && gv.getAttributeValue("class", token.Attr) == "filename" {
				tokenType = z.Next()
				token = z.Token()

				if token.Data == "a" && gv.getAttributeValue("class", token.Attr) == "download" {
					// Go installer url
					goVersion.url = gv.getAttributeValue("href", token.Attr)
				}

				z.Next()
				token = z.Token()

				// Go installer filename
				goVersion.fileName = token.String()

				var i int = 0
				for {
					isEnd := false
					tokenType = z.Next()

					switch tokenType {
					case html.StartTagToken:
						token = z.Token()

						if token.Data == "td" {
							i++
							tokenType = z.Next()
							token = z.Token()

							if tokenType == html.TextToken {
								// Kind, OS, Arch, Size
								switch i {
								case 1: goVersion.kind = token.String()
								case 2: goVersion.os = token.String()
								case 3: goVersion.arch = token.String()
								case 4: goVersion.size = token.String()
								}
							} else if tokenType == html.EndTagToken {
								// Kind, OS, Arch, Size
								switch i {
								case 1: goVersion.kind = ""
								case 2: goVersion.os = ""
								case 3: goVersion.arch = ""
								case 4: goVersion.size = ""
								}
							} else if tokenType == html.StartTagToken {
								if token.Data == "tt" {
									z.Next()
									token = z.Token()
									// Go installation file's checksum and type
									goVersion.checksum = token.String()
									isEnd = true
								}
							}
						}
					}

					if isEnd {
						goVersionList = append(goVersionList, goVersion)
						break
					}
				}
			}
		}

		if isEnd {
			break
		}
	}

	fmt.Println(goVersionList)

	return goVersionList
}

func (gv *GolangListParser) getAttributeValue(key string, attr []html.Attribute) string {
	for _, a := range attr {
		if a.Key == key {
			return a.Val
		}
	}

	return ""
}
