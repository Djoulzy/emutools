module github.com/Djoulzy/emutools/mos6510

go 1.20

replace github.com/Djoulzy/mmu v0.0.0-20221015154434-3927fedd1199 => ../../mmu

replace github.com/Djoulzy/chip v0.0.0-20221015154434-3927fedd1199 => ../../chip

require (
	github.com/Djoulzy/chip v0.0.0-20221015154434-3927fedd1199
	github.com/Djoulzy/mmu v0.0.0-20221015154434-3927fedd1199
	github.com/albenik/bcd v0.0.0-20170831201648-635201416bc7
	golang.org/x/exp v0.0.0-20230519143937-03e91628a987
)

require github.com/stretchr/testify v1.8.3 // indirect
