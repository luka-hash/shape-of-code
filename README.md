# shape-of-code

~Generate `png` out of your source code file~

## Dependencies

- [Go](https://go.dev/)
- [gg](https://github.com/fogleman/gg)
  - Go Graphics - 2d rendering in Go with simple API

## Installation

```
git clone github.com/luka-hash/shape-of-code
make # this will just build the executable
sudo make install PREFIX=/usr/local # this will install it globally
make install PREFIX="$HOME"/.local # will install it locally
```

Also, you can set `DESTDIR` to install it into a different system root, e.g. use
`make install DESTDIR="$pkgdir" PREFIX=/usr` on Arch.

## Usage

```
$ shape-of-code input.txt # will output input.txt.png on success
$ shape-of-code input1.txt input2.txt # will output input1.txt.png and input2.txt.png
$ shape-of-code help (or --help or -h) # prints usage string
$ shape-of-code version (or --version or -v) # prints version number
```

## Is it any good?

Yes.

## TODO

- add some spacing
- make categories for types of lines (docs, comments, regular lines, etc.)
- create shell completions

## Licence

This code is licensed under MIT licence (see LICENCE for details).