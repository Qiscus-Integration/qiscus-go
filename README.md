# Qiscus API for Golang
[![Go Reference](https://pkg.go.dev/badge/github.com/Qiscus-Integration/qiscus-go.svg)](https://pkg.go.dev/github.com/Qiscus-Integration/qiscus-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/Qiscus-Integration/qiscus-go)](https://goreportcard.com/report/github.com/Qiscus-Integration/qiscus-go)

This library is the abstraction of [Qiscus](https://www.qiscus.com) SDK & Multichannel API for access from applications written with Go.

## 1. Installation
Install qiscus-go with:

```sh
go get -u github.com/Qiscus-Integration/qiscus-go
```

Then, import it using:

```go
import (
    "github.com/Qiscus-Integration/qiscus-go"
    "github.com/Qiscus-Integration/qiscus-go/$product$"
)
```
with `$product$` is the product of Qiscus such as `sdk` and `multichannel`.

## 2. Usage
```go
package main

import (
	"fmt"

	"github.com/Qiscus-Integration/qiscus-go"
	"github.com/Qiscus-Integration/qiscus-go/multichannel"
	"github.com/Qiscus-Integration/qiscus-go/sdk"
)

func main() {
	qiscus.DefaultHttpOutboundLog = true

	// Initiate client for Multichannel.
	multichannelClient := multichannel.NewMultichannel("qiscus-app-id", "qiscus-secret-key")

	// Initiate client for Multichannel using creadential email and password admin.
	multichannelClient, err := multichannel.NewMultichannelFromCredential("example@mail.com", "12345678")
	if err != nil {
		panic(err)
	}

	// Initiate client for Multichannel using environment variable.
	// QISCUS_APP_ID, QISCUS_SECRET_KEY and MULTICHANNEL_API_BASE --optional
	multichannelClient, err := multichannel.NewMultichannelFromEnv()
	if err != nil {
		panic(err)
	}

	// Sample Multichannel method.
	resp, _ := multichannelClient.CreateRoomTag(&multichannel.CreateRoomTagReq{
		RoomID: "12345678",
		Tag:    "test",
	})
	fmt.Println(resp)


	// Initiate client for SDK.
	sdkClient := sdk.NewSDK("qiscus-app-id", "qiscus-secret-key")
	
	// Initiate client for SDK using environment variable.
	// QISCUS_APP_ID, QISCUS_SECRET_KEY and QISCUS_API_BASE --optional
	sdkClient, err := sdk.NewSDKFromEnv()
	if err != nil {
		panic(err)
	}

	// Sample SDK method.
	resp, _ := sdkClient.LoginOrRegister(&sdk.LoginOrRegisterReq{
		UserID:   "guest@qiscus.com",
		Password: "12345678",
		Username: "User Demo",
	})
	fmt.Println(resp)

}

```

### 2.1 Multichannel Client
Available methods for `Multichannel`
```go
// CreateRoomTag create room tag
func (m *MultichannelImpl) CreateRoomTag(req *CreateRoomTagReq) (*CreateRoomTagResponse, *qiscus.Error)

// CreateAdditionalInfoRoomWithReplace create additional info room with replace exisiting data
func (m *MultichannelImpl) CreateAdditionalInfoRoomWithReplace(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error)

// GetAdditionalInfoRoom get additional info room by room ID
func (m *MultichannelImpl) GetAdditionalInfoRoom(roomID string) (*GetAdditionalInfoRoomResponse, *qiscus.Error)

// CreateAdditionalInfoRoom create additional info room without replace existing data
func (m *MultichannelImpl) CreateAdditionalInfoRoom(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error)

// SendMessageTextByBot send message text by bot
func (m *MultichannelImpl) SendMessageTextByBot(req *SendMessageTextByBotReq) *qiscus.Error

// SetToggleBotInRoom set tootle bot in room
func (m *MultichannelImpl) SetToggleBotInRoom(roomID string, isActive bool) (*SetToggleBotInRoomResponse, *qiscus.Error)

// GetAllAgents get all agent with scope search included
func (m *MultichannelImpl) GetAllAgents(req *GetAllAgentsReq) (*GetAllAgentsResponse, *qiscus.Error)

// AssignAgent assign agent
func (m *MultichannelImpl) AssignAgent(req *AssignAgentReq) (*AssignAgentResponse, *qiscus.Error)

// GetAgentsByDivision get agents by division
func (m *MultichannelImpl) GetAgentsByDivision(req *GetAgentsByDivisionReq) (*GetAgentsByDivisionResponse, *qiscus.Error)

// GetAllDivision get all division
func (m *MultichannelImpl) GetAllDivision(req *GetAllDivisionReq) (*GetAllDivisionResponse, *qiscus.Error)

// MarkAsResolved mark as resolved room
func (m *MultichannelImpl) MarkAsResolved(req *MarkAsResolvedReq) (*MarkAsResolvedResponse, *qiscus.Error)

// GetAllChannels get all channels
func (m *MultichannelImpl) GetAllChannels() (*GetAllChannelsResponse, *qiscus.Error)

// GetRoomByRoomID get room by room id
func (m *MultichannelImpl) GetRoomByRoomID(roomID string) (*GetRoomByRoomIDResponse, *qiscus.Error)
```

### 2.2 SDK Client
Available methods for `SDK`
```go
// LoginOrRegister Login or register
func (s *SDKImpl) LoginOrRegister(req *LoginOrRegisterReq) (*LoginOrRegisterResponse, *qiscus.Error)

// GetUserProfile Get user profile by user ID
func (s *SDKImpl) GetUserProfile(userID string) (*GetUserProfileResponse, *qiscus.Error)

// GetUserToken Get user profile by user ID
func (s *SDKImpl) GetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error)

// ResetUserToken Reset user token by user ID
func (s *SDKImpl) ResetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error)

// CreateRoom Create new room
func (s *SDKImpl) CreateRoom(req *CreateRoomReq) (*CreateRoomResponse, *qiscus.Error)

// GetOrCreateRoomWithTarget Get or create new room with target
func (s *SDKImpl) GetOrCreateRoomWithTarget(req *GetOrCreateRoomWithTargetReq) (*CreateRoomResponse, *qiscus.Error)

// GetRoomsInfo Get rooms info by room IDs
func (s *SDKImpl) GetRoomsInfo(roomIDs []string) (*GetRoomsInfoResponse, *qiscus.Error)

// UpdateRoom Update room
func (s *SDKImpl) UpdateRoom(req *UpdateRoomReq) (*UpdateRoomResponse, *qiscus.Error)

// GetRoomParticipants is Represent Get room participant
func (s *SDKImpl) GetRoomParticipants(req *GetRoomParticipantsReq) (*GetRoomParticipantsResponse, *qiscus.Error)

// AddRoomParticipants Add room participants
func (s *SDKImpl) AddRoomParticipants(req *AddRoomParticipantsReq) (*AddRoomParticipantsResponse, *qiscus.Error)

// RemoveRoomParticipants Remove room participants
func (s *SDKImpl) RemoveRoomParticipants(req *RemoveRoomParticipantsReq) (*RemoveRoomParticipantsResponse, *qiscus.Error)

// GetUserRooms Get user rooms
func (s *SDKImpl) GetUserRooms(req *GetUserRoomsReq) (*GetUserRoomsResponse, *qiscus.Error)

// PostComment Post comment
func (s *SDKImpl) PostComment(req *PostCommentReq) (*PostCommentResponse, *qiscus.Error)

// LoadComments load comments
func (s *SDKImpl) LoadComments(req *LoadCommentsReq) (*LoadCommentsResponse, *qiscus.Error)

// PostSystemEventMessage post system event message
func (s *SDKImpl) PostSystemEventMessage(req *PostSystemEventMessageReq) (*PostSystemEventMessageResponse, *qiscus.Error)

// GetUnreadCount get unread count in room
func (s *SDKImpl) GetUnreadCount(req *GetUnreadCountReq) (*GetUnreadCountResponse, *qiscus.Error)

// GetUsers get users
func (s *SDKImpl) GetUsers(req *GetUsersReq) (*GetUsersResponse, *qiscus.Error)

// LoadCommentsWithRange load comments with range
func (s *SDKImpl) LoadCommentsWithRange(req *LoadCommentsWithRangeReq) (*LoadCommentsWithRangeResponse, *qiscus.Error)

// GetOrCreateChannel get or create channel
func (s *SDKImpl) GetOrCreateChannel(req *GetOrCreateChannelReq) (*GetOrCreateChannelResponse, *qiscus.Error)

// GetAverageReplyTimeUser get average reply time user
func (s *SDKImpl) GetAverageReplyTimeUser(req *GetAverageReplyTimeUserReq) (*GetAverageReplyTimeUserResponse, *qiscus.Error)

// GetWebhookLogs get webhook logs
func (s *SDKImpl) GetWebhookLogs(req *GetWebhookLogsReq) (*GetWebhookLogsResponse, *qiscus.Error)

// DeactivateUser deactivate user
func (s *SDKImpl) DeactivateUser(req *DeactivateUserReq) (*DeactivateUserResponse, *qiscus.Error)

// ReactivateUser deactivate user
func (s *SDKImpl) ReactivateUser(req *ReactivateUserReq) (*ReactivateUserResponse, *qiscus.Error)


```

## 3. Advance Usage
### 3.1. Override Base API URL
```go
sdkClient,_ := sdk.NewSDKFromEnv()
// Default SDK base is https://api.qiscus.com, you can use SetAPIBase() to override.
sdkClient.SetAPIBase("https://api2.qiscus.com")

multichannelClient, _ := multichannel.NewMultichannelFromEnv()
// Default Multichannel base is https://multichannel.qiscus.com, you can use SetAPIBase() to override.
multichannelClient.SetAPIBase("https://multichannel2.qiscus.com")
```

### 3.2. Override HTTP Client timeout
By default, timeout value for HTTP Client 80 seconds. But you can override the HTTP client default config from global variable `qiscus.DefaultHttpClient`:
```go
t := 100 * time.Second
qiscus.DefaultHttpClient = &http.Client{
	Timeout: t,
}
```

### 3.3. HTTP Outbound Log Configuration
By default, the outbound log is `false`. You have option to change the default outbound log configuration with global variable `qiscus.DefaultHttpOutboundLog`:
```go
qiscus.DefaultHttpOutboundLog = true

// Details HTTP Outbound Log
{
  "level": "info",
  "method": "POST",
  "url": "https://multichannel.qiscus.com/api/v1/room_tag/create",
  "body": "{\"room_id\":\"12345678\",\"tag\":\"test\"}",
  "status": 200,
  "response": "{\"data\":{\"id\":1,\"name\":\"test\"}}",
  "latency": 774.9559,
  "time": "2021-09-20T14:32:24+07:00",
  "message": "OUTBOUND LOG"
}
```

## 4. Error Handling
Several functions in the product allow to throw an error, below is an qiscus error object you can use:
```go
_, err := multichannelClient.GetRoomTags("12345678")
if err != nil {
	message := err.GetMessage()               // general message error
	statusCode := err.GetStatusCode()         // HTTP status code e.g: 400, 401, etc.
	rawApiResponse := err.GetRawApiResponse() // raw Go HTTP response object
	rawError := err.GetRawError()             // raw Go err object
}
```
