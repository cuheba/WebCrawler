package reader

import (
  "fmt"
  "net/url"
  "strings"
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

  u.Host = strings.ToLower(u.Host))

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
