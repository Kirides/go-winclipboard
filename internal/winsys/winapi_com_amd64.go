package winsys

type FORMATETC struct {
	ClipFormat     uint16
	_              [6]byte
	DvTargetDevice uintptr
	Aspect         uint32
	Index          int32
	Tymed          uint32
	_              [4]byte
}
type STGMEDIUM struct {
	Tymed          Tymed
	_              [4]byte
	UnionMember    uintptr
	PUnkForRelease *IUnknown
}