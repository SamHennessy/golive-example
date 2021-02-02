module github.com/SamHennessy/golive-example

go 1.15

require (
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/brendonmatos/golive v0.0.2-0.20210201120951-3fb9d34fa79a
	github.com/gofiber/fiber/v2 v2.4.0
	github.com/gofiber/websocket/v2 v2.0.2
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/savsgio/gotils v0.0.0-20210120114113-f9d780dcbd93 // indirect
	github.com/valyala/fasthttp v1.19.0 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
)

//replace github.com/brendonmatos/golive => github.com/SamHennessy/golive <branch>

//replace github.com/brendonmatos/golive => ../golive
