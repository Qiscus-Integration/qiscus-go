package multichannel

// CreateRoomTagReq is Represent Create room tag request payload
type CreateRoomTagReq struct {
	RoomID string `json:"room_id"`
	Tag    string `json:"tag"`
}

type UserProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CreateAdditionalInfoRoomReq is Represent Create additional room request payload
type CreateAdditionalInfoRoomReq struct {
	UserProperties []UserProperty `json:"user_properties"`
}

// SendMessageTextByBotReq is Represent Send message text by Bot request payload
type SendMessageTextByBotReq struct {
	SenderEmail string
	Message     string
	RoomID      string
}

// SetToggleBotInRoomReq is Represent Set toggle room request payload
type SetToggleBotInRoomReq struct {
	IsActive bool `json:"is_active"`
}

// LoginAdminReq is Represent Login admin request payload
type LoginAdminReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetAllAgentsReq is Represent Get all agents request payload
type GetAllAgentsReq struct {
	Page   int // default 1
	Limit  int // default 20
	Search string
	Scope  string // either `division`, `name`, or `email`, or default
}

// AssignAgentReq is Represent Assign agent request payload
type AssignAgentReq struct {
	RoomID             string `json:"room_id"`
	AgentID            string `json:"agent_id"`
	ReplaceLatestAgent bool   `json:"replace_latest_agent"`
	MaxAgent           int    `json:"max_agent"` // default max agent is 5
}

// GetAgentsByDivisionReq is Represent Get agents by division request payload
type GetAgentsByDivisionReq struct {
	Page        int // default 1
	Limit       int // default 20
	DivisionIDs []string
	IsAvailable bool   // online availability filter, default all, can be true or false
	Sort        string // default asc (less customer count) can be desc
}

// GetAllDivisionReq is Represent Get all division request payload
type GetAllDivisionReq struct {
	Page  int // default 1
	Limit int // default 20
}

// MarkAsResolvedReq is Represent Mark as resolved request payload
type MarkAsResolvedReq struct {
	RoomID        string `json:"room_id"`
	Notes         string `json:"notes"`
	LastCommentID string `json:"last_comment_id"`
}
