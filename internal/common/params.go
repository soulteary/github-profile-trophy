package common

import (
	"net/url"
	"strconv"
	"strings"
)

// Params represents parsed query parameters
type Params struct {
	values url.Values
}

// NewParams creates a new Params from URL query string
func NewParams(query string) *Params {
	values, _ := url.ParseQuery(query)
	return &Params{values: values}
}

// GetStringValue returns a string parameter value or default
func (p *Params) GetStringValue(key string, defaultValue string) string {
	if p.values.Has(key) {
		return p.values.Get(key)
	}
	return defaultValue
}

// GetNumberValue returns a number parameter value or default
func (p *Params) GetNumberValue(key string, defaultValue int) int {
	if p.values.Has(key) {
		val := p.values.Get(key)
		parsed, err := strconv.Atoi(val)
		if err == nil {
			return parsed
		}
	}
	return defaultValue
}

// GetBooleanValue returns a boolean parameter value or default
func (p *Params) GetBooleanValue(key string, defaultValue bool) bool {
	if p.values.Has(key) {
		val := p.values.Get(key)
		return strings.ToLower(val) == "true"
	}
	return defaultValue
}

// GetAll returns all values for a key
func (p *Params) GetAll(key string) []string {
	values := p.values[key]
	result := []string{}
	for _, v := range values {
		// Split by comma and trim
		parts := strings.Split(v, ",")
		for _, part := range parts {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				result = append(result, trimmed)
			}
		}
	}
	return result
}
