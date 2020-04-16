package cmd

import (
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/imthaghost/shareit/upload"
)

func shareit(args []string) {
	// anonfiles api for now :)
	anonurl := "https://api.anonfiles.com/upload"
	// first argument should be the file path that we are uploading
	filepath := args[0]
	// upload the file to anonfiles
	data := upload.FileUpload(anonurl, filepath, "file")
	// get short link
	link := data.Data.File.URL.Short
	color.Green(link + " " + "copied to clipboard!")
	// copy to clipboard
	clipboard.WriteAll(link)
}
