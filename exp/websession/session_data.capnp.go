// Code generated by capnpc-go. DO NOT EDIT.

package websession

import (
	strconv "strconv"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	identity "zenhack.net/go/sandstorm/capnp/identity"
	powerbox "zenhack.net/go/sandstorm/capnp/powerbox"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type SessionData struct{ capnp.Struct }
type SessionData_request SessionData
type SessionData_offer SessionData
type SessionData_Which uint16

const (
	SessionData_Which_normal  SessionData_Which = 0
	SessionData_Which_request SessionData_Which = 1
	SessionData_Which_offer   SessionData_Which = 2
)

func (w SessionData_Which) String() string {
	const s = "normalrequestoffer"
	switch w {
	case SessionData_Which_normal:
		return s[0:6]
	case SessionData_Which_request:
		return s[6:13]
	case SessionData_Which_offer:
		return s[13:18]

	}
	return "SessionData_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// SessionData_TypeID is the unique identifier for the type SessionData.
const SessionData_TypeID = 0x8c13d90395f6e861

func NewSessionData(s *capnp.Segment) (SessionData, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 6})
	return SessionData{st}, err
}

func NewRootSessionData(s *capnp.Segment) (SessionData, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 6})
	return SessionData{st}, err
}

func ReadRootSessionData(msg *capnp.Message) (SessionData, error) {
	root, err := msg.Root()
	return SessionData{root.Struct()}, err
}

func (s SessionData) String() string {
	str, _ := text.Marshal(0x8c13d90395f6e861, s.Struct)
	return str
}

func (s SessionData) Which() SessionData_Which {
	return SessionData_Which(s.Struct.Uint16(0))
}
func (s SessionData) Context() grain.SessionContext {
	p, _ := s.Struct.Ptr(0)
	return grain.SessionContext{Client: p.Interface().Client()}
}

func (s SessionData) HasContext() bool {
	return s.Struct.HasPtr(0)
}

func (s SessionData) SetContext(v grain.SessionContext) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s SessionData) UserInfo() (identity.UserInfo, error) {
	p, err := s.Struct.Ptr(1)
	return identity.UserInfo{Struct: p.Struct()}, err
}

func (s SessionData) HasUserInfo() bool {
	return s.Struct.HasPtr(1)
}

func (s SessionData) SetUserInfo(v identity.UserInfo) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated identity.UserInfo struct, preferring placement in s's segment.
func (s SessionData) NewUserInfo() (identity.UserInfo, error) {
	ss, err := identity.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return identity.UserInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s SessionData) SessionParams() (capnp.Ptr, error) {
	return s.Struct.Ptr(2)
}

func (s SessionData) HasSessionParams() bool {
	return s.Struct.HasPtr(2)
}

func (s SessionData) SetSessionParams(v capnp.Ptr) error {
	return s.Struct.SetPtr(2, v)
}

func (s SessionData) TabId() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s SessionData) HasTabId() bool {
	return s.Struct.HasPtr(3)
}

func (s SessionData) SetTabId(v []byte) error {
	return s.Struct.SetData(3, v)
}

func (s SessionData) SetNormal() {
	s.Struct.SetUint16(0, 0)

}

func (s SessionData) Request() SessionData_request { return SessionData_request(s) }

func (s SessionData) SetRequest() {
	s.Struct.SetUint16(0, 1)
}

func (s SessionData_request) RequestInfo() (powerbox.PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(4)
	return powerbox.PowerboxDescriptor_List{List: p.List()}, err
}

func (s SessionData_request) HasRequestInfo() bool {
	return s.Struct.HasPtr(4)
}

func (s SessionData_request) SetRequestInfo(v powerbox.PowerboxDescriptor_List) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewRequestInfo sets the requestInfo field to a newly
// allocated powerbox.PowerboxDescriptor_List, preferring placement in s's segment.
func (s SessionData_request) NewRequestInfo(n int32) (powerbox.PowerboxDescriptor_List, error) {
	l, err := powerbox.NewPowerboxDescriptor_List(s.Struct.Segment(), n)
	if err != nil {
		return powerbox.PowerboxDescriptor_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s SessionData) Offer() SessionData_offer { return SessionData_offer(s) }

func (s SessionData) SetOffer() {
	s.Struct.SetUint16(0, 2)
}

func (s SessionData_offer) Offer() (capnp.Ptr, error) {
	return s.Struct.Ptr(4)
}

func (s SessionData_offer) HasOffer() bool {
	return s.Struct.HasPtr(4)
}

func (s SessionData_offer) SetOffer(v capnp.Ptr) error {
	return s.Struct.SetPtr(4, v)
}

func (s SessionData_offer) Descriptor() (powerbox.PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(5)
	return powerbox.PowerboxDescriptor{Struct: p.Struct()}, err
}

func (s SessionData_offer) HasDescriptor() bool {
	return s.Struct.HasPtr(5)
}

func (s SessionData_offer) SetDescriptor(v powerbox.PowerboxDescriptor) error {
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated powerbox.PowerboxDescriptor struct, preferring placement in s's segment.
func (s SessionData_offer) NewDescriptor() (powerbox.PowerboxDescriptor, error) {
	ss, err := powerbox.NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s SessionData) SessionType() uint64 {
	return s.Struct.Uint64(8)
}

func (s SessionData) SetSessionType(v uint64) {
	s.Struct.SetUint64(8, v)
}

// SessionData_List is a list of SessionData.
type SessionData_List struct{ capnp.List }

// NewSessionData creates a new list of SessionData.
func NewSessionData_List(s *capnp.Segment, sz int32) (SessionData_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 6}, sz)
	return SessionData_List{l}, err
}

func (s SessionData_List) At(i int) SessionData { return SessionData{s.List.Struct(i)} }

func (s SessionData_List) Set(i int, v SessionData) error { return s.List.SetStruct(i, v.Struct) }

func (s SessionData_List) String() string {
	str, _ := text.MarshalList(0x8c13d90395f6e861, s.List)
	return str
}

// SessionData_Future is a wrapper for a SessionData promised by a client call.
type SessionData_Future struct{ *capnp.Future }

func (p SessionData_Future) Struct() (SessionData, error) {
	s, err := p.Future.Struct()
	return SessionData{s}, err
}

func (p SessionData_Future) Context() grain.SessionContext {
	return grain.SessionContext{Client: p.Future.Field(0, nil).Client()}
}

func (p SessionData_Future) UserInfo() identity.UserInfo_Future {
	return identity.UserInfo_Future{Future: p.Future.Field(1, nil)}
}

func (p SessionData_Future) SessionParams() *capnp.Future {
	return p.Future.Field(2, nil)
}

func (p SessionData_Future) Request() SessionData_request_Future {
	return SessionData_request_Future{p.Future}
}

// SessionData_request_Future is a wrapper for a SessionData_request promised by a client call.
type SessionData_request_Future struct{ *capnp.Future }

func (p SessionData_request_Future) Struct() (SessionData_request, error) {
	s, err := p.Future.Struct()
	return SessionData_request{s}, err
}

func (p SessionData_Future) Offer() SessionData_offer_Future {
	return SessionData_offer_Future{p.Future}
}

// SessionData_offer_Future is a wrapper for a SessionData_offer promised by a client call.
type SessionData_offer_Future struct{ *capnp.Future }

func (p SessionData_offer_Future) Struct() (SessionData_offer, error) {
	s, err := p.Future.Struct()
	return SessionData_offer{s}, err
}

func (p SessionData_offer_Future) Offer() *capnp.Future {
	return p.Future.Field(4, nil)
}

func (p SessionData_offer_Future) Descriptor() powerbox.PowerboxDescriptor_Future {
	return powerbox.PowerboxDescriptor_Future{Future: p.Future.Field(5, nil)}
}

const schema_9cd7bc7194780111 = "x\xda\x84R\xcfK\x14Q\x1c\xff|\xde\x9b\x9dQh" +
	"\xc6}\xeeD\x17e;t\xa8\xc80\xa5(\x0f\xadH" +
	"?0:\xf8\xb0CH\x11\xe3:\x0bB\xce\xac3#" +
	"\x19\x1d\x84nA\xd0Eou\xaa\xbf\xa0 \x08\x0c\x8f" +
	"v\x08bI)A\xc1C\x82AR\x10t\x08\xea\xc5" +
	"\xb4\xfd\x90%\x88w\xfb|\xbf_>\xbf^\xefyk" +
	"P\x1c+\xdc\xeb\x02F/\xb2`\x9b\xad\xc7W\x9e\xae" +
	"\xd4n\xdd\x86.S\x98`\xfb\xcb\x82\\+\xdd\xc5^" +
	"\xdb!\xd0\xdf\xe3\x1d&X:\xee\xdd\x00\xcd|\xf7\xb9" +
	"\xf5}\xa7\x1e\xdd\xf9\xf7\xea\x827\x90\xaf>\xf0*\xe0" +
	"\xdf\xa1.Q\x18\xc5\xd9\xf9\xe9\xc57\xf7q\xd6v$" +
	"Pjx_K\x1b\x9e\x03\xf4\xafy\xcb\x84\xf9\xf9\xa6" +
	"M\x1a\xa6\xe9d\x1c]\x93\x13A\x16\x1c\xad\x06\xf5\xa8" +
	">0\xda\xc4\xce\xe4H\xdcQ\xab\x85\x89n\x93V\xd1" +
	"\xf2i\x03\xeaP\x1f\xa0\x0fH\xea^A\xc5\x82O\x07" +
	"P=c\x80>\"\xa9O\x0a\x96\xe3\xfc\x84\x9d\x96\x04" +
	"\xd9\x09\x9a\x890\xad&\x93\xf5\x0c2NX4\x97\xbb" +
	"\x96\x97\x1ev\xbc\xda\x01\xc8\"\xf8_\x09I%\x9c\x9e" +
	"\x09\xd3L[M\x11\x05@\xb9\xe3\x80\xde#\xa9\x0f\x0a" +
	"\x9a\xa49\x1f\x86\x13\xd5bz\xe0\x88d\x0b\x8f\xb7\x8b" +
	"G\xb4\xf2T\x9aD#\xa4\xde/-\xc0\"\xa0\x1aC" +
	"\x80~)\xa9\xdf\xe66\xe93\x07W/\x00zER" +
	"o\x0a*!|\x0a@m$\x80^\x97\xd4\xdb\x82J" +
	"J\x9f\x12P[yJ\x9b\x92\xfa\x83\xa0k\x19\xe3\xd3" +
	"\x02\xd4\xfb\x01@\xbf\x93\xd4\x9f\x04\xdd\xc2w\xc3]\x05" +
	"\xab\x9d!\x08\xd7\xfe\x96\x83\x7f>\x88Z\xed\x83P\x0e" +
	"}\xb6\x01\xeay\xeezQR\xbf\x10\x9c\xab\xc6Q\x16" +
	"\xcefT\xe6\xc4\xe7\xeaX\xf7\xe0\xe9\xa5\xdc\xa8\x02\xcd" +
	"L\x1a&\xc3Q-\x06\xc0\xa2y\xad\xae>\xf9\xd8x" +
	"6\xdf\x12\xf7\x08\xcaA\x12L\xa5\xec\x84\xc8[*g" +
	"\xc1\xf8\xf0\x04]\x08\xba`%\x8a\x93\xa9\xe0:\xec\xb9" +
	"_\xe16K\xfd}|\x09\xce\xcdz\xc8v\x08\xb6\x83" +
	"?\x02\x00\x00\xff\xffe\"\xcc\xf8"

func init() {
	schemas.Register(schema_9cd7bc7194780111,
		0x827b66d4b65cb2e5,
		0x89a43917dc461d94,
		0x8c13d90395f6e861)
}
