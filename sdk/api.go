package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Qiscus-Integration/qiscus-go"
)

// LoginOrRegister Login or register
func (s *SDKImpl) LoginOrRegister(req *LoginOrRegisterReq) (*LoginOrRegisterResponse, *qiscus.Error) {
	resp := &LoginOrRegisterResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/login_or_register", s.APIBase())

	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetUserProfile Get user profile by user ID
func (s *SDKImpl) GetUserProfile(userID string) (*GetUserProfileResponse, *qiscus.Error) {
	resp := &GetUserProfileResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/user_profile", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("user_id", userID)
	err := r.DoRequest()

	return resp, err
}

// GetUserToken Get user profile by user ID
func (s *SDKImpl) GetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error) {
	resp := &GetUserTokenResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_user_token", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("user_id", userID)
	err := r.DoRequest()

	return resp, err
}

// ResetUserToken Reset user token by user ID
func (s *SDKImpl) ResetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error) {
	resp := &GetUserTokenResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/reset_user_token", s.APIBase())

	req := &ResetUserTokenReq{UserID: userID}
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// CreateRoom Create new room
func (s *SDKImpl) CreateRoom(req *CreateRoomReq) (*CreateRoomResponse, *qiscus.Error) {
	resp := &CreateRoomResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/create_room", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetOrCreateRoomWithTarget Get or create new room with target
func (s *SDKImpl) GetOrCreateRoomWithTarget(req *GetOrCreateRoomWithTargetReq) (*CreateRoomResponse, *qiscus.Error) {
	resp := &CreateRoomResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_or_create_room_with_target", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetRoomsInfo Get rooms info by room IDs
func (s *SDKImpl) GetRoomsInfo(roomIDs []string) (*GetRoomsInfoResponse, *qiscus.Error) {
	resp := &GetRoomsInfoResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_rooms_info", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())

	for _, roomID := range roomIDs {
		r.AddParameter("room_ids[]", roomID)
	}

	err := r.DoRequest()

	return resp, err
}

// UpdateRoom Update room
func (s *SDKImpl) UpdateRoom(req *UpdateRoomReq) (*UpdateRoomResponse, *qiscus.Error) {
	resp := &UpdateRoomResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/update_room", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetRoomParticipants is Represent Get room participant
func (s *SDKImpl) GetRoomParticipants(req *GetRoomParticipantsReq) (*GetRoomParticipantsResponse, *qiscus.Error) {
	resp := &GetRoomParticipantsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_room_participants", s.APIBase())

	// Set default page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 20
	}

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("room_id", req.RoomID)
	r.AddParameter("page", strconv.Itoa(req.Page))
	r.AddParameter("limit", strconv.Itoa(req.Limit))
	err := r.DoRequest()

	return resp, err
}

// AddRoomParticipants Add room participants
func (s *SDKImpl) AddRoomParticipants(req *AddRoomParticipantsReq) (*AddRoomParticipantsResponse, *qiscus.Error) {
	resp := &AddRoomParticipantsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/add_room_participants", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// RemoveRoomParticipants Remove room participants
func (s *SDKImpl) RemoveRoomParticipants(req *RemoveRoomParticipantsReq) (*RemoveRoomParticipantsResponse, *qiscus.Error) {
	resp := &RemoveRoomParticipantsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/remove_room_participants", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetUserRooms Get user rooms
func (s *SDKImpl) GetUserRooms(req *GetUserRoomsReq) (*GetUserRoomsResponse, *qiscus.Error) {
	resp := &GetUserRoomsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_user_rooms", s.APIBase())

	// Set default page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 20
	}

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("user_id", req.UserID)
	r.AddParameter("page", strconv.Itoa(req.Page))
	r.AddParameter("limit", strconv.Itoa(req.Limit))
	err := r.DoRequest()

	return resp, err

}

// PostComment Post comment
func (s *SDKImpl) PostComment(req *PostCommentReq) (*PostCommentResponse, *qiscus.Error) {
	resp := &PostCommentResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/post_comment", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// LoadComments load comments
func (s *SDKImpl) LoadComments(req *LoadCommentsReq) (*LoadCommentsResponse, *qiscus.Error) {
	resp := &LoadCommentsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/load_comments", s.APIBase())

	// Set default page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 20
	}

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("room_id", req.RoomID)
	r.AddParameter("page", strconv.Itoa(req.Page))
	r.AddParameter("limit", strconv.Itoa(req.Limit))
	err := r.DoRequest()

	return resp, err
}

// PostSystemEventMessage post system event message
func (s *SDKImpl) PostSystemEventMessage(req *PostSystemEventMessageReq) (*PostSystemEventMessageResponse, *qiscus.Error) {
	resp := &PostSystemEventMessageResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/post_system_event_message", s.APIBase())

	newReq := struct {
		SystemEventType string      `json:"system_event_type"`
		RoomID          string      `json:"room_id"`
		Message         string      `json:"message"`
		Payload         interface{} `json:"payload"`
		Extras          interface{} `json:"extras"`
	}{
		SystemEventType: "custom",
		RoomID:          req.RoomID,
		Message:         req.Message,
		Payload:         req.Payload,
		Extras:          req.Extras,
	}

	jsonReq, _ := json.Marshal(newReq)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetUnreadCount get unread count in room
func (s *SDKImpl) GetUnreadCount(req *GetUnreadCountReq) (*GetUnreadCountResponse, *qiscus.Error) {
	resp := &GetUnreadCountResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_unread_count", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("user_id", req.UserID)

	for _, roomID := range req.RoomIDs {
		r.AddParameter("room_ids[]", roomID)
	}

	err := r.DoRequest()

	return resp, err
}

// GetUsers get users
func (s *SDKImpl) GetUsers(req *GetUsersReq) (*GetUsersResponse, *qiscus.Error) {
	resp := &GetUsersResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_user_list", s.APIBase())

	// Set default page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 20
	}

	// Set default order query
	if req.OrderQuery == "" {
		req.OrderQuery = "created_at desc nulls last"
	}

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("page", strconv.Itoa(req.Page))
	r.AddParameter("limit", strconv.Itoa(req.Limit))
	r.AddParameter("show_all", strconv.FormatBool(req.ShowAll))
	r.AddParameter("order_query", req.OrderQuery)
	err := r.DoRequest()

	return resp, err
}

// LoadCommentsWithRange load comments with range
func (s *SDKImpl) LoadCommentsWithRange(req *LoadCommentsWithRangeReq) (*LoadCommentsWithRangeResponse, *qiscus.Error) {
	resp := &LoadCommentsWithRangeResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/load_comments_with_range", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("room_id", req.RoomID)
	r.AddParameter("first_comment_id", req.FirstCommentID)
	r.AddParameter("last_comment_id", req.LastCommentID)
	err := r.DoRequest()

	return resp, err
}

// GetOrCreateChannel get or create channel
func (s *SDKImpl) GetOrCreateChannel(req *GetOrCreateChannelReq) (*GetOrCreateChannelResponse, *qiscus.Error) {
	resp := &GetOrCreateChannelResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_or_create_channel", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetAverageReplyTimeUser get average reply time user
func (s *SDKImpl) GetAverageReplyTimeUser(req *GetAverageReplyTimeUserReq) (*GetAverageReplyTimeUserResponse, *qiscus.Error) {
	resp := &GetAverageReplyTimeUserResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_average_reply_time_user", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("user_id", req.UserID)
	r.AddParameter("start_time", req.StartTime)
	r.AddParameter("end_time", req.EndTime)
	err := r.DoRequest()

	return resp, err
}

// GetWebhookLogs get webhook logs
func (s *SDKImpl) GetWebhookLogs(req *GetWebhookLogsReq) (*GetWebhookLogsResponse, *qiscus.Error) {
	resp := &GetWebhookLogsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/webhook_logs", s.APIBase())

	// Set default page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Set default limit
	if req.Limit <= 0 {
		req.Limit = 20
	}

	// Set default type
	if req.Type == "" {
		req.Type = "all"
	}

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	r.AddParameter("page", strconv.Itoa(req.Page))
	r.AddParameter("limit", strconv.Itoa(req.Limit))
	r.AddParameter("type", req.Type)
	err := r.DoRequest()

	return resp, err
}

// DeactivateUser deactivate user
func (s *SDKImpl) DeactivateUser(req *DeactivateUserReq) (*DeactivateUserResponse, *qiscus.Error) {
	resp := &DeactivateUserResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/deactivate_users", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodDelete, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}

// ReactivateUser deactivate user
func (s *SDKImpl) ReactivateUser(req *ReactivateUserReq) (*ReactivateUserResponse, *qiscus.Error) {
	resp := &ReactivateUserResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/reactivate_users", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.QiscusAppID())
	r.AddHeader("QISCUS_SDK_SECRET", s.QiscusSecretKey())
	err := r.DoRequest()

	return resp, err
}
