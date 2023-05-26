module github.com/Djoulzy/emutools/mos6510

go 1.20

replace github.com/Djoulzy/mmu => ../../mmu

require (
	github.com/Djoulzy/mmu v0.0.0-20221015154434-3927fedd1199
	github.com/albenik/bcd v0.0.0-20170831201648-635201416bc7
	golang.org/x/exp v0.0.0-20230519143937-03e91628a987
)

require (
	github.com/Djoulzy/Tools/clog v0.0.0-20220609190146-71af779f6ddc // indirect
	github.com/Djoulzy/emutools/charset v0.0.0-20230526064438-5a4686e43142 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
)
