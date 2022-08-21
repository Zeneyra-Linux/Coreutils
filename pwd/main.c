#include <stdio.h>

#if defined(__unix__)
#include <unistd.h>
#include <limits.h>
#elif defined(_WIN32)
#include <windows.h>
#endif

int main() {
    #if defined(__unix__)
        char path[PATH_MAX + 1] = {0};
        getcwd(path, sizeof(path));
        puts(path);
    #elif defined(_WIN32)
        char path[MAX_PATH + 1] = {0};
        GetCurrentDirectory(MAX_PATH, path);
        puts(path);
    #else
        return 1;
    #endif

    return 0;
}
