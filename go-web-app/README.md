# Go Web App Example

This is a very simple Go web app that you can use as an example on how to setup
Task as part of the build pipeline of a Go app.

Features:

- Build everything with a single command;
- It's fast since it just run the necessary steps;
- Allow live-reloading using the `--watch` flag.

How to use:

- This requires [Go][go] to be installed;
- [Install Task][installtask];
- Run `task dl-deps` to install the tools required to build this app;
- Run `task` to build and run the web app;
- Optionally you can use `task --watch` (or `task -w`) to re-compile and re-run
  the app on any file change.
- Open http://localhost:8383

[go]: https://golang.org/
[installtask]: https://taskfile.dev/installation/
