package op

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type Opsoft struct {
	op *ole.IDispatch
	uk *ole.IUnknown
}

//加载op
func Load() (*Opsoft, error) {
	var err error
	com := new(Opsoft)
	ole.CoInitialize(0)
	com.uk, err = oleutil.CreateObject("op.opsoft")
	if err != nil {
		return nil, err
	}
	com.op, err = com.uk.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return com, nil
}

//释放
func (com *Opsoft) Release() {
	com.uk.Release()
	com.op.Release()
	ole.CoUninitialize()
}
