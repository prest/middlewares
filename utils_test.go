package middlewares

import (
	"testing"

	"github.com/prest/config"
)

func TestMatchURL(t *testing.T) {
	test := []struct {
		Label        string
		URL          string
		JWTWhiteList []string
		match        bool
	}{
		{
			Label:        "/auth/login false",
			URL:          "/login",
			JWTWhiteList: []string{`^\/auth\/(.*)`, `\/(.*)\/test$`},
			match:        false,
		},
		{
			Label:        "/auth/login true",
			URL:          "/auth/login",
			JWTWhiteList: []string{`^\/auth\/(.*)`, `\/(.*)\/test$`},
			match:        true,
		},
		{
			Label:        "/auth/token true",
			URL:          "/auth/token",
			JWTWhiteList: []string{`^\/auth\/(.*)`, `\/(.*)\/test$`},
			match:        true,
		},
		{
			Label:        "/123/test true",
			URL:          "/123/test",
			JWTWhiteList: []string{`^\/auth\/(.*)`, `\/(.*)\/test$`},
			match:        true,
		},
		{
			Label:        "/hi/user/test true",
			URL:          "/hi/user/test",
			JWTWhiteList: []string{`^\/auth\/(.*)`, `\/(.*)\/test$`},
			match:        true,
		},
	}

	for _, tt := range test {
		t.Run(tt.Label, func(t *testing.T) {
			config.PrestConf.JWTWhiteList = tt.JWTWhiteList
			match, err := MatchURL(tt.URL)
			if err != nil {
				t.Error(err)
			}
			if match != tt.match {
				t.Errorf("expected %v, but got %v\n", tt.match, match)
			}
		})
	}
}
