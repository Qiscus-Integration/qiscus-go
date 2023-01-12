package multichannel

import "time"

// RoomTagsResponse is Represent Get room tags response payload
type RoomTagsResponse struct {
	Data []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

// CreateRoomTagResponse is Represent Create room tag response payload
type CreateRoomTagResponse struct {
	Data struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

// CreateAdditionalInfoRoomResponse is Represents Create additional info room response payload
type CreateAdditionalInfoRoomResponse struct {
	Data struct {
		Extras struct {
			UserProperties []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"user_properties"`
		} `json:"extras"`
		FirstInitiated struct {
			Date         string `json:"date"`
			TimezoneType int    `json:"timezone_type"`
			Timezone     string `json:"timezone"`
		} `json:"first_initiated"`
		FirstAgentResponseTime struct {
			Date         string `json:"date"`
			TimezoneType int    `json:"timezone_type"`
			Timezone     string `json:"timezone"`
		} `json:"first_agent_response_time"`
		UserID string `json:"user_id"`
	} `json:"data"`
}

// GetAdditionalInfoRoomResponse is Represent Get additional info room response payload
type GetAdditionalInfoRoomResponse struct {
	Data struct {
		Extras struct {
			UserProperties []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"user_properties"`
		} `json:"extras"`
		FirstInitiated         time.Time `json:"first_initiated"`
		FirstAgentResponseTime time.Time `json:"first_agent_response_time"`
		UserID                 string    `json:"user_id"`
		ChannelID              int       `json:"channel_id"`
		IsBlocked              bool      `json:"is_blocked"`
		ChannelName            string    `json:"channel_name"`
		Channel                struct {
			ID                  int         `json:"id"`
			AppCode             string      `json:"app_code"`
			SecretKey           string      `json:"secret_key"`
			CreatedAt           string      `json:"created_at"`
			UpdatedAt           string      `json:"updated_at"`
			IsActive            bool        `json:"is_active"`
			AppID               int         `json:"app_id"`
			ForwardURL          interface{} `json:"forward_url"`
			ForwardEnabled      bool        `json:"forward_enabled"`
			Name                string      `json:"name"`
			BadgeURL            string      `json:"badge_url"`
			UseChannelResponder bool        `json:"use_channel_responder"`
		} `json:"channel"`
	} `json:"data"`
}

// SetToggleBotInRoomResponse is Represent Set toggle bot in room response payload
type SetToggleBotInRoomResponse struct {
	Data struct {
		ID                    int         `json:"id"`
		AppID                 int         `json:"app_id"`
		UserID                string      `json:"user_id"`
		RoomID                string      `json:"room_id"`
		Source                string      `json:"source"`
		CreatedAt             string      `json:"created_at"`
		UpdatedAt             string      `json:"updated_at"`
		IsHandledByBot        bool        `json:"is_handled_by_bot"`
		StartServiceCommentID string      `json:"start_service_comment_id"`
		UserAvatarURL         string      `json:"user_avatar_url"`
		Name                  string      `json:"name"`
		HasNoMessage          bool        `json:"has_no_message"`
		Extras                string      `json:"extras"`
		CheckWaContact        bool        `json:"check_wa_contact"`
		Origin                string      `json:"origin"`
		RoomBadge             string      `json:"room_badge"`
		IsWaiting             bool        `json:"is_waiting"`
		SubSource             interface{} `json:"sub_source"`
		ChannelID             int         `json:"channel_id"`
		Resolved              bool        `json:"resolved"`
		ResolvedTs            interface{} `json:"resolved_ts"`
		Type                  interface{} `json:"type"`
		DeletedAt             interface{} `json:"deleted_at"`
		CustomerID            interface{} `json:"customer_id"`
	} `json:"data"`
}

// LoginAdminResponse is Represent Login admin response payload
type LoginAdminResponse struct {
	Data struct {
		User struct {
			ID                  int           `json:"id"`
			Name                string        `json:"name"`
			Email               string        `json:"email"`
			AuthenticationToken string        `json:"authentication_token"`
			CreatedAt           string        `json:"created_at"`
			UpdatedAt           string        `json:"updated_at"`
			SdkEmail            string        `json:"sdk_email"`
			SdkKey              string        `json:"sdk_key"`
			IsAvailable         bool          `json:"is_available"`
			Type                int           `json:"type"`
			AvatarURL           string        `json:"avatar_url"`
			AppID               int           `json:"app_id"`
			IsVerified          bool          `json:"is_verified"`
			NotificationsRoomID string        `json:"notifications_room_id"`
			BubbleColor         interface{}   `json:"bubble_color"`
			QismoKey            string        `json:"qismo_key"`
			DirectLoginToken    interface{}   `json:"direct_login_token"`
			LastLogin           string        `json:"last_login"`
			ForceOffline        bool          `json:"force_offline"`
			DeletedAt           interface{}   `json:"deleted_at"`
			TypeAsString        string        `json:"type_as_string"`
			AssignedRules       []interface{} `json:"assigned_rules"`
			App                 struct {
				ID                             int         `json:"id"`
				Name                           string      `json:"name"`
				AppCode                        string      `json:"app_code"`
				SecretKey                      string      `json:"secret_key"`
				CreatedAt                      string      `json:"created_at"`
				UpdatedAt                      string      `json:"updated_at"`
				BotWebhookURL                  string      `json:"bot_webhook_url"`
				IsBotEnabled                   bool        `json:"is_bot_enabled"`
				AllocateAgentWebhookURL        string      `json:"allocate_agent_webhook_url"`
				IsAllocateAgentWebhookEnabled  bool        `json:"is_allocate_agent_webhook_enabled"`
				MarkAsResolvedWebhookURL       string      `json:"mark_as_resolved_webhook_url"`
				IsMarkAsResolvedWebhookEnabled bool        `json:"is_mark_as_resolved_webhook_enabled"`
				IsMobilePnEnabled              bool        `json:"is_mobile_pn_enabled"`
				IsActive                       bool        `json:"is_active"`
				IsSessional                    bool        `json:"is_sessional"`
				IsAgentAllocationEnabled       bool        `json:"is_agent_allocation_enabled"`
				IsAgentTakeoverEnabled         bool        `json:"is_agent_takeover_enabled"`
				IsTokenExpiring                bool        `json:"is_token_expiring"`
				PaidChannelApproved            interface{} `json:"paid_channel_approved"`
				UseLatest                      bool        `json:"use_latest"`
				AppConfig                      struct {
					ID                     int         `json:"id"`
					AppID                  int         `json:"app_id"`
					Widget                 string      `json:"widget"`
					CreatedAt              string      `json:"created_at"`
					UpdatedAt              string      `json:"updated_at"`
					OfflineMessage         interface{} `json:"offline_message"`
					OnlineMessage          string      `json:"online_message"`
					Timezone               string      `json:"timezone"`
					EnableBulkAssign       bool        `json:"enable_bulk_assign"`
					SendOnlineIfResolved   bool        `json:"send_online_if_resolved"`
					SendOfflineEachMessage bool        `json:"send_offline_each_message"`
				} `json:"app_config"`
				AgentRoles []struct {
					ID            int    `json:"id"`
					AppID         int    `json:"app_id"`
					Name          string `json:"name"`
					IsDefaultRole bool   `json:"is_default_role"`
					CreatedAt     string `json:"created_at"`
					UpdatedAt     string `json:"updated_at"`
				} `json:"agent_roles"`
			} `json:"app"`
		} `json:"user"`
		Details struct {
			IsIntegrated bool `json:"is_integrated"`
			SdkUser      struct {
				ID          int    `json:"id"`
				Token       string `json:"token"`
				Email       string `json:"email"`
				Password    string `json:"password"`
				DisplayName string `json:"display_name"`
				AvatarURL   string `json:"avatar_url"`
				Extras      struct {
					Type            string      `json:"type"`
					UserBubbleColor interface{} `json:"user_bubble_color"`
				} `json:"extras"`
			} `json:"sdk_user"`
			App struct {
				AppCode                        string `json:"app_code"`
				SecretKey                      string `json:"secret_key"`
				Name                           string `json:"name"`
				BotWebhookURL                  string `json:"bot_webhook_url"`
				IsBotEnabled                   bool   `json:"is_bot_enabled"`
				IsAllocateAgentWebhookEnabled  bool   `json:"is_allocate_agent_webhook_enabled"`
				AllocateAgentWebhookURL        string `json:"allocate_agent_webhook_url"`
				MarkAsResolvedWebhookURL       string `json:"mark_as_resolved_webhook_url"`
				IsMarkAsResolvedWebhookEnabled bool   `json:"is_mark_as_resolved_webhook_enabled"`
				IsActive                       bool   `json:"is_active"`
				IsSessional                    bool   `json:"is_sessional"`
				IsAgentAllocationEnabled       bool   `json:"is_agent_allocation_enabled"`
				IsAgentTakeoverEnabled         bool   `json:"is_agent_takeover_enabled"`
				UseLatest                      bool   `json:"use_latest"`
				IsBulkAssignmentEnabled        bool   `json:"is_bulk_assignment_enabled"`
			} `json:"app"`
		} `json:"details"`
		LongLivedToken string `json:"long_lived_token"`
		UserConfigs    struct {
			Notifagentjoining           interface{} `json:"notifagentjoining"`
			IsNotifagentjoiningEnabled  bool        `json:"is_notifagentjoining_enabled"`
			Notifmessagecoming          interface{} `json:"notifmessagecoming"`
			IsNotifmessagecomingEnabled bool        `json:"is_notifmessagecoming_enabled"`
		} `json:"user_configs"`
	} `json:"data"`
}

// GetAllAgentsResponse is Represent Get all agents response payload
type GetAllAgentsResponse struct {
	Data struct {
		Agents []struct {
			AvatarURL            string    `json:"avatar_url"`
			CreatedAt            string    `json:"created_at"`
			CurrentCustomerCount int       `json:"current_customer_count"`
			Email                string    `json:"email"`
			ForceOffline         bool      `json:"force_offline"`
			ID                   int       `json:"id"`
			IsAvailable          bool      `json:"is_available"`
			LastLogin            time.Time `json:"last_login"`
			Name                 string    `json:"name"`
			SdkEmail             string    `json:"sdk_email"`
			SdkKey               string    `json:"sdk_key"`
			Type                 int       `json:"type"`
			TypeAsString         string    `json:"type_as_string"`
			UserChannels         []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"user_channels"`
			UserRoles []interface{} `json:"user_roles"`
		} `json:"agents"`
	} `json:"data"`
	Meta struct {
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
	} `json:"meta"`
	Status int `json:"status"`
}

// AssignAgentResponse is Represent Assign agent response payload
type AssignAgentResponse struct {
	Data struct {
		AddedAgent struct {
			ID                  int           `json:"id"`
			Name                string        `json:"name"`
			Email               string        `json:"email"`
			AuthenticationToken string        `json:"authentication_token"`
			CreatedAt           string        `json:"created_at"`
			UpdatedAt           string        `json:"updated_at"`
			SdkEmail            string        `json:"sdk_email"`
			SdkKey              string        `json:"sdk_key"`
			IsAvailable         bool          `json:"is_available"`
			Type                int           `json:"type"`
			AvatarURL           string        `json:"avatar_url"`
			AppID               int           `json:"app_id"`
			IsVerified          bool          `json:"is_verified"`
			NotificationsRoomID string        `json:"notifications_room_id"`
			BubbleColor         interface{}   `json:"bubble_color"`
			QismoKey            string        `json:"qismo_key"`
			DirectLoginToken    interface{}   `json:"direct_login_token"`
			LastLogin           string        `json:"last_login"`
			ForceOffline        bool          `json:"force_offline"`
			DeletedAt           interface{}   `json:"deleted_at"`
			TypeAsString        string        `json:"type_as_string"`
			AssignedRules       []interface{} `json:"assigned_rules"`
		} `json:"added_agent"`
	} `json:"data"`
}

// GetAgentsByDivisionResponse is Represent get agents by divisions response payload
type GetAgentsByDivisionResponse struct {
	Data []struct {
		AvatarURL            string    `json:"avatar_url"`
		CurrentCustomerCount int       `json:"current_customer_count"`
		Email                string    `json:"email"`
		ForceOffline         bool      `json:"force_offline"`
		ID                   int       `json:"id"`
		IsAvailable          bool      `json:"is_available"`
		LastLogin            time.Time `json:"last_login"`
		Name                 string    `json:"name"`
		SdkEmail             string    `json:"sdk_email"`
		SdkKey               string    `json:"sdk_key"`
		Type                 int       `json:"type"`
		TypeAsString         string    `json:"type_as_string"`
		UserChannels         []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"user_channels"`
		UserRoles []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"user_roles"`
	} `json:"data"`
	Meta struct {
		Limit     int `json:"limit"`
		Page      int `json:"page"`
		Total     int `json:"total"`
		TotalPage int `json:"total_page"`
	} `json:"meta"`
}

// GetAllDivisionResponse is Represent Get all division response payload
type GetAllDivisionResponse struct {
	Data []struct {
		AppID         int    `json:"app_id"`
		CreatedAt     string `json:"created_at"`
		ID            int    `json:"id"`
		IsDefaultRole bool   `json:"is_default_role"`
		Name          string `json:"name"`
		UpdatedAt     string `json:"updated_at"`
	} `json:"data"`
	Meta struct {
		Limit     int `json:"limit"`
		Page      int `json:"page"`
		Total     int `json:"total"`
		TotalPage int `json:"total_page"`
	} `json:"meta"`
}

// MarkAsResolvedResponse is Represent Mark as resolved response payload
type MarkAsResolvedResponse struct {
	Data struct {
		Service struct {
			Notes      string `json:"notes"`
			IsResolved bool   `json:"is_resolved"`
			ResolvedAt struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"resolved_at"`
			UserID      int    `json:"user_id"`
			AppID       int    `json:"app_id"`
			RoomLogID   int    `json:"room_log_id"`
			RoomID      string `json:"room_id"`
			RetrievedAt struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"retrieved_at"`
			FirstCommentID string `json:"first_comment_id"`
			LastCommentID  string `json:"last_comment_id"`
			UpdatedAt      string `json:"updated_at"`
			CreatedAt      string `json:"created_at"`
			ID             int    `json:"id"`
			User           struct {
				ID                  int           `json:"id"`
				Name                string        `json:"name"`
				Email               string        `json:"email"`
				AuthenticationToken string        `json:"authentication_token"`
				CreatedAt           string        `json:"created_at"`
				UpdatedAt           string        `json:"updated_at"`
				SdkEmail            string        `json:"sdk_email"`
				SdkKey              string        `json:"sdk_key"`
				IsAvailable         bool          `json:"is_available"`
				Type                int           `json:"type"`
				AvatarURL           string        `json:"avatar_url"`
				AppID               int           `json:"app_id"`
				IsVerified          bool          `json:"is_verified"`
				NotificationsRoomID string        `json:"notifications_room_id"`
				BubbleColor         interface{}   `json:"bubble_color"`
				QismoKey            string        `json:"qismo_key"`
				DirectLoginToken    interface{}   `json:"direct_login_token"`
				LastLogin           string        `json:"last_login"`
				ForceOffline        bool          `json:"force_offline"`
				DeletedAt           interface{}   `json:"deleted_at"`
				TypeAsString        string        `json:"type_as_string"`
				AssignedRules       []interface{} `json:"assigned_rules"`
			} `json:"user"`
		} `json:"service"`
		RoomInfo struct {
			Room struct {
				RoomAvatarURL string `json:"room_avatar_url"`
				RoomChannelID string `json:"room_channel_id"`
				RoomID        string `json:"room_id"`
				RoomName      string `json:"room_name"`
				RoomOptions   string `json:"room_options"`
				RoomType      string `json:"room_type"`
			} `json:"room"`
		} `json:"room_info"`
	} `json:"data"`
}

// GetAllChannelsResponse is Represent Get all channels response payload
type GetAllChannelsResponse struct {
	Data struct {
		CustomChannels []struct {
			ID                  int    `json:"id"`
			WebhookURL          string `json:"webhook_url"`
			LogoURL             string `json:"logo_url"`
			IdentifierKey       string `json:"identifier_key"`
			Name                string `json:"name"`
			IsActive            bool   `json:"is_active"`
			UseChannelResponder bool   `json:"use_channel_responder"`
		} `json:"custom_channels"`
		FbChannels []struct {
			ID                  int    `json:"id"`
			IsActive            bool   `json:"is_active"`
			LongLivedToken      string `json:"long_lived_token"`
			ProfileName         string `json:"profile_name"`
			PageID              string `json:"page_id"`
			BadgeURL            string `json:"badge_url"`
			AppID               int    `json:"app_id"`
			UseChannelResponder bool   `json:"use_channel_responder"`
		} `json:"fb_channels"`
		IgChannels []struct {
			ID                  int    `json:"id"`
			IsActive            bool   `json:"is_active"`
			LongLivedToken      string `json:"long_lived_token"`
			Name                string `json:"name"`
			PageID              string `json:"page_id"`
			BadgeURL            string `json:"badge_url"`
			AppID               int    `json:"app_id"`
			UseChannelResponder bool   `json:"use_channel_responder"`
			IgID                string `json:"ig_id"`
		} `json:"ig_channels"`
		LineChannels []struct {
			ID                  int    `json:"id"`
			IsActive            bool   `json:"is_active"`
			AccessToken         string `json:"access_token"`
			SecretKey           string `json:"secret_key"`
			BadgeURL            string `json:"badge_url"`
			Name                string `json:"name"`
			AppID               int    `json:"app_id"`
			UseChannelResponder bool   `json:"use_channel_responder"`
		} `json:"line_channels"`
		QiscusChannels []struct {
			ID                  int         `json:"id"`
			IsActive            bool        `json:"is_active"`
			AppCode             string      `json:"app_code"`
			SecretKey           string      `json:"secret_key"`
			Name                string      `json:"name"`
			BadgeURL            interface{} `json:"badge_url"`
			AppID               int         `json:"app_id"`
			UseChannelResponder bool        `json:"use_channel_responder"`
		} `json:"qiscus_channels"`
		TelegramChannels []struct {
			ID                  int         `json:"id"`
			IsActive            bool        `json:"is_active"`
			Name                string      `json:"name"`
			Username            string      `json:"username"`
			BotToken            string      `json:"bot_token"`
			BadgeURL            interface{} `json:"badge_url"`
			UseChannelResponder bool        `json:"use_channel_responder"`
			AppID               int         `json:"app_id"`
		} `json:"telegram_channels"`
		WaChannels []struct {
			AllowIntlHsm               bool        `json:"allow_intl_hsm"`
			AppID                      int         `json:"app_id"`
			BadgeURL                   string      `json:"badge_url"`
			BaseURL                    string      `json:"base_url"`
			BusinessID                 interface{} `json:"business_id"`
			BusinessVerificationStatus interface{} `json:"business_verification_status"`
			CreatedAt                  string      `json:"created_at"`
			EncodedToken               string      `json:"encoded_token"`
			ForwardEnabled             bool        `json:"forward_enabled"`
			ForwardURL                 interface{} `json:"forward_url"`
			Hsm24Enabled               bool        `json:"hsm_24_enabled"`
			ID                         int         `json:"id"`
			IsActive                   bool        `json:"is_active"`
			IsSslEnabled               bool        `json:"is_ssl_enabled"`
			Name                       string      `json:"name"`
			OnSync                     bool        `json:"on_sync"`
			PhoneNumber                string      `json:"phone_number"`
			PhoneNumberStatus          interface{} `json:"phone_number_status"`
			Platform                   string      `json:"platform"`
			ReadEnabled                bool        `json:"read_enabled"`
			UpdatedAt                  string      `json:"updated_at"`
			UseChannelResponder        bool        `json:"use_channel_responder"`
		} `json:"wa_channels"`
	} `json:"data"`
}

// GetRoomByRoomIDResponse is Represent Get room by room id response payload
type GetRoomByRoomIDResponse struct {
	Data struct {
		CustomerRoom struct {
			ChannelID               int         `json:"channel_id"`
			ContactID               interface{} `json:"contact_id"`
			ID                      int         `json:"id"`
			IsHandledByBot          bool        `json:"is_handled_by_bot"`
			IsResolved              bool        `json:"is_resolved"`
			IsWaiting               bool        `json:"is_waiting"`
			LastCommentSender       string      `json:"last_comment_sender"`
			LastCommentSenderType   string      `json:"last_comment_sender_type"`
			LastCommentText         string      `json:"last_comment_text"`
			LastCommentTimestamp    time.Time   `json:"last_comment_timestamp"`
			LastCustomerCommentText interface{} `json:"last_customer_comment_text"`
			LastCustomerTimestamp   time.Time   `json:"last_customer_timestamp"`
			Name                    string      `json:"name"`
			RoomBadge               string      `json:"room_badge"`
			RoomID                  string      `json:"room_id"`
			RoomType                string      `json:"room_type"`
			Source                  string      `json:"source"`
			UserAvatarURL           string      `json:"user_avatar_url"`
			UserID                  string      `json:"user_id"`
		} `json:"customer_room"`
	} `json:"data"`
	Status int `json:"status"`
}
