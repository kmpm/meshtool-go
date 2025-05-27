package sec

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

// message represents a test message structure
type message struct {
	Subject string `json:"subject"`
	Data    struct {
		Packet struct {
			From      uint32 `json:"from"`
			To        uint32 `json:"to"`
			Channel   uint32 `json:"channel"`
			Encrypted string `json:"encrypted"`
			ID        uint32 `json:"id"`
		} `json:"packet"`
		ChannelID string `json:"channelId"`
		GatewayID string `json:"gatewayId"`
	} `json:"data"`
}

func Test_createNonce(t *testing.T) {
	type args struct {
		id   uint32
		from uint32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"first", args{1146368586, 3244483888},
			[]byte{0x4a, 0x32, 0x54, 0x44, 0, 0, 0, 0, 0x30, 0xe5, 0x62, 0xc1, 0, 0, 0, 0},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createNonce(tt.args.id, tt.args.from)
			if len(got) != len(tt.want) {
				t.Errorf("createNonce() length = %d, want %d", len(got), len(tt.want))
				return
			}
			assert.Equal(t, len(tt.want), len(got), "createNonce() length mismatch")
			assert.Equal(t, tt.want, got, "createNonce() byte mismatch")
			// for i := range got {
			// 	if got[i] != tt.want[i] {
			// 		t.Errorf("createNonce() byte %d = %v, want %v", i, got[i], tt.want[i])
			// 	}
			// }

		})
	}
}

// https://github.com/FriendlyDev/meshtastic-decode/blob/master/tests/Unit/MeshtasticDecodeTest.php

func TestDecryptString(t *testing.T) {
	type args struct {
		dec             StringDecoder
		encryptedBase64 string
		id              uint32
		from            uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"first", args{base64.RawStdEncoding, "3VFsFwqgIOQq5l9cxMbrF7ZVzgMpfrIhLWDGF40Pn3Bl", 1146368586, 3244483888}, mustBase64DecodeRawStd("asdf"), false},
		{"from_php", args{base64.StdEncoding, "FIft4Bw0T4N8i+WiFCCncEz1PhpqzQ==", 1994624120, 1978765680}, mustBase64DecodeStd("CAMSEg2Uu9kXFchnos4leIkQZrgBIA=="), false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecryptString(tt.args.dec, tt.args.encryptedBase64, tt.args.id, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "DecryptString() byte mismatch")

		})
	}
}
