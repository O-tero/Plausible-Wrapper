package plausible

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

// Site represents a site added to plausible and implements a client
// for all stats requests related with the site.
//
// Site is safe for concurrent use.
type Site struct {
	token           string
	id              string
	httpClient      *fasthttp.Client
	plausibleClient *Client
}

// ID returns the ID of the site.
func (s *Site) ID() string {
	return s.id
}

