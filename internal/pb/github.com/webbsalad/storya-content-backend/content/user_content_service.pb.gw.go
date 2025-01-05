// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/content/user_content_service.proto

/*
Package content is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package content

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_UserContentService_GetUserItems_0(ctx context.Context, marshaler runtime.Marshaler, client UserContentServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetUserItemsRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["content_type"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "content_type")
	}

	e, err = runtime.Enum(val, ContentType_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "content_type", err)
	}

	protoReq.ContentType = ContentType(e)

	val, ok = pathParams["user_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "user_id")
	}

	protoReq.UserId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "user_id", err)
	}

	msg, err := client.GetUserItems(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_UserContentService_GetUserItems_0(ctx context.Context, marshaler runtime.Marshaler, server UserContentServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetUserItemsRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["content_type"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "content_type")
	}

	e, err = runtime.Enum(val, ContentType_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "content_type", err)
	}

	protoReq.ContentType = ContentType(e)

	val, ok = pathParams["user_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "user_id")
	}

	protoReq.UserId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "user_id", err)
	}

	msg, err := server.GetUserItems(ctx, &protoReq)
	return msg, metadata, err

}

func request_UserContentService_GetValued_0(ctx context.Context, marshaler runtime.Marshaler, client UserContentServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetValuedRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["content_type"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "content_type")
	}

	e, err = runtime.Enum(val, ContentType_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "content_type", err)
	}

	protoReq.ContentType = ContentType(e)

	val, ok = pathParams["user_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "user_id")
	}

	protoReq.UserId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "user_id", err)
	}

	val, ok = pathParams["value"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "value")
	}

	e, err = runtime.Enum(val, Value_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "value", err)
	}

	protoReq.Value = Value(e)

	msg, err := client.GetValued(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_UserContentService_GetValued_0(ctx context.Context, marshaler runtime.Marshaler, server UserContentServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetValuedRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["content_type"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "content_type")
	}

	e, err = runtime.Enum(val, ContentType_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "content_type", err)
	}

	protoReq.ContentType = ContentType(e)

	val, ok = pathParams["user_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "user_id")
	}

	protoReq.UserId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "user_id", err)
	}

	val, ok = pathParams["value"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "value")
	}

	e, err = runtime.Enum(val, Value_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "value", err)
	}

	protoReq.Value = Value(e)

	msg, err := server.GetValued(ctx, &protoReq)
	return msg, metadata, err

}

func request_UserContentService_Add_0(ctx context.Context, marshaler runtime.Marshaler, client UserContentServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["content_type"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "content_type")
	}

	e, err = runtime.Enum(val, ContentType_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "content_type", err)
	}

	protoReq.ContentType = ContentType(e)

	msg, err := client.Add(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_UserContentService_Add_0(ctx context.Context, marshaler runtime.Marshaler, server UserContentServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["content_type"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "content_type")
	}

	e, err = runtime.Enum(val, ContentType_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "content_type", err)
	}

	protoReq.ContentType = ContentType(e)

	msg, err := server.Add(ctx, &protoReq)
	return msg, metadata, err

}

func request_UserContentService_Remove_0(ctx context.Context, marshaler runtime.Marshaler, client UserContentServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RemoveItemRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["content_type"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "content_type")
	}

	e, err = runtime.Enum(val, ContentType_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "content_type", err)
	}

	protoReq.ContentType = ContentType(e)

	val, ok = pathParams["item_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "item_id")
	}

	protoReq.ItemId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "item_id", err)
	}

	msg, err := client.Remove(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_UserContentService_Remove_0(ctx context.Context, marshaler runtime.Marshaler, server UserContentServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RemoveItemRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["content_type"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "content_type")
	}

	e, err = runtime.Enum(val, ContentType_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "content_type", err)
	}

	protoReq.ContentType = ContentType(e)

	val, ok = pathParams["item_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "item_id")
	}

	protoReq.ItemId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "item_id", err)
	}

	msg, err := server.Remove(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterUserContentServiceHandlerServer registers the http handlers for service UserContentService to "mux".
// UnaryRPC     :call UserContentServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterUserContentServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterUserContentServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server UserContentServiceServer) error {

	mux.Handle("GET", pattern_UserContentService_GetUserItems_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/content.UserContentService/GetUserItems", runtime.WithHTTPPathPattern("/content/{content_type}/{user_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_UserContentService_GetUserItems_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserContentService_GetUserItems_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_UserContentService_GetValued_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/content.UserContentService/GetValued", runtime.WithHTTPPathPattern("/content/{content_type}/{user_id}/{value}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_UserContentService_GetValued_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserContentService_GetValued_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PATCH", pattern_UserContentService_Add_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/content.UserContentService/Add", runtime.WithHTTPPathPattern("/content/{content_type}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_UserContentService_Add_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserContentService_Add_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_UserContentService_Remove_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/content.UserContentService/Remove", runtime.WithHTTPPathPattern("/content/{content_type}/{item_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_UserContentService_Remove_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserContentService_Remove_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterUserContentServiceHandlerFromEndpoint is same as RegisterUserContentServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterUserContentServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterUserContentServiceHandler(ctx, mux, conn)
}

// RegisterUserContentServiceHandler registers the http handlers for service UserContentService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterUserContentServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterUserContentServiceHandlerClient(ctx, mux, NewUserContentServiceClient(conn))
}

// RegisterUserContentServiceHandlerClient registers the http handlers for service UserContentService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "UserContentServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "UserContentServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "UserContentServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterUserContentServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client UserContentServiceClient) error {

	mux.Handle("GET", pattern_UserContentService_GetUserItems_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/content.UserContentService/GetUserItems", runtime.WithHTTPPathPattern("/content/{content_type}/{user_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_UserContentService_GetUserItems_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserContentService_GetUserItems_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_UserContentService_GetValued_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/content.UserContentService/GetValued", runtime.WithHTTPPathPattern("/content/{content_type}/{user_id}/{value}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_UserContentService_GetValued_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserContentService_GetValued_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PATCH", pattern_UserContentService_Add_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/content.UserContentService/Add", runtime.WithHTTPPathPattern("/content/{content_type}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_UserContentService_Add_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserContentService_Add_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_UserContentService_Remove_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/content.UserContentService/Remove", runtime.WithHTTPPathPattern("/content/{content_type}/{item_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_UserContentService_Remove_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserContentService_Remove_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_UserContentService_GetUserItems_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 1, 0, 4, 1, 5, 2}, []string{"content", "content_type", "user_id"}, ""))

	pattern_UserContentService_GetValued_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 1, 0, 4, 1, 5, 2, 1, 0, 4, 1, 5, 3}, []string{"content", "content_type", "user_id", "value"}, ""))

	pattern_UserContentService_Add_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"content", "content_type"}, ""))

	pattern_UserContentService_Remove_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 1, 0, 4, 1, 5, 2}, []string{"content", "content_type", "item_id"}, ""))
)

var (
	forward_UserContentService_GetUserItems_0 = runtime.ForwardResponseMessage

	forward_UserContentService_GetValued_0 = runtime.ForwardResponseMessage

	forward_UserContentService_Add_0 = runtime.ForwardResponseMessage

	forward_UserContentService_Remove_0 = runtime.ForwardResponseMessage
)
