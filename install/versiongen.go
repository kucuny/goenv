package install

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"regexp"
	"strings"
)

const GolangInstallPackageURL = "https://golang.org/dl/"

type GolangListParser struct {}

type GoVersion struct {
	version      string
	fileName     string
	kind         string
	os           string
	arch         string
	url          string
	size         string
	checksumType string
	checksum     string
}

type GoVersionList struct {
	Latest string
	Version string

}

type GoVersionDetail struct {

}

func NewGoVersion() *GolangListParser {
	return &GolangListParser{}
}

func (gv *GolangListParser) CreateGoVersionListFile() {
	goVersionList := gv.getGoVersionListFromGolangOrg()

	fmt.Println(goVersionList)
}

func (gv *GolangListParser) getGoVersionListFromGolangOrg() []GoVersion {
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

				// Go installer filename and version
				goVersion.fileName = token.String()
				verReg := regexp.MustCompile(`go(\d+\.\d+(?:\.\d+)?)\.(\S+)\.(tar\.gz|pkg|zip|msi)`)
				versionInfo := verReg.FindStringSubmatch(goVersion.fileName)
				goVersion.version = versionInfo[1]

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
								case 1: goVersion.kind = strings.ToLower(token.String())
								case 2: goVersion.os = strings.ToLower(token.String())
								case 3: goVersion.arch = strings.ToLower(token.String())
								case 4: goVersion.size = strings.ToLower(token.String())
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

									if len(goVersion.checksum) == 40 {
										goVersion.checksumType = "sha1"
									} else {
										goVersion.checksumType = "sha256"
									}

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
