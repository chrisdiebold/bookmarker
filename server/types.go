package main

import "math/rand"

type Bookmark struct {
	Link      string `json:"link"`
	Icon      string `json:"icon"`
	ShortName string `json:"shortName"`
	Visited   int    `json:"visited"`
}

func NewBookmark(link, shortName string) *Bookmark {
	return &Bookmark{
		Link:      link,
		Icon:      "Youtube",
		ShortName: shortName,
		Visited:   rand.Intn(1000),
	}
}
