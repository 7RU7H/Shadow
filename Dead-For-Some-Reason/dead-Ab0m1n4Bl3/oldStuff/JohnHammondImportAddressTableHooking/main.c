#include <structs.h>
#


if (argc < 3) {
	wprintf(L"[!] Usage : \"%s\" <Complete Dll Payload Path> <Process Name> \n", argv[0]);
	return -1;
}

wprintf(L"[i] Searching For Process Id Of \"%s\" ... ", argv[2]);
if (!GetRemoteProcessHandle(argv[2], &dwProcessId, &hProcess)) {
	printf("[!] Process is Not Found \n");
	return -1;
}

wprintf(L"[+] Done Getting handle of a remote process\n");