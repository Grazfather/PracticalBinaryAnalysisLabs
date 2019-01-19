# Chapter 4: Building a binary loader using libbfd

Because I want to write my tools in Go, I won't be using BFD.

Golang has a pretty nice ELF parser already ("debug/elf"), and the code here is
mostly just an abstraction on top of the messy bfd API. The golang elf api is
pretty nice already, so until I decide to add support for PE or mach or
something, I don't feel the need to make my own interface, and so my code
samples will probably use the `elf.File` object directly. See _loader_demo.go_.

## 1. Dumping Section Contents

> For brevity, the current version of the loader_demo program doesn’t display
> section contents. Expand it with the ability to take a binary and the name of
> a section as input. Then dump the contents of that section to the screen in
> hexadecimal format.

Easily done

```go
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
```

## 2. Overriding Weak Symbols

> Some symbols are _weak_, which means that their value may be overridden by
> another symbol that isn’t weak. Currently, the binary loader doesn’t take
> this into account and simply stores all symbols. Expand the binary loader so
> that if a weak symbol is later overridden by another symbol, only the latest
> version is kept. Take a look at _/usr/include/bfd.h_ to figure out the flags
> to check for.

This is already done by the elf library in Golang. Looking at how this is done,
it appears to filter on the top 4 bits of the `Info` field, which is the `BIND`
field. A value of 2 here means that the symbol is weak, and we could just
filter on that.

## 3. Printing Data Symbols

> Expand the binary loader and the `loader_demo` program so that they can
> handle local and global data symbols as well as function symbols. You’ll need
> to add handling for data symbols in the loader, add a new `SymbolType` in the
> `Symbol` class, and add code to the `loader_demo` program to print the data
> symbols to screen. Be sure to test your modifications on a nonstripped binary
> to ensure the presence of some data symbols. Note that data items are called
> objects in symbol terminology. If you’re unsure about the correctness of your
> output, use `readelf` to verify it.

The `BIND` part of the `Info` field has this information, so it's just a matter
of adding a column that shows this information.
