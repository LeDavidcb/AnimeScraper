package pkg

import (
	"github.com/gocolly/colly"
)

func ReturnEntries(c *colly.Collector, url string) (Entries, error) {
	var list Entries = make(Entries, 0)
	c.OnHTML("tr.default", func(e *colly.HTMLElement) {
		var entry Entry
		e.ForEach("td", func(index int, el *colly.HTMLElement) {
			switch index {
			case 1:
				// Name
				entry.Name = el.Text
			case 2:
				val, exist := el.DOM.Children().Eq(1).Attr("href")
				if exist {
					entry.Magnet = val
				}
			case 3:
				entry.Size = el.Text
			}
		})
		list = append(list, entry)
	})
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	return list, nil
}
