module github.com/SamHennessy/golive-example

go 1.15

require (
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/brendonmatos/golive v0.0.0-20210117233607-d67133e0b23c
	github.com/gofiber/fiber/v2 v2.3.3
	github.com/gofiber/websocket/v2 v2.0.2
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/savsgio/gotils v0.0.0-20210105085219-0567298fdcac // indirect
	github.com/valyala/fasthttp v1.19.0 // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78 // indirect
)

//replace github.com/brendonmatos/golive => github.com/SamHennessy/golive <branch>

//replace github.com/brendonmatos/golive => ../golive
