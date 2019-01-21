# Chapter 5: Anatomy of a Binary

* [Notes](notes.md)

## 1. A New CTF Challenge

> Complete the new CTF challenge unlocked by the oracle program! You can
> complete the entire challenge using only the tools discussed in this chapter
> and what you learned in Chapter 2. After completing the challenge, don’t
> forget to give the flag you found to the oracle to unlock the next challenge.

Running `oracle` with the flag from level 1, a new binary is dumped from
_levels.db_. The flag is probably used both to lookup and decrypt the challenge
from the database, but I don't need to figure this out yet. A new binary,
`lvl2` is dropped to disk.

```bash
[root@gpwn:~]$ ls -lh lvl2
-rwxr-xr-x 1 root root 6.4K Jan 21 20:35 lvl2
[root@gpwn:~]$ ldd ./lvl2
        linux-vdso.so.1 =>  (0x00007fff579d2000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fc927364000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fc92772e000)
[root@gpwn:~]$ nm -D lvl2
                 w __gmon_start__
                 U __libc_start_main
                 U puts
                 U rand
                 U srand
                 U time
```

Running it a few times on its own it seems to just print a random hex-encoded
byte. Looking at it with `ltrace` we see probably where this comes from.

```bash
[root@gpwn:~]$ ltrace ./lvl2
__libc_start_main(0x400500, 1, 0x7ffecd79bd68, 0x400640 <unfinished ...>
time(0)                                                                                          = 1548103348
srand(0x5c462eb4, 0x7ffecd79bd68, 0x7ffecd79bd78, 0)                                             = 0
rand(0x7fea0eb18620, 0x7ffecd79bc4c, 0x7fea0eb180a4, 0x7fea0eb1811c)                             = 0x273d4d7c
puts("81"81
)                                                                                       = 3
+++ exited (status 0) +++
```

Although `ltrace` seems to get the arguments wrong (`srand` only takes one
argument), we can see that the return value from `time(0)` is fed to `srand`:
1548103348 == 0x5c462eb4.

Disassembling `main` (at 0x400500, the first arg to `__libc_start_main`), we
see it's very simple:

```asm
0000000000400500 <.text>:
  400500:       48 83 ec 08             sub    rsp,0x8
  400504:       31 ff                   xor    edi,edi
  400506:       e8 c5 ff ff ff          call   4004d0 <time@plt>
  40050b:       89 c7                   mov    edi,eax
  40050d:       e8 ae ff ff ff          call   4004c0 <srand@plt>
  400512:       e8 c9 ff ff ff          call   4004e0 <rand@plt>
  400517:       99                      cdq
  400518:       c1 ea 1c                shr    edx,0x1c
  40051b:       01 d0                   add    eax,edx
  40051d:       83 e0 0f                and    eax,0xf
  400520:       29 d0                   sub    eax,edx
  400522:       48 98                   cdqe
  400524:       48 8b 3c c5 60 10 60    mov    rdi,QWORD PTR [rax*8+0x601060]
  40052b:       00
  40052c:       e8 6f ff ff ff          call   4004a0 <puts@plt>
  400531:       31 c0                   xor    eax,eax
  400533:       48 83 c4 08             add    rsp,0x8
  400537:       c3                      ret
```

The string printed is just a random index into the string table at 0x601060.
Let's look at that table:

```bash
[root@gpwn:~]$ objdump -s --section .data lvl2

lvl2:     file format elf64-x86-64

Contents of section .data:
 601040 00000000 00000000 00000000 00000000  ................
 601050 00000000 00000000 00000000 00000000  ................
 601060 c4064000 00000000 c7064000 00000000  ..@.......@.....
 601070 ca064000 00000000 cd064000 00000000  ..@.......@.....
 601080 d0064000 00000000 d3064000 00000000  ..@.......@.....
 601090 d6064000 00000000 d9064000 00000000  ..@.......@.....
 6010a0 dc064000 00000000 df064000 00000000  ..@.......@.....
 6010b0 e2064000 00000000 e5064000 00000000  ..@.......@.....
 6010c0 e8064000 00000000 eb064000 00000000  ..@.......@.....
 6010d0 ee064000 00000000 f1064000 00000000  ..@.......@.....
```

These are all pointers to strings at 0x004006c0+, which, looking at the section
header, we see is the `.rodata` section. Let's dump that.

```bash
[root@gpwn:~/code/PracticalBinaryAnalysisLabs/code/chapter5]$ objdump -s --section .rodata lvl2

lvl2:     file format elf64-x86-64

Contents of section .rodata:
 4006c0 01000200 30330034 66006334 00663600  ....03.4f.c4.f6.
 4006d0 61350033 36006632 00626600 37340066  a5.36.f2.bf.74.f
 4006e0 38006436 00643300 38310036 63006466  8.d6.d3.81.6c.df
 4006f0 00383800                             .88.
```

Flag: 034fc4f6a536f2bf74f8d6d3816cdf88
