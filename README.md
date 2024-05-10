# go-learn

## Installing go: https://go.dev/doc/install , https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04

### Linux:
1. ```cd ~```
2. ```curl -OL https://golang.org/dl/go1.22.3.linux-amd64.tar.gz```
3. ```sha256sum go1.22.3.linux-amd64.tar.gz```
4. ```sudo tar -C /usr/local -xvf go1.22.3.linux-amd64.tar.gz```
5. ```sudo nano ~/.profile```
6. Add "export PATH=$PATH:/usr/local/go/bin" at the end of the file
7. ```source ~/.profile```
8. Validate: ```go version```

### Running hello world
1. ```mkdir hello```
2. ```cd hello```
3. ```go mod init your_domain/hello```
2. ```nano hello.go```
3. Add following code to the file:<br>
    ```
        package main

        import "fmt"

        func main() {
            fmt.Println("Hello, World!")
        }
    ```
4. ```go run .```