#include <stdio.h>

int main(int argc, char **argv) {
    for (size_t idx = 1; idx < argc; idx++) {
        fputs(argv[idx], stdout);
        putc(' ', stdout);
    }

    putc('\n', stdout);
    return 0;
}