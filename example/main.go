package main

import (
	"fmt"

	"github.com/Qiscus-Integration/qiscus-go"

	"github.com/Qiscus-Integration/qiscus-go/multichannel"
	"github.com/Qiscus-Integration/qiscus-go/sdk"
)

func main() {
	qiscus.DefaultHttpOutboundLog = true

	// Initiate client for Multichannel
	multichannelClient := multichannel.NewMultichannel("qiscus-app-id", "qiscus-secret-key")

	// Initiate client for Multichannel using credential email and password admin
	// multichannelClient, err := multichannel.NewMultichannelFromCredential("example@mail.com", "12345678")
	// if err != nil {
	// 	panic(err)
	// }

	// Initiate client for Multichannel using environment variable
	// QISCUS_APP_ID, QISCUS_SECRET_KEY and MULTICHANNEL_API_BASE
	// multichannelClient, err := multichannel.NewMultichannelFromEnv()
	// if err != nil {
	// 	panic(err)
	// }

	// Default Multichannel base is https://multichannel.qiscus.com, use SetAPIBase() to override.
	// multichannelClient.SetAPIBase("https://multichannel2.qiscus.com")

	// Sample Multichannel method
	resp, _ := multichannelClient.CreateRoomTag(&multichannel.CreateRoomTagReq{
		RoomID: "12345678",
		Tag:    "test",
	})
	fmt.Println(resp)

	// Initiate client for SDK
	sdkClient := sdk.NewSDK("qiscus-app-id", "qiscus-secret-key")

	// Initiate client for SDK using environment variable
	// QISCUS_APP_ID, QISCUS_SECRET_KEY and QISCUS_API_BASE
	// sdkClient, err := sdk.NewSDKFromEnv()
	// if err != nil {
	// 	panic(err)
	// }

	// Default SDK base is https://api.qiscus.com, use SetAPIBase() to override.
	// sdkClient.SetAPIBase("https://api2.qiscus.com")

	// Sample SDK method
	resp2, _ := sdkClient.LoginOrRegister(&sdk.LoginOrRegisterReq{
		UserID:   "guest2@qiscus.com",
		Password: "12345678",
		Username: "User Demo",
	})
	fmt.Println(resp2)
}
