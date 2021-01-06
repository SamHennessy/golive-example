# Example GoLive Application

If you're interested in using GoLive, you can use this code to see how to use it in a full app.

This project is only one way of doing things and only presented here to act as inspiration.

Often the features are needlessly complicated (for example, the real-time clock) but that is by choice to demonstrate different technical features.

GoLive is a new project with lots of opportunities to contribute code. As such,  I'm also using this project to help push the development of GoLive by giving me inspiration for new features.

## Technology

- [GoLive](https://github.com/brendonferreira/golive)
    - Reactive HTML Server Side Rendered by GoLang over WebSockets
    - Acts as the view layer
- [Bulma](https://bulma.io/)
    - CSS framework based on Flexbox
    - Has good components, like a modal, that doesn't need its own JavaScript to work
- [Animate.css](https://animate.style/)
    - Animate.css is a library of ready-to-use, cross-browser animations.
    - No JavaScript required.
- [Material Design Icons](https://materialdesignicons.com/)
    - Icon library
    - No JavaScript required. While working on this project, I learned some Icon libraries, like FontAwesome, will change your DOM using JavaScript.
- [Fiber](https://github.com/gofiber/fiber)
    - Express inspired web framework written in Go
    - Used as the HTTP router
    - GoLive is currently using Fiber's WebSocket implementation, so I'm using it as a router because it's already there. You could use GoLive with any router you are not limited to Fiber.

### JavaScript
You can use JavaScript libraries and GoLive together. If a library makes changes to the DOM, you may have issues with the DOM diffing that GoLive does on the server-side. This is due to the issue where what GoLive thinks you have in the browser will be different from what is there.

As a personal goal, I wanted to see what I could do without using any JavaScript for the app. GoLive uses JavaScript to enable all of this, so it's not about being anti-JavaScript, just a preference for writing all my logic in Go.

## Working With The Project

To see how to work with the code in the project, build, run, etc. Link in the [Makefile](./Makefile)

## TODO

- Config from flags and env vars
- Move stuff in main out to a web server
- Auto-reset all data with a global count down
- Add a structured logger like ZeroLog?
- Global stats? E.g. active sessions, messages in and out, memory usage, etc.
- Use Pub Sub to show server pushed updates

### More Apps

Possible other apps to build:

- Chat / Slack clone
- Twitter clone
- WYSIWYG editor
- Server "real-time" metrics
