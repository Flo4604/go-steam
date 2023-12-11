// A simple example that uses the modules from the gsbot package and go-steam to log on
// to the Steam network.
//
// Use the right options for your account settings:
// Normal login: username + password
//
// Email code:   username + password
//               username + password + authcode
//
// Mobile code:  username + password + twofactorcode
//               username + loginkey
//
//     example [username] [-p password] [-a authcode] [-t twofactorcode] [-l loginkey]

package main

import (
	"context"
	"time"

	"github.com/Flo4604/go-steam/v4/session"
)

// const usage string = "usage: example [username] [-p password] [-a authcode] [-t twofactorcode] [-l loginkey] [-anon]"

// func main() {
// 	details := &steam.LogOnDetails{
// 		Username:               "",
// 		Password:               "",
// 		ShouldRememberPassword: true,
// 		Anonymous:              false,
// 	}

// 	// for i := 2; i < len(os.Args)-1; i += 2 {
// 	// 	switch os.Args[i] {
// 	// 	case "-p":
// 	// 		details.Password = os.Args[i+1]
// 	// 	case "-a":
// 	// 		details.AuthCode = os.Args[i+1]
// 	// 	case "-t":
// 	// 		details.TwoFactorCode = os.Args[i+1]
// 	// 	case "-l":
// 	// 		details.LoginKey = os.Args[i+1]
// 	// 	case "-anon":
// 	// 		details.Anonymous = true
// 	// 	default:
// 	// 		fmt.Println(usage)
// 	// 		return
// 	// 	}
// 	// }

// 	// cs := csgo.New(client)

// 	for event := range client.Events() {

// 		switch e := event.(type) {
// 		case *steam.ConnectedEvent:
// 			client.Auth.LogOn(details)
// 			client.Auth.GetRSAKey("thaumic_horizons")
// 		case *steam.LoggedOnEvent:
// 			fmt.Println("Logged on!")
// 			client.Social.SetPersonaState(steamlang.EPersonaState_Online)

// 			// cs.SetPlaying(true)

// 			println("GC is ready!")

// 			// for i := 0; i < 3; i++ {
// 			// 	go func(ii int) {
// 			// 		response, err := cs.GetPlayerProfile(845829334)
// 			// 		if err != nil {
// 			// 			fmt.Println("Error:", err, fmt.Sprintf("i: %d", ii))
// 			// 		} else {
// 			// 			fmt.Println("Received response:", response, fmt.Sprintf("i: %d", ii))
// 			// 		}
// 			// 	}(i)
// 			// }
// 		case *csgo.PlayerProfileEvent:
// 			fmt.Printf("!!!!!!!!!!Player profile: %+v\n", e)
// 		case error:
// 			fmt.Printf("Error: %v", e)
// 		default:
// 			fmt.Printf("Event: %v", e)
// 		}
// 	}
// }

func main() {
	s := session.New(&session.Options{
		PlatformType: "k_EAuthTokenPlatformType_SteamClient", // TODO: make this a constant
		Proxy:        "",
		Ctx:          context.Background(),
	})

	time.Sleep(5 * time.Second)

	s.StartWithCredentials(&session.CredentialOptions{AccountName: "", Password: ""})

	// dont stop
	for {
	}
}
