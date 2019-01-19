package main

import (
	"debug/elf"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

var section = flag.String("section", "", "Show section contents")

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
	}
	filename := flag.Arg(0)

	f, err := elf.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open binary: %s", err)
		os.Exit(1)
	}

	fmt.Printf("loaded binary '%s' %s/%s (%d bits) entry@%#x\n", filename,
		f.Type, f.Machine, 666, f.Entry)

	for _, s := range f.Sections {
		fmt.Printf("  %#016x %-8x %-20.20s %s\n", s.Addr, s.Size, s.Name, s.Type)
	}

	dumpSymbols := func(symbols []elf.Symbol) {
		for _, s := range symbols {
			var Type string
			if elf.ST_TYPE(s.Info) == elf.STT_FUNC {
				Type = "FUNC"
			} else {
				Type = ""
			}
			var bind string
			if elf.ST_BIND(s.Info) == elf.STB_LOCAL {
				bind = "LOCAL"
			} else if elf.ST_BIND(s.Info) == elf.STB_GLOBAL {
				bind = "GLOBAL"
			}
			fmt.Printf("  %-30.30s %#016x %-8s %-8s\n", s.Name, s.Value, Type, bind)
		}
	}
	symbols, err := f.Symbols()
	if err == nil {
		fmt.Printf("scanned symbol tables\n")
		dumpSymbols(symbols)
	}
	symbols, err = f.DynamicSymbols()
	if err == nil {
		dumpSymbols(symbols)
	}

	if *section != "" {
		err := dumpSection(*section, f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't dump section: %s\n", err)
		}
	}
}

func dumpSection(name string, f *elf.File) error {
	section := f.Section(name)
	if section == nil {
		return fmt.Errorf("no section named %s", name)
	}
	data, err := section.Data()
	if err != nil {
		return fmt.Errorf("cannot read section data from %s: %s", name, err)
	}
	fmt.Printf("dump of section %s\n", name)
	fmt.Println(hex.Dump(data))

	return nil
}
