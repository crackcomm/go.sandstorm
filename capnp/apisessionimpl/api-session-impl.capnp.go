// Code generated by capnpc-go. DO NOT EDIT.

package apisessionimpl

import (
	context "context"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	websession "zenhack.net/go/sandstorm/capnp/websession"
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
	persistent "zombiezen.com/go/capnproto2/std/capnp/persistent"
)

type PersistentApiSession struct{ Client *capnp.Client }

// PersistentApiSession_TypeID is the unique identifier for the type PersistentApiSession.
const PersistentApiSession_TypeID = 0xaca1efb8eb6cd18a

func (c PersistentApiSession) Get(ctx context.Context, params func(websession.WebSession_get_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "get",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_get_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Post(ctx context.Context, params func(websession.WebSession_post_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_post_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) OpenWebSocket(ctx context.Context, params func(websession.WebSession_openWebSocket_Params) error) (websession.WebSession_openWebSocket_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_openWebSocket_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_openWebSocket_Results_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Put(ctx context.Context, params func(websession.WebSession_put_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      3,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "put",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_put_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Delete(ctx context.Context, params func(websession.WebSession_delete_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_delete_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) PostStreaming(ctx context.Context, params func(websession.WebSession_postStreaming_Params) error) (websession.WebSession_postStreaming_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_postStreaming_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_postStreaming_Results_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) PutStreaming(ctx context.Context, params func(websession.WebSession_putStreaming_Params) error) (websession.WebSession_putStreaming_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      6,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "putStreaming",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_putStreaming_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_putStreaming_Results_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Propfind(ctx context.Context, params func(websession.WebSession_propfind_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      7,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "propfind",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_propfind_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Proppatch(ctx context.Context, params func(websession.WebSession_proppatch_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      8,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "proppatch",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_proppatch_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Mkcol(ctx context.Context, params func(websession.WebSession_mkcol_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      9,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "mkcol",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_mkcol_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Copy(ctx context.Context, params func(websession.WebSession_copy_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      10,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "copy",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_copy_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Move(ctx context.Context, params func(websession.WebSession_move_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      11,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "move",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_move_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Lock(ctx context.Context, params func(websession.WebSession_lock_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      12,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "lock",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_lock_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Unlock(ctx context.Context, params func(websession.WebSession_unlock_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      13,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "unlock",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_unlock_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Acl(ctx context.Context, params func(websession.WebSession_acl_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      14,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "acl",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_acl_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Report(ctx context.Context, params func(websession.WebSession_report_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      15,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "report",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_report_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Options(ctx context.Context, params func(websession.WebSession_options_Params) error) (websession.WebSession_Options_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      16,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "options",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_options_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Options_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Patch(ctx context.Context, params func(websession.WebSession_patch_Params) error) (websession.WebSession_Response_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      17,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "patch",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		s.PlaceArgs = func(s capnp.Struct) error { return params(websession.WebSession_patch_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return websession.WebSession_Response_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) AddRequirements(ctx context.Context, params func(supervisor.SystemPersistent_addRequirements_Params) error) (supervisor.SystemPersistent_addRequirements_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		s.PlaceArgs = func(s capnp.Struct) error {
			return params(supervisor.SystemPersistent_addRequirements_Params{Struct: s})
		}
	}
	ans, release := c.Client.SendCall(ctx, s)
	return supervisor.SystemPersistent_addRequirements_Results_Future{Future: ans.Future()}, release
}
func (c PersistentApiSession) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error) (persistent.Persistent_SaveResults_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "capnp/persistent.capnp:Persistent",
			MethodName:    "save",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(persistent.Persistent_SaveParams{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return persistent.Persistent_SaveResults_Future{Future: ans.Future()}, release
}

// A PersistentApiSession_Server is a PersistentApiSession with a local implementation.
type PersistentApiSession_Server interface {
	Get(context.Context, websession.WebSession_get) error

	Post(context.Context, websession.WebSession_post) error

	OpenWebSocket(context.Context, websession.WebSession_openWebSocket) error

	Put(context.Context, websession.WebSession_put) error

	Delete(context.Context, websession.WebSession_delete) error

	PostStreaming(context.Context, websession.WebSession_postStreaming) error

	PutStreaming(context.Context, websession.WebSession_putStreaming) error

	Propfind(context.Context, websession.WebSession_propfind) error

	Proppatch(context.Context, websession.WebSession_proppatch) error

	Mkcol(context.Context, websession.WebSession_mkcol) error

	Copy(context.Context, websession.WebSession_copy) error

	Move(context.Context, websession.WebSession_move) error

	Lock(context.Context, websession.WebSession_lock) error

	Unlock(context.Context, websession.WebSession_unlock) error

	Acl(context.Context, websession.WebSession_acl) error

	Report(context.Context, websession.WebSession_report) error

	Options(context.Context, websession.WebSession_options) error

	Patch(context.Context, websession.WebSession_patch) error

	AddRequirements(context.Context, supervisor.SystemPersistent_addRequirements) error

	Save(context.Context, persistent.Persistent_save) error
}

// PersistentApiSession_NewServer creates a new Server from an implementation of PersistentApiSession_Server.
func PersistentApiSession_NewServer(s PersistentApiSession_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(PersistentApiSession_Methods(nil, s), s, c, policy)
}

// PersistentApiSession_ServerToClient creates a new Client from an implementation of PersistentApiSession_Server.
// The caller is responsible for calling Release on the returned Client.
func PersistentApiSession_ServerToClient(s PersistentApiSession_Server, policy *server.Policy) PersistentApiSession {
	return PersistentApiSession{Client: capnp.NewClient(PersistentApiSession_NewServer(s, policy))}
}

// PersistentApiSession_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func PersistentApiSession_Methods(methods []server.Method, s PersistentApiSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 20)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "get",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Get(ctx, websession.WebSession_get{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Post(ctx, websession.WebSession_post{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.OpenWebSocket(ctx, websession.WebSession_openWebSocket{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      3,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "put",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Put(ctx, websession.WebSession_put{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Delete(ctx, websession.WebSession_delete{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.PostStreaming(ctx, websession.WebSession_postStreaming{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      6,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "putStreaming",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.PutStreaming(ctx, websession.WebSession_putStreaming{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      7,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "propfind",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Propfind(ctx, websession.WebSession_propfind{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      8,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "proppatch",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Proppatch(ctx, websession.WebSession_proppatch{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      9,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "mkcol",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Mkcol(ctx, websession.WebSession_mkcol{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      10,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "copy",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Copy(ctx, websession.WebSession_copy{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      11,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "move",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Move(ctx, websession.WebSession_move{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      12,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "lock",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Lock(ctx, websession.WebSession_lock{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      13,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "unlock",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Unlock(ctx, websession.WebSession_unlock{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      14,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "acl",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Acl(ctx, websession.WebSession_acl{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      15,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "report",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Report(ctx, websession.WebSession_report{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      16,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "options",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Options(ctx, websession.WebSession_options{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      17,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "patch",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Patch(ctx, websession.WebSession_patch{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.AddRequirements(ctx, supervisor.SystemPersistent_addRequirements{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "capnp/persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Save(ctx, persistent.Persistent_save{call})
		},
	})

	return methods
}

const schema_a949cfa7be07085b = "x\xda2\xc8`q`2d\xdd\xaf\xc4\xc0\x10<\x85" +
	"\x91\x95\xed\x7f\xd7\xc5\x9c\xd7;\xde/\\\xc3 (\xce" +
	"\xfc?\x9a\x83}\xdf\xf2\xf3\x9e+\x19\x18\x18\x8d]\x05" +
	"\x93\x18\x85#\x05\xd9\x19\x18\x84C\x05\xd9\x85C\x05\xd5" +
	"\x19\xfeC\xa1\xcf\xff\xc4\x82L\xdd\xe2\xd4\xe2b\x96\xcc" +
	"\xfc<\xdd\xcc\xdc\x82\x1c\xbd\xe4\xc4\x82\xbc\x02\xab\x80\xd4" +
	"\xa2\xe2\xcc\xe2\x92\xd4\xbc\x12\xc7\x82\xcc\xe0\xd4\xe2b\xf6" +
	"\xcc\xfc\xbc\x00F\xc6\x00f\xd6@\x0eF\xc6\xff\xc7\xcf" +
	"\xaa\x1e\xab|\\y\x82\x81\x81\xe1\xff\x96\xab\xfbj\xae" +
	"\xbf\xed9\xcc\xc0\xc0\x00\x08\x00\x00\xff\xff\x83R5\x13"

func init() {
	schemas.Register(schema_a949cfa7be07085b,
		0xaca1efb8eb6cd18a)
}
