package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"unicode"

	"github.com/BurntSushi/toml"
	flag "github.com/spf13/pflag"

	"github.com/fanaticscripter/HorribleOrganizer/thetvdb"
)

type tomlConfig struct {
	Shows map[string]showConfig
}

type showConfig struct {
	ID      int
	Name    string
	Mapping []string
}

// TODO: configure alternative show name used in organized filenames

// Config wraps parsed app configuration.
type Config struct {
	Shows map[string]Show
}

// Show is the interface that contains metadata of a HorribleSubs show
// and how episodes map to TheTVDB episodes.
type Show struct {
	Name      string
	TheTvDbID int
	Mapping   map[int]Episode
	Continued bool
}

// Episode contains metadata of an episode.
type Episode struct {
	Season   int
	SeasonEp int
	TotalEp  int
	Name     string
}

// insertMappingRange expects a valid, matching range; otherwise the behavior is undefined.
func (s *Show) insertMappingRange(totalEpFrom, totalEpTo, season, seasonEpFrom, seasonEpTo int) {
	if s.Mapping == nil {
		s.Mapping = make(map[int]Episode)
	}
	if totalEpTo == -1 {
		s.Continued = true
		totalEpTo = totalEpFrom
	}
	for totalEp := totalEpFrom; totalEp <= totalEpTo; totalEp++ {
		seasonEp := totalEp - totalEpFrom + seasonEpFrom
		s.Mapping[totalEp] = Episode{season, seasonEp, totalEp, ""}
	}
}

// TotalEpToEpisode retrieves the details of an episode given a total episode number.
func (s *Show) TotalEpToEpisode(totalEp int) (*Episode, error) {
	if s.Mapping == nil {
		return &Episode{1, totalEp, totalEp, ""}, nil
	}
	episode, ok := s.Mapping[totalEp]
	if ok {
		return &episode, nil
	}
	if s.Continued {
		lastMappedEp := 0
		for ep := range s.Mapping {
			if ep > lastMappedEp {
				lastMappedEp = ep
			}
		}
		lastMappedEpisode := s.Mapping[lastMappedEp]
		if totalEp > lastMappedEpisode.TotalEp {
			season := lastMappedEpisode.Season
			seasonEp := lastMappedEpisode.SeasonEp + totalEp - lastMappedEpisode.TotalEp
			return &Episode{season, seasonEp, totalEp, ""}, nil
		}
	}
	return nil, fmt.Errorf("'%s' has no declared mapping for episode %d", s.Name, totalEp)
}

// ResolveName resolves episode name using data from TheTVDB API.
func (e *Episode) ResolveName(theTvDbShow *thetvdb.Show) error {
	if e.Name != "" {
		return nil
	}
	theTvDbEpisode, err := theTvDbShow.GetEpisode(e.Season, e.SeasonEp)
	if err != nil {
		return err
	}
	e.Name = theTvDbEpisode.Name
	return nil
}

// ParseConfigFile parses the TOML application config file and returns the Config object.
func ParseConfigFile(configFile string) (*Config, error) {
	var tomlConf tomlConfig
	var parsedConf Config
	if _, err := toml.DecodeFile(configFile, &tomlConf); err != nil {
		return nil, err
	}
	validSpec := regexp.MustCompile(`^(E(?P<totalEpFrom>\d+)-(?P<totalEpTo>\d+)?)\s*:\s*` +
		`(S(?P<season>\d+)E(?P<seasonEpFrom>\d+)-(?P<seasonEpTo>\d+)?)$`)
	parsedConf.Shows = make(map[string]Show)
	for name, showConf := range tomlConf.Shows {
		var showName string
		if showConf.Name != "" {
			showName = showConf.Name
		} else {
			showName = name
		}
		show := Show{Name: showName, TheTvDbID: showConf.ID}
		var lastSpec string
		var lastTotalEp int
		for _, spec := range showConf.Mapping {
			match := validSpec.FindStringSubmatch(spec)
			if match == nil {
				return nil, fmt.Errorf("malformed spec: %s", spec)
			}
			params := make(map[string]int)
			for i, groupName := range validSpec.SubexpNames() {
				if i != 0 && groupName != "" {
					valueStr := match[i]
					if valueStr == "" {
						valueStr = "-1"
					}
					value, err := strconv.Atoi(valueStr)
					if err != nil {
						return nil, fmt.Errorf("malformed spec: %s", spec)
					}
					params[groupName] = value
				}
			}
			totalEpFrom := params["totalEpFrom"]
			totalEpTo := params["totalEpTo"]
			season := params["season"]
			seasonEpFrom := params["seasonEpFrom"]
			seasonEpTo := params["seasonEpTo"]
			if (totalEpTo > 0 && totalEpTo < totalEpFrom) || (seasonEpTo > 0 && seasonEpTo < seasonEpFrom) {
				return nil, fmt.Errorf("invalid range: %s", spec)
			}
			if ((totalEpTo > 0) != (seasonEpTo > 0)) ||
				(totalEpTo > 0 && seasonEpTo > 0 && totalEpTo-totalEpFrom != seasonEpTo-seasonEpFrom) {
				return nil, fmt.Errorf("range mismatch: %s", spec)
			}
			if totalEpFrom <= lastTotalEp || lastTotalEp == -1 {
				return nil, fmt.Errorf("range out of order: '%s' after '%s'", spec, lastSpec)
			}
			lastSpec = spec
			lastTotalEp = totalEpTo
			show.insertMappingRange(totalEpFrom, totalEpTo, season, seasonEpFrom, seasonEpTo)
		}
		parsedConf.Shows[name] = show
	}
	return &parsedConf, nil
}

// GetOrganizedPath converts a HorribleSubs filename to an organized path consumable by Plex, Infuse, etc.
func (conf *Config) GetOrganizedPath(horribleSubsFilename string) (string, error) {
	horribleSubsFilenameSpec := regexp.MustCompile(
		`^\[HorribleSubs\] (?P<showName>.*) - (?P<ep>[\d.]+) \[\d+p\]\.\w+$`)
	match := horribleSubsFilenameSpec.FindStringSubmatch(horribleSubsFilename)
	if match == nil {
		return "", fmt.Errorf("unable to parse filename: %s", horribleSubsFilename)
	}
	horribleSubsShowName := match[1]
	epStr := match[2]
	extension := filepath.Ext(horribleSubsFilename)

	show, ok := conf.Shows[horribleSubsShowName]
	if !ok {
		return "", fmt.Errorf("show '%s' not configured", horribleSubsShowName)
	}

	ep, err := strconv.Atoi(epStr)
	if err != nil {
		// Non-integer episode number, likely a special, cannot be processed automatically.
		return filepath.Join(show.Name, fmt.Sprintf("%s - E%s%s", show.Name, epStr, extension)), nil
	}

	episode, err := show.TotalEpToEpisode(ep)
	if err != nil {
		return "", err
	}

	theTvDbShow := thetvdb.GetStore().LoadShow(show.TheTvDbID, show.Name)
	if err := episode.ResolveName(theTvDbShow); err != nil {
		return "", err
	}

	// TODO: option to not include total episode number as No. # in filenames (think about WUG)
	seasonDirname := fmt.Sprintf("Season %02d", episode.Season)
	sanitizedEpisodeName := sanitizeNameForFilename(episode.Name)
	var filename string
	if show.Mapping == nil {
		if sanitizedEpisodeName != "" {
			filename = fmt.Sprintf("%s - S%02dE%03d - %s%s",
				show.Name, episode.Season, episode.SeasonEp, sanitizedEpisodeName, extension)
		} else {
			filename = fmt.Sprintf("%s - S%02dE%03d%s",
				show.Name, episode.Season, episode.SeasonEp, extension)
		}
	} else {
		if sanitizedEpisodeName != "" {
			filename = fmt.Sprintf("%s - S%02dE%03d - No.%03d %s%s",
				show.Name, episode.Season, episode.SeasonEp, episode.TotalEp, sanitizedEpisodeName, extension)
		} else {
			filename = fmt.Sprintf("%s - S%02dE%03d - No.%03d%s",
				show.Name, episode.Season, episode.SeasonEp, episode.TotalEp, extension)
		}
	}
	return filepath.Join(show.Name, seasonDirname, filename), nil
}

func sanitizeNameForFilename(s string) string {
	runes := make([]rune, 0, len(s))
	for _, ch := range s {
		var repl rune
		switch {
		// Non-printable characters
		case ch <= '\x1f' || ch == '\x7f':
			break
		// Replace whitespace characters with space.
		case unicode.IsSpace(ch):
			repl = ' '
		// Replace illegal characters in exFAT/NTFS/etc. with full-width variants
		// " => U+FF02 FULLWIDTH QUOTATION MARK (＂)
		// * => U+FF0A FULLWIDTH ASTERISK (＊)
		// / => U+FF0F FULLWIDTH SOLIDUS (／)
		// : => U+FF1A FULLWIDTH COLON (：)
		// < => U+FF1C FULLWIDTH LESS-THAN SIGN (＜)
		// > => U+FF1E FULLWIDTH GREATER-THAN SIGN (＞)
		// ? => U+FF1F FULLWIDTH QUESTION MARK (？)
		// \ => U+FF3C FULLWIDTH REVERSE SOLIDUS (＼)
		// | => U+FF5C FULLWIDTH VERTICAL LINE (｜)
		case ch == '"':
			repl = '\uFF02'
		case ch == '*':
			repl = '\uFF0A'
		case ch == '/':
			repl = '\uFF0F'
		case ch == ':':
			repl = '\uFF1A'
		case ch == '<':
			repl = '\uFF1C'
		case ch == '>':
			repl = '\uFF1E'
		case ch == '?':
			repl = '\uFF1F'
		case ch == '\\':
			repl = '\uFF3C'
		case ch == '|':
			repl = '\uFF5C'
		// Replace non-Basic Multilingual Plane characters (e.g. emojis) with
		// U+FFFD REPLACEMENT CHARACTER. Non-BMP characters are not supported
		// by old filesystems like FAT32 with only UCS-2 support.
		case ch > '\uFFFF':
			repl = '\uFFFD'
		default:
			repl = ch
		}
		if repl > 0 {
			runes = append(runes, repl)
		}
	}
	return string(runes)
}

var dry bool
var files []string

func init() {
	// TODO: --destdir option
	flag.Usage = func() {
		fmt.Printf(`Usage: %s [OPTIONS] FILE [FILE...]

FILE is the path to a HorribleSubs-downloaded video file.

Options:
`, os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&dry, "dry", false, "print what would be done but do not move or rename files")
	flag.Parse()
	files = flag.Args()
}

func main() {
	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	conf, err := ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	errorCount := 0
	for _, file := range files {
		if _, err := os.Stat(file); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			errorCount++
			continue
		}
		filename := filepath.Base(file)
		relpath, err := conf.GetOrganizedPath(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s: %s\n", filename, err)
			errorCount++
			continue
		}
		fmt.Printf("'%s'\t=>\t'%s'\n", filename, relpath)

		if dry {
			continue
		}

		destfile := filepath.Join(filepath.Dir(file), relpath)
		if _, err := os.Stat(destfile); err == nil {
			fmt.Fprintf(os.Stderr, "Error: '%s' already exists\n", destfile)
			errorCount++
			continue
		}
		destfileParentDir := filepath.Dir(destfile)
		if err := os.MkdirAll(destfileParentDir, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			errorCount++
			continue
		}
		if err := os.Rename(file, destfile); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			errorCount++
			continue
		}
	}
	if errorCount > 0 {
		log.Fatalf("Error: failed to organize %d files", errorCount)
	}
}
