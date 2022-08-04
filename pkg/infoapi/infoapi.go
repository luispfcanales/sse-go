package infoapi

import "fmt"

const (
	userInfo    = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	userEmail   = "https://www.googleapis.com/auth/userinfo.email"
	userProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

//GetAPIGoogleUserInfo return url api userinfo
func GetAPIGoogleUserInfo(token string) string {
	return fmt.Sprintf("%s%s", userInfo, token)
}

//GetAPIGoogleEmail return url api oauth2 userEmail
func GetAPIGoogleEmail() string {
	return userEmail
}

//GetAPIGoogleProfile return url api oauth2 userEmail
func GetAPIGoogleProfile() string {
	return userProfile
}
