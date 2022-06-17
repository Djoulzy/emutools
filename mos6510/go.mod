module github.com/Djoulzy/emutools/mos6510

go 1.18

require (
	github.com/Djoulzy/emutools/mem v0.0.0-20220616174737-0b8bae1ef37f
	github.com/albenik/bcd v0.0.0-20170831201648-635201416bc7
	golang.org/x/exp v0.0.0-20220613132600-b0d781184e0d
)

require (
	github.com/Djoulzy/Tools/clog v0.0.0-20220609190146-71af779f6ddc // indirect
	github.com/Djoulzy/emutools/charset v0.0.0-20220616174737-0b8bae1ef37f // indirect
	github.com/stretchr/testify v1.7.2 // indirect
)

replace github.com/Djoulzy/emutools/mem v0.0.0-20220614171243-2cd571b9749d => ../mem
