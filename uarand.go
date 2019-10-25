package uarand

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	userAgentList = "./user-agents.txt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomUserAgent() (string, error) {
	var lines []string
	file, err := os.Open(userAgentList)
	if err != nil {
		return "", fmt.Errorf("can't readfile %s, got %v", userAgentList, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
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
