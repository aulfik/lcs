# lcs
Lcs is a CLI application for easly creating/generating LICENSE text.

Lcs providing a simple interface similar to git & go tools. Project using [Cobra](https://github.com/spf13/cobra) and [Survey](github.com/AlecAivazis/survey/v2).

Lcs provides Generate LICENSE text based [ChooseALicense.com](https://choosealicense.com/):
*  "MIT License",
*  "GNU AGPLv3",
*  "GNU GPLv3",
*  "GNU LGPLv3",
*  "Mozilla Public License 2.0",
*  "Apache License 2.0",
*  "Boost Software License 1.0",
*  "The Unlicense",

# Run it on your machine

Prerequisites an installation of Go 1.16 or later. For installation instructions, see [Installing Go](https://go.dev/doc/install).

    # Clone this repo
    git clone https://github.com/aulfik/lcs.git
    cd lcs

    # From the command line in the directory containing main.go, build and run the code. Use a dot argument to mean “build and run code in the current directory.”
    go build .
    ./lcs generate  

# Author

[Aulia Fikri](https://aulfikri.engineer): [@aulfik](https://github.com/aulfik)

# License

[MIT](LICENSE)
