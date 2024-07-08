#include "utilities.h"
#include <stdio.h>

void EnableDebugConsole() {
	if (AllocConsole()) {
		FILE* fpstdout = stdout;
		FILE* fpstderr = stderr;
		freopen_s(&fpstdout, "CONOUT$", "w", stdout);
		freopen_s(&fpstderr, "CONOUT$", "w", stderr);
		SetWindowText(GetConsoleWindow(), L"Our Malware");
	}
}


BOOL IsEqualCStr(IN LPCSTR Str1, IN LPCSTR Str2) {

	WCHAR lStr1[MAX_PATH], lStr2[MAX_PATH];
	int lenStr1 = lstrlenW(Str1), lenStr2 = lstrlenW(Str2);
	int i = 0, j = 0;

	if (lenStr1 >= MAX_PATH || lenStr2 >= MAX_PATH)
		return FALSE;

	for (i = 0; i < lenStr1; i++) {
		lStr1[i] = (CHAR)tolower(Str1[i]);
	}
	lStr1[i++] = L'\0';

	for (j = 0; j < lenStr2; j++) {
		lStr2[j] = (CHAR)tolower(Str2[j]);
	}
	lStr2[j++] = L'\0';

	if (strcmp(lStr1, lStr2) == 0)
		return TRUE;

	return FALSE;
}


BOOL IsWStrEqualCStr(IN LPCWSTR Str1, IN LPCWSTR Str2) {

	WCHAR lStr1[MAX_PATH], lStr2[MAX_PATH];
	int lenStr1 = lstrlenW(Str1), lenStr2 = lstrlenW(Str2);
	int i = 0, j = 0;

	if (lenStr1 >= MAX_PATH || lenStr2 >= MAX_PATH)
		return FALSE;

	for (i = 0; i < lenStr1; i++ ) {
		lStr1[i] = (WCHAR)tolower(Str1[i]);
	}
	lStr1[i++] = L'\0';

	for (j = 0; j < lenStr2; j++) {
		lStr2[j] = (WCHAR)tolower(Str2[j]);
	}
	lStr2[j++] = L'\0';

	if (lstrcmpiW(lStr1, lStr2) == 0)
		return TRUE;

	return FALSE;

}