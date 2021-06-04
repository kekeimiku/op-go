package op

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type Opsoft struct {
	op *ole.IDispatch
	uk *ole.IUnknown
}

func Load() (*Opsoft, error) {
	var err error
	com := new(Opsoft)
	ole.CoInitialize(0)
	com.uk, err = oleutil.CreateObject("op.Opsoft")
	if err != nil {
		return nil, err
	}
	com.op, err = com.uk.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return com, nil
}

func (com *Opsoft) Release() {
	com.uk.Release()
	com.op.Release()
	ole.CoUninitialize()
}
