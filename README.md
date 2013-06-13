# Z80 - A Zilog Z80 emulator written in Go

The z80 package implements a Zilog Z80 emulator written in Go. The CPU
emulation code is generated using a modified version of the
<tt>z80.pl</tt> script shipped with
[FUSE](http://fuse-emulator.sourceforge.net/). The script has been
hacked to generate Go code rather than C code.

# Features

* Machine independent (i.e. the same package could be used to emulate
  several Z80-based machines).
* Tested against the excellent test-suite shipped with FUSE.
* It includes a couple of disassembling/debugging functions.

# Emulators

The z80 package is at the core of a couple of emulator written in
Go. Namely:

* [GoSpeccy](https://github.com/remogatto/gospeccy) - A Spectrum ZX 48k Emulator
* [SMS](https://github.com/remogatto/sms) - A Sega Master System Emulator
* [TRS80](https://github.com/lkesteloot/trs80) - A TRS-80 Model III Emulator
 
# Contributors

* [Atom](https://github.com/0xe2-0x9a-0x9b)
* [Lawrence Kesteloot](https://github.com/lkesteloot)

# License

Copyright (c) 2013 Andrea Fazzi

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
