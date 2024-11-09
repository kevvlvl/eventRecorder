package main

import "eventRecorder/server"

func main() {

	api := new(server.Server)

	err := api.Echo()
	if err != nil {
		return
	}
}
