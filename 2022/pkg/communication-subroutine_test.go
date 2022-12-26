package pkg

import "testing"

func TestFindStartOfPacketMarker(t *testing.T) {
    tests := []struct {
        name             string
        datastreamBuffer string
        want             int
    }{
        {
            name:             "test case 1",
            datastreamBuffer: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
            want:             7,
        },
        {
            name:             "test case 2",
            datastreamBuffer: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
            want:             10,
        },
        {
            name:             "test case 3",
            datastreamBuffer: "bvwajplbgvbhsrlpgdmjqwftvncz",
            want:             4,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FindStartOfPacketMarker(tt.datastreamBuffer); got != tt.want {
                t.Errorf("FindStartOfPacketMarker() = %v, want %v", got, tt.want)
            }
        })
    }
}
