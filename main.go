package main

import (
    "fmt"
    "log"

    "strings"

    "flag"

    "net"

    "github.com/andreadipersio/ds2key-srv/parser"
    "github.com/andreadipersio/ds2key-srv/kbd"
)

var (
    port int
    verbose bool

    status map[string] bool // used to determine when to release keys
)

func init() {
    flag.IntVar(&port, "port", 9501, "DS2KEY Port")
    flag.BoolVar(&verbose, "verbose", false, "Enable logging of keystrokes on stderr")

    flag.Parse()

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

func released(keys []string, key string) bool {
    for _, newKey := range keys {
        if newKey == key {
            return false
        }
    }

    return true
}

func logKeyStatus() {
    var s = []string{}

    for key, isPressed := range status {
        if !isPressed {
            continue
        }

        s = append(s, fmt.Sprintf("[%v]", key))
    }

    log.Print(strings.Join(s, " --- "))
}


func main() {
    fullAddr := fmt.Sprintf(":%d", port)
    addr, err := net.ResolveUDPAddr("udp", fullAddr)

    log.Print(addr)

    if err != nil {
        log.Panicf("Wrong address %v: %v", fullAddr, err);
    }

    sock, err := net.ListenUDP("udp", addr)

    if err != nil {
        log.Panicf("Cannot listen from %v: %v", fullAddr, err);
    }

    buf := [11]byte{}

    for {
        if _, err := sock.Read(buf[0:]); err != nil {
            log.Printf("ERROR::%v", err)
            continue
        }

        // first 4 bytes contains status of pad buttons
        payload := buf[:4]

        keys := parser.DetectKeys(payload);

        // all buttons on gamepad released
        if len(keys) == 0 {
            releaseAll()
            continue
        }

        for _, key := range keys {
            stillDown, wasPressed := status[key]

            if wasPressed && stillDown {
                continue
            } else {
                kbd.KeyDown(key)
            }

            status[key] = true
        }

        for key, wasPressed := range status {
            if wasPressed && released(keys, key) {
                kbd.KeyUp(key)
                status[key] = false
            }
        }

        if verbose {
            logKeyStatus()
        }
    }
}
