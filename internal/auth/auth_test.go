package auth

import (
	"net/http"
	"testing"
)

func TestAuthFail(t *testing.T) {
	header := http.Header(map[string][]string{
		"Authorization": {"asuh", "dud"},
	})
	res, err := GetAPIKey(header)
	if err == nil {
		t.Error("there should be an error", res, err)
	}
}
