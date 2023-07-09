package musicbrainz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// MusicBrainzAPIEndpoint represents the base URL of the MusicBrainz API
const MusicBrainzAPIEndpoint = "https://musicbrainz.org/ws/2/"

// Artist represents an artist in the MusicBrainz database
type Artist struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	SortName  string     `json:"sort-name"`
	Type      string     `json:"type"`
	Country   string     `json:"country"`
	Area      string     `json:"area"`
	BeginDate string     `json:"begin_date"`
	EndDate   string     `json:"end_date"`
	Disambig  string     `json:"disambiguation"`
	Aliases   []Alias    `json:"aliases"`
	Relations []Relation `json:"relations"`
	Tags      []Tag      `json:"tags"`
}

// Alias represents an artist's alias in the MusicBrainz database
type Alias struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Relation represents a relation between artists in the MusicBrainz database
type Relation struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Artist Artist `json:"artist"`
}

// Tag represents a tag associated with an artist in the MusicBrainz database
type Tag struct {
	Name string `json:"name"`
}

// Release represents a release in the MusicBrainz database
type Release struct {
	ID                string             `json:"id"`
	Title             string             `json:"title"`
	Status            string             `json:"status"`
	TextRepresetation TextRepresentation `json:"text-representation"`
	ArtistCredit      []ArtistCredit     `json:"artist-credit"`
	ReleaseGroup      ReleaseGroup       `json:"release-group"`
	Relations         []Relation         `json:"relations"`
	Tags              []Tag              `json:"tags"`
	CoverArtURL       []CoverArtURL      `json:"cover-art-archive"`
}
type CoverArtURL struct {
	Artwork bool      `json:"artwork"`
	Front   bool      `json:"front"`
	Back    bool      `json:"back"`
	Count   int       `json:"count"`
	Images  []GBImage `json:"images"`
}
type GBImage struct {
	ImageURL string   `json:"image"`
	Types    []string `json:"types"`
}

// TextRepresentation represents the text representation of a release in the MusicBrainz database
type TextRepresentation struct {
	Language string `json:"language"`
	Script   string `json:"script"`
}

// ArtistCredit represents the artist credit of a release in the MusicBrainz database
type ArtistCredit struct {
	Name string `json:"name"`
}

// ArtistName represents the name of an artist in the MusicBrainz database
type ArtistName struct {
	Name string `json:"name"`
}

// ReleaseGroup represents the release group of a release in the MusicBrainz database
type ReleaseGroup struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Recording represents a recording in the MusicBrainz database
type Recording struct {
	ID           string       `json:"id"`
	Title        string       `json:"title"`
	Length       int          `json:"length"`
	ReleaseDate  string       `json:"first-release-date"`
	Relations    []Relation   `json:"relations"`
	Tags         []Tag        `json:"tags"`
	ArtistCredit []ArtistName `json:"artist-credit"`
	Releases     []Release    `json:"releases"`
}

// SearchArtists searches for artists by their name
func SearchArtists(name string, limit int) ([]Artist, error) {
	params := url.Values{}
	params.Set("query", name)
	params.Set("limit", strconv.Itoa(limit))
	params.Set("fmt", "json")

	url := fmt.Sprintf("%sartist/?%s", MusicBrainzAPIEndpoint, params.Encode())
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Artists []Artist `json:"artists"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Artists, nil
}

// GetArtistByID retrieves an artist by their ID
func GetArtistByID(id string) (*Artist, error) {
	url := fmt.Sprintf("%sartist/%s?fmt=json", MusicBrainzAPIEndpoint, id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var artist Artist
	err = json.Unmarshal(body, &artist)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

// SearchReleases searches for releases by their title
func SearchReleases(title string, limit int) ([]Release, error) {
	params := url.Values{}
	params.Set("query", title)
	params.Set("limit", strconv.Itoa(limit))
	params.Set("fmt", "json")

	url := fmt.Sprintf("%srelease/?%s", MusicBrainzAPIEndpoint, params.Encode())
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Releases []Release `json:"releases"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Releases, nil
}

// GetReleaseByID retrieves a release by its ID
func GetReleaseByID(id string) (*Release, error) {
	url := fmt.Sprintf("%srelease/%s?fmt=json", MusicBrainzAPIEndpoint, id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		return nil, err
	}

	return &release, nil
}

// SearchRecordings searches for recordings by their title
func SearchRecordings(title string, limit int) ([]Recording, error) {
	params := url.Values{}
	params.Set("query", title)
	params.Set("limit", strconv.Itoa(limit))
	params.Set("fmt", "json")

	url := fmt.Sprintf("%srecording/?%s", MusicBrainzAPIEndpoint, params.Encode())
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Recordings []Recording `json:"recordings"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Recordings, nil
}

// GetRecordingByID retrieves a recording by its ID
func GetRecordingByID(id string) (*Recording, error) {
	url := fmt.Sprintf("%srecording/%s?fmt=json", MusicBrainzAPIEndpoint, id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var recording Recording
	err = json.Unmarshal(body, &recording)
	if err != nil {
		return nil, err
	}

	return &recording, nil
}

// searchRecordings searches for recordings by song title and artist name
func SearchRecordingsByTitleAndArtist(title, artist string) ([]Recording, error) {
	query := url.QueryEscape(fmt.Sprintf("recording:%s artist:%s", title, artist))
	url := fmt.Sprintf("%srecording/?query=%s&limit=20&fmt=json", MusicBrainzAPIEndpoint, query)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Recordings []Recording `json:"recordings"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Recordings, nil
}

func GetTagsByTitleAndArtistAndAlbum(title, artist string, album string) ([]Tag, string, error) {
	query := url.QueryEscape(fmt.Sprintf("recording:%s artist:%s release:%s", title, artist, album))
	url := fmt.Sprintf("%srecording/?query=%s&limit=1&fmt=json", MusicBrainzAPIEndpoint, query)

	response, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	var result struct {
		Recordings []Recording `json:"recordings"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, "", err
	}
	if len(result.Recordings) == 1 {
		recording, err := GetRecordingByIDWithTags(result.Recordings[0].ID)
		if err != nil {
			return nil, "", err
		}
		return recording.Tags, recording.ReleaseDate, nil
	} else {
		err = errors.New("MusicBrainz didn't find the song")
	}

	return nil, "", err
}
func GetRecordingByIDWithTags(id string) (*Recording, error) {
	url := fmt.Sprintf("%srecording/%s?fmt=json", MusicBrainzAPIEndpoint, id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var recording Recording
	err = json.Unmarshal(body, &recording)
	if err != nil {
		return nil, err
	}

	return &recording, nil
}
