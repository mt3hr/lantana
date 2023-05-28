// ˅
package lantana

// ˄

type LantanaSearchQuery struct {
	// ˅

	// ˄

	Tags []string `json:"tags"`

	Words string `json:"words"`

	Mood int `json:"mood"`

	LantanaSearchType *LantanaSearchType `json:"lantana_search_type"`

	// ˅

	// ˄
}

// ˅

// ˄
