# Chapter 2: The Elf Format

## 1. Manual Header Inspection

> Use a hex viewer such as `xxd` to view the bytes in an ELF binary in
> hexadecimal format. For example, you can use the command `xxd /bin/ls | head
> -n 30` to view the first 30 lines of bytes for the `/bin/ls` program. Can you
> identify the bytes representing the ELF header? Try to find all of the ELF
> header fields in the `xxd` output and see whether the contents of those
> fields make sense to you.

Using a dump from my `dummy` binary:

```
[ELF HEADER]
00000000: 7f45 4c46 0201 0100 0000 0000 0000 0000  .ELF............ ELF magic
00000010: 0200 3e00 0100 0000 7004 4000 0000 0000  ..>.....p.@..... Type = 2 (ELFCLASS64)
                                                                    Machine = 0x3e EM_X86_64
                                                                    Version = 1 (current)
                                                                    Entrypoint = 0x400470
00000020: 4000 0000 0000 0000 481a 0000 0000 0000  @.......H....... Program header offset = 0x40
                                                                    Section header offset = 0x1a48
00000030: 0000 0000 4000 3800 0900 4000 1f00 1c00  ....@.8...@..... Flags = 0
                                                                    ELF header size = 0x40
                                                                    Program header table entry size = 0x38
                                                                    Program header table entry count = 9
                                                                    Section header table entry size = 0x40
                                                                    Section header table entry count = 0x1f
                                                                    Section header string table index = 0x1c
[PROGRAM HEADERS] - Starts at 0x40 as e_phoff shows
00000040: 0600 0000 0500 0000 4000 0000 0000 0000  ........@....... Type = 0x06 (PHDR)
                                                                    Flags = 0x05 (RX)
                                                                    File offset = 0x40
00000050: 4000 4000 0000 0000 4000 4000 0000 0000  @.@.....@.@..... Virtual address = 0x400040
                                                                    Physical address = 0x400040
00000060: f801 0000 0000 0000 f801 0000 0000 0000  ................ Size in file = 0x01f8
                                                                    Size in memory = 0x01f8
00000070: 0800 0000 0000 0000                                       Alignment = 0x08

00000070:                     0300 0000 0400 0000  ................ Type = 0x03 (INTERP)
                                                                    Flags = 0x04 (R)
00000080: 3802 0000 0000 0000 3802 4000 0000 0000  8.......8.@..... File offset = 0x238
                                                                    Virtual address = 0x400238
00000090: 3802 4000 0000 0000 1c00 0000 0000 0000  8.@............. Physical address =0x400238
                                                                    Size in file = 0x1c
000000a0: 1c00 0000 0000 0000 0100 0000 0000 0000  ................ Size in memory = 0x1c
                                                                    Alignment = 0x01

000000b0: 0100 0000 0500 0000 0000 0000 0000 0000  ................ Type = 0x01 (LOAD)
                                                                    Flags = 0x05 (RX)
                                                                    File offset = 0
000000c0: 0000 4000 0000 0000 0000 4000 0000 0000  ..@.......@..... Virtual address = 0x400000
                                                                    Physical address = 0x400000
000000d0: ec07 0000 0000 0000 ec07 0000 0000 0000  ................ Size in file = 0x7ec
                                                                    Size in memory = 0x7ec
000000e0: 0000 2000 0000 0000                                       Alignment = 0x02
...
[SECTION HEADERS] - Starts at 0x1a48 as e_shoff shows
00001a40:                     0000 0000 0000 0000  ................ First section is all zeroes
00001a50: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00001a60: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00001a70: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00001a80: 0000 0000 0000 0000
00001a80:                     1b00 0000 0100 0000  ................ Name offset = 0x1b (.interp)
                                                                    Type = 0x01 (PROGBITS)
00001a90: 0200 0000 0000 0000 3802 4000 0000 0000  ........8.@..... Flags = 0x02 (ALLOC)
                                                                    Virtual address = 0x400238
00001aa0: 3802 0000 0000 0000 1c00 0000 0000 0000  8............... File offset = 0x238
                                                                    Size = 0x1c
00001ab0: 0000 0000 0000 0000 0100 0000 0000 0000  ................ Link = 0
                                                                    Info = 0
                                                                    Alignment = 0x1
00001ac0: 0000 0000 0000 0000                                       Entry size = 0

00001ac0:                     2300 0000 0700 0000  ........#....... Name offset = 0x23 (.note.ABI-tag)
                                                                    Type = 0x07 (NOTE)
00001ad0: 0200 0000 0000 0000 5402 4000 0000 0000  ........T.@..... Flags = 0x02 (ALLOC)
                                                                    Virtual address = 0x400254
00001ae0: 5402 0000 0000 0000 2000 0000 0000 0000  T....... ....... File offset = 0x254
                                                                    Size = 0x20
00001af0: 0000 0000 0000 0000 0400 0000 0000 0000  ................ Link = 0
                                                                    Info = 0
                                                                    Alignment = 0x4
00001b00: 0000 0000 0000 0000                                       Entry size = 0
```

## 2. Sections and Segments

> Use `readelf` to view the sections and segments in an ELF binary. How are the
> sections mapped into segments? Make an illustration of the binary’s on-disk
> representation versus its representation in memory. What are the major
> differences?

The segments are mapped based on their address and flags, and the loader only
has to worry about mapping and copying them as appropriate. The sections were
already mapped into the proper location in the file at link time to overlap
with the segment to which they belong.

## 3. C and C++ Binaries

> Use `readelf` to disassemble two binaries, namely a binary produced from C
> source and one produced from C++ source. What differences are there?

Simply converting _dummy.c_ from chapter 1 to C++, to make the comparison easy.

```cpp
#include <iostream>

void func1();
void func2(std::string s);

int main(int argc, char* argv[])
{
    func1();
    func2("World");
}

void func1()
{
    std::cout << "This is func 1" << std::endl;
}

void func2(std::string s)
{
    std::cout << "Hello " << s << std::endl;
}
```

The code is mostly the same, same sections and segments. The real only
different I noticed is that the imported function are different, and because
C++ support polymorphism, the names are mangled. Using `ldd` to see which
libraries are imported, the C++ version imports a few more: libstdc++, libgcc,
and libc.

## 4. Lazy Binding
> Use 'objdump` to disassemble the PLT section of an ELF binary. Which GOT
> entries do the PLT stubs use? Now view the contents of those GOT entries
> (again with objdump) and analyze their relationship with the PLT.

Each PLT entry starts with a jump a different 'slot' of the GOT, and the GOT is
prepopulated with the address immediately after the jump, which jumps to a
special slot at the start of the PLT, after an index arg is pushed onto the
stack. This special slot jumps to a resolving function, which both patches the
GOT with the correct address of the function, and finally jumps there.
Subsequent calls to this PLT entry does the same first step, jumping to the
appropriate step in the GOT, but at this point the GOT has been modified, and
the GOT entry now points directly to the function. The push and jump after the
jump in this PLT entry should never be executed again.
