// ˅
package lantana

import "github.com/mt3hr/lantana/src/app/lantana/lantana/api_request_response"

// ˄

type Config struct {
	// ˅

	// ˄

	ApplicationConfig *api_request_response.ApplicationConfig `json:"application_config"`

	ServerConfig *ServerConfig `json:"server_config"`

	Reps *Reps `json:"reps"`

	// ˅

	// ˄
}

// ˅

// ˄
