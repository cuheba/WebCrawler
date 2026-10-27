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

	// TODO 4: Убери слэш в конце пути, если он есть (и если путь не равен "/").
	// Подсказка: используй strings.TrimSuffix(u.Path, "/")
	// Внимание: если u.Path пустой, то он автоматически станет "/" после нормализации.

	
	// TODO 5: Убери якорь (u.Fragment = "").
	
	// TODO 6: Убери UTM-метки из query-параметров. 
	// Подсказка: это сложно, поэтому вот шпаргалка:
	// q := u.Query()
	// for key := range q {
	//     if strings.HasPrefix(key, "utm_") {
	//         q.Del(key)
	//     }
	// }
	// u.RawQuery = q.Encode()

	// TODO 7: Верни нормализованный URL в виде строки (u.String()) и nil.
	
	return "", nil
}
