# http-echo
Simple HTTP server that just echos everything to the console.  For debugging POST requests.

# Installation

```
go get -u github.com/robert-wallis/http-echo
```

# Usage

## Basic
Starts listening at http://localhost:8000/ and prints all requests to the console.
```
http-echo
```

## Custom Port
Starts listening on your IP from anywhere (on your LAN), and prints all requests to the console.
```
sudo http-echo -address 0.0.0.0:80
```
