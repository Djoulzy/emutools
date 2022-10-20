module github.com/Djoulzy/emutools/mos6510

go 1.18

require (
	github.com/Djoulzy/emutools/mem/v2 v2.0.0-20221019172136-ccfe72c37753
	github.com/albenik/bcd v0.0.0-20170831201648-635201416bc7
	golang.org/x/exp v0.0.0-20221018221608-02f3b879a704
)

require (
	github.com/Djoulzy/Tools/clog v0.0.0-20220609190146-71af779f6ddc // indirect
	github.com/Djoulzy/emutools/charset v0.0.0-20221019155951-08ad875a524a // indirect
	github.com/stretchr/testify v1.7.2 // indirect
)

replace github.com/Djoulzy/emutools/mem v0.0.0-20221015154434-3927fedd1199 => ../mem

replace github.com/Djoulzy/emutools/mem/v2 v2.0.0-20221019172136-ccfe72c37753 => ../mem/v2
