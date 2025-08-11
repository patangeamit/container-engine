import json
import sys
import os
import subprocess

BUNDLES_DIR = "./bundles"

def create_bundle(name, args):
    bundle_path = os.path.join(BUNDLES_DIR, name)
    os.makedirs(bundle_path, exist_ok=True)
    rootfs_path = "rootfs"
    config = {
        "hostname": "container",
        "mounts": [{"destination": "/proc", "type": "proc", "source": "proc"}],
        "root": {"path": rootfs_path},
        "process": {"args": args},
        "linux": {
            "resources": {"pids": {"limit": 5}},
        },
        "name": name
    }
    with open(os.path.join(bundle_path, "config.json"), "w") as f:
        json.dump(config, f, indent=2)
    print(f"[daemon] Created bundle at {bundle_path}")
    return bundle_path


def run_container(name, args):
    bundle = create_bundle(name, args)
    print(f"[daemon] Starting runtime for {name}")
    p = subprocess.Popen(["python3", "runtime.py", bundle])
    exit_code = p.wait()
    print(f"[daemon] Stopped {name} with exit code {exit_code}")


if __name__ == "__main__":

    if len(sys.argv) < 4:
        print("Usage: python3 daemon.py run <name> <cmd...>")
        sys.exit(1)

    if sys.argv[1] == "run":
        name = sys.argv[2]
        cmd = sys.argv[3:]
        run_container(name, cmd)
    else:
        print(f"Command {sys.argv[1]} not found.")