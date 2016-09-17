package main

import (
  "golang.org/x/net/websocket"
  "log"
  "flag"
  "crypto/tls"
  "net"
  "strconv"
  "strings"
  "time"
  "os"
)

const (
  message       = "Ping"
  StopCharacter = "\r\n\r\n"
)

func TestTcp(ip string, port int) {
  log.Printf("Try to connect to %s on port %d with message %s", ip, port, message)
  timeout := time.Duration(5) * time.Second
  addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
  conn, err := net.DialTimeout("tcp", addr, timeout)

  defer conn.Close()

  if err != nil {
    log.Print(err)
    os.Exit(2)
    return
  }

  conn.Write([]byte(message))
  conn.Write([]byte(StopCharacter))
  log.Printf("Send: %s", message)

  buff := make([]byte, 1024)
  n, _ := conn.Read(buff)
  log.Printf("Receive: %s", buff[:n])
}

func TestWebsocket(url string) {
  origin := "http://localhost/"

  config, err := websocket.NewConfig(url, origin)
  config.TlsConfig = &tls.Config{InsecureSkipVerify: true}

  ws, err := websocket.DialConfig(config)

  if err != nil {
    log.Fatal(err)
    os.Exit(2)
  }

  if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
      log.Print(err)
  }

  var msg = make([]byte, 512)
  var n int

  if n, err = ws.Read(msg); err != nil {
    log.Print(err)
  }

  log.Printf("Connecting to <%s> Received <%s>\n", url, msg[:n])
}

func main() {
  urlStr := flag.String("url", "wss://echo.websocket.org:8080", "the url of the websocket server")
  flag.Parse()

  url := *urlStr

  var (
    ip = "127.0.0.1"
    port = 22
  )

  TestTcp(ip, port)
  TestWebsocket(url)
}
