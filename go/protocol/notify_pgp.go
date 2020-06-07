// Auto-generated by avdl-compiler v1.3.4 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/notify_pgp.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PGPKeyInSecretStoreFileArg struct {
}

type NotifyPGPInterface interface {
	PGPKeyInSecretStoreFile(context.Context) error
}

func NotifyPGPProtocol(i NotifyPGPInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyPGP",
		Methods: map[string]rpc.ServeHandlerDescription{
			"pgpKeyInSecretStoreFile": {
				MakeArg: func() interface{} {
					ret := make([]PGPKeyInSecretStoreFileArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.PGPKeyInSecretStoreFile(ctx)
					return
				},
				MethodType: rpc.MethodNotify,
			},
		},
	}
}

type NotifyPGPClient struct {
	Cli rpc.GenericClient
}

func (c NotifyPGPClient) PGPKeyInSecretStoreFile(ctx context.Context) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifyPGP.pgpKeyInSecretStoreFile", []interface{}{PGPKeyInSecretStoreFileArg{}})
	return
}
