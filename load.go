package op

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type opsoft struct {
	op *ole.IDispatch
	uk *ole.IUnknown
}

func Load() (*opsoft, error) {
	var err error
	com := new(opsoft)
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

func (com *opsoft) Release() {
	com.uk.Release()
	com.op.Release()
	ole.CoUninitialize()
}
