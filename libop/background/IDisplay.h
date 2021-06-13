#pragma once
//#include <boost/interprocess/sync/named_mutex.hpp> 
//#include <boost/interprocess/shared_memory_object.hpp>
//#include <boost/interprocess/windows_shared_memory.hpp> 
//#include <boost/interprocess/mapped_region.hpp>
//#include <boost/interprocess/sync/interprocess_mutex.hpp>
#include <exception>
#include "./include/promutex.h"
#include "./include/sharedmem.h"
#include "frameInfo.h"
struct Image;
class IDisplay
{
public:
	IDisplay();
	 ~IDisplay();
	//bind window
	long Bind(HWND hwnd, long flag);
	//unbind window
	long UnBind();
	//unbind window
	//virtual long UnBind(HWND hwnd) = 0;
	virtual bool requestCapture(int x1, int y1, int w, int h, Image& img) = 0;
	

	promutex* get_mutex() {
		return _pmutex;
	}

	long get_height() {
		return _height;
	}

	long get_width() {
		return _width;
	}
	void getFrameInfo(FrameInfo& info);
private:
	//��Ϊ���ֽ�ͼ��ʽ�Ĳ��죬�Ƿ�ɹ��жϽϸ��ӣ����ڴ�ʵ����Դ��������ͷţ��������
	//��Դ����
	long bind_init();
	//��Դ�ͷ�
	long bind_release();
	//virtual byte* get_data();
		//��״̬
	long _bind_state;
protected:
	virtual long BindEx(HWND hwnd, long flag)=0;
	virtual long UnBindEx() = 0;
	//���ھ��
	HWND _hwnd;
	//�����ڴ�
	sharedmem* _shmem;
	//���̻�����
	promutex* _pmutex;

	wchar_t _shared_res_name[256];

	wchar_t _mutex_name[256];

	//
	int _render_type;
	//���
	long _width;
	long _height;
	//�ͻ���ƫ��
	int _client_x, _client_y;
	//need capture rect
	//RECT rect;
	
};

