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
//     gsbot [username] [-p password] [-a authcode] [-t twofactorcode] [-l loginkey]

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Flo4604/go-steam/v4"
	"github.com/Flo4604/go-steam/v4/csgo"
	"github.com/Flo4604/go-steam/v4/protocol/steamlang"
)

// const usage string = "usage: example [username] [-p password] [-a authcode] [-t twofactorcode] [-l loginkey] [-anon]"

func main() {
	details := &steam.LogOnDetails{
		Username:               "",
		Password:               "",
		ShouldRememberPassword: true,
		Anonymous:              false,
	}

	// for i := 2; i < len(os.Args)-1; i += 2 {
	// 	switch os.Args[i] {
	// 	case "-p":
	// 		details.Password = os.Args[i+1]
	// 	case "-a":
	// 		details.AuthCode = os.Args[i+1]
	// 	case "-t":
	// 		details.TwoFactorCode = os.Args[i+1]
	// 	case "-l":
	// 		details.LoginKey = os.Args[i+1]
	// 	case "-anon":
	// 		details.Anonymous = true
	// 	default:
	// 		fmt.Println(usage)
	// 		return
	// 	}
	// }

	client := steam.NewClient(&steam.ClientOptions{
		Debug: &steam.DebugOptions{Enabled: true, Base: "debug"},
		App:   &steam.AppOptions{EnablePicsCache: false, ChangelistUpdateInterval: 15, PicsCacheAll: false},
	})

	client.Connect()
	cs := csgo.New(client)

	for event := range client.Events() {

		switch e := event.(type) {
		case *steam.ConnectedEvent:
			client.Auth.LogOn(details)
		case *steam.LoggedOnEvent:
			fmt.Println("Logged on!")
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)

			cs.SetPlaying(true)
		case *steam.SteamGuardEvent:
			// create input reader
			reader := bufio.NewReader(os.Stdin)

			// print out if we want the email code or the mobile code by checking if there is a domain in the event
			if e.Domain != "" {
				fmt.Printf("Enter email code ending in %s: ", e.Domain)
			} else {
				fmt.Printf("Enter mobile code: ")
			}

			// read the code from the console
			code, _ := reader.ReadString('\n')

			// set AuthCode if email and two factor code if mobile
			if e.Domain != "" {
				details.AuthCode = code[:5]
			} else {
				details.TwoFactorCode = code[:5]
			}

			// print AuthCode and TwoFactorCode
			fmt.Printf("AuthCode: %s\nTwoFactorCode: %s\n", details.AuthCode, details.TwoFactorCode)

			// log on again
			client.Connect()
		case *csgo.GCReadyEvent:
			println("GC is ready!")

			for i := 0; i < 3; i++ {
				go func(ii int) {
					response, err := cs.GetPlayerProfile(845829334)
					if err != nil {
						fmt.Println("Error:", err, fmt.Sprintf("i: %d", ii))
					} else {
						fmt.Println("Received response:", response, fmt.Sprintf("i: %d", ii))
					}
				}(i)
			}
		case *csgo.PlayerProfileEvent:
			fmt.Printf("!!!!!!!!!!Player profile: %+v\n", e)
		case error:
			fmt.Printf("Error: %v", e)
		default:
			fmt.Printf("Event: %v", e)
		}
	}
}
