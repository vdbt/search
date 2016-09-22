// Auto-generated by avdl-compiler v1.3.6 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_fs.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type FSActivityArg struct {
	Notification FSNotification `codec:"notification" json:"notification"`
}

type FSSyncActivityArg struct {
	Status FSPathSyncStatus `codec:"status" json:"status"`
}

type FSEditListResponseArg struct {
	Edits     []FSNotification `codec:"edits" json:"edits"`
	RequestID int              `codec:"requestID" json:"requestID"`
}

type FSSyncStatusResponseArg struct {
	Status    FSSyncStatus `codec:"status" json:"status"`
	RequestID int          `codec:"requestID" json:"requestID"`
}

type NotifyFSInterface interface {
	FSActivity(context.Context, FSNotification) error
	FSSyncActivity(context.Context, FSPathSyncStatus) error
	FSEditListResponse(context.Context, FSEditListResponseArg) error
	FSSyncStatusResponse(context.Context, FSSyncStatusResponseArg) error
}

func NotifyFSProtocol(i NotifyFSInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyFS",
		Methods: map[string]rpc.ServeHandlerDescription{
			"FSActivity": {
				MakeArg: func() interface{} {
					ret := make([]FSActivityArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FSActivityArg)
					if !ok {
						err = rpc.NewTypeError((*[]FSActivityArg)(nil), args)
						return
					}
					err = i.FSActivity(ctx, (*typedArgs)[0].Notification)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"FSSyncActivity": {
				MakeArg: func() interface{} {
					ret := make([]FSSyncActivityArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FSSyncActivityArg)
					if !ok {
						err = rpc.NewTypeError((*[]FSSyncActivityArg)(nil), args)
						return
					}
					err = i.FSSyncActivity(ctx, (*typedArgs)[0].Status)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"FSEditListResponse": {
				MakeArg: func() interface{} {
					ret := make([]FSEditListResponseArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FSEditListResponseArg)
					if !ok {
						err = rpc.NewTypeError((*[]FSEditListResponseArg)(nil), args)
						return
					}
					err = i.FSEditListResponse(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"FSSyncStatusResponse": {
				MakeArg: func() interface{} {
					ret := make([]FSSyncStatusResponseArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FSSyncStatusResponseArg)
					if !ok {
						err = rpc.NewTypeError((*[]FSSyncStatusResponseArg)(nil), args)
						return
					}
					err = i.FSSyncStatusResponse(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
		},
	}
}

type NotifyFSClient struct {
	Cli rpc.GenericClient
}

func (c NotifyFSClient) FSActivity(ctx context.Context, notification FSNotification) (err error) {
	__arg := FSActivityArg{Notification: notification}
	err = c.Cli.Notify(ctx, "keybase.1.NotifyFS.FSActivity", []interface{}{__arg})
	return
}

func (c NotifyFSClient) FSSyncActivity(ctx context.Context, status FSPathSyncStatus) (err error) {
	__arg := FSSyncActivityArg{Status: status}
	err = c.Cli.Notify(ctx, "keybase.1.NotifyFS.FSSyncActivity", []interface{}{__arg})
	return
}

func (c NotifyFSClient) FSEditListResponse(ctx context.Context, __arg FSEditListResponseArg) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifyFS.FSEditListResponse", []interface{}{__arg})
	return
}

func (c NotifyFSClient) FSSyncStatusResponse(ctx context.Context, __arg FSSyncStatusResponseArg) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifyFS.FSSyncStatusResponse", []interface{}{__arg})
	return
}
