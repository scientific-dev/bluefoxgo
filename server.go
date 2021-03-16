package bluefox

// Bluefox server object
type Server struct {
	// Boolean stating is the current user owner or not
	Owner bool `json:"server_owner"`
	// ID of the server
	ID string `json:"identifier"`
	// Internal id of ther server
	InternalID string `json:"internal_id"`
	// UUId of the server
	UUID string `json:"uuid"`
	// Name of the server
	Name string `json:"name"`
	// Sftp details of the server
	Sftp struct {
		// IP of the server
		IP string `json:"ip"`
		// Port of the server
		Port int `json:"port"`
	} `json:"sftp_details"`
	// Limits of the server
	Limits struct {
		// Memory limit of the server
		Memory int `json:"memory"`
		// Swap limit of the server
		Swap int `json:"swap"`
		// Disk limit of the server
		Disk int `json:"disk"`
		// IO limit of the server
		IO int `json:"io"`
		// CPU limit of the server
		CPU int `json:"cpu"`
	} `json:"limits"`
	// Invocations of the server
	Invocation string `json:"invocation"`
	// Docker image url of the server
	DockerImage string `json:"docker_image"`
	// Eggfeatures of the server
	EggFeatures interface{} `json:"egg_features"`
	// Featured limits of the server
	FeatureLimits struct {
		// Database featured limit of the server
		Databases int64 `json:"databases"`
		// Allocation featured limit of the server
		Allocations int64 `json:"allocations"`
		// Featured backup limit of the server
		Backups int64 `json:"backups"`
	} `json:"feature_limits"`
	// Boolean stating is the server suspended or not
	Supended bool `json:"is_suspended"`
	// Boolean stating is the server installed or not
	Installed bool `json:"is_installed"`
	// Boolean stating is the server transferring or not
	Transferring bool `json:"is_tranferring"`
	// Meta data of the server
	Meta struct {
		// Boolean stating is the user server owner or not
		ServerOwner bool `json:"is_server_owner"`
		// Permissions of the current user in the server
		Permissions []string `json:"user_permissions"`
	} `json:"meta"`
	// Relationships of the server
	Relationships struct {
		// Allocation relationships of the server
		Allocations RelationshipData `json:"allocations"`
		// Variable relationships of the server
		Variables RelationshipData `json:"variables"`
	} `json:"relationships"`
	// Your bluefox client
	Client Client `json:"-"`
}

// Starts your server
func (self Server) Start() {
	self.SetPowerMode("start")
}

// Stops your server
func (self Server) Stop() {
	self.SetPowerMode("stop")
}

// Retstarts your server
func (self Server) Restart() {
	self.SetPowerMode("restart")
}

// Kills your server
func (self Server) Kill() {
	self.SetPowerMode("kill")
}

// Set power mode of the current server
func (self Server) SetPowerMode(action string) error {
	return self.Client.SetPowerMode(self.ID, action)
}

// Sets server name
func (self Server) SetName(name string) error {
	err := self.Client.SetServerName(self.ID, name)

	if err != nil {
		self.Name = name
	}

	return err
}

// Sends a command to the server
func (self Server) Send(command string) error {
	return self.Client.SendCommand(self.ID, command)
}

// Reinstalls the server
func (self Server) Reinstall() error {
	return self.Client.ReinstallServer(self.ID)
}

// Deletes the server
func (self Server) Delete(force bool) error {
	return self.Client.DeleteServer(self.ID, force)
}
