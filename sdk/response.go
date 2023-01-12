package sdk

import "time"

// LoginOrRegisterResponse is Represent Login or register response payload
type LoginOrRegisterResponse struct {
	Results struct {
		User struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetUserProfileResponse is Represent Get user profile response payload
type GetUserProfileResponse struct {
	Results struct {
		User struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetUserTokenResponse is Represent Get user token response payload
type GetUserTokenResponse struct {
	Results struct {
		Token string `json:"token"`
	} `json:"results"`
	Status int `json:"status"`
}

// CreateRoomResponse is Represent Create room response payload
type CreateRoomResponse struct {
	Results struct {
		Room struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"room"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetRoomsInfoResponse is Represent Get rooms info response payload
type GetRoomsInfoResponse struct {
	Results struct {
		Rooms []struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"rooms"`
	} `json:"results"`
	Status int `json:"status"`
}

// UpdateRoomResponse is Represent Update room response payload
type UpdateRoomResponse struct {
	Results struct {
		Changed bool `json:"changed"`
		Room    struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"room"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetRoomParticipantsResponse is Represent Get room participants response payload
type GetRoomParticipantsResponse struct {
	Results struct {
		Participants []struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"participants"`
	} `json:"results"`
	Status int `json:"status"`
}

// AddRoomParticipantsResponse is Represent Add room participants response payload
type AddRoomParticipantsResponse struct {
	Results struct {
		ParticipantsAdded []struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"participants_added"`
	} `json:"results"`
	Status int `json:"status"`
}

// RemoveRoomParticipantsResponse is Represent Remove room participants response payload
type RemoveRoomParticipantsResponse struct {
	Results struct {
		ParticipantsRemoved []struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"participants_removed"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetUserRoomsResponse is Represent Get user rooms response payload
type GetUserRoomsResponse struct {
	Results struct {
		Meta struct {
			CurrentPage int `json:"current_page"`
			TotalRoom   int `json:"total_room"`
		} `json:"meta"`
		Rooms []struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"rooms"`
	} `json:"results"`
	Status int `json:"status"`
}

// PostCommentResponse is Represent Post comment response payload
type PostCommentResponse struct {
	Results struct {
		Comment struct {
			Extras struct {
			} `json:"extras"`
			ID      int    `json:"id"`
			Message string `json:"message"`
			Payload struct {
			} `json:"payload"`
			Timestamp time.Time `json:"timestamp"`
			Type      string    `json:"type"`
			User      struct {
				Active    bool   `json:"active"`
				AvatarURL string `json:"avatar_url"`
				Extras    struct {
				} `json:"extras"`
				UserID   string `json:"user_id"`
				Username string `json:"username"`
			} `json:"user"`
		} `json:"comment"`
	} `json:"results"`
	Status int `json:"status"`
}

// LoadCommentsResponse is Represent Load comments response payload
type LoadCommentsResponse struct {
	Results struct {
		Comments []struct {
			Extras struct {
				Action string `json:"action"`
			} `json:"extras,omitempty"`
			ID        int       `json:"id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			Type      string    `json:"type"`
			User      struct {
				Active    bool   `json:"active"`
				AvatarURL string `json:"avatar_url"`
				Extras    struct {
					Type            string      `json:"type"`
					UserBubbleColor interface{} `json:"user_bubble_color"`
				} `json:"extras"`
				UserID   string `json:"user_id"`
				Username string `json:"username"`
			} `json:"user"`
			Payload struct {
				ObjectEmail        string        `json:"object_email"`
				ObjectEmailList    []interface{} `json:"object_email_list"`
				ObjectUsername     string        `json:"object_username"`
				ObjectUsernameList []interface{} `json:"object_username_list"`
				Payload            struct {
					Type string `json:"type"`
				} `json:"payload"`
				RoomName        string `json:"room_name"`
				SubjectEmail    string `json:"subject_email"`
				SubjectUsername string `json:"subject_username"`
				Type            string `json:"type"`
			} `json:"payload,omitempty"`
		} `json:"comments"`
	} `json:"results"`
	Status int `json:"status"`
}

// PostSystemEventMessageResponse is Represent Post system event message response payload
type PostSystemEventMessageResponse struct {
	Results struct {
		Comment struct {
			Extras struct {
				QiscusIosPn struct {
					Aps struct {
						ContentAvaibility int `json:"content-avaibility"`
					} `json:"aps"`
				} `json:"qiscus_ios_pn"`
			} `json:"extras"`
			ID      int    `json:"id"`
			Message string `json:"message"`
			Payload struct {
				ObjectEmail        string        `json:"object_email"`
				ObjectEmailList    []interface{} `json:"object_email_list"`
				ObjectUsername     string        `json:"object_username"`
				ObjectUsernameList []interface{} `json:"object_username_list"`
				Payload            struct {
					AdminEmail string `json:"admin_email"`
					Type       string `json:"type"`
				} `json:"payload"`
				RoomName        string `json:"room_name"`
				SubjectEmail    string `json:"subject_email"`
				SubjectUsername string `json:"subject_username"`
				Type            string `json:"type"`
			} `json:"payload"`
			Timestamp time.Time `json:"timestamp"`
			Type      string    `json:"type"`
			User      struct {
				Active    bool   `json:"active"`
				AvatarURL string `json:"avatar_url"`
				Extras    struct {
				} `json:"extras"`
				UserID   string `json:"user_id"`
				Username string `json:"username"`
			} `json:"user"`
		} `json:"comment"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetUnreadCountResponse is Represent Get unread count response payload
type GetUnreadCountResponse struct {
	Results struct {
		UnreadCounts []struct {
			RoomID      string `json:"room_id"`
			UnreadCount int    `json:"unread_count"`
		} `json:"unread_counts"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetUsersResponse is Represent Get users response payload
type GetUsersResponse struct {
	Results struct {
		Meta struct {
			TotalData int `json:"total_data"`
			TotalPage int `json:"total_page"`
		} `json:"meta"`
		Users []struct {
			Active    bool      `json:"active"`
			AvatarURL string    `json:"avatar_url"`
			CreatedAt time.Time `json:"created_at"`
			Email     string    `json:"email"`
			Extras    struct {
				Type            string      `json:"type"`
				UserBubbleColor interface{} `json:"user_bubble_color"`
			} `json:"extras,omitempty"`
			ID        int       `json:"id"`
			Name      string    `json:"name"`
			UpdatedAt time.Time `json:"updated_at"`
			Username  string    `json:"username"`
		} `json:"users"`
	} `json:"results"`
	Status int `json:"status"`
}

// LoadCommentsWithRangeResponse is Represent Load comments with range response payload
type LoadCommentsWithRangeResponse struct {
	Results struct {
		Comments []struct {
			Extras struct {
				Action string `json:"action"`
			} `json:"extras,omitempty"`
			ID        int         `json:"id"`
			Message   string      `json:"message"`
			Payload   interface{} `json:"payload"`
			Timestamp string      `json:"timestamp"`
			Type      string      `json:"type"`
			UniqueID  string      `json:"unique_id"`
			User      struct {
				Active    bool   `json:"active"`
				AvatarURL string `json:"avatar_url"`
				Extras    struct {
					Type            string      `json:"type"`
					UserBubbleColor interface{} `json:"user_bubble_color"`
				} `json:"extras,omitempty"`
				UserID   string `json:"user_id"`
				Username string `json:"username"`
			} `json:"user"`
		} `json:"comments"`
	} `json:"results"`
}

// GetOrCreateChannelResponse is Represent Get or create channel response payload
type GetOrCreateChannelResponse struct {
	Results struct {
		Changed bool `json:"changed"`
		Room    struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"room"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetAverageReplyTimeUserResponse is Represent Get average reply time user response payload
type GetAverageReplyTimeUserResponse struct {
	Results struct {
		Data struct {
			Duration struct {
				Average  int `json:"average"`
				Longest  int `json:"longest"`
				Shortest int `json:"shortest"`
			} `json:"duration"`
			UserID string `json:"user_id"`
		} `json:"data"`
		EndTime   string `json:"end_time"`
		StartTime string `json:"start_time"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetWebhookLogsResponse is Represent Get webhook logs response payload
type GetWebhookLogsResponse struct {
	Results struct {
		WebhookLogs []struct {
			AttemptedAt  time.Time `json:"attempted_at"`
			Endpoint     string    `json:"endpoint"`
			ErrorMessage string    `json:"error_message"`
			ID           int       `json:"id"`
			IsSuccess    bool      `json:"is_success"`
			RequestBody  time.Time `json:"request_body"`
			ResponseBody string    `json:"response_body"`
			ResponseCode int       `json:"response_code"`
		} `json:"webhook_logs"`
	} `json:"results"`
	Status int `json:"status"`
}

// DeactivateUserResponse is Represent Deactivate user response payload
type DeactivateUserResponse struct {
	Results struct {
		Message string `json:"message"`
	} `json:"results"`
	Status int `json:"status"`
}

// ReactivateUserResponse is Represent Reactivate user response payload
type ReactivateUserResponse struct {
	Results struct {
		Message string `json:"message"`
	} `json:"results"`
	Status int `json:"status"`
}
