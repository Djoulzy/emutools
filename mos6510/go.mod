module github.com/Djoulzy/emutools/mos6510

go 1.18

require (
	github.com/Djoulzy/emutools/mem v0.0.0-20220614171243-2cd571b9749d
	github.com/albenik/bcd v0.0.0-20170831201648-635201416bc7
)

require (
	github.com/Djoulzy/Tools/clog v0.0.0-20220609190146-71af779f6ddc // indirect
	github.com/Djoulzy/emutools/charset v0.0.0-20220614171243-2cd571b9749d // indirect
	github.com/stretchr/testify v1.7.1 // indirect
)

replace (
	github.com/Djoulzy/emutools/mem v0.0.0-20220614171243-2cd571b9749d => ../mem
)