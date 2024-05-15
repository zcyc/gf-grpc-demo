package interceptor

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SetCtxLanguage set the language of the context from grpc metadata.
func SetCtxLanguage(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (res interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	lang := md.Get("lang")
	if len(lang) == 0 {
		lang = []string{"en"}
	}
	ctx = gi18n.WithLanguage(ctx, lang[0])
	return handler(ctx, req)
}
