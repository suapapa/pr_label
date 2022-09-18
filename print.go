package main

type Orient int

const (
	Landscape Orient = iota
	Portrait

	ql800MaxWidth = 696 // 62mm endless
)

func printAddrFrom(addr *Addr) error {
	// return printFromAddr(addr)
	return nil
}

func printAddrTo(addr *Addr) error {
	// return printAddr(addr, Landscape)
	return nil
}
