module github.com/Djoulzy/emutools/mos6510

go 1.18

require (
	github.com/Djoulzy/emutools/mem v0.0.0-20221011195259-f223a2cc27f3
	github.com/albenik/bcd v0.0.0-20170831201648-635201416bc7
	golang.org/x/exp v0.0.0-20220713135740-79cabaa25d75
)

require (
	github.com/Djoulzy/Tools/clog v0.0.0-20220609190146-71af779f6ddc // indirect
	github.com/Djoulzy/emutools/charset v0.0.0-20220714172143-04f22bef44b8 // indirect
	github.com/stretchr/testify v1.7.2 // indirect
)

replace github.com/Djoulzy/emutools/mem v0.0.0-20221011195259-f223a2cc27f3 => ../mem
