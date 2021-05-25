package main

import (
	"bufio"
	"fmt"
	"hw3_bench/user"
	"io"
	"os"
	"strings"
)

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seenBrowsers := make(map[string]bool)
	var foundUsers strings.Builder

	sc := bufio.NewScanner(file)
	for i := 0; sc.Scan(); i++ {
		user := user.User{}
		err := user.UnmarshalJSON([]byte(sc.Text()))
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "Android") {
				isAndroid = true
				_, notSeenBefore := seenBrowsers[browser]
				if !notSeenBefore {
					seenBrowsers[browser] = true
				}
			}

			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				_, notSeenBefore := seenBrowsers[browser]
				if !notSeenBefore {
					seenBrowsers[browser] = true
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email, "@", " [at] ", -1)
		foundUsers.WriteString(fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email))
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers.String())
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
