#include "IATHook.h"
#include "GetModuleHandle.h"
#include "GetProcAddress.h"

PVOID HookIAT(PBYTE pTarget, PCSTR lpModuleName, PCWSTR lpApiName, PVOID replacement) {
	// Is it a real executable
	PIMAGE_DOS_HEADER pImgDosHdr = (PIMAGE_DOS_HEADER)pTarget;
	if (pImgDosHdr->e_magic != IMAGE_DOS_SIGNATURE) {
		return NULL;
	}

	PIMAGE_NT_HEADERS pImgNtHdrs = (PIMAGE_NT_HEADERS)(pTarget + pImgDosHdr->e_lfanew);
	if (pImgNtHdrs->Signature != IMAGE_NT_SIGNATURE) {
		return NULL;
	}

	IMAGE_OPTIONAL_HEADER ImgOptHdr = pImgNtHdrs->OptionalHeader;

	IMAGE_DATA_DIRECTORY impDataDir = ImgOptHdr.DataDirectory[IMAGE_DIRECTORY_ENTRY_IMPORT];

	PIMAGE_IMPORT_DESCRIPTOR pImportAddressTable = (PIMAGE_IMPORT_DESCRIPTOR)(pTarget + impDataDir.VirtualAddress);

	SIZE_T iatSize = impDataDir.Size / sizeof(IMAGE_IMPORT_DESCRIPTOR);

	BOOL found = FALSE;

	for (SIZE_T i = 0; i < iatSize; i++) {
		char* pModuleName = pTarget + pImportAddressTable[i].Name;
		if (IsEqualCStr(lpModuleName, pModuleName)) {
			if (PatchIATEntry(pTarget,lpApiName, &pImportAddressTable[i], replacement)) {
				found = TRUE;
		}
	}

	if (found) {
		return GetProcAddressReplacement(GetModuleHandle(lpModuleName), lpApiName);
	}

	return NULL;
}

const int PAGE_SIZE = 4096;

BOOL PatchIATEntry(PBYTE pTarget, PCSTR lpApiName, PIMAGE_IMPORT_DESCRIPTOR pModuleEntry, PVOID replacement) {

	PULONG_PTR originalThunk = pTarget + pModuleEntry->OriginalFirstThunk;
	PULONG_PTR thunk = pTarget + pModuleEntry->FirstThunk;

	BOOL found = FALSE;
	while (*originalThunk != NULL) {
		PIMAGE_IMPORT_BY_NAME importByName = pTarget + *originalThunk;

		if (strcmp(importByName->Name, lpApiName) == 0) {
			found = TRUE;

			DWORD protect = 0;
			VirtualProtect(thunk, PAGE_SIZE, PAGE_READWRITE, &protect);
			*thunk = replacement;
			VirtualProtect(thunk, PAGE_SIZE, protect, &protect);
		}

		originalThunk++;
		thunk++;
	}

	return found;
}