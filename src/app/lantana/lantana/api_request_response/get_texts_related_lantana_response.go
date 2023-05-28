// ˅
package api_request_response

import "github.com/mt3hr/rykv/text"

// ˄

type GetTextsRelatedLantanaResponse struct {
	// ˅

	// ˄

	Errors []string `json:"errors"`

	Texts []*text.Text `json:"texts"`

	// ˅

	// ˄
}

// ˅

// ˄
