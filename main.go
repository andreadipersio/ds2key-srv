package main

import (
    "net"
    "fmt"
    "flag"
    "log"
)

var port int

func init() {
    flag.IntVar(&port, "port", 9501, "DS2KEY Port")
}

const (
    KEY_UP = iota
    KEY_DOWN
    KEY_LEFT
    KEY_RIGHT

    KEY_A
    KEY_B

    KEY_SELECT
    KEY_START


    KEY_X
    KEY_Y

    KEY_L
    KEY_R
)

type Pad struct {
    Key int32
    Value int32
}

var keyValues = map[int32] []Pad {
    2: []Pad{
        Pad{KEY_DOWN, 128},  // 2^7 1000 0000
        Pad{KEY_UP,    64},  // 2^6 0100 0000
        Pad{KEY_LEFT,  32},  // 2^5 0010 0000
        Pad{KEY_RIGHT, 16},  // 2^4 0001 0000
        Pad{KEY_START,  8},  // 2^3 0000 1000
        Pad{KEY_SELECT, 4},  // 2^2 0000 0100
        Pad{KEY_B,      2},  //     0000 0010
        Pad{KEY_A,      1},  //     0000 0001
    },

    3: []Pad{
        Pad{KEY_X, 4},
        Pad{KEY_Y, 8},
        Pad{KEY_L, 2},
        Pad{KEY_R, 1},
    },
}

func detectKeys(payload []byte) {
    keys := []int32{}

    for offset, pads  := range keyValues {
        fullValue := int32(payload[offset])

        for _, pad := range pads {
            value := fullValue & pad.Value

            if (value & pad.Value) == pad.Value  {
                keys = append(keys, pad.Key)
            }
        }
    }

    fmt.Printf("%v --- %v\n", payload, keys)
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

        detectKeys(payload);
    }
}
