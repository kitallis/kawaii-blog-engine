package config

import (
	"time"
)

type CookieConfig struct {
	Name string
	Value string
	Domain string
	Path string
	Secure   bool
	HTTPOnly bool
	SameSite string
}

func ExpirationTime(ttl time.Duration) time.Time {
	return time.Now().Add(time.Hour * ttl)
}

func DefaultCookieConfig() CookieConfig {
	return CookieConfig{
		Name: "_" + Config("APP_NAME") + "_token",
		Value: "",
		Domain: "",
		Path: "",
		Secure: true,
		HTTPOnly: true,
		SameSite: "Strict",
	}
}
