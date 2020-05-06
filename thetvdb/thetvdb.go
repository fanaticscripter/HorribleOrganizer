package thetvdb

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
)

// APIVersion is the version of TheTVDB API in use in this package.
// https://api.thetvdb.com/swagger
const APIVersion = "3.0.0"

// Store wraps all TheTVDB API config and data.
type Store struct {
	APIKey string
	Token  string
	Shows  map[int]Show
}

// Show wraps API data of a show on TheTVDB.
type Show struct {
	ID      int
	Name    string
	Seasons map[int]Season
}

// Season wraps API data of a season on TheTVDB.
type Season struct {
	Episodes map[int]Episode
}

// Episode wraps API data of an episode on TheTVDB.
type Episode struct {
	Season   int    `json:"airedSeason"`
	SeasonEp int    `json:"airedEpisodeNumber"`
	Name     string `json:"episodeName"`
	Date     string `json:"firstAired"`
}

type links struct {
	First int
	Last  int
	Next  int
	Prev  int
}

var singleton *Store
var once sync.Once

// GetStore returns the global store for TheTVDB API data.
func GetStore() *Store {
	once.Do(func() {
		singleton = &Store{
			Shows: make(map[int]Show),
		}
	})
	return singleton
}

// LoadAuth loads API auth credentials from disk and retrieve new token as
// necessary. If token has almost expired and API key is not found, the user is
// prompted to enter the key on stdin.
func (s *Store) LoadAuth() {
	var authConf struct {
		APIKey string
		Token  string
	}
	_, _ = toml.DecodeFile("auth.toml", &authConf)
	s.APIKey = authConf.APIKey
	s.Token = authConf.Token

	if s.Token != "" {
		expiration, _ := getTokenExpirationTime(s.Token)
		if time.Until(expiration).Minutes() >= 5 {
			// Token at least 5 minutes away from expiration, no need to get new
			// token.
			return
		}
	}

	for s.APIKey == "" {
		fmt.Fprint(os.Stderr, "Enter API key <https://thetvdb.com/dashboard/account/apikey>: ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			log.Fatal("Error: failed to read API key.")
		}
		s.APIKey = strings.TrimSpace(scanner.Text())
	}

	// Get a new token from the /login API. Forget refreshing, I'm not sure if
	// expired tokens can be refreshed.
	// TODO: logging
	body, _ := json.Marshal(map[string]string{"apikey": s.APIKey})
	req, _ := http.NewRequest("POST", "https://api.thetvdb.com/login", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.thetvdb.v"+APIVersion)
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Fprintf(
			os.Stderr, "Error: login failed with HTTP %d: %s\n",
			resp.StatusCode, string(response))
		os.Exit(1)
	}
	var r struct {
		Token string
	}
	if err := json.Unmarshal(response, &r); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s: %s\n", err, string(response))
		os.Exit(1)
	}
	s.Token = r.Token

	buf := new(bytes.Buffer)
	_ = toml.NewEncoder(buf).Encode(map[string]string{
		"apikey": s.APIKey,
		"token":  s.Token,
	})
	if err := ioutil.WriteFile("auth.toml", buf.Bytes(), 0600); err != nil {
		log.Printf("Error: failed to persist credentials: %s", err)
	}
}

// LoadShow loads a TheTVDB show into the store, if it isn't loaded already.
func (s *Store) LoadShow(id int, name string) *Show {
	show, ok := s.Shows[id]
	if ok {
		return &show
	}
	show = Show{ID: id, Name: name, Seasons: make(map[int]Season)}
	s.Shows[id] = show
	return &show
}

func (s *Store) apiGetRequest(path string, query url.Values) ([]byte, error) {
	if s.Token == "" {
		s.LoadAuth()
	}
	// TODO: logging
	u := url.URL{
		Scheme:   "https",
		Host:     "api.thetvdb.com",
		Path:     path,
		RawQuery: query.Encode(),
	}
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/vnd.thetvdb.v"+APIVersion)
	req.Header.Add("Authorization", "Bearer "+s.Token)
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Fprintf(
			os.Stderr, "Error: %s failed with HTTP %d: %s\n",
			u.String(), resp.StatusCode, string(response))
		// TODO: custom error type
		return nil, errors.New("API request failed")
	}
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetSeason returns a cached season or retrieves it from the API.
func (s *Show) GetSeason(seasonNumber int) (*Season, error) {
	season, ok := s.Seasons[seasonNumber]
	if ok {
		return &season, nil
	}
	season = Season{Episodes: make(map[int]Episode)}
	path := fmt.Sprintf("/series/%d/episodes/query", s.ID)
	v := url.Values{}
	v.Set("airedSeason", strconv.Itoa(seasonNumber))
	page := 1
	for {
		v.Set("page", strconv.Itoa(page))
		response, err := GetStore().apiGetRequest(path, v)
		// TODO: custom season does not exist error, and return the empty episodes map.
		if err != nil {
			return nil, err
		}

		var r struct {
			Links links
			Data  []Episode
		}
		if err := json.Unmarshal(response, &r); err != nil {
			return nil, err
		}
		for _, episode := range r.Data {
			episode.Name = strings.TrimSpace(episode.Name)
			season.Episodes[episode.SeasonEp] = episode
		}
		if r.Links.Next == 0 {
			break
		}
		page = r.Links.Next
	}
	s.Seasons[seasonNumber] = season
	return &season, nil
}

// GetEpisode returns a cached episode or retrieves it from the API.
func (s *Show) GetEpisode(seasonNumber, episodeNumber int) (*Episode, error) {
	season, err := s.GetSeason(seasonNumber)
	if err != nil {
		return nil, err
	}
	episode, err := season.GetEpisode(episodeNumber)
	if err != nil {
		return nil, err
	}
	return episode, nil
}

// GetEpisode returns an episode from the season.
func (s *Season) GetEpisode(episodeNumber int) (*Episode, error) {
	episode, ok := s.Episodes[episodeNumber]
	if ok {
		return &episode, nil
	}
	// TODO: custom error type
	return nil, errors.New("episode does not exist")
}

func getTokenExpirationTime(token string) (time.Time, error) {
	var err error
	parsed, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		err = fmt.Errorf("failed to parse token %s", token)
		log.Print(err)
		return time.Unix(0, 0), err
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("failed to extract claims from token %s", token)
		log.Print(err)
		return time.Unix(0, 0), err
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		err = fmt.Errorf("invalid claims from token %s: %s", token, claims)
		log.Print(err)
		return time.Unix(0, 0), err
	}
	return time.Unix(int64(exp), 0), nil
}
