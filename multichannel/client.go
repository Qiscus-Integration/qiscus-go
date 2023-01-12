package multichannel

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/syahidfrd/qiscus-unofficial-go"
)

// APIBase is base Url the library uses to contact multichannel. Use SetAPIBase() to override
const APIBase = "https://multichannel.qiscus.com"

// Multichannel defines the supported subset of the Multichannel API.
type Multichannel interface {
	APIBase() string
	QiscusAppID() string
	QiscusSecretKey() string
	SetAPIBase(address string)

	GetRoomTags(roomID string) (*RoomTagsResponse, *qiscus.Error)
	CreateRoomTag(req *CreateRoomTagReq) (*CreateRoomTagResponse, *qiscus.Error)
	CreateAdditionalInfoRoomWithReplace(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error)
	GetAdditionalInfoRoom(roomID string) (*GetAdditionalInfoRoomResponse, *qiscus.Error)
	CreateAdditionalInfoRoom(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error)
	MarkAsResolved(req *MarkAsResolvedReq) (*MarkAsResolvedResponse, *qiscus.Error)
	GetRoomByRoomID(roomID string) (*GetRoomByRoomIDResponse, *qiscus.Error)
	SendMessageTextByBot(req *SendMessageTextByBotReq) *qiscus.Error
	SetToggleBotInRoom(roomID string, isActive bool) (*SetToggleBotInRoomResponse, *qiscus.Error)
	GetAllAgents(req *GetAllAgentsReq) (*GetAllAgentsResponse, *qiscus.Error)
	AssignAgent(req *AssignAgentReq) (*AssignAgentResponse, *qiscus.Error)
	GetAgentsByDivision(req *GetAgentsByDivisionReq) (*GetAgentsByDivisionResponse, *qiscus.Error)
	GetAllDivision(req *GetAllDivisionReq) (*GetAllDivisionResponse, *qiscus.Error)
	GetAllChannels() (*GetAllChannelsResponse, *qiscus.Error)
}

// MultichannelImpl bundles data needed by a large number of methods in order to interact with the Multichannel API.
type MultichannelImpl struct {
	apiBase         string
	qiscusAppID     string
	qiscusSecretKey string
}

// NewMultichannel creates a new client instance.
func NewMultichannel(qiscusAppID, qiscusSecretKey string) Multichannel {
	return &MultichannelImpl{
		apiBase:         APIBase,
		qiscusAppID:     qiscusAppID,
		qiscusSecretKey: qiscusSecretKey,
	}
}

// NewMultichannelFromEnv returns a new Multichannel client using the environment variables
// QISCUS_APP_ID, QISCUS_SECRET_KEY and MULTICHANNEL_API_BASE
func NewMultichannelFromEnv() (Multichannel, error) {
	qiscusAppID := os.Getenv("QISCUS_APP_ID")
	if qiscusAppID == "" {
		return nil, errors.New("required environment variable QISCUS_APP_ID not defined")
	}

	qiscusSecretKey := os.Getenv("QISCUS_SECRET_KEY")
	if qiscusSecretKey == "" {
		return nil, errors.New("required environment variable QISCUS_SECRET_KEY not defined")
	}

	m := NewMultichannel(qiscusAppID, qiscusSecretKey)

	url := os.Getenv("MULTICHANNEL_API_BASE")
	if url != "" {
		m.SetAPIBase(url)
	}

	return m, nil

}

func NewMultichannelFromCredential(email, password string) (Multichannel, error) {
	resp := &LoginAdminResponse{}
	url := fmt.Sprintf("%s/api/v1/auth", APIBase)

	req := &LoginAdminReq{Email: email, Password: password}
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	if err := r.DoRequest(); err != nil {
		return nil, fmt.Errorf("initiate client for multichannel failed. %s", err.Message)
	}

	m := NewMultichannel(resp.Data.User.App.AppCode, resp.Data.User.App.SecretKey)
	return m, nil

}

// APIBase returns the API Base URL configured for this client
func (m *MultichannelImpl) APIBase() string {
	return m.apiBase
}

// QiscusAppID returns the App ID configured for this client
func (m *MultichannelImpl) QiscusAppID() string {
	return m.qiscusAppID
}

// QiscusSecretKey returns the Secret Key configured for this client
func (m *MultichannelImpl) QiscusSecretKey() string {
	return m.qiscusSecretKey
}

// SetAPIBase updates the API Base URL for this client.
// Set a custom base API: m.SetAPIBase("https://multichannel-test.qiscus.com")
func (m *MultichannelImpl) SetAPIBase(address string) {
	m.apiBase = address
}
