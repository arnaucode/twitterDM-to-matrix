# twitterDM-to-matrix [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/twitterDM-to-matrix)](https://goreportcard.com/report/github.com/arnaucode/twitterDM-to-matrix)

bridge to send the received twitter Direct Messages (https://twitter.com) to a Matrix room (https://matrix.org), written in Go


needs a twitterConfig.json file on the /build folder with the content:
```
{
    "consumer_key": "xxxxxxxxxxxxxxxx",
    "consumer_secret": "xxxxxxxxxxxxxxxx",
    "access_token_key": "xxxxxxxxxxxxxxxx",
    "access_token_secret": "xxxxxxxxxxxxxxxx"
}

```
and a matrixConfig.json file on the /build folder with the content:
```
{
    "room_id": "xxxxx",
    "user": "xxxxx",
    "password": "xxxxx",
    "server": "https://xxxxx.xxxxx"
}

```

to run it:
- go to the /build folder
- open terminal
- execute the script with:
```
./twitterDM-to-matrix
```
