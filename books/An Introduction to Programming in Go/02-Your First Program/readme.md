**How to upgrade version of golang on Mac OSX?**

via Brew from [stackoverflow](https://stackoverflow.com/questions/34587978/how-to-upgrade-go-from-1-4-to-1-5-via-homebrew)
```
brew update
brew upgrade

brew install go --cross-compile-common
```

`brew upgrade` might upgrade go alone without having to try to install it, after you run brew update/upgrade try running go version and see which version you have.


via Manual installation https://medium.com/golang-learn/quick-go-setup-guide-on-mac-os-x-956b327222b8


**Basic command**

- `go version` for version checking
- `go help` for show all command