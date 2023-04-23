#ifndef _COMMON_H_
#define _COMMON_H_

#include <sys/socket.h> 
#include <sys/types.h>
#include <signal.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <stdarg.h> // for variadic function
#include <errno.h>
#include <fcntl.h>
#include <sys/time.h>
#include <sys/ioctl.h>
#include <netdb.h>

//Thanks Jacob Sorber

#define SERVER_PORT 18000
#define MAXLINE 4096
#define SOCKADDR struct sockaddr

void err_n_die(const char *fmt, ...);
char *bin2hex(const unsigned char *input, size_t len);

#endif
