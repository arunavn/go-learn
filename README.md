# go-learn

## Installing go: https://go.dev/doc/install , https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04

### Linux:

```bash 
cd ~ 
curl -OL https://golang.org/dl/go1.22.3.linux-amd64.tar.gz
sha256sum go1.22.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.22.3.linux-amd64.tar.gz
sudo nano ~/.profile
## Add "export PATH=$PATH:/usr/local/go/bin" at the end of the file
source ~/.profile
## 
go version

```

### Running hello world
```bash
mkdir hello
cd hello
go mod init your_domain/hello
nano hello.go
```

Add following code to the file:<br>
``` go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    }
```
    
```bash
go run .
```