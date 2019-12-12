# go-dns-client
Basic DNS Client example, written in GoLang


### Install Golang on Mac OSX
```
$ brew install go
```

### Configure and prepare Golang
```
$ go env GOPATH
/Users/user/go
$ mkdir /Users/user/go
$ mkdir /Users/user/go/src
```

### Install dns-client from src.
```
$ cd /Users/user/go/src
$ git commit https://github.com/BerginLabs/go-dns-client.git
$ cd go-dns-client
$ go build
```

### Test go-dns-client
```
$ ./go-dns-client yahoo.com 
[+] Starting DNS Lookup for domain: yahoo.com
[+] Query result: 72.30.35.10
[+] Query result: 72.30.35.9
[+] Query result: 98.137.246.7
[+] Query result: 98.137.246.8
[+] Query result: 98.138.219.231
[+] Query result: 98.138.219.232
[+] DNS Resolution Complete.
```