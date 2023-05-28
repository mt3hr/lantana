// ˅
package api_request_response

import "github.com/mt3hr/kmemo"

// ˄

type GetKmemosRelatedLantanaResponse struct {
	// ˅

	// ˄

	Errors []string `json:"errors"`

	Kmemos []*kmemo.Kmemo `json:"kmemos"`

	// ˅

	// ˄
}

// ˅

// ˄
