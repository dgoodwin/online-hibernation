package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_Group = map[string]string{
	"":         "Group represents a referenceable set of Users",
	"metadata": "Standard object's metadata.",
	"users":    "Users is the list of users in this group.",
}

func (Group) SwaggerDoc() map[string]string {
	return map_Group
}

var map_GroupList = map[string]string{
	"":         "GroupList is a collection of Groups",
	"metadata": "Standard object's metadata.",
	"items":    "Items is the list of groups",
}

func (GroupList) SwaggerDoc() map[string]string {
	return map_GroupList
}

var map_Identity = map[string]string{
	"":                 "Identity records a successful authentication of a user with an identity provider. The information about the source of authentication is stored on the identity, and the identity is then associated with a single user object. Multiple identities can reference a single user. Information retrieved from the authentication provider is stored in the extra field using a schema determined by the provider.",
	"metadata":         "Standard object's metadata.",
	"providerName":     "ProviderName is the source of identity information",
	"providerUserName": "ProviderUserName uniquely represents this identity in the scope of the provider",
	"user":             "User is a reference to the user this identity is associated with Both Name and UID must be set",
	"extra":            "Extra holds extra information about this identity",
}

func (Identity) SwaggerDoc() map[string]string {
	return map_Identity
}

var map_IdentityList = map[string]string{
	"":         "IdentityList is a collection of Identities",
	"metadata": "Standard object's metadata.",
	"items":    "Items is the list of identities",
}

func (IdentityList) SwaggerDoc() map[string]string {
	return map_IdentityList
}

var map_User = map[string]string{
	"":           "Upon log in, every user of the system receives a User and Identity resource. Administrators may directly manipulate the attributes of the users for their own tracking, or set groups via the API. The user name is unique and is chosen based on the value provided by the identity provider - if a user already exists with the incoming name, the user name may have a number appended to it depending on the configuration of the system.",
	"metadata":   "Standard object's metadata.",
	"fullName":   "FullName is the full name of user",
	"identities": "Identities are the identities associated with this user",
	"groups":     "Groups are the groups that this user is a member of",
}

func (User) SwaggerDoc() map[string]string {
	return map_User
}

var map_UserIdentityMapping = map[string]string{
	"":         "UserIdentityMapping maps a user to an identity",
	"metadata": "Standard object's metadata.",
	"identity": "Identity is a reference to an identity",
	"user":     "User is a reference to a user",
}

func (UserIdentityMapping) SwaggerDoc() map[string]string {
	return map_UserIdentityMapping
}

var map_UserList = map[string]string{
	"":         "UserList is a collection of Users",
	"metadata": "Standard object's metadata.",
	"items":    "Items is the list of users",
}

func (UserList) SwaggerDoc() map[string]string {
	return map_UserList
}
