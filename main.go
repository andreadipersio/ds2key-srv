package main

import (
    "net"
    "fmt"
    "flag"
    "log"
    "github.com/andreadipersio/ds2key-srv/parser"
    "github.com/andreadipersio/ds2key-srv/kbd"
)

var (
    port int
    status map[string] bool
)

func init() {
    flag.IntVar(&port, "port", 9501, "DS2KEY Port")
    status = make(map[string] bool)
}

func releaseAll() {
    for key, isPressed := range status {
        if isPressed {
            kbd.KeyUp(key)
            status[key] = false
        }
    }
}

func main() {
    addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))

    log.Print(addr)

    if err != nil {
        panic(err);
    }

    sock, err := net.ListenUDP("udp", addr)

    if err != nil {
        panic(err);
    }

    buf := [11]byte{}

    for {
        _, err := sock.Read(buf[0:]) 
        if err != nil {
            panic(err)
        }

        payload := buf[:4]

        keys := parser.DetectKeys(payload);

        if len(keys) == 0 {
            releaseAll()
            continue
        }

        log.Printf("%v --- %v", payload, keys)

        for _, key := range keys {
            stillDown, wasPressed := status[key]

            if wasPressed && stillDown {
                continue
            } else {
                kbd.KeyDown(key)
            }

            status[key] = true
        }
    }
}
