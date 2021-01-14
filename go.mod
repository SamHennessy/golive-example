module github.com/SamHennessy/golive-example

go 1.15

require (
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/brendonferreira/golive v0.0.0-20210106184349-68417b98b395
	github.com/gofiber/fiber/v2 v2.3.2
	github.com/gofiber/websocket/v2 v2.0.2
	github.com/klauspost/compress v1.11.6 // indirect
	github.com/savsgio/gotils v0.0.0-20210105085219-0567298fdcac // indirect
	github.com/valyala/fasthttp v1.19.0 // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20210105210732-16f7687f5001 // indirect
)

//replace github.com/brendonferreira/golive => github.com/SamHennessy/golive v0.0.0-20210107112545-cc34e325b33d

replace github.com/brendonferreira/golive => ../golive
