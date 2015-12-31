# nsqrepl

A REPL for sending requests to the nsq HTTP API, using a simple script.

![Progress](http://progressed.io/bar/30?title=Underway)

## Dependencies
```
go get github.com/russmack/nsqscript
```

## Install
```
go get github.com/russmack/nsqscript
git clone https://github.com/russmack/nsqrepl.git
go build
```

## Usage
```
./nsqrepl

>ping ip 127.0.0.1
>pause ip 127.0.0.1 topic thetopic channel thechannel
```

## License
BSD 3-Clause: [LICENSE.txt](LICENSE.txt)

[<img alt="LICENSE" src="http://img.shields.io/pypi/l/Django.svg?style=flat-square"/>](LICENSE.txt)
