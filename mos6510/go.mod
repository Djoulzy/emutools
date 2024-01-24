module github.com/Djoulzy/emutools/mos6510

go 1.21

replace github.com/Djoulzy/mmu => ../../mmu

require (
	github.com/Djoulzy/mmu v0.0.0-20230605062009-e48b6d54957a
	github.com/albenik/bcd v0.0.0-20170831201648-635201416bc7
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a
)

require (
	github.com/Djoulzy/Tools/clog v0.0.0-20220609190146-71af779f6ddc // indirect
	github.com/Djoulzy/emutools/charset v0.0.0-20240123173627-4140dd715cad // indirect
	github.com/stretchr/testify v1.8.3 // indirect
)
