This project aims to implement an editor and language agnostic backend with the initial
scope of supporting:

* Code completion
* Going to definitions/implementations
* Going to the documentation

The project was born out of a growing frustration with the badly architectured Sublime Text
code completion plugins [SublimeClang](https://github.com/quarnster/SublimeClang),
[SublimeJava](https://github.com/quarnster/SublimeJava) and
[CompleteSharp](https://github.com/quarnster/CompleteSharp). Either large code rewrites
would have to be made or the code had to be written from scratch to amend their
weaknesses. I decided to go with the latter approach.

[Go](http://golang.org/) was chosen as the one language to rule them all because put
simply, it has rocked my socks off and continue to do so as I discover more about it.
It's a joy to work with and its expressiveness, yet simplicity results in beautiful
code that reads like a haiku.

Collaborators wanting to aid the project are welcome to submit pull requests and/or
open up a new issue number for discussing implementation details not already discussed
in the [existing issues](https://github.com/quarnster/completion/issues?state=open) list.

The license of the project is the 2-clause BSD license:

```
Copyright (c) 2013, project authors (https://github.com/quarnster/completion/contributors)
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
