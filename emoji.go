package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
	"golang.org/x/net/html"
)

var emojis []string

func getTitle(n *html.Node) (ok bool, title string) {
	for _, a := range n.Attr {
		if a.Key == "name" {
			title = a.Val
			ok = true
		}
	}

	return
}

func getCodes(s string) (codes []string) {
	codes = strings.Split(s, "_")

	return
}

func getEmoji(codes []string) (emoji string) {

	for _, code := range codes {
		padded := fmt.Sprintf("%08s", code)
		unicode := "'\\U" + padded + "'"
		character, _ := strconv.Unquote(unicode)
		emoji += character
	}

	return
}

func parseHTML(filepath string) (count int, err error) {
	file, err := os.Open(filepath)

	if err != nil {
		return 0, err
	}

	doc, err := html.Parse(file)
	if err != nil {
		return 0, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			ok, title := getTitle(n)
			if !ok {
				return
			}

			codes := getCodes(title)
			emoji := getEmoji(codes)

			emojis = append(emojis, emoji)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return len(emojis), nil
}

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run emoji parser",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "source",
					Usage: "Source file for emoji",
				},
				&cli.StringFlag{
					Name:  "destination",
					Usage: "Destination file for emoji",
				},
				&cli.DurationFlag{
					Name:  "interval",
					Usage: "Interval of emoji generation in seconds",
				},
			},
			Action: func(c *cli.Context) error {
				source, _ := filepath.Abs(c.String("source"))
				fmt.Printf("Reading emoji from %s\n", source)

				destination, _ := filepath.Abs(c.String("destination"))
				fmt.Printf("Writing emoji to %s\n", destination)

				fmt.Printf("Ticking every %s\n", c.Duration("interval").String())

				count, _ := parseHTML(source)
				fmt.Printf("Parsed %d emojis\n", count)

				s := rand.NewSource(time.Now().UnixNano())
				r := rand.New(s)

				ticker := time.NewTicker(c.Duration("interval"))
				for range ticker.C {
					emoji := emojis[r.Intn(len(emojis))]
					fmt.Printf("Writing emoji to file: %s\n", emoji)
					ioutil.WriteFile(destination, []byte(emoji), 0644)
				}

				ticker.Stop()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
