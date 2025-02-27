package main

import "testing"

func TestUnescapeRSSFeed(t *testing.T) {
	cases := map[string]struct {
		input    *RSSFeed
		expected *RSSFeed
	}{
		"nil feed": {
			input:    nil,
			expected: nil,
		},
		"has escaped entities": {
			input: &RSSFeed{struct {
				Title       string    "xml:\"title\""
				Link        string    "xml:\"link\""
				Description string    "xml:\"description\""
				Item        []RSSItem "xml:\"item\""
			}{
				Title:       "&lt;channel title&gt;",
				Description: "&lt;channel description&gt;",
				Item: []RSSItem{
					RSSItem{Title: "&lt;rss item title&gt;", Description: "&lt;rss item description&gt;"},
				},
			}},
			expected: &RSSFeed{struct {
				Title       string    "xml:\"title\""
				Link        string    "xml:\"link\""
				Description string    "xml:\"description\""
				Item        []RSSItem "xml:\"item\""
			}{
				Title:       "<channel title>",
				Description: "<channel description>",
				Item: []RSSItem{
					RSSItem{Title: "<rss item title>", Description: "<rss item description>"},
				},
			}},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			output := unescapeRSSFeed(c.input)

			if c.expected == nil {
				if output == nil {
					return
				} else {
					t.Errorf("%v != %v", output, c.expected)
				}
			}

			if output.Channel.Title != c.expected.Channel.Title {
				t.Errorf("%v != %v", output.Channel.Title, c.expected.Channel.Title)
			}

			if output.Channel.Description != c.expected.Channel.Description {
				t.Errorf("%v != %v", output.Channel.Description, c.expected.Channel.Description)
			}

			for i := range output.Channel.Item {
				if output.Channel.Item[i].Title != c.expected.Channel.Item[i].Title {
					t.Errorf("%v != %v", output.Channel.Item[i].Title, c.expected.Channel.Item[i].Title)
				}

				if output.Channel.Item[i].Description != c.expected.Channel.Item[i].Description {
					t.Errorf("%v != %v", output.Channel.Item[i].Description, c.expected.Channel.Item[i].Description)
				}
			}
		})
	}
}
