package main

import (
	"bujem/gateway/access"
	"encoding/json"
)

func main() {
	access.CreateUser(make(map[string]*json.RawMessage))
}
