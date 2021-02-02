package unparser

import "net/url"

type Href struct {
	Protocol string
	User     url.Userinfo
}
