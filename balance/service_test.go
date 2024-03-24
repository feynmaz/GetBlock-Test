package balance

import (
	"math/big"
	"testing"

	"github.com/feynmaz/GetBlock-Test/block"
	"github.com/stretchr/testify/assert"
)

var service = NewService()

func Test_GetBalances(t *testing.T) {
	t.Run("single transaction", func(t *testing.T) {
		blocks := []block.Block{{
			Transactions: []block.Transaction{
				{From: "0x229548ea8bb086ce2c3c40c6852d029ba8549b2c", To: "0x616713b662b0a597db3d67583c48a6ec29ef2c0f", Value: big.NewInt(87000000000)},
			},
		}}
		balances := service.GetBalances(blocks)

		assert.Equal(t,
			balances[address("0x229548ea8bb086ce2c3c40c6852d029ba8549b2c")],
			big.NewInt(-87000000000),
		)
		assert.Equal(t,
			balances[address("0x616713b662b0a597db3d67583c48a6ec29ef2c0f")],
			big.NewInt(87000000000),
		)
	})

	t.Run("sigle block, miltiple transactions, same addresses", func(t *testing.T) {
		blocks := []block.Block{{
			Transactions: []block.Transaction{
				{From: "0x229548ea8bb086ce2c3c40c6852d029ba8549b2c", To: "0x616713b662b0a597db3d67583c48a6ec29ef2c0f", Value: big.NewInt(87000000000)},
				{From: "0x616713b662b0a597db3d67583c48a6ec29ef2c0f", To: "0x2b9aa475ecfa65275ebe15bb3dda776e77664a29", Value: big.NewInt(56000000000)},
				{From: "0x229548ea8bb086ce2c3c40c6852d029ba8549b2c", To: "0x2b9aa475ecfa65275ebe15bb3dda776e77664a29", Value: big.NewInt(1430000)},
			},
		}}
		balances := service.GetBalances(blocks)
		assert.Equal(t,
			balances[address("0x229548ea8bb086ce2c3c40c6852d029ba8549b2c")],
			big.NewInt(-87001430000),
		)
		assert.Equal(t,
			balances[address("0x616713b662b0a597db3d67583c48a6ec29ef2c0f")],
			big.NewInt(31000000000),
		)
		assert.Equal(t,
			balances[address("0x2b9aa475ecfa65275ebe15bb3dda776e77664a29")],
			big.NewInt(56001430000),
		)
	})

	t.Run("multiple blocks, same addresses", func(t *testing.T) {
		blocks := []block.Block{{
			Transactions: []block.Transaction{
				{From: "0x229548ea8bb086ce2c3c40c6852d029ba8549b2c", To: "0x2b9aa475ecfa65275ebe15bb3dda776e77664a29", Value: big.NewInt(1430000)},
			},
		}, {
			Transactions: []block.Transaction{
				{From: "0x229548ea8bb086ce2c3c40c6852d029ba8549b2c", To: "0x616713b662b0a597db3d67583c48a6ec29ef2c0f", Value: big.NewInt(87000000000)},
				{From: "0x616713b662b0a597db3d67583c48a6ec29ef2c0f", To: "0x2b9aa475ecfa65275ebe15bb3dda776e77664a29", Value: big.NewInt(56000000000)},
			},
		}}
		balances := service.GetBalances(blocks)
		assert.Equal(t,
			balances[address("0x229548ea8bb086ce2c3c40c6852d029ba8549b2c")],
			big.NewInt(-87001430000),
		)
		assert.Equal(t,
			balances[address("0x616713b662b0a597db3d67583c48a6ec29ef2c0f")],
			big.NewInt(31000000000),
		)
		assert.Equal(t,
			balances[address("0x2b9aa475ecfa65275ebe15bb3dda776e77664a29")],
			big.NewInt(56001430000),
		)
	})

	t.Run("sigle block, miltiple transactions, different addresses", func(t *testing.T) {
		blocks := []block.Block{{
			Transactions: []block.Transaction{
				{From: "0x229548ea8bb086ce2c3c40c6852d029ba8549b2c", To: "0x616713b662b0a597db3d67583c48a6ec29ef2c0f", Value: big.NewInt(87000000000)},
				{From: "0xf77787f4ef3e3c442805c39efc27dbf9da07a86e", To: "0x2b9aa475ecfa65275ebe15bb3dda776e77664a29", Value: big.NewInt(56000000000)},
				{From: "0x4200bd5dc856fc0be0cb5a235199262c94748b57", To: "0x1aac2278c6462f5d33349ec62274ecd399cd371e", Value: big.NewInt(90000000000)},
			},
		}}
		balances := service.GetBalances(blocks)
		assert.Equal(t,
			balances[address("0x229548ea8bb086ce2c3c40c6852d029ba8549b2c")],
			big.NewInt(-87000000000),
		)
		assert.Equal(t,
			balances[address("0x616713b662b0a597db3d67583c48a6ec29ef2c0f")],
			big.NewInt(87000000000),
		)
		assert.Equal(t,
			balances[address("0xf77787f4ef3e3c442805c39efc27dbf9da07a86e")],
			big.NewInt(-56000000000),
		)
		assert.Equal(t,
			balances[address("0x2b9aa475ecfa65275ebe15bb3dda776e77664a29")],
			big.NewInt(56000000000),
		)
		assert.Equal(t,
			balances[address("0x4200bd5dc856fc0be0cb5a235199262c94748b57")],
			big.NewInt(-90000000000),
		)
		assert.Equal(t,
			balances[address("0x1aac2278c6462f5d33349ec62274ecd399cd371e")],
			big.NewInt(90000000000),
		)
	})

	t.Run("miltiple blocks, different addresses", func(t *testing.T) {
		blocks := []block.Block{{
			Transactions: []block.Transaction{
				{From: "0x229548ea8bb086ce2c3c40c6852d029ba8549b2c", To: "0x616713b662b0a597db3d67583c48a6ec29ef2c0f", Value: big.NewInt(87000000000)},
				{From: "0xf77787f4ef3e3c442805c39efc27dbf9da07a86e", To: "0x2b9aa475ecfa65275ebe15bb3dda776e77664a29", Value: big.NewInt(56000000000)},
			},
		}, {
			Transactions: []block.Transaction{
				{From: "0x4200bd5dc856fc0be0cb5a235199262c94748b57", To: "0x1aac2278c6462f5d33349ec62274ecd399cd371e", Value: big.NewInt(90000000000)},
			},
		}}
		balances := service.GetBalances(blocks)
		assert.Equal(t,
			balances[address("0x229548ea8bb086ce2c3c40c6852d029ba8549b2c")],
			big.NewInt(-87000000000),
		)
		assert.Equal(t,
			balances[address("0x616713b662b0a597db3d67583c48a6ec29ef2c0f")],
			big.NewInt(87000000000),
		)
		assert.Equal(t,
			balances[address("0xf77787f4ef3e3c442805c39efc27dbf9da07a86e")],
			big.NewInt(-56000000000),
		)
		assert.Equal(t,
			balances[address("0x2b9aa475ecfa65275ebe15bb3dda776e77664a29")],
			big.NewInt(56000000000),
		)
		assert.Equal(t,
			balances[address("0x4200bd5dc856fc0be0cb5a235199262c94748b57")],
			big.NewInt(-90000000000),
		)
		assert.Equal(t,
			balances[address("0x1aac2278c6462f5d33349ec62274ecd399cd371e")],
			big.NewInt(90000000000),
		)
	})
}
