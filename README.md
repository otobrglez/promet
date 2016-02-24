# promet

Experimental shared library for Slovenian traffic services.

- [Oto Brglez](https://github.com/otobrglez)

# Requirements

```bash
# Python
mkvirtualenv --no-site-packages promet
workon promet
pip install --upgrade -r python/requirements.txt


```

# Compiling (OSX)

```bash
make

# Inspecting object
nm -gU promet.so
```


# Resources
- [golang / go / cgo](https://github.com/golang/go/wiki/cgo)
- [c-go-cgo](http://blog.golang.org/c-go-cgo)
- [nm](https://sourceware.org/binutils/docs/binutils/nm.html)
- [cgo-slides](http://akrennmair.github.io/golang-cgo-slides/#8)

# License
- [MIT](LICENSE)
