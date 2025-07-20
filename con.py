import ctypes
import sys
import os
import subprocess

# flag for hostname namespace. from /usr/include/linux/sched.h
CLONE_NEWUTS = 0x04000000
#flag for mount namespace. from /usr/include/linux/sched.h
CLONE_NEWNS = 0x00020000
#flag for network namespace. from /usr/include/linux/sched.h
CLONE_NEWNET = 0x40000000

# flag for changing propagation mode of existing mount, saying do not share mounts with parent (host)
MS_PRIVATE = 0x00040000
# flag for applying the rule recursively for all sub-mounts
MS_REC = 0x00004000

# Load libc
libc = ctypes.CDLL("libc.so.6", use_errno=True)

def setup_namespace():

    # isolate hostname
    libc.unshare(CLONE_NEWUTS)
    name = b"container"
    libc.sethostname(name, len(name))
    
    # isolate mounts
    libc.unshare(CLONE_NEWNS)
    libc.mount(None, b"/", None, MS_REC | MS_PRIVATE, None)

    # isolate network
    libc.unshare(CLONE_NEWNET)

    
def mount():
    ret = libc.mount(b"proc", b"/proc", b"proc", 0, None)
    
def setup_chroot():
    os.chroot("./rootfs")
    os.chdir("/")
    
def setup():
    try:
        setup_namespace()
        setup_chroot()
        mount()
    except Exception as e:
        print(e)
    
p = subprocess.Popen(["/bin/sh"], preexec_fn=setup)
p.wait()

