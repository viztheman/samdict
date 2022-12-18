// Entry point package.
package main

import (
	"fmt"
	"os"
	"strings"
	"samdict/cmu"
)

const CMUDICT_URL string = "https://raw.githubusercontent.com/Alexir/CMUdict/master/cmudict-0.7b"
const MAIN_DICT_FILE string = "cmudict.txt"
const SUPP_DICT_FILE string = "userdict.txt"
const PITCH_ADJUSTMENT int = 3

func mainDict() (main cmu.Dictionary, err error) {
	main, err = cmu.Read(MAIN_DICT_FILE, PITCH_ADJUSTMENT)
	if err != nil && !os.IsNotExist(err) {
		return
	}

	if os.IsNotExist(err) {
		err = cmu.Download(CMUDICT_URL, MAIN_DICT_FILE)
		if err == nil {
			main, err = cmu.Read(MAIN_DICT_FILE, PITCH_ADJUSTMENT)
		}
	}

	return
}

func transformWords(dict cmu.Dictionary, words []string) (result []string) {
	for _, word := range words {
		translated, ok := dict.Lookup(word)
		if !ok {
			result = append(result, "["+word+"]")
			continue
		}
		result = append(result, translated)
	}

	return
}

func main() {
	main, err := mainDict()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	supp, err := cmu.Read(SUPP_DICT_FILE, 0)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
	main.Merge(supp)

	translation := transformWords(main, os.Args[1:])
	fmt.Print("say \"", strings.Join(translation, " "), "\"\n")
}
