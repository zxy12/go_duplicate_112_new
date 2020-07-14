# asm 命令到底干了什么

lexer:

parser:

ctxt

## 依赖的包介绍

### cmd/asm/internal/arch
arch.go:

    定义伪寄存器	RFP = -(iota + 1) RSB、RSP、RPC
    asm.internal.arch -> obj.LinkArch -> sys.Arch
    设置指令集 x86.Anames
    寄存器 x86.Register
    architecture(Arch).linkarch x86.Linkamd64
    ctxt: Link
    Arch.Init -> cmd.internal.obj.x86.asm6.instinit
    Plist:Curfn, Prog 一条机器指令
    As: Opcode
    For example, MOVL R1, R2 encodes using only As=MOVL, From=R1, To=R2.

### cmd/internal/sys/ 
arch.go: 
    
    ArchFamily: byte 环境，指针长度、寄存器长度、大小端



