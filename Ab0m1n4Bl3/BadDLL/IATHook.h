#pragma once
#include <windows.h>

PVOID HookIAT(PBYTE pTarget, HMODULE hModule, PCWSTR lpApiName, PVOID replacement);

BOOL PatchIATEntry(PBYTE pTarget, PCSTR lpApiName, PIMAGE_IMPORT_DESCRIPTOR pModuleEntry, PVOID replacement);