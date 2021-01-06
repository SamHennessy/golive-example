module github.com/SamHennessy/golive-example

go 1.15

require (
	github.com/brendonferreira/golive v0.0.0-20210105171218-45446051266e
	github.com/gofiber/fiber/v2 v2.2.3
	github.com/gofiber/websocket/v2 v2.0.2
)

replace github.com/brendonferreira/golive => github.com/SamHennessy/golive v0.0.0-20210106120408-d978afac08fc

//replace github.com/brendonferreira/golive => ../golive
