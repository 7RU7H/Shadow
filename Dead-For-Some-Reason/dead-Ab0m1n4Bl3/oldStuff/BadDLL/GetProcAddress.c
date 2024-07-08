#include "GetProcAddress.h"

FARPROC GetProcAddressReplacement(IN HMODULE hModule, IN LPCSTR lpApiName) {

	PBYTE pBase = (PBYTE)hModule;

	// Is it a real executable
	PIMAGE_DOS_HEADER pImgDosHdr = (PIMAGE_DOS_HEADER)pBase;
	if (pImgDosHdr->e_magic != IMAGE_DOS_SIGNATURE) {
		return NULL;
	}

	PIMAGE_NT_HEADERS pImgNtHdrs = (PIMAGE_NT_HEADERS)(pBase + pImgDosHdr->e_lfanew);
	if (pImgNtHdrs->Signature != IMAGE_NT_SIGNATURE) {
		return NULL;
	}

	IMAGE_OPTIONAL_HEADER ImgOptHdr = pImgNtHdrs->OptionalHeader;

	IMAGE_DATA_DIRECTORY exportDataDir = ImgOptHdr.DataDirectory[IMAGE_DIRECTORY_ENTRY_EXPORT];
	PIMAGE_EXPORT_DIRECTORY pImgExportDir = (PIMAGE_EXPORT_DIRECTORY)(pBase + exportDataDir.VirtualAddress);

	PDWORD FunctionAddressArray = (PDWORD)(pBase + pImgExportDir->AddressOfFunctions);
	
	PVOID pFunctionAddress = NULL;

	if (lpApiName <= 0xFFFF) {
		WORD ordinal = (WORD)lpApiName & 0xFFFF; // ordinals are word size so we get the first two bytes
		DWORD base = pImgExportDir->Base;

		if (ordinal < base || (ordinal >= base + pImgExportDir->NumberOfFunctions)) {
			return NULL;
		}

		pFunctionAddress = (PVOID)(pBase + FunctionAddressArray[ordinal - base]);

	} 
	else {
		PWORD FunctionNameArray = (PWORD)(pBase + pImgExportDir->AddressOfNames);
		PWORD FunctionOrdinalArray = (PWORD)(pBase + pImgExportDir->AddressOfNameOrdinals);

		for (DWORD i = 0; i < pImgExportDir->NumberOfFunctions; i++) {
			CHAR* pFunctionName = (CHAR*)(pBase + FunctionNameArray[i]);

			if (strcmp(lpApiName, pFunctionName) == 0) {
				pFunctionAddress = (PVOID)(pBase + FunctionAddressArray[FunctionOrdinalArray[i]]);
			}
		}
	}

	if (pFunctionAddress >= (PBYTE)pImgExportDir && pFunctionAddress < (PBYTE)pImgExportDir + exportDataDir.Size) {
		SIZE_T len = strlen(pFunctionAddress) + 1;

		char* moduleName = calloc(len, sizeof(char));
		strncpy_s(moduleName, len, pFunctionAddress, len);

		char* function = strchr(moduleName, '.');
		*function = 0;
		function++;

		pFunctionAddress = GetProcAddressReplacement(GetModuleHandleAReplacement(moduleName), lpApiName);

		free(moduleName);
	}

	return pFunctionAddress;
}
