import ctypes
import sys
import os
import subprocess

# from linux/sched.h
CLONE_NEWUTS = 0x04000000

# Load libc
libc = ctypes.CDLL("libc.so.6", use_errno=True)

def setup_namespace():
    # call unshare(CLONE_NEWUTS)
    if libc.unshare(CLONE_NEWUTS) != 0:
        err = ctypes.get_errno()
        sys.stderr.write(f"unshare failed: {os.strerror(err)}\n")
        sys.stderr.flush()
        raise OSError(err, "unshare failed")
    # set hostname
    # os.sethostname(b"isolated")
    os.chroot("/root/con")
    os.chdir("/");


p = subprocess.Popen(["/bin/bash", "-c", "hostname container && exec bash"], preexec_fn=setup_namespace)
p.wait()

