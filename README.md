![Go Report](https://goreportcard.com/badge/github.com/SerjRamone/steam-cli-autentificator)
![Repository Top Language](https://img.shields.io/github/languages/top/SerjRamone/steam-cli-autentificator)
![Github Repository Size](https://img.shields.io/github/repo-size/SerjRamone/steam-cli-autentificator)
![Github Open Issues](https://img.shields.io/github/issues/SerjRamone/steam-cli-autentificator)
![Lines of code](https://img.shields.io/tokei/lines/github/SerjRamone/steam-cli-autentificator)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/SerjRamone/steam-cli-autentificator)
![GitHub contributors](https://img.shields.io/github/contributors/SerjRamone/steam-cli-autentificator)

<img align="right" width="30%" src="./images/gopher.png">

### Steam CLI Autentificator
This is a command-line tool written in Go that generates a Steam Guard authorization code (2FA) 
by providing a maFile (Mobile Authenticator) file. The maFile file contains a shared secret 
that is used to generate the Steam Guard code. The tool uses the HMAC-SHA1 algorithm to generate the code.

### Installation
1. Clone the repository to your local machine.
```
git clone https://github.com/SerjRamone/steam-cli-autentificator
```

2. Change directory to the cloned repository.
```
cd steam-cli-autentificator
```

3. Install the dependencies.
```
go mod download
```

4. Build the tool.
```
go build -o sca
```

### Usage
```
./sca 2fa /path/to/maFile
```
The path to the maFile is required.
<img src="./images/1.png">

### License
This tool is licensed under the [MIT License](LICENSE.md).
