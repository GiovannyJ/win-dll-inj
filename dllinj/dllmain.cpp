#include <stdio.h>
#include <Windows.h>

BOOL WINAPI DllMain(
    HINSTANCE hModule,
    DWORD fwdReason,
    LPVOID lpvReason
){
    const int UNLEN = 20;
    wchar_t username[UNLEN + 1];
    DWORD size = UNLEN + 1;
    wchar_t client[MAX_PATH];

    switch(fwdReason){
        case DLL_PROCESS_ATTACH:
            // MessageBoxW(NULL,L"Click it you idiot", L"THIS IS NOT MALWARE", MB_OK | MB_ICONEXCLAMATION);
            STARTUPINFOW si = { 0 };
            PROCESS_INFORMATION pi = { 0 };
            if(GetUserNameW(username, &size)){
                if (wcscmp(username, L"vagrant") == 0){
                    wchar_t client[MAX_PATH] = L"C:\\Users\\vagrant\\Desktop\\build\\client.exe";
                }else{
                    //CHANGE FOR LOCAL ENV
                    wchar_t client[MAX_PATH] = L"G:\\...\\...\\win-dll-inj\\build\\client.exe";
                }
            }

            if(!CreateProcessW(
                client,
                NULL,
                NULL,
                NULL,
                FALSE,
                0,
                NULL,
                NULL,
                &si,
                &pi
            ))
            {
                printf("(-) process failed to start, error: %ld", GetLastError());
            }
            printf("(+) process started, pid: %ld", pi.dwProcessId);
            break;
    }

    return TRUE;
}
