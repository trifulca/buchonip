package main
import "testing"

func TestCanParseIP(t *testing.T) {
  ip := parse_ip("12.12.12.12:80")

  if ip != "12.12.12.12" {
    t.Error("Expected ip but got ", ip)
  }
}
