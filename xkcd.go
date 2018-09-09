package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
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

func (c Comic) getFilename() string {
	return fmt.Sprintf("%d_%s_%s_%s.png", c.ID, c.Year, c.Month, c.Day)
}

// Update :update comic id and fetch the meta
func (c *Comic) Update(id int) error {
	c.ID = id
	return c.fetchMeta()
}

func fetchURL(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("could not fetch url %s: %v", URL, err)
	}
	defer resp.Body.Close()

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

// SaveTo :save image to directory
func (c *Comic) SaveTo(dirPath string) error {
	dir, err := os.Stat(dirPath)
	if err != nil {
		return err
	}
	if !dir.IsDir() {
		return fmt.Errorf("could not save, the path is not directory")
	}

	d, err := filepath.Abs(dirPath)
	if err != nil {
		return err
	}

	filePath := path.Join(d, c.getFilename())

	img, err := fetchURL(c.ImageURL)
	if err != nil {
		return fmt.Errorf("fetch image fail: %v", err)
	}

	return ioutil.WriteFile(filePath, img, 0666)
	// f, err := os.Create(filePath)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	// _, err = f.Write(img)
	// return err
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
