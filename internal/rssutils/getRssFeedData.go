package rssutils

import (
	"encoding/xml"
	"log"
	"net/http"
	"time"
)

const bootDevBlog = "https://blog.boot.dev/index.xml"

// Custom date format based on your RSS feed
const timeFormat = "Mon, 02 Jan 2006 15:04:05 MST"

type Item struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	PubDate     RssTime `xml:"pubDate"`
	Description string  `xml:"description"`
}

type Rss struct {
	Channel struct {
		Item []Item `xml:"item"`
	} `xml:"channel"`
}

// RssTime is a custom type to handle the date format in the RSS feed
type RssTime struct {
	time.Time
}

// UnmarshalXML implements the xml.Unmarshaller interface for RssTime
func (c *RssTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsedTime, err := time.Parse(timeFormat, v)
	if err != nil {
		return err
	}
	c.Time = parsedTime
	return nil
}

func GetRssFeedData(url string) (v Rss) {
	res, err := http.Get(bootDevBlog)
	if err != nil {
		log.Printf("error fetching content for: %s", url)
	}
	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)
	decoder.Decode(&v)

	return
}
