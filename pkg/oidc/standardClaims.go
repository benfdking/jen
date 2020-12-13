package oidc

// StandardClaims returns a map of standard claims as defined by https://openid.net/specs/openid-connect-core-1_0.html.
func StandardClaims() map[string]string {
	return map[string]string{
		"sub":                "90e5763f-7088-426f-941a-984dda3248c8",
		"name":               "Dr. Michael Smith",
		"given_name":         "Michael",
		"middle_name:":       "",
		"family_name":        "Smith",
		"nickname":           "Mike",
		"preferred_username": "m.smith",
		"email":              "mikesmith@example.com",
		"website":            "example.com",
		"profile":            "example.com/Michael.jpg",
		"gender":             "male",
		"birthdate":          "1950-12-12",
		"zoneinfo":           "America/Los_Angeles",
		"locale":             "en-US",
		"updated_at":         "1970-01-01T0:0:0Z",
	}
}
