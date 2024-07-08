#pragma once
#include <windows.h>

/*
	header file that contains our own definition of a LDR_DATA_TABLE_ENTRY & PEB and required 
	This was made with phind, because clicking like a madman I could not get VS code to show me the goods to copy and paste in.
*/

#define STRUCTS

typedef PVOID PRTL_USER_PROCESS_PARAMETERS;
typedef PRTL_USER_PROCESS_PARAMETERS;
typedef PVOID PAPI_SET_NAMESPACE; 
typedef PVOID PPS_POST_PROCESS_INIT_ROUTINE;

typedef void (*PPS_POST_PROCESS_INIT_ROUTINE)();

typedef struct _RTL_USER_PROCESS_PARAMETERS {
	BYTE Reserved1[16];
	PVOID Reserved2[10];
	UNICODE_STRING ImagePathName;
	UNICODE_STRING CommandLine;
} RTL_USER_PROCESS_PARAMETERS, * PRTL_USER_PROCESS_PARAMETERS;

typedef struct _TRUNCATED_LDR_DATA_TABLE_ENTRY {
	PVOID DllBase;
	PVOID EntryPoint;
	LIST_ENTRY InMemoryOrderLinks;
} TRUNCATED_LDR_DATA_TABLE_ENTRY, * PTRUNCATED_LDR_DATA_TABLE_ENTRY;



typedef struct _UNICODE_STRING {
	USHORT Length;
	USHORT MaximumLength;
	PWSTR Buffer;
} UNICODE_STRING, * PUNICODE_STRING;


typedef struct _PEB {
	BYTE Reserved1[2];
	BYTE BeingDebugged;
	BYTE Reserved2[21];
	PPEB_LDR_DATA LoaderData;
	PRTL_USER_PROCESS_PARAMETERS ProcessParameters;
	BYTE Reserved3[520];
	PPS_POST_PROCESS_INIT_ROUTINE PostProcessInitRoutine;
	BYTE Reserved4[136];
	ULONG SessionId;
} PEB;


typedef struct _PEB_LDR_DATA {
	BYTE Reserved1[8];
	PVOID Reserved2[3];
	LIST_ENTRY InMemoryOrderModuleList;
} PEB_LDR_DATA, * PPEB_LDR_DATA;

typedef struct _LDR_DATA_TABLE_ENTRY {
	PVOID Reserved1[2];
	LIST_ENTRY InMemoryOrderLinks;
	PVOID Reserved2[2];
	PVOID DllBase;
	PVOID EntryPoint;
	PVOID Reserved3;
	UNICODE_STRING FullDllName;
	BYTE Reserved4[8];
	PVOID Reserved5[3];
	union {
		ULONG CheckSum;
		PVOID Reserved6;
	};
	ULONG TimeDateStamp;
} LDR_DATA_TABLE_ENTRY, * PLDR_DATA_TABLE_ENTRY;

typedef struct _LIST_ENTRY {
	struct _LIST_ENTRY* Flink;
	struct _LIST_ENTRY* Blink;
} LIST_ENTRY, * PLIST_ENTRY;
