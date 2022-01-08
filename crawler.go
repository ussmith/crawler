package crawler

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

type MatchType int

const (
	Exact MatchType = iota
	Fuzzy
)

func Find(start, pattern string, matchType MatchType) []string {

	result := followDir(start, pattern, matchType)

	log.Infof("Found %d results", len(result))
	return result
}

func addIfMatches(file fs.FileInfo, pattern string, result *[]string) {
	log.Info(file.Name())
}

func followDir(start, pattern string, matchType MatchType) []string {
	log.Infof("Following the dir %s", start)

	var results []string
	filepath.Walk(start,
		func(path string, info os.FileInfo, err error) error {
			var match bool

			if matchType == Exact {
				match = strings.TrimSpace(info.Name()) == strings.TrimSpace(pattern)
			} else {
				match, err = regexp.Match(pattern, []byte(info.Name()))

				if err != nil {
					log.WithError(err).Error("Match attempt failed")
					return err
				}
			}

			if match {
				log.Info("####################  MATCH ###############")
				results = append(results, path)
			}
			return nil
		},
	)

	log.Infof("Returning %d results", len(results))
	return results
}
