# 72 Godziny

## Prerequisites

- [Go](https://golang.google.cn/doc/install)

Install templ and air:

```
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/air-verse/air@latest
```

Install dependencies

```
go mod tidy
```

If using VSCode, you propably want to use Go, Templ, and Tailwind plugins.

Install Tailwindcss binary, [standalone](https://tailwindcss.com/blog/standalone-cli), or by installing it as a global NPM package.

## Running the project

One process is set up to run the backend server, and another to recompile
Tailwind, so you need two commands running at the same time (for example
in VSCode's split terminals):

```
air
tailwindcss -i input.css -o ./static/styles.css --watch
```

After that, the local website is available at `localhost:8383`.

Using `8080` will make the hot reload not work!
