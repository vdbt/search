// Auto-generated by avdl-compiler v1.3.6 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/remote.avdl

package chat1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type MessageBoxed struct {
	ServerHeader    *MessageServerHeader `codec:"serverHeader,omitempty" json:"serverHeader,omitempty"`
	ClientHeader    MessageClientHeader  `codec:"clientHeader" json:"clientHeader"`
	HeaderSignature SignatureInfo        `codec:"headerSignature" json:"headerSignature"`
	BodyCiphertext  EncryptedData        `codec:"bodyCiphertext" json:"bodyCiphertext"`
	BodySignature   SignatureInfo        `codec:"bodySignature" json:"bodySignature"`
	KeyGeneration   int                  `codec:"keyGeneration" json:"keyGeneration"`
}

type ThreadViewBoxed struct {
	Messages   []MessageBoxed `codec:"messages" json:"messages"`
	Pagination *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetInboxRemoteArg struct {
	Pagination *Pagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetInboxByTLFIDRemoteArg struct {
	TLFID TLFID `codec:"TLFID" json:"TLFID"`
}

type GetThreadRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MarkAsRead     bool           `codec:"markAsRead" json:"markAsRead"`
	Pagination     *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetConversationMetadataRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
}

type PostRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MessageBoxed   MessageBoxed   `codec:"messageBoxed" json:"messageBoxed"`
}

type NewConversationRemoteArg struct {
	IdTriple ConversationIDTriple `codec:"idTriple" json:"idTriple"`
}

type NewConversationRemote2Arg struct {
	IdTriple   ConversationIDTriple `codec:"idTriple" json:"idTriple"`
	TLFMessage MessageBoxed         `codec:"TLFMessage" json:"TLFMessage"`
}

type GetMessagesRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MessageIDs     []MessageID    `codec:"messageIDs" json:"messageIDs"`
}

type MarkAsReadArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MsgID          MessageID      `codec:"msgID" json:"msgID"`
}

type RemoteInterface interface {
	GetInboxRemote(context.Context, *Pagination) (InboxView, error)
	GetInboxByTLFIDRemote(context.Context, TLFID) ([]Conversation, error)
	GetThreadRemote(context.Context, GetThreadRemoteArg) (ThreadViewBoxed, error)
	GetConversationMetadataRemote(context.Context, ConversationID) (Conversation, error)
	PostRemote(context.Context, PostRemoteArg) (MessageID, error)
	NewConversationRemote(context.Context, ConversationIDTriple) (ConversationID, error)
	NewConversationRemote2(context.Context, NewConversationRemote2Arg) (ConversationID, error)
	GetMessagesRemote(context.Context, GetMessagesRemoteArg) ([]MessageBoxed, error)
	MarkAsRead(context.Context, MarkAsReadArg) error
}

func RemoteProtocol(i RemoteInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.remote",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getInboxRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetInboxRemote(ctx, (*typedArgs)[0].Pagination)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getInboxByTLFIDRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxByTLFIDRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxByTLFIDRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxByTLFIDRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetInboxByTLFIDRemote(ctx, (*typedArgs)[0].TLFID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getThreadRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetThreadRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetThreadRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetThreadRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetThreadRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getConversationMetadataRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetConversationMetadataRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetConversationMetadataRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetConversationMetadataRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetConversationMetadataRemote(ctx, (*typedArgs)[0].ConversationID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"postRemote": {
				MakeArg: func() interface{} {
					ret := make([]PostRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PostRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]PostRemoteArg)(nil), args)
						return
					}
					ret, err = i.PostRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationRemote": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationRemoteArg)(nil), args)
						return
					}
					ret, err = i.NewConversationRemote(ctx, (*typedArgs)[0].IdTriple)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationRemote2": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationRemote2Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationRemote2Arg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationRemote2Arg)(nil), args)
						return
					}
					ret, err = i.NewConversationRemote2(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getMessagesRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetMessagesRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetMessagesRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetMessagesRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetMessagesRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"markAsRead": {
				MakeArg: func() interface{} {
					ret := make([]MarkAsReadArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]MarkAsReadArg)
					if !ok {
						err = rpc.NewTypeError((*[]MarkAsReadArg)(nil), args)
						return
					}
					err = i.MarkAsRead(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RemoteClient struct {
	Cli rpc.GenericClient
}

func (c RemoteClient) GetInboxRemote(ctx context.Context, pagination *Pagination) (res InboxView, err error) {
	__arg := GetInboxRemoteArg{Pagination: pagination}
	err = c.Cli.Call(ctx, "chat.1.remote.getInboxRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetInboxByTLFIDRemote(ctx context.Context, TLFID TLFID) (res []Conversation, err error) {
	__arg := GetInboxByTLFIDRemoteArg{TLFID: TLFID}
	err = c.Cli.Call(ctx, "chat.1.remote.getInboxByTLFIDRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetThreadRemote(ctx context.Context, __arg GetThreadRemoteArg) (res ThreadViewBoxed, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.getThreadRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetConversationMetadataRemote(ctx context.Context, conversationID ConversationID) (res Conversation, err error) {
	__arg := GetConversationMetadataRemoteArg{ConversationID: conversationID}
	err = c.Cli.Call(ctx, "chat.1.remote.getConversationMetadataRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) PostRemote(ctx context.Context, __arg PostRemoteArg) (res MessageID, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.postRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) NewConversationRemote(ctx context.Context, idTriple ConversationIDTriple) (res ConversationID, err error) {
	__arg := NewConversationRemoteArg{IdTriple: idTriple}
	err = c.Cli.Call(ctx, "chat.1.remote.newConversationRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) NewConversationRemote2(ctx context.Context, __arg NewConversationRemote2Arg) (res ConversationID, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.newConversationRemote2", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetMessagesRemote(ctx context.Context, __arg GetMessagesRemoteArg) (res []MessageBoxed, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.getMessagesRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) MarkAsRead(ctx context.Context, __arg MarkAsReadArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.markAsRead", []interface{}{__arg}, nil)
	return
}
