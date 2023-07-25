# repl.uptime

repl.uptime is a Go package that keeps your Replit (repl) project continuously uptimed, ensuring 24/7 availability with no rest. It is a Go version inspired by the npm package [repl.uptime](https://www.npmjs.com/package/repl.uptime).

## Usage

To use repl.uptime, follow these two simple steps:

1. Create an `uptime.json` file with your desired options. For example:

```json
{
  "debug": false,
  "api": true,
  "port": "3000",
  "path": "/"
}
```

2. Import the package in your Go code and call `replup.Uptime()` in your `main()` function:

```go
package main

import (
  "github.com/abdlmutii/repl.uptime"
)

func main() {
  // Your other code

  replup.Uptime()

  // Your other code
}
```

## Installation

To install the repl.uptime library, follow these steps:

1. Open your command line interface and run the following command:
```
go get github.com/abdlmutii/repl.uptime
```

2. Import the library in your Go code using the import statement:
```go
import "github.com/abdlmutii/repl.uptime"
```

## License and Credits

repl.uptime is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.

Credits go to Abdlmu'tii for creating the Go version of repl.uptime and to Shuruhatik for maintaining the original npm package and the uptime API.

## Contact

For any inquiries or help, you can reach us:

- Email: abdlmutii@outlook.com
- Discord Server: [dsc.gg/abdlmutii](https://dsc.gg/abdlmutii)
- Shuruhatik Discord Server: [discord.gg/YEdA2BpQrd](https://discord.gg/YEdA2BpQrd)

Feel free to contact us with any questions, suggestions, or issues you may have.
