// Package websession_pogs defines pogs-style types for the websesion schema.
// This is WIP.
package websession_pogs

// TODO: default values?

import (
	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/util/pogs"
)

type AcceptedEncoding struct {
	ContentCoding string
	QValue        float32
}

type AcceptedType struct {
	MimeType string
	QValue   float32
}

type CachePolicy struct {
	WithCheck      websession.WebSession_CachePolicy_Scope
	Permanent      websession.WebSession_CachePolicy_Scope
	VariesOnCookie bool
	VariesOnAccept bool
}

type Context struct {
	Cookies          []util_pogs.KeyValue
	ResponseStream   util.ByteStream
	Accept           []AcceptedType
	AcceptEncoding   []AcceptedEncoding
	ETagPrecondition struct {
		Which         websession.WebSession_Context_eTagPrecondition_Which
		MatchesOneOf  []ETag
		MatchesNoneOf []ETag
	}
	AdditionalHeaders []Header
}

type Cookie struct {
	Name    string
	Value   string
	Expires struct {
		Which    websession.WebSession_Cookie_expires_Which
		Absolute int64
		Relative uint64
	}
	HttpOnly bool
	Path     string
}

type ETag struct {
	Value string
	Weak  bool
}

type Header struct {
	Name  string
	Value string
}

type Response_content struct {
	StatusCode  websession.WebSession_Response_SuccessCode
	Encoding    string
	Language    string
	MimeType    string
	ETag        ETag
	Body        Response_content_body
	Disposition struct {
		Which    websession.WebSession_Response_content_disposition_Which
		Download string
	}
}

type Response_content_body struct {
	Which  websession.WebSession_Response_content_body_Which
	Bytes  []byte
	Stream util.Handle
}

type Response struct {
	SetCookies  []Cookie
	CachePolicy CachePolicy
	Which       websession.WebSession_Response_Which
	Content     Response_content
	NoContent   struct {
		ShouldResetForm bool
		ETag            ETag
	}
	PreconditionFailed struct {
		MatchingETag ETag
	}
	Redirect struct {
		IsPermanent bool
		SwitchToGet bool
		Location    string
	}
	ClientError struct {
		StatusCode      websession.WebSession_Response_ClientErrorCode
		DescriptionHtml string
	}
	ServerError struct {
		DescriptionHtml string
	}
	AdditionalHeaders []Header
}

type Get_args struct {
	Path       string
	Context    Context
	IgnoreBody bool
}

type Post_args struct {
	Path    string
	Content PContent
	Context Context
}

// Post/Put Content
type PContent struct {
	MimeType string
	Content  []byte
	Encoding string
}

type (
	Put_args   Post_args
	Patch_args Post_args
)

type Delete_args struct {
	Path    string
	Context Context
}
