# Setup

## Requirements

1. go1.14

## How to run

1. Clone and setup the project

2. >> go run main.go or >> go build >> ./m

3. open localhost:8000 to two browser tabs

4. Click on New Chat and Video Conferencing should start immediately

## How to do video chat over internet

1. deploy it to server or use something like ngrok.com to expose your localhost to internet.

2. Share the ngrok url to someone who you want to connect.

3. Click on New Chat and Video Conferencing should start immediately

## In case of failure

You might need to setup STUN and TURN server. There are many free STUN server you can user them. As for TURN server you might not get those very easily for free. User https://numb.viagenie.ca/ they are free but need signup.

### Note: Put STUN and TURN server config inside pcConfig variable
