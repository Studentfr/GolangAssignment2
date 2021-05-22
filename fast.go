package asd

import (
	"bufio"
	"bytes"

	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "log"
)

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Browsers []string `json:"browsers"`
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	user := User{}
	if err != nil {
		panic(err)
	}
	input := bufio.NewScanner(file)
	foundUsers := ""
	seenBrowsers := []string{}
	uniqueBrowsers := 0
	//var users []User
	for i:=0; input.Scan(); i++ {
		line := input.Bytes()
		if !(bytes.Contains(line, []byte("Android")) || bytes.Contains(line, []byte("MSIE"))) {
			continue
		}

		// fmt.Printf("%v %v\n", err, line)
		err := user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}
		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {

			if ok := strings.Contains(browser, "Android"); ok {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		for _, browser := range user.Browsers {

			if ok := strings.Contains(browser, "MSIE"); ok {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := strings.Replace(user.Email, "@", " [at] ", -1)
		foundUsers += "["+strconv.Itoa(i)+"] "+user.Name+" <"+email+">\n"
		//users = append(users, user)
	}


	//for i, user := range users {
	//	fmt.Println(user)
	//
	//
	//}
	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}