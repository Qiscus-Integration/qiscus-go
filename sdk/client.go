package sdk

import (
	"errors"
	"os"

	"github.com/Qiscus-Integration/qiscus-go"
)

// APIBase is base url the library uses to contact multichannel. Use SetAPIBase() to override
const APIBase = "https://api.qiscus.com"

// SDK defines the supported subset of the SDK API.
type SDK interface {
	APIBase() string
	QiscusAppID() string
	QiscusSecretKey() string
	SetAPIBase(address string)

	LoginOrRegister(req *LoginOrRegisterReq) (*LoginOrRegisterResponse, *qiscus.Error)
	GetUserProfile(userID string) (*GetUserProfileResponse, *qiscus.Error)
	GetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error)
	ResetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error)
	CreateRoom(req *CreateRoomReq) (*CreateRoomResponse, *qiscus.Error)
	GetOrCreateRoomWithTarget(req *GetOrCreateRoomWithTargetReq) (*CreateRoomResponse, *qiscus.Error)
	GetRoomsInfo(roomIDs []string) (*GetRoomsInfoResponse, *qiscus.Error)
	UpdateRoom(req *UpdateRoomReq) (*UpdateRoomResponse, *qiscus.Error)
	GetRoomParticipants(req *GetRoomParticipantsReq) (*GetRoomParticipantsResponse, *qiscus.Error)
	AddRoomParticipants(req *AddRoomParticipantsReq) (*AddRoomParticipantsResponse, *qiscus.Error)
	RemoveRoomParticipants(req *RemoveRoomParticipantsReq) (*RemoveRoomParticipantsResponse, *qiscus.Error)
	GetUserRooms(req *GetUserRoomsReq) (*GetUserRoomsResponse, *qiscus.Error)
	PostComment(req *PostCommentReq) (*PostCommentResponse, *qiscus.Error)
	LoadComments(req *LoadCommentsReq) (*LoadCommentsResponse, *qiscus.Error)
	PostSystemEventMessage(req *PostSystemEventMessageReq) (*PostSystemEventMessageResponse, *qiscus.Error)
	GetUnreadCount(req *GetUnreadCountReq) (*GetUnreadCountResponse, *qiscus.Error)
	GetUsers(req *GetUsersReq) (*GetUsersResponse, *qiscus.Error)
	LoadCommentsWithRange(req *LoadCommentsWithRangeReq) (*LoadCommentsWithRangeResponse, *qiscus.Error)
	GetOrCreateChannel(req *GetOrCreateChannelReq) (*GetOrCreateChannelResponse, *qiscus.Error)
	GetAverageReplyTimeUser(req *GetAverageReplyTimeUserReq) (*GetAverageReplyTimeUserResponse, *qiscus.Error)
	GetWebhookLogs(req *GetWebhookLogsReq) (*GetWebhookLogsResponse, *qiscus.Error)
	DeactivateUser(req *DeactivateUserReq) (*DeactivateUserResponse, *qiscus.Error)
	ReactivateUser(req *ReactivateUserReq) (*ReactivateUserResponse, *qiscus.Error)
}

// SDKImpl bundles data needed by a large number of methods in order to interact with the SDK API.
type SDKImpl struct {
	apiBase         string
	qiscusAppID     string
	qiscusSecretKey string
}

// NewSDK creates a new client instance
func NewSDK(qiscusAppID, qiscusSecretKey string) SDK {
	return &SDKImpl{
		apiBase:         APIBase,
		qiscusAppID:     qiscusAppID,
		qiscusSecretKey: qiscusSecretKey,
	}
}

// NewSDKFromEnv returns a new SDK client using the environment variables
// QISCUS_APP_ID, QISCUS_SECRET_KEY and QISCUS_API_BASE
func NewSDKFromEnv() (SDK, error) {
	qiscusAppID := os.Getenv("QISCUS_APP_ID")
	if qiscusAppID == "" {
		return nil, errors.New("required environment variable QISCUS_APP_ID not defined")
	}

	qiscusSecretKey := os.Getenv("QISCUS_SECRET_KEY")
	if qiscusSecretKey == "" {
		return nil, errors.New("required environment variable QISCUS_SECRET_KEY not defined")
	}

	s := NewSDK(qiscusAppID, qiscusSecretKey)

	url := os.Getenv("QISCUS_API_BASE")
	if url != "" {
		s.SetAPIBase(url)
	}

	return s, nil
}

// APIBase returns the API Base URL configured for this client
func (s *SDKImpl) APIBase() string {
	return s.apiBase
}

// QiscusAppID returns the App ID configured for this client
func (s *SDKImpl) QiscusAppID() string {
	return s.qiscusAppID
}

// QiscusSecretKey returns the Secret Key configured for this client
func (s *SDKImpl) QiscusSecretKey() string {
	return s.qiscusSecretKey
}

// SetAPIBase updates the API Base URL for this client.
// Set a custom base API: m.SetAPIBase("https://api3.qiscus.com")
func (s *SDKImpl) SetAPIBase(address string) {
	s.apiBase = address
}
