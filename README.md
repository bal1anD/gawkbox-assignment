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
                            "name":"dallas", 
                            "display_name":"dallas", 
                            "bio":"Just a gamer playing games and chatting. :)", 
                            "created_at":1370286722580,  
                            "_id":"44322889"
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
                        "views":6497,
                        "followers":108,
                        "game":"Battle Chasers: Nightwar", 
                        "streamingStatus":true,
                        "language":"en",  
                        "_id":"44322889"
                       }

Error codes:  404 - if the username (hence, the channel) is not found,
              500 - for other internal server errors.

Your instructions go here.

INSTRUCTIONS
========================================================================================================================

The app can be run as a docker container or it can be run without docker.

1 )Running on Docker
-----------------

Note: These steps were performed with Docker on Mac.

1.1) Get the source code from github https://github.com/bal1anD/gawkbox-assignment , by running the following
command:
    go get github.com/bal1anD/gawkbox-assignment
    
     OR
     
    git clone https://github.com/bal1anD/gawkbox-assignment
     
    Note: use the "master" branch 

1.2) Navigate to the location where go downloads the code from step #1

Example: 
    cd $HOME/go/src/github.com/bal1anD/gawkbox-assignment
    
    Note: My $GOPATH=$HOME/go

1.3) There is a Dockerfile. Now we need to build this file by running the following:
    docker build -t <docker_image_name> .
    
1.4) After step#3 finishes successfully you can run the docker image using the following command

  docker run -p 8080:8080 <docker_image_name> ./gawkbox-assignment
    
you will see that the service is started at port 8080. 


2) Running without Docker
----------------------

2.1) The first two steps are same as 1.1 and 1.2 above.

2.2) Once in the root folder (after following step# 1.2 above), build the code using the following
   
   go build

2.3) Run the service by running the following command

   go run main.go
   
    
This will start up the service at port 8080





