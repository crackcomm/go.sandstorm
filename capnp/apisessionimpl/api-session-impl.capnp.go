// Code generated by capnpc-go. DO NOT EDIT.

package apisessionimpl

import (
	context "golang.org/x/net/context"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	websession "zenhack.net/go/sandstorm/capnp/websession"
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
	persistent "zombiezen.com/go/capnproto2/std/capnp/persistent"
)

type PersistentApiSession struct{ Client capnp.Client }

// PersistentApiSession_TypeID is the unique identifier for the type PersistentApiSession.
const PersistentApiSession_TypeID = 0xaca1efb8eb6cd18a

func (c PersistentApiSession) Get(ctx context.Context, params func(websession.WebSession_get_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "get",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_get_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Post(ctx context.Context, params func(websession.WebSession_post_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_post_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) OpenWebSocket(ctx context.Context, params func(websession.WebSession_openWebSocket_Params) error, opts ...capnp.CallOption) websession.WebSession_openWebSocket_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_openWebSocket_Params{Struct: s}) }
	}
	return websession.WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Put(ctx context.Context, params func(websession.WebSession_put_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      3,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "put",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_put_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Delete(ctx context.Context, params func(websession.WebSession_delete_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_delete_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) PostStreaming(ctx context.Context, params func(websession.WebSession_postStreaming_Params) error, opts ...capnp.CallOption) websession.WebSession_postStreaming_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_postStreaming_Params{Struct: s}) }
	}
	return websession.WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) PutStreaming(ctx context.Context, params func(websession.WebSession_putStreaming_Params) error, opts ...capnp.CallOption) websession.WebSession_putStreaming_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      6,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "putStreaming",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_putStreaming_Params{Struct: s}) }
	}
	return websession.WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Propfind(ctx context.Context, params func(websession.WebSession_propfind_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      7,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "propfind",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_propfind_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Proppatch(ctx context.Context, params func(websession.WebSession_proppatch_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      8,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "proppatch",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_proppatch_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Mkcol(ctx context.Context, params func(websession.WebSession_mkcol_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      9,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "mkcol",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_mkcol_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Copy(ctx context.Context, params func(websession.WebSession_copy_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      10,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "copy",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_copy_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Move(ctx context.Context, params func(websession.WebSession_move_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      11,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "move",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_move_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Lock(ctx context.Context, params func(websession.WebSession_lock_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      12,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "lock",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_lock_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Unlock(ctx context.Context, params func(websession.WebSession_unlock_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      13,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "unlock",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_unlock_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Acl(ctx context.Context, params func(websession.WebSession_acl_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      14,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "acl",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_acl_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Report(ctx context.Context, params func(websession.WebSession_report_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      15,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "report",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_report_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Options(ctx context.Context, params func(websession.WebSession_options_Params) error, opts ...capnp.CallOption) websession.WebSession_Options_Promise {
	if c.Client == nil {
		return websession.WebSession_Options_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      16,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "options",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_options_Params{Struct: s}) }
	}
	return websession.WebSession_Options_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Patch(ctx context.Context, params func(websession.WebSession_patch_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      17,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "patch",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_patch_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) AddRequirements(ctx context.Context, params func(supervisor.SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) supervisor.SystemPersistent_addRequirements_Results_Promise {
	if c.Client == nil {
		return supervisor.SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error {
			return params(supervisor.SystemPersistent_addRequirements_Params{Struct: s})
		}
	}
	return supervisor.SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentApiSession) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
	if c.Client == nil {
		return persistent.Persistent_SaveResults_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "capnp/persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(persistent.Persistent_SaveParams{Struct: s}) }
	}
	return persistent.Persistent_SaveResults_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type PersistentApiSession_Server interface {
	Get(websession.WebSession_get) error

	Post(websession.WebSession_post) error

	OpenWebSocket(websession.WebSession_openWebSocket) error

	Put(websession.WebSession_put) error

	Delete(websession.WebSession_delete) error

	PostStreaming(websession.WebSession_postStreaming) error

	PutStreaming(websession.WebSession_putStreaming) error

	Propfind(websession.WebSession_propfind) error

	Proppatch(websession.WebSession_proppatch) error

	Mkcol(websession.WebSession_mkcol) error

	Copy(websession.WebSession_copy) error

	Move(websession.WebSession_move) error

	Lock(websession.WebSession_lock) error

	Unlock(websession.WebSession_unlock) error

	Acl(websession.WebSession_acl) error

	Report(websession.WebSession_report) error

	Options(websession.WebSession_options) error

	Patch(websession.WebSession_patch) error

	AddRequirements(supervisor.SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentApiSession_ServerToClient(s PersistentApiSession_Server) PersistentApiSession {
	c, _ := s.(server.Closer)
	return PersistentApiSession{Client: server.New(PersistentApiSession_Methods(nil, s), c)}
}

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
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_get{c, opts, websession.WebSession_get_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Get(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_post{c, opts, websession.WebSession_post_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Post(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_openWebSocket{c, opts, websession.WebSession_openWebSocket_Params{Struct: p}, websession.WebSession_openWebSocket_Results{Struct: r}}
			return s.OpenWebSocket(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      3,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "put",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_put{c, opts, websession.WebSession_put_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Put(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_delete{c, opts, websession.WebSession_delete_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Delete(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_postStreaming{c, opts, websession.WebSession_postStreaming_Params{Struct: p}, websession.WebSession_postStreaming_Results{Struct: r}}
			return s.PostStreaming(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      6,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "putStreaming",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_putStreaming{c, opts, websession.WebSession_putStreaming_Params{Struct: p}, websession.WebSession_putStreaming_Results{Struct: r}}
			return s.PutStreaming(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      7,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "propfind",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_propfind{c, opts, websession.WebSession_propfind_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Propfind(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      8,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "proppatch",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_proppatch{c, opts, websession.WebSession_proppatch_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Proppatch(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      9,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "mkcol",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_mkcol{c, opts, websession.WebSession_mkcol_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Mkcol(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      10,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "copy",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_copy{c, opts, websession.WebSession_copy_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Copy(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      11,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "move",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_move{c, opts, websession.WebSession_move_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Move(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      12,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "lock",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_lock{c, opts, websession.WebSession_lock_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Lock(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      13,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "unlock",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_unlock{c, opts, websession.WebSession_unlock_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Unlock(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      14,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "acl",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_acl{c, opts, websession.WebSession_acl_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Acl(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      15,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "report",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_report{c, opts, websession.WebSession_report_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Report(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      16,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "options",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_options{c, opts, websession.WebSession_options_Params{Struct: p}, websession.WebSession_Options{Struct: r}}
			return s.Options(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      17,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "patch",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_patch{c, opts, websession.WebSession_patch_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Patch(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := supervisor.SystemPersistent_addRequirements{c, opts, supervisor.SystemPersistent_addRequirements_Params{Struct: p}, supervisor.SystemPersistent_addRequirements_Results{Struct: r}}
			return s.AddRequirements(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "capnp/persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := persistent.Persistent_save{c, opts, persistent.Persistent_SaveParams{Struct: p}, persistent.Persistent_SaveResults{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
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