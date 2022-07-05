package server

//
//import (
//    "context"
//    v1 "github.com/BingguWang/bingfood-client-micro/api/order/service/v1"
//    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/conf"
//    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/service"
//    "github.com/BingguWang/bingfood-client-micro/app/user/service/global"
//    "github.com/go-kratos/kratos/v2/log"
//    "github.com/go-kratos/kratos/v2/middleware/selector"
//    "github.com/go-kratos/kratos/v2/transport/http"
//    "github.com/gorilla/handlers"
//)
//
//func NewSkipAuthMatcher() selector.MatchFunc {
//    whiteList := make(map[string]struct{})
//    whiteList["/user.v1.BingfoodService/LoginOrRegister"] = struct{}{}
//
//    return func(ctx context.Context, operation string) bool {
//        if _, ok := whiteList[operation]; ok {
//            return false // 无需auth
//        }
//        return true
//    }
//}
//
//// NewHTTPServer new a HTTP configs.
//func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, svc *service.OrderServiceImpl, logger log.Logger) *http.Server {
//    var opts = []http.ServerOption{
//
//        //http.Middleware(
//        //    recovery.Recovery(),
//        //    selector.Server(middleware.AuthMiddleware()).Match(NewSkipAuthMatcher()).Build(),
//        //),
//
//        // Filter在middleware之前执行
//        http.Filter(handlers.CORS(
//            handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
//            handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
//            handlers.AllowedOrigins([]string{"*"}), // 资源可以被任意外域访问
//            //handlers.AllowedOrigins([]string{"https://foo.example"}), // 资源仅可以被https://foo.example访问
//        )),
//        // 可以写自己的filter handler
//    }
//    InitGlobalValue(c, jwtc, svc)
//
//    if c.Http.Network != "" {
//        opts = append(opts, http.Network(c.Http.Network))
//    }
//    if c.Http.Addr != "" {
//        opts = append(opts, http.Address(c.Http.Addr))
//    }
//    if c.Http.Timeout != nil {
//        opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
//    }
//    srv := http.NewServer(opts...)
//    v1.RegisterOrderServiceHTTPServer(srv, svc)
//    return srv
//}
//func InitGlobalValue(c *conf.Server, jwtc *conf.JWT, svc *service.OrderServiceImpl) {
//    global.JWT_SECRET = []byte(jwtc.Secret)
//}
