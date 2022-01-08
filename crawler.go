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
	var result []string

	followDir(start, pattern, result, matchType)

	//for _, i := range result {
	//log.Infof("Found %s", i)
	//}

	log.Infof("Found %d results", len(result))
	return result
}

func addIfMatches(file fs.FileInfo, pattern string, result *[]string) {
	log.Info(file.Name())
}

func followDir(start, pattern string, results []string, matchType MatchType) {
	log.Infof("Following the dir %s", start)

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

	for _, i := range results {
		log.Infof("Found %s", i)
	}
	// for _, f := range dir {
	// 	if f.IsDir() {
	// 		if root == "" {
	// 			log.Info("Emtpy root")
	// 			root = start + "/" + f.Name()
	// 		} else {
	// 			log.Infof("root base: %s", root)
	// 			root = root + "/" + f.Name()
	// 			log.Infof("root : %s", root)
	// 		}

	// 		log.Infof("Following: %s", root)
	// 		go followDir(root, pattern, results)
	// 	} else {
	// 		log.Infof("Adding if matches %s", f.Name())
	// 		addIfMatches(f, pattern, results)
	// 	}
	// }

}
