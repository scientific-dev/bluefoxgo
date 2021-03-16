package bluefox

// Bluefox user object
type User struct {
	// Type of object sent as response!
	ObjectType string `json:"object"`
	// Attributes of the object
	Attributes struct {
		// Id of the bluefox user
		ID int64 `json:"id"`
		// Boolean stating is the user admin or not
		Admin bool `json:"admin"`
		// Username of the bluefox user
		Username string `json:"username"`
		// Email of the bluefox user
		Email string `json:"email"`
		// First name of the bluefox user
		FirstName string `json:"first_name"`
		// Last name of the bluefox user
		LastName string `json:"last_name"`
		// Language of the bluefox user
		Language string `json:"language"`
	} `json:"attributes"`
}

// Bluefox server api object
type BluefoxServer struct {
	// Type of object sent as response!
	ObjectType string `json:"object"`
	// Attributes of the object
	Attributes Server `json:"attributes"`
}

// Relation ship data of bluefox server object!
type RelationshipData struct {
	// Type of object sent as response!
	ObjectType string `json:"object"`
	// Data of relationships
	Data []interface{} `json:"data"`
}
