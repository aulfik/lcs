# lcs

Lcs is a CLI application to easily creating/generating LICENSE text. Lcs provide a simple interface similar to git & go tools. This project is using [Cobra](https://github.com/spf13/cobra) and [Survey](github.com/AlecAivazis/survey/v2).

Lcs generate a LICENSE text based on [Choose a License](https://choosealicense.com/):
  - MIT License
  - GNU AGPLv3
  - GNU GPLv3
  - GNU LGPLv3
  - Mozilla Public License 2.0
  - Apache License 2.0
  - Boost Software License 1.0
  - The Unlicense

# Run it on your machine

This tool requires you to install Go v1.16 or later. For installation instructions, see [Installing Go](https://go.dev/doc/install).

```bash
# Clone this repo
git clone https://github.com/aulfik/lcs
cd lcs

# From the command line in the directory containing main.go, build and run the code. Use a dot argument to mean “build and run code in the current directory.”
go build .
./lcs generate  
```

# Author

[Aulia Fikri](https://aulfikri.engineer): [@aulfik](https://github.com/aulfik)

# License

[MIT](LICENSE)
