package reader

import (
  "bufio"
  "os"
  "fmt"
  "net/url"
  "strings"

  "github.com/cuheba/webcrawler/internal/models"
)


func NormalizeURL(raw string) (string, error) {
  raw = strings.TrimSpace(raw)

  u, err := url.Parse(raw)
  if err != nil {
    return "", fmt.Errorf("failed to parse url: %w", err)
  }


  if u.Scheme != "http" && u.Scheme != "https" {
    return "", fmt.Errorf("invalid scheme: %s", u.Scheme)
  }

  u.Host = strings.ToLower(u.Host)

  if u.Path == "" {
    u.Path = "/"
  } else if u.Path != "/" {
    u.Path = strings.TrimSuffix(u.Path, "/")
  }
	
  u.Fragment = ""
	
  q := u.Query()
  for key := range q {
    if strings.HasPrefix(key, "utm_") {
      q.Del(key)
    }
  }
  u.RawQuery = q.Encode()
	
	return u.String(), nil
}


func ReadURL(filepath string) ([]models.URL, error) {
  f, err := os.Open(filepath)
  if err != nil {
    return nil, fmt.Errorf("failed to open file: %w", err)
  }
  defer f.Close()

  seen := make(map[string]bool)
  var urls []models.URL

  scanner := bufio.NewScanner(f)
  
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" { continue }

    norm, err := NormalizeURL(line)
    if err != nil { continue }

    if seen[norm] == true { continue }
    if seen[norm] = true
    
    append(urls, models.URL{Raw: norm, Priority: 1})
  }

  if err := scanner.Err(); err != nil {
    return nil, fmt.Errorf("scanner error: %w", err)
  }

  return urls, nil
} 



