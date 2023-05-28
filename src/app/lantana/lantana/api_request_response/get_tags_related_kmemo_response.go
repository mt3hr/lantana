// ˅
package api_request_response

import (
	"github.com/mt3hr/rykv/tag"
)

// ˄

type GetTagsRelatedKmemoResponse struct {
	// ˅

	// ˄

	Errors []string `json:"errors"`

	Tags []*tag.Tag `json:"tags"`

	// ˅

	// ˄
}

// ˅

// ˄
