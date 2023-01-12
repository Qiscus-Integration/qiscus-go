package sdk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const roomID = "123123"

func TestLoginOrRegister(t *testing.T) {
	const (
		userID    = "guest@mail.com"
		userName  = "Guest"
		password  = "12345678"
		avatarURL = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/login_or_register")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"user":{"user_id":"%s","username":"%s","avatar_url":"%s"}}}`, userID, userName, avatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.LoginOrRegister(&LoginOrRegisterReq{
		UserID:    userID,
		Username:  userName,
		Password:  password,
		AvatarURL: avatarURL,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.User.UserID, userID)
	assert.Equal(t, result.Results.User.Username, userName)
	assert.Equal(t, result.Results.User.AvatarURL, avatarURL)
}

func TestGetUserProfile(t *testing.T) {
	const (
		userID    = "guest@mail.com"
		userName  = "Guest"
		password  = "12345678"
		avatarURL = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/user_profile")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"user":{"user_id":"%s","username":"%s","avatar_url":"%s"}}}`, userID, userName, avatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetUserProfile(userID)
	assert.Nil(t, err)
	assert.Equal(t, result.Results.User.UserID, userID)
	assert.Equal(t, result.Results.User.Username, userName)
	assert.Equal(t, result.Results.User.AvatarURL, avatarURL)
}

func TestGetUserToken(t *testing.T) {
	const (
		userToken = "token-123"
		userID    = "guest@mail.com"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_user_token")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"token":"%s"}}`, userToken)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetUserToken(userID)
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Token, userToken)
}

func TestResetUserToken(t *testing.T) {
	const (
		newUserToken = "token-123"
		userID       = "guest@mail.com"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/reset_user_token")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"token":"%s"}}`, newUserToken)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.ResetUserToken(userID)
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Token, newUserToken)
}

func TestCreateRoom(t *testing.T) {
	const (
		roomName      = "Room sample"
		creator       = "guest@mail.com"
		participant   = "guest@mail.com"
		roomAvatarURL = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/create_room")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"room":{"room_name":"%s","room_avatar_url":"%s"}}}`, roomName, roomAvatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.CreateRoom(&CreateRoomReq{
		RoomName:      roomName,
		Creator:       creator,
		Participants:  []string{participant},
		RoomAvatarURL: roomAvatarURL,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Room.RoomName, roomName)
	assert.Equal(t, result.Results.Room.RoomAvatarURL, roomAvatarURL)
}

func TestGetOrCreateRoomWithTarget(t *testing.T) {
	const userID = "guest@mail.com"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_or_create_room_with_target")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		fmt.Fprint(w, `{"results":{"room":{"room_name":"","room_avatar_url":""}}}`)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetOrCreateRoomWithTarget(&GetOrCreateRoomWithTargetReq{
		UserIDs: []string{userID},
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Room.RoomName, "")
	assert.Equal(t, result.Results.Room.RoomAvatarURL, "")
}

func TestGetRoomsInfo(t *testing.T) {
	const (
		roomName      = "Room sample"
		roomAvatarURL = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_rooms_info")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"rooms":[{"room_id":"%s","room_name":"%s","room_avatar_url":"%s"}]}}`, roomID, roomName, roomAvatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetRoomsInfo([]string{roomID})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Rooms[0].RoomID, roomID)
	assert.Equal(t, result.Results.Rooms[0].RoomName, roomName)
	assert.Equal(t, result.Results.Rooms[0].RoomAvatarURL, roomAvatarURL)
}

func TestUpdateRoom(t *testing.T) {
	const newRoomName = "New room sample"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/update_room")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"changed":true,"room":{"room_id":"%s","room_name":"%s"}}}`, roomID, newRoomName)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.UpdateRoom(&UpdateRoomReq{
		RoomID:   roomID,
		RoomName: newRoomName,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Changed, true)
	assert.Equal(t, result.Results.Room.RoomID, roomID)
	assert.Equal(t, result.Results.Room.RoomName, newRoomName)
}

func TestGetRoomParticipants(t *testing.T) {
	const (
		participantID   = "guest@mail.com"
		participantName = "Guest"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_room_participants")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"participants":[{"user_id":"%s","username":"%s"}]}}`, participantID, participantName)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetRoomParticipants(&GetRoomParticipantsReq{RoomID: roomID})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Participants[0].UserID, participantID)
	assert.Equal(t, result.Results.Participants[0].Username, participantName)
}

func TestAddRoomParticipants(t *testing.T) {
	const (
		userID    = "guest@mail.com"
		userName  = "Guest"
		avatarURL = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/add_room_participants")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"participants_added":[{"user_id":"%s","username":"%s","avatar_url":"%s"}]}}`, userID, userName, avatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.AddRoomParticipants(&AddRoomParticipantsReq{
		RoomID:  roomID,
		UserIDs: []string{userID},
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.ParticipantsAdded[0].UserID, userID)
	assert.Equal(t, result.Results.ParticipantsAdded[0].Username, userName)
	assert.Equal(t, result.Results.ParticipantsAdded[0].AvatarURL, avatarURL)
}

func TestRemoveRoomParticipants(t *testing.T) {
	const (
		userID    = "guest@mail.com"
		userName  = "Guest"
		avatarURL = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/remove_room_participants")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"participants_removed":[{"user_id":"%s","username":"%s","avatar_url":"%s"}]}}`, userID, userName, avatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.RemoveRoomParticipants(&RemoveRoomParticipantsReq{})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.ParticipantsRemoved[0].UserID, userID)
	assert.Equal(t, result.Results.ParticipantsRemoved[0].Username, userName)
	assert.Equal(t, result.Results.ParticipantsRemoved[0].AvatarURL, avatarURL)
}

func TestGetUserRooms(t *testing.T) {
	const (
		userID        = "guest@mail.com"
		roomName      = "Room sample"
		roomAvatarURL = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_user_rooms")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"rooms":[{"room_id":"%s","room_name":"%s","room_avatar_url":"%s"}]}}`, roomID, roomName, roomAvatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetUserRooms(&GetUserRoomsReq{UserID: userID})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Rooms[0].RoomID, roomID)
	assert.Equal(t, result.Results.Rooms[0].RoomName, roomName)
	assert.Equal(t, result.Results.Rooms[0].RoomAvatarURL, roomAvatarURL)
}

func TestPostComment(t *testing.T) {
	const (
		message = "test"
		userID  = "guest@mail.com"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/post_comment")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"comment":{"message":"%s","type":"text","user":{"user_id":"%s"}}}}`, message, userID)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.PostComment(&PostCommentReq{
		UserID:  userID,
		Message: message,
		RoomID:  roomID,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Comment.Message, message)
	assert.Equal(t, result.Results.Comment.Type, "text")
	assert.Equal(t, result.Results.Comment.User.UserID, userID)
}

func TestLoadComments(t *testing.T) {
	const (
		commentID      = 1
		commentMessage = "test"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/load_comments")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"comments":[{"id":%d,"message":"%s"}]}}`, commentID, commentMessage)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.LoadComments(&LoadCommentsReq{RoomID: roomID})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Comments[0].ID, commentID)
	assert.Equal(t, result.Results.Comments[0].Message, commentMessage)
}

func TestPostSystemEventMessage(t *testing.T) {
	const (
		commentID      = 1
		commentMessage = "test"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/post_system_event_message")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"comment":{"id":%d,"message":"%s","type":"custom"}}}`, commentID, commentMessage)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.PostSystemEventMessage(&PostSystemEventMessageReq{
		RoomID:  roomID,
		Message: commentMessage,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Comment.ID, commentID)
	assert.Equal(t, result.Results.Comment.Message, commentMessage)
	assert.Equal(t, result.Results.Comment.Type, "custom")
}

func TestGetUnreadCount(t *testing.T) {
	const (
		userID      = "guest@mail.com"
		unreadCount = 10
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_unread_count")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"unread_counts":[{"room_id":"%s","unread_count":%d}]}}`, roomID, unreadCount)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetUnreadCount(&GetUnreadCountReq{
		RoomIDs: []string{roomID},
		UserID:  userID,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.UnreadCounts[0].RoomID, roomID)
	assert.Equal(t, result.Results.UnreadCounts[0].UnreadCount, unreadCount)
}

func TestGetUsers(t *testing.T) {
	const (
		userID        = 1
		userEmail     = "guest@mail.com"
		userName      = "Guest"
		userAvatarURL = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_user_list")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"users":[{"id":%d,"email":"%s","name":"%s","avatar_url":"%s"}]}}`, userID, userEmail, userName, userAvatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetUsers(&GetUsersReq{})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Users[0].ID, userID)
	assert.Equal(t, result.Results.Users[0].Email, userEmail)
	assert.Equal(t, result.Results.Users[0].Name, userName)
	assert.Equal(t, result.Results.Users[0].AvatarURL, userAvatarURL)
}

func TestLoadCommentsWithRange(t *testing.T) {
	const (
		firstCommentID      = 1
		firstCommentMessage = "Message 1"
		lastCommentID       = 2
		lastCommentMessage  = "Message 2"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/load_comments_with_range")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"comments":[{"id":%d,"message":"%s"},{"id":%d,"message":"%s"}]}}`, firstCommentID, firstCommentMessage, lastCommentID, lastCommentMessage)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.LoadCommentsWithRange(&LoadCommentsWithRangeReq{
		RoomID:         roomID,
		FirstCommentID: strconv.Itoa(firstCommentID),
		LastCommentID:  strconv.Itoa(lastCommentID),
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Comments[0].ID, firstCommentID)
	assert.Equal(t, result.Results.Comments[0].Message, firstCommentMessage)
	assert.Equal(t, result.Results.Comments[1].ID, lastCommentID)
	assert.Equal(t, result.Results.Comments[1].Message, lastCommentMessage)
}

func TestGetOrCreateChannel(t *testing.T) {
	const (
		roomChannelID = "channel-123"
		roomName      = "Channel sample"
		participant   = "guest@mail.com"
		avatarURL     = "https://example.com/avatar.svg"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_or_create_channel")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"room":{"room_channel_id":"%s","room_id":"%s","room_name":"%s","room_avatar_url":"%s","room_type":"channel"}}}`, roomChannelID, roomID, roomName, avatarURL)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetOrCreateChannel(&GetOrCreateChannelReq{
		UniqueID:      roomChannelID,
		RoomName:      roomName,
		Participants:  []string{participant},
		RoomAvatarURL: avatarURL,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Room.RoomChannelID, roomChannelID)
	assert.Equal(t, result.Results.Room.RoomID, roomID)
	assert.Equal(t, result.Results.Room.RoomName, roomName)
	assert.Equal(t, result.Results.Room.RoomAvatarURL, avatarURL)
	assert.Equal(t, result.Results.Room.RoomType, "channel")
}

func TestGetAverageReplyTimeUser(t *testing.T) {
	const (
		userID  = "guest@mail.com"
		average = 10
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/get_average_reply_time_user")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"data":{"user_id":"%s","duration":{"average":%d}}}}`, userID, average)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetAverageReplyTimeUser(&GetAverageReplyTimeUserReq{UserID: userID})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Data.UserID, userID)
	assert.Equal(t, result.Results.Data.Duration.Average, average)
}

func TestGetWebhookLogs(t *testing.T) {
	const (
		webhookID    = 1
		responseCode = 200
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/webhook_logs")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"webhook_logs":[{"id":%d,"response_code":%d}]}}`, webhookID, responseCode)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.GetWebhookLogs(&GetWebhookLogsReq{})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.WebhookLogs[0].ID, webhookID)
	assert.Equal(t, result.Results.WebhookLogs[0].ResponseCode, responseCode)
}

func TestDeactivateUser(t *testing.T) {
	const (
		userID      = "guest@mail.com"
		respMessage = "successfully deactivate user"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodDelete)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/deactivate_users")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"message":"%s"}}`, respMessage)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.DeactivateUser(&DeactivateUserReq{
		UserIDs: []string{userID},
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Message, respMessage)
}

func TestReactivateUser(t *testing.T) {
	const (
		userID      = "guest@mail.com"
		respMessage = "successfully reactivate user"
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.Path, "/api/v2.1/rest/reactivate_users")
		assert.Equal(t, req.Header.Get("QISCUS_SDK_APP_ID"), qiscusAppID)
		assert.Equal(t, req.Header.Get("QISCUS_SDK_SECRET"), qiscusSecretKey)

		rsp := fmt.Sprintf(`{"results":{"message":"%s"}}`, respMessage)
		fmt.Fprint(w, rsp)
	}))

	defer srv.Close()

	c := NewSDK(qiscusAppID, qiscusSecretKey)
	c.SetAPIBase(srv.URL)

	result, err := c.ReactivateUser(&ReactivateUserReq{
		UserIDs: []string{userID},
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Results.Message, respMessage)
}
