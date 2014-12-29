// Client for betaseries.com using
// the JSON api.
//
// Request responses.
//
// Copyright © 2014 - Rémy MATHIEU

package model

import "time"

type Subtitle struct {
    Id int `json:"id"`
    Language string `json:"language"`
    Source string `json:"source"`
    Quality int `json:"quality"`
    File string `json:"file"`
    Content []string `json:"content"`
    URL string `json:"url"`
    Episode Episode `json:"episode"`
    Date time.time `json:"date"`
}

type Episode struct {
    ShowId int `json:"show_id"`
    EpisodeId int `json:"episode_id"`
    Season int `json:"season"`
    Episode int `json:"episode"`
}

