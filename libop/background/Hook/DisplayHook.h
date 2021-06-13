//#pragma once
#ifndef __DX9HOOK_H_
#define __DX9HOOK_H_
#include "../../core/globalVar.h"


namespace DisplayHook {
	
	/*target window hwnd*/
	extern HWND render_hwnd;
	extern int render_type;
	/*name of ...*/
	extern wchar_t shared_res_name[256];
	extern wchar_t mutex_name[256];
	extern void* old_address;
	//
	int setup(HWND hwnd_, int render_type_);
	int release();


};
//以下函数用于HOOK DX9

//此函数做以下工作
/*
1.hook相关函数
2.设置共享内存,互斥量
3.截图(hook)至共享内存
*/

//返回值:1 成功，0失败
DLL_API long __stdcall SetDisplayHook(HWND hwnd_,int render_type_);

DLL_API long __stdcall ReleaseDisplayHook();
#endif // !__DX9HOOK_H_