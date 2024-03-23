package app

type app struct {
	blockDataGetter BlockDataGetter
}

func NewApp(blockDataGetter BlockDataGetter) *app {
	return &app{blockDataGetter: blockDataGetter}
}

