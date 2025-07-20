import ctypes
import sys
import os
import subprocess

# from linux/sched.h
# flag for hostname namespace
CLONE_NEWUTS = 0x04000000
#flag for mount namespace
CLONE_NEWNS = 0x00020000

# flag for changing propagation mode of existing mount, saying do not share mounts with parent (host)
MS_PRIVATE = 0x00040000
# flag for applying the rule recursively for all sub-mounts
MS_REC = 0x00004000

# Load libc
libc = ctypes.CDLL("libc.so.6", use_errno=True)

def setup_namespace():
    # call unshare(CLONE_NEWUTS)
    libc.unshare(CLONE_NEWUTS)
    libc.unshare(CLONE_NEWNS)

    libc.mount(None, b"/", None, MS_REC | MS_PRIVATE, None)

    # set hostname
    # os.sethostname(b"isolated")
    os.chroot("./rootfs")
    os.chdir("/");


p = subprocess.Popen(["/bin/sh", "-c", "hostname container && exec sh"], preexec_fn=setup_namespace)
p.wait()

