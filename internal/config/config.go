package config

import (
  "flag"
  "strings"
)

type Config struct {
  URLFile      string
  DBPath    string
  Keywords  []string
  Fetchers  int
  Parsers   int
  RateLimit int
}


func ParseFlags() *Config {
  cfg := &Config{}

  var wordsRaw string

  flag.StringVar(&cfg.URLFile, "urls", "urls.txt", "path to url file")
  flag.StringVar(&cfg.DBPath, "db", "crawler.db", "path to sqlite database")
  flag.StringVar(&wordsRaw, "words", "go,http", "comma-separated Keywords")
  flag.IntVar(&cfg.Fetchers, "fetchers", 5, "number of fetcher workers")
  flag.IntVar(&cfg.Parsers, "parsers", 5, "number of parser workers")
  flag.IntVar(&cfg.RateLimit, "rate", 10, "requests per second")

  flag.Parse()


  rawWords := strings.Split(wordsRaw, ",")

  for _, w := range rawWords {
    trimmed := strings.TrimSpace(w)
    if trimmed != "" {
      cfg.Keywords = append(cfg.Keywords, trimmed)
    }
  }

  return cfg
}

