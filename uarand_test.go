package uarand

import (
  "testing"
)

func TestFetchUserAgentList(t *testing.T) {
  _, err := fetchUserAgentList()
  if err != nil {
    t.Fatal(err)
  }
}

func TestGetRandomUserAgent(t *testing.T) {
  ua, err := GetRandomUserAgent()
  if err != nil {
    t.Fatal(err)
  }

  if ua == "" {
    t.Fatalf("ua should be define, but got %s", ua)
  }
}

func TestRandomRange(t *testing.T) {
  rnd := randomRange(0, 5)
  if rnd < 0 || rnd > 5 {
    t.Fatalf("rnd should be between 0 and 5, but got %d", rnd)
  }
}
