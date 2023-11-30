package wget

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"slices"

	"golang.org/x/net/html"
)

var urlAttrs = []string{
	"src",
	"srcset",
	"href",
}

type Options struct {
	Recursive bool
}

type Client struct {
	seen map[string]struct{}
}

func NewClient() *Client {
	return &Client{
		seen: make(map[string]struct{}),
	}
}

// 1. идем по урлу и получаем ресурс
// 2. сохраняем его в файл
// 3. отмечаем, что этот урл просмотрен
// 4. пытаемся распарсить содержимое как html и находим все ссылки
// 5. проходимся по всем ссылкам на ресурсы сайта
// 6. повторить с шага 1

func (c *Client) Get(rawURL string, opts Options) error {
	// parsing url and checking validity
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	//

	// fetching resource
	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error Status Code", resp.StatusCode, rawURL)
		return errors.New("bad request")
	}
	//

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	// saving resource
	pathname := urlToPathname(parsedURL, opts.Recursive)

	file, err := createFile(pathname)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, tee)
	if err != nil {
		return err
	}
	//

	c.seen[rawURL] = struct{}{}

	rawRefs, err := parseRefs(&buf)
	if err != nil {
		return err
	}

	// fetch related resources
	for i, rawRef := range rawRefs {
		parsedRef, err := url.Parse(rawRef)
		if err != nil {
			fmt.Println("invalid ref")
			continue
		}

		resolvedRef := parsedURL.ResolveReference(parsedRef)
		if resolvedRef.Hostname() != parsedURL.Hostname() {
			continue
		}

		ref := resolvedRef.String()
		fmt.Println(i, rawRef, resolvedRef.String(), resolvedRef.Hostname())

		if _, ok := c.seen[ref]; !ok {
			c.Get(ref, opts)
		}
	}
	//

	return nil
}

func urlToPathname(u *url.URL, full bool) string {
	dir, filename := path.Split(u.Path)
	if filename == "" {
		filename = "index.html"
	}

	if full {
		return filepath.Join(u.Hostname(), dir, filename)
	}

	return filename
}

func createFile(pathname string) (file *os.File, err error) {
	dir, _ := filepath.Split(pathname)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return file, err
	}

	file, err = os.Create(pathname)
	if err != nil {
		return file, err
	}

	return file, nil
}

func parseRefs(r io.Reader) (rawRefs []string, err error) {
	doc, err := html.Parse(r)
	if err != nil {
		return rawRefs, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if slices.Contains(urlAttrs, attr.Key) {
					rawRefs = append(rawRefs, attr.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return rawRefs, err
}
