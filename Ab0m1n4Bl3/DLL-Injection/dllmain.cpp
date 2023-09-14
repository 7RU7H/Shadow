// From https://tryhackme.com/room/abusingwindowsinternals
// dllmain.cpp : Defines the entry point for the DLL application.
#include "pch.h"
#define EXPORTING_DLL

BOOL APIENTRY DllMain( HMODULE hModule,
                       DWORD  ul_reason_for_call,
                       LPVOID lpReserved
                     )
{
    MessageBox(NULL, TEXT("THM{}"), TEXT("flag"), MB_OK);
    return TRUE;
}
