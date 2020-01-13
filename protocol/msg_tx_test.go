package protocol_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/Jeiwan/tinybit/protocol"
	"github.com/google/go-cmp/cmp"
)

func TestMsgTxMarshalBinary(t *testing.T) {
	tests := []struct {
		name  string
		txmsg string
		err   error
	}{
		{name: "legacy",
			txmsg: "0100000001317c144ae5b5a224370bd68c928b9f9e152d9829235ffbecec5ee64113662fc4000000006a47304402203c6ef3cba423365b37c031d235a674a10cf06b14fccda68bb5c35cbda5a2969b02207da3f69ea61c4a98eb488dac9d8a421dda9000e8afdc4a90cc2ebf93fbefb84f012102e248c2b8e9a5b78f2406c60b75ef1c4e88a06c7c36ad31e009db256505e27e79ffffffff0388270c00000000001976a914fe46ec55e937e584005b337495d76464b6b1cdba88ac22020000000000001976a914bdcccc7ce08a732ce55dcc3c1d8890e372bf7c1d88ac0000000000000000166a146f6d6e69000000000000001f0000886c98b7600000000000",
			err:   nil},
		{name: "segwit",
			txmsg: "0100000000010145b87f940bc57475403a3928ecf4cb3b86d2ba192039d4d703126edad14487ca0100000000ffffffff0200093d000000000017a91469f375f23b3d5d37bd942f3c31d7ae5a0cb61f5e87c8db030000000000220020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d0400473044022025863cfe71648bc8703f9f0607558cb7e79fcbebadc080ef1f0d7bfdd6ab1afa0220101ffaeb01b70e3360e87d6b3616886e547593e41c8a09d00cf8803601a9cc7901473044022031caba2ba6b079bc0d995e04f3651f977b5fc22dacceab1046a311fa2fb83898022030f5a852b425bdeb156be3a0a3de5bbc764302fbc1b7ca77058740cd511d49a9016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae00000000",
			err:   nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tx := protocol.MsgTx{}

			txmsg, err := hex.DecodeString(test.txmsg)
			if err != nil {
				t.Fatal(err)
			}

			r := bytes.NewBuffer(txmsg)
			err = tx.UnmarshalBinary(r)
			if err != nil {
				tt.Fatalf("unexpected error: %+v", err)
			}

			got, err := tx.MarshalBinary()
			if err != nil && test.err == nil {
				tt.Errorf("unexpected error: %+v", err)
				return
			}

			if err == nil && test.err != nil {
				tt.Errorf("expected error: %+v, got:%+v", test.err, tx)
				return
			}

			if diff := cmp.Diff(txmsg, got); diff != "" {
				tt.Errorf("MarshalBinary() mismatch(-want +got):\n%s", diff)
				return
			}
		})
	}

}

func TestMsgTxUnmarshalBinary(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *protocol.MsgTx
		err      error
	}{
		{name: "legacy",
			input: "0100000001317c144ae5b5a224370bd68c928b9f9e152d9829235ffbecec5ee64113662fc4000000006a47304402203c6ef3cba423365b37c031d235a674a10cf06b14fccda68bb5c35cbda5a2969b02207da3f69ea61c4a98eb488dac9d8a421dda9000e8afdc4a90cc2ebf93fbefb84f012102e248c2b8e9a5b78f2406c60b75ef1c4e88a06c7c36ad31e009db256505e27e79ffffffff0388270c00000000001976a914fe46ec55e937e584005b337495d76464b6b1cdba88ac22020000000000001976a914bdcccc7ce08a732ce55dcc3c1d8890e372bf7c1d88ac0000000000000000166a146f6d6e69000000000000001f0000886c98b7600000000000",
			expected: &protocol.MsgTx{
				Version:   1,
				Flag:      0,
				TxInCount: 1,
				TxIn: []protocol.TxInput{
					{PreviousOutput: protocol.OutPoint{
						Hash:  [32]byte{0x31, 0x7c, 0x14, 0x4a, 0xe5, 0xb5, 0xa2, 0x24, 0x37, 0x0b, 0xd6, 0x8c, 0x92, 0x8b, 0x9f, 0x9e, 0x15, 0x2d, 0x98, 0x29, 0x23, 0x5f, 0xfb, 0xec, 0xec, 0x5e, 0xe6, 0x41, 0x13, 0x66, 0x2f, 0xc4},
						Index: 0},
						ScriptLength: 106,
						SignatureScript: []byte{
							0x47, 0x30, 0x44, 0x02, 0x20, 0x3c, 0x6e, 0xf3, 0xcb, 0xa4, 0x23, 0x36, 0x5b, 0x37, 0xc0, 0x31,
							0xd2, 0x35, 0xa6, 0x74, 0xa1, 0x0c, 0xf0, 0x6b, 0x14, 0xfc, 0xcd, 0xa6, 0x8b, 0xb5, 0xc3, 0x5c,
							0xbd, 0xa5, 0xa2, 0x96, 0x9b, 0x02, 0x20, 0x7d, 0xa3, 0xf6, 0x9e, 0xa6, 0x1c, 0x4a, 0x98, 0xeb,
							0x48, 0x8d, 0xac, 0x9d, 0x8a, 0x42, 0x1d, 0xda, 0x90, 0x00, 0xe8, 0xaf, 0xdc, 0x4a, 0x90, 0xcc,
							0x2e, 0xbf, 0x93, 0xfb, 0xef, 0xb8, 0x4f, 0x01, 0x21, 0x02, 0xe2, 0x48, 0xc2, 0xb8, 0xe9, 0xa5,
							0xb7, 0x8f, 0x24, 0x06, 0xc6, 0x0b, 0x75, 0xef, 0x1c, 0x4e, 0x88, 0xa0, 0x6c, 0x7c, 0x36, 0xad,
							0x31, 0xe0, 0x09, 0xdb, 0x25, 0x65, 0x05, 0xe2, 0x7e, 0x79,
						},
						Sequence: 4294967295}},
				TxOutCount: 3,
				TxOut: []protocol.TxOutput{
					{Value: 796552,
						PkScriptLength: 25,
						PkScript: []byte{
							0x76, 0xa9, 0x14, 0xfe, 0x46, 0xec, 0x55, 0xe9, 0x37, 0xe5, 0x84, 0x00, 0x5b, 0x33, 0x74, 0x95,
							0xd7, 0x64, 0x64, 0xb6, 0xb1, 0xcd, 0xba, 0x88, 0xac,
						},
					},
					{Value: 546,
						PkScriptLength: 25,
						PkScript: []byte{
							0x76, 0xa9, 0x14, 0xbd, 0xcc, 0xcc, 0x7c, 0xe0, 0x8a, 0x73, 0x2c, 0xe5, 0x5d, 0xcc, 0x3c, 0x1d,
							0x88, 0x90, 0xe3, 0x72, 0xbf, 0x7c, 0x1d, 0x88, 0xac,
						},
					},
					{Value: 0,
						PkScriptLength: 22,
						PkScript: []byte{
							0x6a, 0x14, 0x6f, 0x6d, 0x6e, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00,
							0x88, 0x6c, 0x98, 0xb7, 0x60, 0x00,
						},
					},
				},
				TxWitness: protocol.TxWitnessData{Count: 0, Witness: nil},
				LockTime:  0},
			err: nil},
		{name: "segwit",
			input: "0100000000010145b87f940bc57475403a3928ecf4cb3b86d2ba192039d4d703126edad14487ca0100000000ffffffff0200093d000000000017a91469f375f23b3d5d37bd942f3c31d7ae5a0cb61f5e87c8db030000000000220020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d0400473044022025863cfe71648bc8703f9f0607558cb7e79fcbebadc080ef1f0d7bfdd6ab1afa0220101ffaeb01b70e3360e87d6b3616886e547593e41c8a09d00cf8803601a9cc7901473044022031caba2ba6b079bc0d995e04f3651f977b5fc22dacceab1046a311fa2fb83898022030f5a852b425bdeb156be3a0a3de5bbc764302fbc1b7ca77058740cd511d49a9016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae00000000",
			expected: &protocol.MsgTx{
				Version:   1,
				Flag:      1,
				TxInCount: 1,
				TxIn: []protocol.TxInput{
					{PreviousOutput: protocol.OutPoint{
						Hash:  [32]uint8{0x45, 0xb8, 0x7f, 0x94, 0x0b, 0xc5, 0x74, 0x75, 0x40, 0x3a, 0x39, 0x28, 0xec, 0xf4, 0xcb, 0x3b, 0x86, 0xd2, 0xba, 0x19, 0x20, 0x39, 0xd4, 0xd7, 0x03, 0x12, 0x6e, 0xda, 0xd1, 0x44, 0x87, 0xca},
						Index: 0x01},
						Sequence: 0xffffffff},
				},
				TxOutCount: 2,
				TxOut: []protocol.TxOutput{
					{Value: 4000000,
						PkScriptLength: 0x17,
						PkScript: []uint8{
							0xa9, 0x14, 0x69, 0xf3, 0x75, 0xf2, 0x3b, 0x3d, 0x5d, 0x37, 0xbd, 0x94, 0x2f, 0x3c, 0x31, 0xd7,
							0xae, 0x5a, 0x0c, 0xb6, 0x1f, 0x5e, 0x87,
						},
					},
					{
						Value:          252872,
						PkScriptLength: 0x22,
						PkScript: []byte{
							0x00, 0x20, 0x70, 0x1a, 0x8d, 0x40, 0x1c, 0x84, 0xfb, 0x13, 0xe6, 0xba, 0xf1, 0x69, 0xd5, 0x96,
							0x84, 0xe1, 0x7a, 0xbd, 0x9f, 0xa2, 0x16, 0xc8, 0xcc, 0x5b, 0x9f, 0xc6, 0x3d, 0x62, 0x2f, 0xf8,
							0xc5, 0x8d,
						},
					},
				},
				TxWitness: protocol.TxWitnessData{
					Count: 0x04,
					Witness: []protocol.TxWitness{
						{},
						{
							Length: 0x47,
							Data: []uint8{
								0x30, 0x44, 0x02, 0x20, 0x25, 0x86, 0x3c, 0xfe, 0x71, 0x64, 0x8b, 0xc8, 0x70, 0x3f, 0x9f, 0x06,
								0x07, 0x55, 0x8c, 0xb7, 0xe7, 0x9f, 0xcb, 0xeb, 0xad, 0xc0, 0x80, 0xef, 0x1f, 0x0d, 0x7b, 0xfd,
								0xd6, 0xab, 0x1a, 0xfa, 0x02, 0x20, 0x10, 0x1f, 0xfa, 0xeb, 0x01, 0xb7, 0x0e, 0x33, 0x60, 0xe8,
								0x7d, 0x6b, 0x36, 0x16, 0x88, 0x6e, 0x54, 0x75, 0x93, 0xe4, 0x1c, 0x8a, 0x09, 0xd0, 0x0c, 0xf8,
								0x80, 0x36, 0x01, 0xa9, 0xcc, 0x79, 0x01,
							},
						},
						{
							Length: 0x47,
							Data: []byte{
								0x30, 0x44, 0x02, 0x20, 0x31, 0xca, 0xba, 0x2b, 0xa6, 0xb0, 0x79, 0xbc, 0x0d, 0x99, 0x5e, 0x04,
								0xf3, 0x65, 0x1f, 0x97, 0x7b, 0x5f, 0xc2, 0x2d, 0xac, 0xce, 0xab, 0x10, 0x46, 0xa3, 0x11, 0xfa,
								0x2f, 0xb8, 0x38, 0x98, 0x02, 0x20, 0x30, 0xf5, 0xa8, 0x52, 0xb4, 0x25, 0xbd, 0xeb, 0x15, 0x6b,
								0xe3, 0xa0, 0xa3, 0xde, 0x5b, 0xbc, 0x76, 0x43, 0x02, 0xfb, 0xc1, 0xb7, 0xca, 0x77, 0x05, 0x87,
								0x40, 0xcd, 0x51, 0x1d, 0x49, 0xa9, 0x01,
							},
						},
						{
							Length: 0x69,
							Data: []uint8{
								0x52, 0x21, 0x03, 0x75, 0xe0, 0x0e, 0xb7, 0x2e, 0x29, 0xda, 0x82, 0xb8, 0x93, 0x67, 0x94, 0x7f,
								0x29, 0xef, 0x34, 0xaf, 0xb7, 0x5e, 0x86, 0x54, 0xf6, 0xea, 0x36, 0x8e, 0x0a, 0xcd, 0xfd, 0x92,
								0x97, 0x6b, 0x7c, 0x21, 0x03, 0xa1, 0xb2, 0x63, 0x13, 0xf4, 0x30, 0xc4, 0xb1, 0x5b, 0xb1, 0xfd,
								0xce, 0x66, 0x32, 0x07, 0x65, 0x9d, 0x8c, 0xac, 0x74, 0x9a, 0x0e, 0x53, 0xd7, 0x0e, 0xff, 0x01,
								0x87, 0x44, 0x96, 0xfe, 0xff, 0x21, 0x03, 0xc9, 0x6d, 0x49, 0x5b, 0xfd, 0xd5, 0xba, 0x41, 0x45,
								0xe3, 0xe0, 0x46, 0xfe, 0xe4, 0x5e, 0x84, 0xa8, 0xa4, 0x8a, 0xd0, 0x5b, 0xd8, 0xdb, 0xb3, 0x95,
								0xc0, 0x11, 0xa3, 0x2c, 0xf9, 0xf8, 0x80, 0x53, 0xae,
							},
						},
					},
				},
				LockTime: 0},
			err: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tx := protocol.MsgTx{}

			input, err := hex.DecodeString(test.input)
			if err != nil {
				t.Fatal(err)
			}

			r := bytes.NewBuffer(input)
			err = tx.UnmarshalBinary(r)

			if err != nil && test.err == nil {
				tt.Errorf("unexpected error:%+v", err)
				return
			}

			if err == nil && test.err != nil {
				tt.Errorf("expected error: %+v, got:%+v", test.err, tx)
				return
			}

			if diff := cmp.Diff(test.expected, &tx); diff != "" {
				tt.Errorf("UnmarshalBinary() mismatch(-want +got):\n%s", diff)
				return
			}
		})
	}

}
