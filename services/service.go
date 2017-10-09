package services

import (
	"gawkbox-assignment/twitch"
	"gawkbox-assignment/store"

	"reflect"
	"gawkbox-assignment/models"
	"fmt"
)

func GetUserInfo(username string) models.User {
	/*
		Note: here, it's not attempting to get the userinfo
		from the local cache because it might result in
		stale info because the user can update their
		bio/display_name at anytime using the twitch UI
	 */
	userInfo,_ := twitch.GetUserInfo(username)
	if !reflect.DeepEqual(models.User{}, userInfo) {
		store.Put(username,userInfo)
	}
	return userInfo
}


func GetChannelViews(username string ) models.Channel {
	var channel models.Channel = models.Channel{}

	/*
		Note: here, first the cache is checked to
		see if the userInfo is for the username is
		present or not. If it's present then the userId is
		obtained from the local cache and is used to
		get the channel info from twitch. This saves us
		1 external call to twitch. If it's absent then
		we call twitch to get the userId.
	 */
	user := store.Get(username);
	var userId string
	if !reflect.DeepEqual(models.User{}, user) {
		fmt.Println("Cache hit")
		userId = user.Id
	} else {
		fmt.Println("Cache miss")
		user,_:= twitch.GetUserInfo(username)
		if !reflect.DeepEqual(models.User{}, user) {
			store.Put(username,user)
		} else {
			return channel
		}
		userId = user.Id
	}

	channel,_ = twitch.GetUserChannelInfo(userId)
	return  channel
}

