package uarand

import (
  "io"
	"bufio"
	"fmt"
	"math/rand"
	"time"
  "net/http"
)

const (
  userAgentList = "https://gist.githubusercontent.com/pzb/b4b6f57144aea7827ae4/raw/cf847b76a142955b1410c8bcef3aabe221a63db1/user-agents.txt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func fetchUserAgentList() (io.ReadCloser, error) {
  res, err := http.Get(userAgentList)
  if err != nil {
    return nil, err
  }

  return res.Body, nil
}

func GetRandomUserAgent() (string, error) {
	var lines []string
	data, err := fetchUserAgentList()
	if err != nil {
		return "", fmt.Errorf("can't fetch user agent lists, got %v", err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	random := randomRange(0, len(lines)-1)
	userAgent := lines[random]

	return userAgent, nil
}

func randomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
