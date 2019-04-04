package cstore

import homedir "github.com/mitchellh/go-homedir"

var FilePath, _ = homedir.Expand("~/.command-storer")
var HistPath, _ = homedir.Expand("~/.zsh_history")
