RotRoh v0.1.0
=============

A Go package and command line tool to apply one of several transforms to text input


Features
--------

* Can transform standard input or multiple quoated command line arguments
* Transforms
	* Base64: use the `-base64` option
	* ROT-13: use the `-rot13` option
	* ROT-47: use the `-rot47` option
	* RotRoh: the default if no other options are used
		* RotRoh = ROT-47 + Base64
	* Custom ROT: use the `-custom-rot`
	* Pipeable: use the `-pipe` option to not add a newline to the output
* RotRoh should be pronounced as an homage to Scooby Doo's "Ruh Roh"
* Relies only on the Go stdlib. No 3rd party dependencies.


Examples
--------

### Help

```bash
$ rotroh -h
RotRoh v0.1.0
Usage: rotroh [OPTIONS] STRINGS

OPTIONS:
  -base64          Use Base64 codex (default: false)
  -help            Display this help info (default: false)
  -pipe            Do not output a newline at the end (default: false)
  -rot-custom SET  Use a custom ROT transform set
  -rot13           Use ROT-13 transform (default: false)
  -rot47           Use ROT-47 transform (default: false)
  -rotroh          Use RotRoh codex (default: true)
  -version         Display version info (default: false)

```

### Processing Command Line Arguments

```bash
$ rotroh 'Where is Waldo?'
KDk2QzYgOkQgKDI9NUBu
```

```bash
$ rotroh 'Where is Waldo?' KDk2QzYgOkQgKDI9NUBu 'dzYgPkZERSAzNiA6PyBFOTYgNDI/NUogNDI/NiA3QEM2REVd'
KDk2QzYgOkQgKDI9NUBu
Where is Waldo?
He must be in the candy cane forest.
```

### Processing Standard Input

```bash
$ echo 'Where is Waldo?' | rotroh | rotroh
enM8YSJLKjh+PCI4enN4aH0mcUZyOGxsCg==
```

Don't use `echo` in a pipeline to `rotroh` unless you really want a newline character to be part of the processing


```bash
$ printf 'Where is Waldo?' | rotroh | rotroh
enM8YSJLKjh+PCI4enN4aH0mcUYK
```

Neet to make sure we're not adding extra newlines when using `rotroh` in a pipeline

```bash
$ printf 'Where is Waldo?' | rotroh -pipe | rotroh
Where is Waldo?
```

### Using a Custom ROT

```bash
$ rotroh -rot-custom '1234567890' '1342 Fast Ln.'
0879 Fast Ln.
```

```bash
$ rotroh -pipe -rot-custom '1234567890' '1342 Fast Ln.' | rotroh -rot-custom '1234567890'
1342 Fast Ln.
```


Test Coverage
-------------

73.7% of statements


ToDo
----

* Add custom RotRoh
* Add Base32, etc.
* Allow use of all options for layered and ordered processing
* Allow a config file to store custom transforms, etc.
* data URL codex?


Warning
-------

While this is technically an encryption and obfuscation tool it is not, I repeat **NOT**, to be used were real security is needed. This project is more for fun than security. In other words, it _might_ slow down Velma for a hot second if she couldn't find her glasses. :wink:


Articles & Reference
--------------------

* [ROT-13 Cipher - ROT13 - Online Text Decoder, Encoder, Translator][]
* [ROT-47 Cipher - ROT47 - Online Decoder, Encoder, Translator][]


[ROT-13 Cipher - ROT13 - Online Text Decoder, Encoder, Translator]: https://www.dcode.fr/rot-13-cipher
[ROT-47 Cipher - ROT47 - Online Decoder, Encoder, Translator]: https://www.dcode.fr/rot-47-cipher

