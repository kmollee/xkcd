// Package xkcd :parse xkcd API, and download comic image
package xkcd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

/*
If you want to fetch comics and metadata automatically,
you can use the JSON interface. The URLs look like this:

http://xkcd.com/info.0.json (current comic)

or:

http://xkcd.com/614/info.0.json (comic #614)

Those files contain, in a plaintext and easily-parsed format: comic titles,
URLs, post dates, transcripts (when available), and other metadata.
*/

const apiURL = `https://xkcd.com`

type Comic struct {
	ID int `json:"num"`

	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`

	Title       string `json:"title"`
	Transcripts string `json:"transcript"`
	ImageURL    string `json:"img"`
}

// GetFilename :using comic id-year-month-day as filename
func (c Comic) GetFilename() string {
	return fmt.Sprintf("%d_%s_%s_%s.png", c.ID, c.Year, c.Month, c.Day)
}

// Update :update comic id and fetch the meta
func (c *Comic) Update(id int) error {
	c.ID = id
	err := c.fetchMeta()
	return err
}

func fetchURL(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("could not fetch url %s: %v", URL, err)
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("reponse is not ok: %v", resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *Comic) fetchMeta() error {
	URL := fmt.Sprintf(apiURL+"/%d/info.0.json", c.ID)
	body, err := fetchURL(URL)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, c); err != nil {
		return fmt.Errorf("could not decode the response: %v", err)
	}

	return nil
}

// SaveTo :save image
func (c *Comic) SaveTo(w io.Writer) error {

	img, err := fetchURL(c.ImageURL)
	if err != nil {
		return fmt.Errorf("fetch image fail: %v", err)
	}
	b := bytes.NewReader(img)
	_, err = io.Copy(w, b)
	return err
}

// NewComic :crete comic
func NewComic() *Comic {
	return new(Comic)
}

// FetchLast :fetch last comic
func FetchLast() (*Comic, error) {
	URL := apiURL + "/info.0.json"
	body, err := fetchURL(URL)
	if err != nil {
		return nil, err
	}

	var comic Comic

	if err := json.Unmarshal(body, &comic); err != nil {
		return nil, fmt.Errorf("could not decode the response: %v", err)
	}
	return &comic, nil
}
