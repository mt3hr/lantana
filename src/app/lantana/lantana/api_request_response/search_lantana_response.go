// ˅
package api_request_response

import "github.com/mt3hr/lantana/src/app/lantana"

// ˄

type SearchLantanaResponse struct {
	// ˅

	// ˄

	Errors []string `json:"errors"`

	Lantanas []*lantana.Lantana `json:"lantanas"`

	// ˅

	// ˄
}

// ˅

// ˄
