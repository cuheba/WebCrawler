package models

import "time"

type URL struct {
  Raw      string
  Priority int
}

type FetchResult struct {
  URL        string
  StatusCode int
  Body       []byte
  Err        string
}

type PageResult struct {
  URL           string
  StatusCode    int
  MatchedWords  []string
  Snippet       string
  Err           string
  Duration      time.Duration
}
