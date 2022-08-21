#include <stdio.h>
#include <stdlib.h>

#define RESULT_SUCCESS        0
#define RESULT_FILE_NOT_FOUND 1
#define RESULT_ALLOC_ERROR    2

typedef struct cat_file {
    char *content;
    unsigned short success;
} cat_file;

void inf_stdin(void);
cat_file read_file(const char* path);

int main(int argc, char **argv) {
    if (argc == 1) {
        inf_stdin();
    } else if (argc > 1) {
        for (size_t idx = 1; idx < argc; idx++) {
            const char *file = argv[idx];
            cat_file result = read_file(file);

            switch (result.success) {
                case RESULT_SUCCESS:
                    fputs(result.content, stdout);
                    free(result.content);
                    break;
                case RESULT_ALLOC_ERROR:
                    return 1;
                default:
                    break;
            }
        }
    }

    return 0;
}

void inf_stdin(void) {
    int got;
    
    for (;;) {
        got = getchar();
        putchar(got);
    }
}

cat_file read_file(const char* path) {
    size_t lpos;

    FILE *file = fopen(path, "r");
    if (file == NULL) {
        cat_file error = { .content = NULL, .success = 1 };
        return error;
    }

    fseek(file, 0, SEEK_END);
    lpos = ftell(file);
    rewind(file);

    char *buffer = calloc(lpos + 1, sizeof(char));
    if (buffer == NULL) {
        fclose(file);

        cat_file error = { .content = NULL, .success = 2 };
        return error;
    }

    fread(buffer, lpos, 1, file);
    fclose(file);

    cat_file error = { .content = buffer, .success = 0 };
    return error;
}
