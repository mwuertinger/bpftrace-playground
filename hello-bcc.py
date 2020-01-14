#!/usr/bin/python

# Simple program to demonstrate how to use BCC.
# Prints "Hello BPF" on all clone() syscalls.

from bcc import BPF

program = """
    int hello(void *ctx) {
        bpf_trace_printk("Hello BPF\\n");
        return 0;
    }
"""

b = BPF(text=program)
b.attach_kprobe(event="__x64_sys_clone", fn_name="hello")
b.trace_print()
