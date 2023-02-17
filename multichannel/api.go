package multichannel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Qiscus-Integration/qiscus-go"
)

// GetRoomTags get room tags by room ID
func (m *MultichannelImpl) GetRoomTags(roomID string) (*RoomTagsResponse, *qiscus.Error) {
	resp := &RoomTagsResponse{}
	url := fmt.Sprintf("%s/api/v1/room_tag/%s", m.APIBase(), roomID)

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// CreateRoomTag create room tag
func (m *MultichannelImpl) CreateRoomTag(req *CreateRoomTagReq) (*CreateRoomTagResponse, *qiscus.Error) {
	resp := &CreateRoomTagResponse{}
	url := fmt.Sprintf("%s/api/v1/room_tag/create", m.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// CreateAdditionalInfoRoomWithReplace create additional info room with replace exisiting data
func (m *MultichannelImpl) CreateAdditionalInfoRoomWithReplace(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error) {
	resp := &CreateAdditionalInfoRoomResponse{}
	url := fmt.Sprintf("%s/api/v1/qiscus/room/%s/user_info", m.APIBase(), roomID)

	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// GetAdditionalInfoRoom get additional info room by room ID
func (m *MultichannelImpl) GetAdditionalInfoRoom(roomID string) (*GetAdditionalInfoRoomResponse, *qiscus.Error) {
	resp := &GetAdditionalInfoRoomResponse{}
	url := fmt.Sprintf("%s/api/v1/qiscus/room/%s/user_info", m.APIBase(), roomID)

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// CreateAdditionalInfoRoom create additional info room without replace existing data
func (m *MultichannelImpl) CreateAdditionalInfoRoom(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error) {
	resp := &CreateAdditionalInfoRoomResponse{}

	res, e := m.GetAdditionalInfoRoom(roomID)
	if e != nil {
		return resp, e
	}

	// Merge existing additional info data (if available).
	// We must do the get and set method when adding data to the additional info room,
	// so as not to replace the existing data.
	var existingAdditionalInfoData []UserProperty
	userProperties := res.Data.Extras.UserProperties
	if len(userProperties) > 0 {
		for _, userProp := range userProperties {
			existingAdditionalInfoData = append(existingAdditionalInfoData, UserProperty{
				Key:   userProp.Key,
				Value: userProp.Value,
			})
		}
	}
	req.UserProperties = append(req.UserProperties, existingAdditionalInfoData...)

	url := fmt.Sprintf("%s/api/v1/qiscus/room/%s/user_info", m.APIBase(), roomID)
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// SendMessageTextByBot send message text by bot
func (m *MultichannelImpl) SendMessageTextByBot(req *SendMessageTextByBotReq) *qiscus.Error {
	url := fmt.Sprintf("%s/%s/bot", m.APIBase(), m.QiscusAppID())

	newReq := struct {
		SenderEmail string `json:"sender_email"`
		Message     string `json:"message"`
		RoomID      string `json:"room_id"`
		Type        string `json:"type"`
	}{
		SenderEmail: req.SenderEmail,
		Message:     req.Message,
		RoomID:      req.RoomID,
		Type:        "text",
	}

	jsonReq, _ := json.Marshal(newReq)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), nil)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return err
}

// SetToggleBotInRoom set tootle bot in room
func (m *MultichannelImpl) SetToggleBotInRoom(roomID string, isActive bool) (*SetToggleBotInRoomResponse, *qiscus.Error) {
	resp := &SetToggleBotInRoomResponse{}
	url := fmt.Sprintf("%s/bot/%s/activate", m.APIBase(), roomID)

	req := SetToggleBotInRoomReq{IsActive: isActive}
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// GetAllAgents get all agent with scope search included
func (m *MultichannelImpl) GetAllAgents(req *GetAllAgentsReq) (*GetAllAgentsResponse, *qiscus.Error) {
	resp := &GetAllAgentsResponse{}
	url := fmt.Sprintf("%s/api/v2/admin/agents", m.APIBase())

	// Set default page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 20
	}

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	r.AddParameter("page", strconv.Itoa(req.Page))
	r.AddParameter("limit", strconv.Itoa(req.Limit))
	r.AddParameter("search", req.Search)
	r.AddParameter("scope", req.Scope)

	err := r.DoRequest()

	return resp, err
}

// AssignAgent assign agent
func (m *MultichannelImpl) AssignAgent(req *AssignAgentReq) (*AssignAgentResponse, *qiscus.Error) {
	resp := &AssignAgentResponse{}
	url := fmt.Sprintf("%s/api/v1/admin/service/assign_agent", m.APIBase())

	// Default max agent
	if req.MaxAgent <= 0 {
		req.MaxAgent = 5
	}

	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// GetAgentsByDivision get agents by division
func (m *MultichannelImpl) GetAgentsByDivision(req *GetAgentsByDivisionReq) (*GetAgentsByDivisionResponse, *qiscus.Error) {
	resp := &GetAgentsByDivisionResponse{}
	url := fmt.Sprintf("%s/api/v2/admin/agents/by_division", m.APIBase())

	// Set default page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 20
	}

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	r.AddParameter("page", strconv.Itoa(req.Page))
	r.AddParameter("limit", strconv.Itoa(req.Limit))
	r.AddParameter("is_available", strconv.FormatBool(req.IsAvailable))
	r.AddParameter("sort", req.Sort)

	for _, divisionID := range req.DivisionIDs {
		r.AddParameter("division_ids[]", divisionID)
	}
	err := r.DoRequest()

	return resp, err
}

// GetAllDivision get all division
func (m *MultichannelImpl) GetAllDivision(req *GetAllDivisionReq) (*GetAllDivisionResponse, *qiscus.Error) {
	resp := &GetAllDivisionResponse{}
	url := fmt.Sprintf("%s/api/v2/divisions", m.APIBase())

	// Set default page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 20
	}

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	r.AddParameter("page", strconv.Itoa(req.Page))
	r.AddParameter("limit", strconv.Itoa(req.Limit))
	err := r.DoRequest()

	return resp, err
}

// MarkAsResolved mark as resolved room
func (m *MultichannelImpl) MarkAsResolved(req *MarkAsResolvedReq) (*MarkAsResolvedResponse, *qiscus.Error) {
	resp := &MarkAsResolvedResponse{}
	url := fmt.Sprintf("%s/api/v1/admin/service/mark_as_resolved", m.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// GetAllChannels get all channels
func (m *MultichannelImpl) GetAllChannels() (*GetAllChannelsResponse, *qiscus.Error) {
	resp := &GetAllChannelsResponse{}
	url := fmt.Sprintf("%s/api/v2/channels", m.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}

// GetRoomByRoomID get room by room id
func (m *MultichannelImpl) GetRoomByRoomID(roomID string) (*GetRoomByRoomIDResponse, *qiscus.Error) {
	resp := &GetRoomByRoomIDResponse{}
	url := fmt.Sprintf("%s/api/v2/customer_rooms/%s", m.APIBase(), roomID)

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Qiscus-App-Id", m.QiscusAppID())
	r.AddHeader("Qiscus-Secret-Key", m.QiscusSecretKey())

	err := r.DoRequest()

	return resp, err
}
