# go-dns-client
Basic DNS Client example, written in GoLang


### Install Golang on Mac OSX
```
$ brew install go
```

### Configure and prepare Golang
```
$ go env GOPATH
# export GOPATH=/Users/user/go
$ mkdir $GOPATH
$ mkdir $GOPATH/src
```

### Install go-dns-client from src.
```
$ cd /Users/user/go/src
$ git commit https://github.com/BerginLabs/go-dns-client.git
$ cd go-dns-client
$ go build
```

### Examples
output as json
``` 
$ ./go-dns-client google.com json | jq .
{
  "queriedHostname": "google.com",
  "resolvedIps": [
    "172.217.5.14"
  ],
  "scriptExecution": "2019-12-11T18:45:47-08:00"
}
$
```
output to stdout
``` 
$ ./go-dns-client yahoo.com stdout
[+] Domain: yahoo.com
[+] Run Time: 2019-12-11T18:48:22-08:00
1. 72.30.35.10
2. 72.30.35.9
3. 98.137.246.7
4. 98.137.246.8
5. 98.138.219.231
6. 98.138.219.232

$
```
pretty json output to file
``` 
$ TARGET="test.com" && \
  touch data/$TARGET.json && \
  ./go-dns-client $TARGET json \
  | jq . > data/$TARGET.json
```