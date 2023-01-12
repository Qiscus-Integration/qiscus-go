package multichannel

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const roomID = "123123"

func TestGetRoomTags(t *testing.T) {
	const (
		tagID   = 1
		tagName = "test"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, fmt.Sprintf("/api/v1/room_tag/%s", roomID))
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":[{"id":%d,"name":"%s"}]}`, tagID, tagName)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetRoomTags(roomID)
	assert.Nil(t, err)
	assert.Equal(t, result.Data[0].ID, tagID)
	assert.Equal(t, result.Data[0].Name, tagName)
}

func TestCreateRoomTag(t *testing.T) {
	const (
		tagID   = 1
		tagName = "test"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v1/room_tag/create")
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"id":%d,"name":"%s"}}`, tagID, tagName)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.CreateRoomTag(&CreateRoomTagReq{RoomID: roomID, Tag: strconv.Itoa(tagID)})
	assert.Nil(t, err)
	assert.Equal(t, result.Data.ID, tagID)
	assert.Equal(t, result.Data.Name, tagName)
}

func TestCreateAdditionalInfoRoomWithReplace(t *testing.T) {
	const (
		additionalInfoRoomKey   = "ping"
		additionalInfoRoomValue = "pong"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, fmt.Sprintf("/api/v1/qiscus/room/%s/user_info", roomID))
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"extras":{"user_properties":[{"key":"%s","value":"%s"}]}}}`, additionalInfoRoomKey, additionalInfoRoomValue)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.CreateAdditionalInfoRoomWithReplace(roomID, &CreateAdditionalInfoRoomReq{
		UserProperties: []UserProperty{
			{
				Key:   additionalInfoRoomKey,
				Value: additionalInfoRoomValue,
			},
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Data.Extras.UserProperties[0].Key, additionalInfoRoomKey)
	assert.Equal(t, result.Data.Extras.UserProperties[0].Value, additionalInfoRoomValue)
}

func TestGetAdditionalInfoRoom(t *testing.T) {
	const (
		additionalInfoRoomKey   = "ping"
		additionalInfoRoomValue = "pong"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, fmt.Sprintf("/api/v1/qiscus/room/%s/user_info", roomID))
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"extras":{"user_properties":[{"key":"%s","value":"%s"}]}}}`, additionalInfoRoomKey, additionalInfoRoomValue)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetAdditionalInfoRoom(roomID)
	assert.Nil(t, err)
	assert.Equal(t, result.Data.Extras.UserProperties[0].Key, additionalInfoRoomKey)
	assert.Equal(t, result.Data.Extras.UserProperties[0].Value, additionalInfoRoomValue)

}

func TestCreateAdditionalInfoRoom(t *testing.T) {
	const (
		additionalInfoRoomKey   = "ping"
		additionalInfoRoomValue = "pong"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, fmt.Sprintf("/api/v1/qiscus/room/%s/user_info", roomID))
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"extras":{"user_properties":[{"key":"%s","value":"%s"}]}}}`, additionalInfoRoomKey, additionalInfoRoomValue)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.CreateAdditionalInfoRoomWithReplace(roomID, &CreateAdditionalInfoRoomReq{
		UserProperties: []UserProperty{
			{
				Key:   additionalInfoRoomKey,
				Value: additionalInfoRoomValue,
			},
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Data.Extras.UserProperties[0].Key, additionalInfoRoomKey)
	assert.Equal(t, result.Data.Extras.UserProperties[0].Value, additionalInfoRoomValue)
}

func TestSendMessageTextByBot(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, fmt.Sprintf("/%s/bot", qiscusAppID))
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	err := c.SendMessageTextByBot(&SendMessageTextByBotReq{Message: "Hello", RoomID: roomID, SenderEmail: "test@mail.com"})
	assert.Nil(t, err)
}

func TestSetToggleBotInRoom(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, fmt.Sprintf("/bot/%s/activate", roomID))
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"room_id":"%s"}}`, roomID)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.SetToggleBotInRoom(roomID, false)
	assert.Nil(t, err)
	assert.Equal(t, result.Data.RoomID, roomID)
}

func TestGetAllAgents(t *testing.T) {
	const (
		agentName  = "Agent Sample"
		agentEmail = "agentsample@mail.com"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2/admin/agents")
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"agents":[{"name":"%s","email":"%s"}]}}`, agentName, agentEmail)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetAllAgents(&GetAllAgentsReq{
		Search: agentEmail,
		Scope:  "email",
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Data.Agents[0].Name, agentName)
	assert.Equal(t, result.Data.Agents[0].Email, agentEmail)
}

func TestAssignAgent(t *testing.T) {
	const (
		agentID    = 1
		agentName  = "Agent Sample"
		agentEmail = "agentsample@mail.com"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v1/admin/service/assign_agent")
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"added_agent":{"id":%d,"name":"%s","email":"%s"}}}`, agentID, agentName, agentEmail)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.AssignAgent(&AssignAgentReq{
		RoomID:  roomID,
		AgentID: strconv.Itoa(agentID),
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Data.AddedAgent.ID, agentID)
	assert.Equal(t, result.Data.AddedAgent.Name, agentName)
	assert.Equal(t, result.Data.AddedAgent.Email, agentEmail)
}

func TestGetAgentsByDivision(t *testing.T) {
	const (
		agentID      = 1
		agentName    = "Agent Sample"
		agentEmail   = "agentsample@mail.com"
		divisionID   = 1
		divisionName = "general"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2/admin/agents/by_division")
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":[{"id":%d,"name":"%s","email":"%s","user_roles":[{"id":%d,"name":"%s"}]}]}`, agentID, agentName, agentEmail, divisionID, divisionName)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetAgentsByDivision(&GetAgentsByDivisionReq{DivisionIDs: []string{strconv.Itoa(divisionID)}})
	assert.Nil(t, err)
	assert.Equal(t, result.Data[0].ID, agentID)
	assert.Equal(t, result.Data[0].Name, agentName)
	assert.Equal(t, result.Data[0].Email, agentEmail)
	assert.Equal(t, result.Data[0].UserRoles[0].ID, divisionID)
	assert.Equal(t, result.Data[0].UserRoles[0].Name, divisionName)
}

func TestGetAllDivision(t *testing.T) {
	const (
		divisionID   = 1
		divisionName = "general"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2/divisions")
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":[{"id":%d,"name":"%s"}]}`, divisionID, divisionName)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetAllDivision(&GetAllDivisionReq{})
	assert.Nil(t, err)
	assert.Equal(t, result.Data[0].ID, divisionID)
	assert.Equal(t, result.Data[0].Name, divisionName)
}

func TestMarkAsResolved(t *testing.T) {
	const (
		lastCommentID = 1
		notes         = "test"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v1/admin/service/mark_as_resolved")
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"service":{"room_id":"%s","notes":"%s","last_comment_id":"%d"}}}`, roomID, notes, lastCommentID)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.MarkAsResolved(&MarkAsResolvedReq{
		LastCommentID: strconv.Itoa(lastCommentID),
		Notes:         notes,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Data.Service.RoomID, roomID)
	assert.Equal(t, result.Data.Service.Notes, notes)
	assert.Equal(t, result.Data.Service.LastCommentID, strconv.Itoa(lastCommentID))
}

func TestGetAllChannels(t *testing.T) {
	const (
		qiscusChannelID   = 1
		qiscusChannelName = "test"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2/channels")
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"qiscus_channels":[{"id":%d,"name":"%s"}]}}`, qiscusChannelID, qiscusChannelName)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetAllChannels()
	assert.Nil(t, err)
	assert.Equal(t, result.Data.QiscusChannels[0].ID, qiscusChannelID)
	assert.Equal(t, result.Data.QiscusChannels[0].Name, qiscusChannelName)
}

func TestGetRoomByRoomID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, fmt.Sprintf("/api/v2/customer_rooms/%s", roomID))
		assert.Equal(t, req.Header.Get("Qiscus-App-Id"), qiscusAppID)
		assert.Equal(t, req.Header.Get("Qiscus-Secret-Key"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"data":{"customer_room":{"room_id":"%s"}}}`, roomID)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewMultichannel(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetRoomByRoomID(roomID)
	assert.Nil(t, err)
	assert.Equal(t, result.Data.CustomerRoom.RoomID, roomID)
}
