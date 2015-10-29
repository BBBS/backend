package variables

const (
	// APIPathAccountServer is the path to the
	APIPathAccountServer = "/account/:id"
	// APIPathAccount is the path to the
	APIPathAccount = "/api" + APIPathAccountServer
	// APIPathUserServer is the path to a users settings
	APIPathUserServer = "/user/:id"
	// APIPathUser is the path to a users settings
	APIPathUser = "/api" + APIPathUserServer
	// APIPathUserSettingsServer is the path to a users settings
	APIPathUserSettingsServer = "/user/settings/:id"
	// APIPathUserSettings is the path to a users settings
	APIPathUserSettings = "/api" + APIPathUserSettingsServer
)

var (
	// BlankResponse is so that we can send something without an EOF
	BlankResponse = []byte("{}")
)
