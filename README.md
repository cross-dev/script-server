It is a trendy way to provide installation scripts to the users like this:

```
curl -sSL https://get.blingy.com/ | sh
```
So, we should keep up with the trend...

## Install

```shell
$ go get github.com/cross-dev/script-server
```

## Use

```shell
$ $GOPATH/bin/script-server -h
Usage: script-server [options] [-|filename]
  -b string
        Base URL pathname (default "/")
  -l string
        Interface and port to listen (default ":41267")
$ $GOPATH/bin/script-server -l 34567 -b /a/b <(echo '{{a}}-{{b}}') >/dev/null 2>&1 &
$ curl -s 'http://localhost:34567/a/b/get?a=oo&b=8'
oo-8
```

## Build from source

The `$GOPATH` has to be assigned.

```shell
$ make
```

Contributions through PR and issues through issues are welcome.

## Why this project

I think next step would be to begin customizing the said scripts and build them
on the fly. Now our `blingy.com` would accept parameters in the URL query and
bake the script content from the template.

## Implementation

The server is done in [Go](https://golang.org/). It is a simple and powerful
language, which can easily create HTTP servers and produce textual output from
a template.

All parameters passed over the GET request query are parsed and then they are fed
to the text processor along with the template. The processor creates output, which
becomes an installer script content returned to the client.

## Notes

* The application does not support HTTPS - proxy it through a secure server
* A customized URL will quickly become hideous and impossible to remember - use
URL shortener
