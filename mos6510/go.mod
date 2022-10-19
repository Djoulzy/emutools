module github.com/Djoulzy/emutools/mos6510

go 1.18

require (
	github.com/Djoulzy/emutools/mem/v1 v0.0.0-20221019154839-c8e57b56aa7d
	github.com/albenik/bcd v0.0.0-20170831201648-635201416bc7
	golang.org/x/exp v0.0.0-20221018221608-02f3b879a704
)

require (
	github.com/Djoulzy/Tools/clog v0.0.0-20220609190146-71af779f6ddc // indirect
	github.com/Djoulzy/emutools/charset v0.0.0-20221019132048-ccd043a2363c // indirect
	github.com/stretchr/testify v1.7.2 // indirect
)

replace github.com/Djoulzy/emutools/mem/v1 v0.0.0-20221019154839-c8e57b56aa7d => ../mem/v1
