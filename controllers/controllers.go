package controllers

import (
	"net/http"
	"fmt"
	"gawkbox-assignment/services"
	"gawkbox-assignment/models"
	"encoding/json"
	"reflect"
)


/*
	Controller for retrieving the User information
	@Param: username from the reaource pathParam
	@Response: A json containing name, bio, display_name,
	          account creation date and Id

 */
func GetUserInfoController(w http.ResponseWriter, username string) {
	fmt.Println("Getting userInfo for userName ", username)
	var userInfo = services.GetUserInfo(username)
	if reflect.DeepEqual(models.User{}, userInfo) {
		w.WriteHeader(http.StatusNotFound)
		var e = fmt.Errorf("Displayname not found")
		prepareResp(nil, e, w)
	} else {
		jsonResp, err := json.Marshal(userInfo)
		prepareResp(jsonResp, err, w)
	}
}

/*
	Controller for retrieving a User's channel's information
	@Param: username from the reaource pathParam
	@Response: A json containing channel_id, view, followers, Game,
		  streaming_status and language

 */
func GetChannelInfoController(w http.ResponseWriter, username string) {
	fmt.Println("Getting channel views ", username)
	channel := services.GetChannelViews(username)
	fmt.Println(channel)
	if reflect.DeepEqual(models.Channel{}, channel) {
		w.WriteHeader(http.StatusNotFound)
		var e = fmt.Errorf("User not found")
		prepareResp(nil, e, w)
	} else {
		jsonResp, err := json.Marshal(channel)
		prepareResp(jsonResp, err, w)
	}
}

/*
	Controller for responding to paths which do not exist
	@Response: Returns a message informing the client about the supported paths
 */
func DefaultPathController(w http.ResponseWriter) {
	allowedPaths := "{ message: Path not supported, " +
		" supported path: [{User info: /users/{username}/info}," +
		" {Channel info: /users/{username}/channel}]}"
	jsonResp, err := json.Marshal(allowedPaths)
	prepareResp(jsonResp, err, w)

}

func prepareResp( b []byte , err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([] byte(b))
}

