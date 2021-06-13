#define OP_API
%module opEx

%{
    #define OP_API
#include "libop.h"
%}

%include "std_string.i"
%include "typemaps.i"
%include "libop.h"
