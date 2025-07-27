#define _GNU_SOURCE
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

int main() {
    // Unshare the UTS (hostname) namespace
    if (unshare(CLONE_NEWUTS) == -1) {
        perror("unshare");
        exit(EXIT_FAILURE);
    }

    printf("[+] UTS namespace unshared!\n");

    // Optional: set a test hostname
    const char *newname = "mytest";
    if (sethostname(newname, strlen(newname)) == -1) {
        perror("sethostname");
        exit(EXIT_FAILURE);
    }
    printf("[+] Hostname inside namespace set to: %s\n", newname);

    // Exec a shell so you can play inside this namespace
    printf("[+] Launching bash shell in new UTS namespace...\n");
    execlp("bash", "bash", NULL);

    // If execlp fails:
    perror("execlp");
    return 1;
}
