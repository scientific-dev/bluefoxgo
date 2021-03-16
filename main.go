package bluefox

import (
	"errors"
)

// Base url of the api
var BaseURL string = "https://panel.bluefoxhost.com/api"

// Power actions of the server
var PowerActions []string = []string{"start", "stop", "restart", "kill"}

// Verifies if the power action is invalid or not
func isInvalidPowerAction(action string) bool {
	for _, act := range PowerActions {
		if act == action {
			return false
		}
	}

	return true
}

// Basic bluefox client structure
type Client struct {
	// Your blueofx auth token
	Token string
}

// Creates a new client with default options!
func NewClient(token string) Client {

	return Client{token}

}

// Returns current user profile information
func (self Client) GetProfile() (User, error) {

	var user User
	err := Fetch("GET", "/client/account", &user, self.Token, map[string]string{})
	return user, err

}

// Returns a bluefox server information by id
func (self Client) GetServer(ID string) (Server, error) {

	var server BluefoxServer
	err := Fetch("GET", "/client/servers/"+ID, &server, self.Token, map[string]string{})
	server.Attributes.Client = self
	return server.Attributes, err

}

// Returns all the servers of the current user
func (self Client) GetServers() ([]Server, error) {

	var servers struct {
		Data []BluefoxServer `json:"data"`
	}

	err := Fetch("GET", "/client", &servers, self.Token, map[string]string{})

	if err != nil {
		return []Server{}, err
	}

	var result []Server

	for _, server := range servers.Data {
		server.Attributes.Client = self
		result = append(result, server.Attributes)
	}

	return result, err

}

// Set power mode of the server!
func (self Client) SetPowerMode(ID string, action string) error {

	if isInvalidPowerAction(action) {
		return errors.New("Invalid action provided! Must be one of 'start', 'stop', 'restart' or 'kill'!")
	}

	_, err := self.GetServer(ID)

	if err != nil {
		return errors.New("Unknown server: " + err.Error())
	}

	var a interface{}
	posterr := Fetch("POST", "/client/servers/"+ID+"/power", &a, self.Token, map[string]string{
		"signal": action,
	})

	return posterr

}

// Sets server name
func (self Client) SetServerName(ID string, name string) error {

	var a interface{}
	err := Fetch("POST", "/client/servers/"+ID+"/settings/rename", &a, self.Token, map[string]string{
		"name": name,
	})
	return err

}

// Sends command to the server
func (self Client) SendCommand(ID string, command string) error {

	var a interface{}
	err := Fetch("POST", "/client/servers/"+ID+"/command", &a, self.Token, map[string]string{
		"command": command,
	})
	return err

}

// Reinstalls the server
func (self Client) ReinstallServer(ID string) error {

	var a interface{}
	err := Fetch("POST", "/client/servers/"+ID+"/settings/reinstall", &a, self.Token, map[string]string{})
	return err

}

// Deletes the server
func (self Client) DeleteServer(ID string, force bool) error {

	forcestr := ""

	if force {
		forcestr = "force/"
	}

	var a interface{}
	err := Fetch("POST", "/applications/servers/"+ID+"/"+forcestr+"/delete", &a, self.Token, map[string]string{})
	return err

}
