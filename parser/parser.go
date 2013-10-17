package parser

var KEYS = map[uint32] []string {
    2: []string{"KEY_A", "KEY_B",
                "KEY_SELECT", "KEY_START",
                "KEY_RIGHT", "KEY_LEFT", "KEY_UP", "KEY_DOWN"},

    3: []string{"KEY_R", "KEY_L", "KEY_X", "KEY_Y"},
}

func DetectKeys(payload []byte) []string {
    pressedKeys := []string{}

    for offset, keys := range KEYS {
        value := uint32(payload[offset])

        for n, keyStr := range keys {
            mask := uint32(1 << uint32(n))

            if (value & mask) == mask {
                pressedKeys = append(pressedKeys, keyStr)
            }
        }
    }

    return pressedKeys
}
