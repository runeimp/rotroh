RotRoh
======

Command line tool to apply one of several transforms to text input


Features
--------

* Can transform standard input or multiple quoated command line arguments
* Transforms
	* Base64: use the `-b64` option
	* ROT-13: use the `-13` option
	* ROT-47: use the `-47` option
	* RotRoh: the default if no other options are used
		* RotRoh = ROT-47 + Base64
* RotRoh should be pronounced as a reference to Scooby Doo's "Ruh Roh"


Examples
--------

```bash
$ rot -h
RotRoh v1.0.0
Usage: rot [OPTIONS] STRINGS

OPTIONS:
  -13       Use ROT-13 transform (default: false)
  -47       Use ROT-47 transform (default: false)
  -b64      Use Base64 codex (default: false)
  -help     Display this help info (default: false)
  -rotroh   Use RotRoh codex (default: true)
  -version  Display version info (default: false)

```



Articles & Reference
--------------------

* [ROT-13 Cipher - ROT13 - Online Text Decoder, Encoder, Translator][]
* [ROT-47 Cipher - ROT47 - Online Decoder, Encoder, Translator][]



[ROT-13 Cipher - ROT13 - Online Text Decoder, Encoder, Translator]: https://www.dcode.fr/rot-13-cipher
[ROT-47 Cipher - ROT47 - Online Decoder, Encoder, Translator]: https://www.dcode.fr/rot-47-cipher


