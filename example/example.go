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
	"fmt"
	"os"

	"github.com/Flo4604/go-steam/v3"
	"github.com/Flo4604/go-steam/v3/protocol/steamlang"
)

const usage string = "usage: example [username] [-p password] [-a authcode] [-t twofactorcode] [-l loginkey] [-anon]"

func main() {
	details := &steam.LogOnDetails{
		Username:               os.Args[1],
		ShouldRememberPassword: false,
		Anonymous:              true,
	}

	for i := 2; i < len(os.Args)-1; i += 2 {
		switch os.Args[i] {
		case "-p":
			details.Password = os.Args[i+1]
		case "-a":
			details.AuthCode = os.Args[i+1]
		case "-t":
			details.TwoFactorCode = os.Args[i+1]
		case "-l":
			details.LoginKey = os.Args[i+1]
		case "-anon":
			details.Anonymous = true
		default:
			fmt.Println(usage)
			return
		}
	}

	client := steam.NewClient(&steam.ClientOptions{
		Debug: &steam.DebugOptions{Enabled: true, Base: "debug"},
		App:   &steam.AppOptions{EnablePicsCache: true, ChangelistUpdateInterval: 15, PicsCacheAll: true},
	})

	client.Connect()

	for event := range client.Events() {

		switch e := event.(type) {
		case error:
			fmt.Printf("Error: %v", e)
		case *steam.LoggedOnEvent:
			fmt.Println("Logged on!")
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
		case *steam.ConnectedEvent:
			client.Auth.LogOn(&steam.LogOnDetails{
				Anonymous: true,
			})
		default:
			fmt.Printf("Event: %v", e)
		}
	}
}
