## Remote Assignment V3

API DESCRIPTION
========================================================================================================================
Currently supports the following endpoints:

1)  /users/{username}/info :-
----------------------------

Description: This endpoint returns the user related information such as , id, display name, account creation date
             name, bio. These information are higly likely to be used together, so they have been combined into a
             single endpoint so that a client can get these in one GET call instead of making separate calls for
             accessing each of these information.

Example CURL Request : curl http://localhost:8080/users/dallas/info

Sample Output (JSON): {
                            "name":"dallas", # name of the user <String>
                            "display_name":"dallas", # display name of the user <String>
                            "bio":"Just a gamer playing games and chatting. :)", # bio of the user <String>
                            "created_at":1370286722580,  # timestamp (since Epoch) of account creation in millisecond <Number>
                            "_id":"44322889" # Id of the user <String>
                      }

Error codes:  404 - if the username is not found,
              500 - for other internal server errors.



2)  /users/{username}/channel :-
----------------------------

Description: This endpoint returns the user's channel related information such as , id, # of views, # of followers,
             game, status of whether the user is streaming or not and language. These information are higly likely
             to be used together, so they have been combined into a single endpoint so that a client can get these
             in one GET call instead of making separate calls for accessing each of these information.

Example CURL Request : curl http://localhost:8080/users/dallas/channel

Sample Output (JSON): {
                        "views":6497, # number of views <Number>
                        "followers":108, # number of followers <Number>
                        "game":"Battle Chasers: Nightwar", # Name of the game <String>
                        "streamingStatus":true, # Whether the user is streaming or not <Boolean>
                        "language":"en",  # language <String>
                        "_id":"44322889" # Id of the user (which also the id of the channel) <String>
                       }

Error codes:  404 - if the username (hence, the channel) is not found,
              500 - for other internal server errors.

Your instructions go here.

INSTRUCTIONS
========================================================================================================================

The app can be run as a docker container or it can be run without docker.

Running on Docker
-----------------


