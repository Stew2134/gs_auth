# gs_auth

2025-07-24 Stuart McNally

Access token parser for use with google sheet authentication (Accessing private google sheets via the terminal)

Please note the scope available in this authorizer is read only and as of right now there is no intention to add additional scopes.

for use with Alpine Linux.

/////////////////////////////
example usage in terminal:
/////////////////////////////

gs_auth ./application_credentials.json

outputs access token for use in Authorization headers

How to enable google sheets api on google cloud:

Step 1:

Create service account on google cloud for the app used to interact with the google sheet

Step 2:

Go into the service account and under the header "Keys" click generate key and download in json format

Step 3:

Go to the details page of the service account and retrieve the email

Step 4:

Click share on the private google sheet and enter the email retrieved from step 3

Step 5:

Go to API's on google cloud and click enable API's

Step 6:

Search for google sheets API and click enable

Step 7:

using the json file downloaded in step 2 run the following curl:

curl -H "Authorization: Bearer $(./gs_auth ./<application_credentials>.json)" "https://sheets.googleapis.com/v4/spreadsheets/{spreadsheetId}/values/Sheet1"

Step 8:

Have fun writing your app using a private google sheet as the source!!!

///////////////////////
install file using wget
///////////////////////

you can install this program on your alpine linux install using the following commands:

apk update

apk add wget

wget -O /tmp/gs_auth https://github.com/Stew2134/gs_auth/releases/download/v0.1/gs_auth

chmod +x /tmp/gs_auth

mv /tmp/gs_auth /usr/local/bin/

//////////////////////////
Dependencies for building from source
//////////////////////////

Linux
Docker
(Must have user have docker perms without sudo)
git

//////////////////////////
Building from source
//////////////////////////

Step 1:

git clone the repo

Step 2: 

Navigate to the repo on terminal and run . ./install.sh

(This will build and run the docker image which proceeds to run a build of the go app and dumps in the folder created called cgo)
