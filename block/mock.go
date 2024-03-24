package block

import "math/big"

type mockBlockGetter struct{}

func NewMockBlockGetter() *mockBlockGetter {
	return &mockBlockGetter{}
}

func (g *mockBlockGetter) GetLatestBlocks(n uint) ([]Block, error) {
	if n == 0 {
		return nil, ErrNoBlocksRequested
	}

	blocks := []Block{
		{
			Transactions: []Transaction{
				{From: "0x229548ea8bb086ce2c3c40c6852d029ba8549b2c", To: "0x616713b662b0a597db3d67583c48a6ec29ef2c0f", Value: big.NewInt(87000000000)},
				{From: "0xf77787f4ef3e3c442805c39efc27dbf9da07a86e", To: "0x2b9aa475ecfa65275ebe15bb3dda776e77664a29", Value: big.NewInt(56000000000)},
			},
		},
		{
			Transactions: []Transaction{
				{From: "0x6179dadf42729694b95287a0be8be3edaaec41f1", To: "0xc290f3b9e56494efc188428b4d183836fa0a4ef8", Value: big.NewInt(91000000000)},
				{From: "0x4200bd5dc856fc0be0cb5a235199262c94748b57", To: "0x1aac2278c6462f5d33349ec62274ecd399cd371e", Value: big.NewInt(90000000000)},
			},
		},
		{
			Transactions: []Transaction{
				{From: "0x510afd08789df71b49bdec3a5ed3e4fb07c5839a", To: "0xc22ac7127c424259751eb35e1744ec829a069bce", Value: big.NewInt(58000000000)},
				{From: "0xbe81625609d3183f4372d8c6fa6cd38716a60067", To: "0x9661f02568691c98ef6eae9ad096efd514838eb7", Value: big.NewInt(85000000000)},
			},
		},
		{
			Transactions: []Transaction{
				{From: "0x878140df0cae0669d67ed21a9de214588ffb5042", To: "0x379a8c467e6ed8a8fb7875ea0a4cd6393e17dab7", Value: big.NewInt(60000000000)},
				{From: "0xa7758f2873cc8077583587bd7a80d42d4215a3cf", To: "0xb9570d47d5a262543a40f154da407b98f8060669", Value: big.NewInt(10000000000)},
			},
		},
	}

	if n < 4 {
		return blocks[:n], nil
	}

	return blocks, nil
}
