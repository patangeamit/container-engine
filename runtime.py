import sys
import json
import os
import ctypes

# flags
CLONE_NEWUTS = 0x04000000 #hostname
CLONE_NEWNS = 0x00020000 #mounts
CLONE_NEWNET = 0x40000000 #network interfaces
CLONE_NEWPID = 0x20000000 #process
MS_PRIVATE = 0x00040000 
MS_REC = 0x00004000

libc = ctypes.CDLL("libc.so.6", use_errno=True)


def setup_namespace(hostname):
    libc.unshare(CLONE_NEWUTS | CLONE_NEWNS | CLONE_NEWNET | CLONE_NEWPID)
    libc.sethostname(hostname.encode(), len(hostname))
    libc.mount(None, b"/", None, MS_REC | MS_PRIVATE, None)

def child(rootfs, mounts, args):
    os.chroot(rootfs)
    os.chdir("/")
    os.makedirs("proc", exist_ok=True)
    for m in mounts:
        src = m["source"].encode()
        dst = m["destination"].encode()
        fstype = m["type"].encode()
        libc.mount(src, dst, fstype, 0, None)
    # run process
    p = os.execvp(args[0], args)

if __name__ == "__main__":
    bundle = sys.argv[1]
    with open(os.path.join(bundle, "config.json")) as f:
        spec = json.load(f)

    hostname = spec["hostname"]
    rootfs = spec["root"]["path"]
    args = spec["process"]["args"]
    mounts = spec["mounts"]
    setup_namespace(hostname)
    pid = os.fork()
    if pid == 0: # child. fork returns 0 in the child process
        child(rootfs, mounts, args)
    else: # parent 
        os.waitpid(pid, 0)
