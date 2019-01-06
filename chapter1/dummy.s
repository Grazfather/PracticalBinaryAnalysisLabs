	.file	"dummy.c"
	.intel_syntax noprefix
	.section	.rodata
.LC0:
	.string	"World"
	.text
	.globl	main
	.type	main, @function
main:
.LFB0:
	.cfi_startproc
	push	rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	mov	rbp, rsp
	.cfi_def_cfa_register 6
	sub	rsp, 16
	mov	DWORD PTR [rbp-4], edi
	mov	QWORD PTR [rbp-16], rsi
	mov	eax, 0
	call	func1
	mov	edi, OFFSET FLAT:.LC0
	call	func2
	mov	eax, 0
	leave
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE0:
	.size	main, .-main
	.section	.rodata
.LC1:
	.string	"This is func 1"
	.text
	.globl	func1
	.type	func1, @function
func1:
.LFB1:
	.cfi_startproc
	push	rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	mov	rbp, rsp
	.cfi_def_cfa_register 6
	mov	edi, OFFSET FLAT:.LC1
	call	puts
	nop
	pop	rbp
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE1:
	.size	func1, .-func1
	.section	.rodata
.LC2:
	.string	"Hello %s\n"
	.text
	.globl	func2
	.type	func2, @function
func2:
.LFB2:
	.cfi_startproc
	push	rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	mov	rbp, rsp
	.cfi_def_cfa_register 6
	sub	rsp, 16
	mov	QWORD PTR [rbp-8], rdi
	mov	rax, QWORD PTR [rbp-8]
	mov	rsi, rax
	mov	edi, OFFSET FLAT:.LC2
	mov	eax, 0
	call	printf
	nop
	leave
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE2:
	.size	func2, .-func2
	.ident	"GCC: (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609"
	.section	.note.GNU-stack,"",@progbits
