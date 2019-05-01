package cachestorage

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// CacheID unique identifier of the Cache object.
type CacheID string

// String returns the CacheID as string value.
func (t CacheID) String() string {
	return string(t)
}

// CachedResponseType type of HTTP response cached.
type CachedResponseType string

// String returns the CachedResponseType as string value.
func (t CachedResponseType) String() string {
	return string(t)
}

// CachedResponseType values.
const (
	CachedResponseTypeBasic          CachedResponseType = "basic"
	CachedResponseTypeCors           CachedResponseType = "cors"
	CachedResponseTypeDefault        CachedResponseType = "default"
	CachedResponseTypeError          CachedResponseType = "error"
	CachedResponseTypeOpaqueResponse CachedResponseType = "opaqueResponse"
	CachedResponseTypeOpaqueRedirect CachedResponseType = "opaqueRedirect"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CachedResponseType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CachedResponseType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CachedResponseType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CachedResponseType(in.String()) {
	case CachedResponseTypeBasic:
		*t = CachedResponseTypeBasic
	case CachedResponseTypeCors:
		*t = CachedResponseTypeCors
	case CachedResponseTypeDefault:
		*t = CachedResponseTypeDefault
	case CachedResponseTypeError:
		*t = CachedResponseTypeError
	case CachedResponseTypeOpaqueResponse:
		*t = CachedResponseTypeOpaqueResponse
	case CachedResponseTypeOpaqueRedirect:
		*t = CachedResponseTypeOpaqueRedirect

	default:
		in.AddError(errors.New("unknown CachedResponseType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CachedResponseType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// DataEntry data entry.
type DataEntry struct {
	RequestURL         string             `json:"requestURL"`         // Request URL.
	RequestMethod      string             `json:"requestMethod"`      // Request method.
	RequestHeaders     []*Header          `json:"requestHeaders"`     // Request headers
	ResponseTime       float64            `json:"responseTime"`       // Number of seconds since epoch.
	ResponseStatus     int64              `json:"responseStatus"`     // HTTP response status code.
	ResponseStatusText string             `json:"responseStatusText"` // HTTP response status text.
	ResponseType       CachedResponseType `json:"responseType"`       // HTTP response type
	ResponseHeaders    []*Header          `json:"responseHeaders"`    // Response headers
}

// Cache cache identifier.
type Cache struct {
	CacheID        CacheID `json:"cacheId"`        // An opaque unique id of the cache.
	SecurityOrigin string  `json:"securityOrigin"` // Security origin of the cache.
	CacheName      string  `json:"cacheName"`      // The name of the cache.
}

// Header [no description].
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CachedResponse cached response.
type CachedResponse struct {
	Body string `json:"body"` // Entry content, base64-encoded.
}
