# Chapter 3: The PE Format

## 1. Manual Header Inspection

> Just as you did for ELF binaries in Chapter 2, use a hex viewer like `xxd` to
> view the bytes in a PE binary. You can use the same command as before, `xxd
> program.exe | head -n 30`, where program.exe is your PE binary. Can you
> identify the bytes representing the PE header and make sense of all of the
> header fields?

```
[MS-DOS HEADER]
00000000: 4d5a 9000 0300 0400 0000 0000 ffff 0000  MZ.............. Magic = "MZ"
00000010: 8b00 0000 0000 0000 4000 0000 0000 0000  ........@.......
00000020: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000030: 0000 0000 0000 0000 0000 0000 8000 0000  ................ PE offset = 0x80
00000040: 0e1f ba0e 00b4 09cd 21b8 014c cd21 5468  ........!..L.!Th
00000050: 6973 2070 726f 6772 616d 2063 616e 6e6f  is program canno
00000060: 7420 6265 2072 756e 2069 6e20 444f 5320  t be run in DOS
00000070: 6d6f 6465 2e0d 0d0a 2400 0000 0000 0000  mode....$.......
[PE headers] Starts at 0x80 as e_lfanew says
[PE signature + PE file header]
00000080: 5045 0000 6486 0e00 0000 0000 00b4 1b00  PE..d........... Magic = "PE\x00\x00"
                                                                    Machine = 0x8664 (x86_64)
                                                                    Number of sections = 0x0e
                                                                    Time/date stamp = 0
                                                                    Pointer to symbol table = 0x001bb400
00000090: 070d 0000 f000 2302                      ......#......... Number of symbols = 0x0d07
                                                                    Size of optional header = 0x00f0
                                                                    Characteristics = 0x0223
[OPTIONAL HEADER]
00000090:                     0b02 0300 00d2 0800  ......#......... Magic = 0x020b (PE32+)
                                                                    Major linker version = 0x03
                                                                    Minor linker version = 0x00
                                                                    Size of code = 0x0008d200
000000a0: 0032 0100 0000 0000 8024 0500 0010 0000  .2.......$...... Size of initialized data = 0x00013200
                                                                    Size of uninitialized data = 0
                                                                    Address of entry point = 0x00052480
                                                                    Base of code = 0x00001000
000000b0: 0000 4000 0000 0000 0010 0000 0002 0000  ..@............. Image base = 0x0000000000400000
                                                                    Section alignment = 0x00001000
                                                                    File alignment = 0x00000200
000000c0: 0400 0000 0100 0000 0400 0000 0000 0000  ................ Major OS system version = 0x0004
                                                                    Minor OS system version = 0x0000
                                                                    Major image version = 0x0001
                                                                    Minor image version = 0x0000
                                                                    Major subsystem version = 0x0004
                                                                    Minor subsystem version = 0x0000
                                                                    Win32 version = 0x00000000
000000d0: 0030 2000 0006 0000 0000 0000 0300 0000  .0 ............. Size of image = 0x00203000
                                                                    Size of headers = 0x00000600
                                                                    Checksum = 0x00000000
                                                                    Subsystem = 0x0003 (Windows CUI)
                                                                    Dll characteristics = 0x0000
000000e0: 0000 2000 0000 0000 0010 0000 0000 0000  .. ............. Size of stack reserve = 0x0000000000200000
                                                                    Size of stack commit = 0x0000000000001000
000000f0: 0000 1000 0000 0000 0010 0000 0000 0000  ................ Size of heap reserve = 0x0000000000100000
                                                                    Size of heap commit = 0x00000000001000
00000100: 0000 0000 1000 0000                      ................ Loader flags = 0x00000000
                                                                    Number of RVA and sizes = 0x00000010
[DATA DIRECTORY]
00000100:                     0000 0000 0000 0000  ................ Entry 0 (Export)
                                                                    RVA = 0x00000000
                                                                    Size = 0x00000000
00000110: 00f0 1d00 1604 0000                      ................ Entry 1 (Import)
                                                                    RVA = 0x001df000
                                                                    Size = 0x00000416
00000110:                     0000 0000 0000 0000  ................ Entry 2 (Resource)
00000120: 0000 0000 0000 0000 0000 0000 0000 0000  ................ Entry 3 (Exception) & 4 (Security)
00000130: 0000 0000 0000 0000 0000 0000 0000 0000  ................ Entry 5 (Reloc) 6...
00000140: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000150: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000160: 0000 0000 0000 0000 0080 1400 2001 0000  ............ ...
00000170: 0000 0000 0000 0000 0000 0000 0000 0000  ................
00000180: 0000 0000 0000 0000                      .........text...
[SECTIONS]
00000180:                     2e74 6578 7400 0000  .........text... Name = ".text"
00000190: 73d1 0800 0010 0000 00d2 0800 0006 0000  s............... VirtualSize = 0x0008d173
                                                                    RVA = 0x00001000
                                                                    Size of raw data = 0x0008d200
                                                                    Pointer to raw data = 0x00000600
000001a0: 0000 0000 0000 0000 0000 0000 6000 0060  ............`..` Pointer to relocations = 0
                                                                    Pointer to line numbers = 0
                                                                    Number of relocations = 0
                                                                    Number of line numbers = 0
                                                                    Characteristics = 0x60000060

000001b0: 2e72 6461 7461 0000 e58b 0b00 00f0 0800  .rdata.......... Name = ".rdata"
                                                                    VirtualSize = 0x000b8be5
                                                                    RVA = 0x0008f000
000001c0: 008c 0b00 00d8 0800 0000 0000 0000 0000  ................ Size of raw data = 0x000b8c00
                                                                    Pointer to raw data = 0x0008d800
                                                                    Pointer to relocations = 0
                                                                    Pointer to line numbers = 0
000001d0: 0000 0000 4000 0060                      ....@..`.data... Number of relocations = 0
                                                                    Number of line numbmers = 0
                                                                    Characteristics = 0x60000040

000001d0:                     2e64 6174 6100 0000  ....@..`.data... ...
000001e0: b80b 0300 0080 1400 0032 0100 0064 1400  .........2...d..
000001f0: 0000 0000 0000 0000 0000 0000 4000 00c0  ............@...
```

## 2. Disk Representation vs. Memory Representation

> Use `readelf` to view the contents of a PE binary. Then make an illustration
> of the binary’s on-disk representation versus its representation in memory.
> What are the major differences?

Unlike with ELF binaries, there is no concept of segments in a PE binary, just sections. Although in the binaries I looked at the section seem to be in the binary in the same order as they would be loaded in memory, that shouldn't be a necessity. Like with ELF, a section does not need to be as large on disk as it will be in memory, for example for the .bss section, which would just be filled with zeroes. This allows the binary to be much smaller than the process loaded in memory.

## 3. PE vs. ELF

> Use `objdump` to disassemble an ELF and a PE binary. Do the binaries use
> different kinds of code and data constructs? Can you identify some code or
> data patterns that are typical for the ELF compiler and the PE compiler
> you’re using, respectively?
