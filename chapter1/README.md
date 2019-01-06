# Chapter 1: Anatomy of a Binary

## 1. Locating Functions

> Write a C program that contains several functions and compile it into an
> assembly file, an object file, and an executable binary, respectively. Try to
> locate the functions you wrote in the assembly file and in the disassembled
> object file and executable. Can you see the correspondence between the C code
> and the assembly code? Finally, strip the executable and try to identify the
> functions again.

```c
#include <stdio.h>

void func1();
void func2(const char *s);

int main(int argc, char* argv[])
{
	func1();
	func2("World");
}

void func1()
{
	printf("This is func 1\n");
}

void func2(const char *s)
{
	printf("Hello %s\n", s);
}
```

```bash
gcc -S -masm=intel dummy.c
```

See _dummy.s_. The functions are easy to identify: Just grep for the function name with a colon after it e.g. `func1:`.

For the object file:

```bash
gcc -c dummy.c
objdump -M intel -d dummy.o
```

```asm
dummy.o:     file format elf64-x86-64


Disassembly of section .text:

0000000000000000 <main>:
   0:   55                      push   rbp
   1:   48 89 e5                mov    rbp,rsp
   4:   48 83 ec 10             sub    rsp,0x10
   8:   89 7d fc                mov    DWORD PTR [rbp-0x4],edi
   b:   48 89 75 f0             mov    QWORD PTR [rbp-0x10],rsi
   f:   b8 00 00 00 00          mov    eax,0x0
  14:   e8 00 00 00 00          call   19 <main+0x19>
  19:   bf 00 00 00 00          mov    edi,0x0
  1e:   e8 00 00 00 00          call   23 <main+0x23>
  23:   b8 00 00 00 00          mov    eax,0x0
  28:   c9                      leave
  29:   c3                      ret

000000000000002a <func1>:
  2a:   55                      push   rbp
  2b:   48 89 e5                mov    rbp,rsp
  2e:   bf 00 00 00 00          mov    edi,0x0
  33:   e8 00 00 00 00          call   38 <func1+0xe>
  38:   90                      nop
  39:   5d                      pop    rbp
  3a:   c3                      ret

000000000000003b <func2>:
  3b:   55                      push   rbp
  3c:   48 89 e5                mov    rbp,rsp
  3f:   48 83 ec 10             sub    rsp,0x10
  43:   48 89 7d f8             mov    QWORD PTR [rbp-0x8],rdi
  47:   48 8b 45 f8             mov    rax,QWORD PTR [rbp-0x8]
  4b:   48 89 c6                mov    rsi,rax
  4e:   bf 00 00 00 00          mov    edi,0x0
  53:   b8 00 00 00 00          mov    eax,0x0
  58:   e8 00 00 00 00          call   5d <func2+0x22>
  5d:   90                      nop
  5e:   c9                      leave
  5f:   c3                      ret
```

Note that the calls to `printf` look incorrect: The static relocations have not happened because this file has not yet been linked.

```bash
gcc -o dummy dummy.c
strip dummy
objdump -M intel -d dummy
```

After stripping the binary it's harder to find the functions, since there's no symbol to mark the start point. We can find the functions by looking for `ret` instructions, though this isn't super reliable since some of the other functions end with `jmp`s.  Since in my code one of the call to `printf` becomes `puts` since there's no format string, we can look for the only call to `printf` to find `func2`. Backtracing to the start of this function and noting the address, we can find a call to this address to identify the call site, which we know happens in main.

## 2. Sections

> As you’ve seen, ELF binaries (and other types of binaries) are divided into
> sections. Some sections contain code, and others contain data. Why do you
> think the distinction between code and data sections exists? How do you think
> the loading process differs for code and data sections? Is it necessary to
> copy all sections into memory when a binary is loaded for execution?

Since there are different types of data, and we don't want executable code to also be writable (in most cases), to prevent attackers from easily patching code directly if they have a mechanism to write arbitrary memory, it makes sense to have different sections mapped with different permissions. Code should be RX, global data should be RW, and constants should be RO. Sections like BSS, which is initialized to 0, do not need to actually have any data in the file in disk, they just need to be given blank memory.

Sections are created during compilation and placed in the ELF during linking. They do not all need to be used in a running process. For example, the `.shrtrtab` section is used by the loader, but probably doesn't need to be used by anything once the process has been mapped into the process memory space.
