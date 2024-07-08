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

#define SERVER_PORT 80
#define MAXLINE 4096
#define SOCKADDR struct sockaddr

void err_n_die(const char *fmt, ...);

int main(int argc, char **argv)
{
	int sockfd,n;
	int sendbytes;
	struct sockaddr_in servaddr;
	char sendline[MAXLINE];
	char recvline[MAXLINE];

	if (argc != 2)
		err_n_die("usage: %s <server address>",argv[0]);

	if ( (sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0) //create a socket, 0 = tcp, stream sockets to stream data back and forth
		err_n_die("Error while creating the socket!");

	bzero(&servaddr, sizeof(servaddr)); //zero out the address
	servaddr.sin_family = AF_INET;
	servaddr.sin_port = htons(SERVER_PORT); //chat server, hton = host to network short - endianess
	
	if (inet_pton(AF_INET, argv[1], &servaddr.sin_addr) <= 0) //conv bin representation
		err_n_die("inet_pton error for %s ",argv[1]);

	if (connect(sockfd, (SOCKADDR *) &servaddr, sizeof(servaddr)) < 0)
		err_n_die("Connection failed!");

	//next section is over simplified and definately refactorable and extendable
	sprintf(sendline, "GET / HTTP/1.1\r\n\r\n"); //the esc are terminator
	sendbytes = strlen(sendline); //needs to be strnlen needs size_t maxlen
	//normally you want retry for the next section, its apperently fragile			    
	if (write(sockfd, sendline, sendbytes) != sendbytes)
		err_n_die("Write error");

	while ( (n = read(sockfd, recvline, MAXLINE-1)) > 0)
	{
		printf("%s",recvline); //very very unsafe
		memset(recvline, 0, MAXLINE);
	}
	if (n < 0)
		err_n_die("Read error");

	exit(0);
}

