package go_english_words

import (
	"bufio"
	"os"
	"path"
	"strings"

	go_file "github.com/pefish/go-file"
)

var words map[string]bool

func init() {
	sourceFileDir, err := go_file.SourceFileDir()
	if err != nil {
		panic(err)
	}
	words, err = loadDictionary(path.Join(sourceFileDir, "words.txt"))
	if err != nil {
		panic(err)
	}
}

// LoadDictionary 从文件中加载字典到 map 中
func loadDictionary(filePath string) (map[string]bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dictionary := make(map[string]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		dictionary[strings.ToLower(word)] = true
	}

	return dictionary, scanner.Err()
}

func IsWord(d string) bool {
	return words[d]
}
