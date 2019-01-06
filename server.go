package main

import (
    "flag"
    "fmt"
    "log"
    "net"
    "os"
)

const (
    CONN_HOST = "0.0.0.0"
    CONN_TYPE = "tcp"
)

var logPtr *string

func main() {
    // Listen for incoming connections.
    portPtr := flag.String("port", "3389", "listening port")
    logPtr = flag.String("logfile", "connlog", "connection log file")
    flag.Parse()

    portValue := *portPtr

    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+portValue)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    f, err := os.OpenFile(*logPtr, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
      log.Fatalf("error opening file: %v", err)
    }
    defer f.Close()
    log.SetOutput(f)
    log.Println("Listening on " + CONN_HOST + ":" + portValue)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  _, err := conn.Read(buf)
  if err != nil {
    log.Println("Error reading:", err.Error())
  }
  // We could do something with the read buffer but we don't
  log.Println("Connection received from: ", conn.RemoteAddr())
  conn.Close()
}	
