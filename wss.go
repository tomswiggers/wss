package main

import (
  "golang.org/x/net/websocket"
  "log"
  "flag"
  "crypto/tls"
  "net"
  "strconv"
  "strings"
)

const (
  message       = "Ping"
  StopCharacter = "\r\n\r\n"
)

func SocketClient(ip string, port int) {
  addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
  conn, err := net.Dial("tcp", addr)

  defer conn.Close()

  if err != nil {
    log.Fatalln(err)
  }

  conn.Write([]byte(message))
  conn.Write([]byte(StopCharacter))
  log.Printf("Send: %s", message)

  buff := make([]byte, 1024)
  n, _ := conn.Read(buff)
  log.Printf("Receive: %s", buff[:n])

}

func main() {
  urlStr := flag.String("url", "wss://echo.websocket.org", "the url of the websocket server")
  flag.Parse()

  origin := "http://localhost/"
  url := *urlStr

  config, err := websocket.NewConfig(url, origin)
  config.TlsConfig = &tls.Config{InsecureSkipVerify: true}

  ws, err := websocket.DialConfig(config)

  if err != nil {
    log.Fatal(err)
  }

  if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
      log.Fatal(err)
  }

  var msg = make([]byte, 512)
  var n int

  if n, err = ws.Read(msg); err != nil {
    log.Fatal(err)
  }

  var (
    ip = "10.10.20.21"
    port = 22
  )

  SocketClient(ip, port)

  log.Printf("Connecting to <%s> Received <%s>\n", url, msg[:n])
}
