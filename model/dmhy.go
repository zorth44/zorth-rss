package model

type RSSFeedDMHY struct {
	Channel struct {
		Title string        `xml:"title"`
		Link  string        `xml:"link"`
		Desc  string        `xml:"description"`
		Item  []RssItemDMHY `xml:"item"`
	} `xml:"channel"`
}

type RssItemDMHY struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Disc      string `xml:"description"`
	Author    string `xml:"author"`
	PubDate   string `xml:"pubDate"`
	Enclosure struct {
		URL    string `xml:"url"`
		Length string `xml:"length"`
		Type   string `xml:"type"`
	} `xml:"enclosure"`
	GUID     string   `xml:"guid"`
	Category []string `xml:"category"`
}
