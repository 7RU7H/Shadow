#include "GetModuleHandle.h"
#include "utilities.h"

const char* DEFAULT_EXT = ".DLL";

HMODULE GetMOduleHandleAReplacement(IN LPCSTR szModuleName) {

#ifdef _WIN64
	PPEB pPeb = (PEB*)(__readgsqword(0x60)); // Process Environment Block
#elif _WIN32
	PPEB pPeb = (PEB*)(__readgsqword(0x30));
#endif

	if (szModuleName = NULL) {
		return pPeb->ImageBaseAddress;
	}

	LPCSTR ext = strchr(szModuleName, '.');
	if (!ext) {

		SIZE_T len = strlen(szModuleName);
		SIZE_T extLen = strlen(DEFAULT_EXT);
		SIZE_T newLen = extLen + len + 1;
		LPCSTR temp = calloc(newLen, sizeof(char));

		if (temp == NULL) {
			printf("calloc failed");
			return NULL;
		}

		errno_t err = strncpy_s(temp, newLen, szModuleName, len);
		if (err) {
			printf("strncpy_s failed");
			return NULL;
		}

		err = strncat_s(temp, newLen, DEFAULT_EXT, extLen);
		if (err) {
			printf("strncat_s failed");
			return NULL;
		}

		szModuleName = temp;
		}

	PPEB_LDR_DATA pLdr = (PPEB_LDR_DATA)pPeb->Ldr;
	PTRUNCATED_LDR_DATA_TABLE_ENTRY pDte = (PTRUNCATED_LDR_DATA_TABLE_ENTRY)pLdr->InMemoryOrderLinks;

	HMODULE hModule = NULL;

	while (pDte) {
		
		if (pDte->BaseDllName.Length != NULL) {

			if (IsWStrEqualCStr(pDte->BaseDllName.Buffer, szModuleName)) {
				wprintf(L"[+] Found Dll \"%s\" \n", pDte->BaseDllName.Buffer);
				hModule = pDte->DllBase;
				break;
			}
			wprintf(L"[i] \"%s\" \n", pDte->FullDllName.Buffer);

		}
		else {
			break;
		}
		pDte = pDte->InMemoryOrderLinks.Flink;
	}

	if (!ext) {
		free(szModuleName);
	}
	return hModule;
}