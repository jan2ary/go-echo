package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func handle(conn net.Conn) {
    defer conn.Close()
    // NewScanner returns a new Scanner to read from r.
    // The split function defaults to ScanLines.
    scanner := bufio.NewScanner(conn)
    // Scan advances the Scanner to the next token, which will then be
    // available through the Bytes or Text method.
    for scanner.Scan() {
        // Text returns the most recent token generated by a call to Scan
        // as a newly allocated string holding its bytes.
        ln := scanner.Text()
        if ln == "" {
            break
        }
        ln = strings.ToLower(ln)
        bs := []byte(ln)
        var rot13 = make([]byte, len(bs))
        for k, v := range bs {
            // ascii 97 - 122
            /// 109 + 13 = 122
            if v <= 109 {
                rot13[k] = v + 13
            } else {
                // 110 + 13 = 123
                //123 - 26 = 97
                rot13[k] = v - 26 + 13
            }
        }
        fmt.Fprintf(conn, "%s\n", rot13)
    }
}

func main() {
    ln, err := net.Listen("tcp", ":9000")
    if err != nil {
        panic(err)
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            panic(err)
        }
        go handle(conn)
    }
}
