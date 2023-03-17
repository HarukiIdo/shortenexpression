package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"sync"
)

type message struct {
	Path      string `json:"path"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

const url = "https://index.golang.org/index"

func main() {

	var wg sync.WaitGroup
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("fail to request")
	}

	// 一行ごとに読み込む
	scanner := bufio.NewScanner(resp.Body)
	pkgLists := make([]string, 0)
	for scanner.Scan() {
		var m message
		if err := json.Unmarshal(scanner.Bytes(), &m); err != nil {
			log.Fatal("Json Unmarshal error:", err)
		}
		pkgLists = append(pkgLists, m.Path)
	}

	for i, pkgName := range pkgLists[0:10] {
		hashDir := sha256.Sum256([]byte(pkgName))
		dir := fmt.Sprintf("%x", hashDir[:8])
		dir = path.Join(".", "tmpdir", dir)
		fmt.Println(i, dir)

		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			if err := prepareDir(pkgName, dir); err != nil {
				fmt.Printf("prepareDir failed: %s", err)
			}
			fmt.Println("prepare directory")

			if err := execVet(pkgName, dir); err != nil {
				fmt.Printf("execVet failed: %s", err)
			}
		}(dir)
		wg.Wait()
	}
}

// mkdir tmpdir/[hash]/go.mod
// go mod init a
func prepareDir(pkgName, dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	if _, err := os.Stat(path.Join(dir, "go.mod")); os.IsNotExist(err) {
		cmd := exec.Command("go", "mod", "init", "a")
		cmd.Dir = path.Join(".", dir)
		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("%s go mod init failed: %s", pkgName, err))
		}
	}
	return nil
}

// go get pkgName/...
// go vet pkgName/...
func execVet(pkgName, dir string) error {
	// TODO: 関数の外に切り出す
	defer func() {
		// remove dir
		if err := os.RemoveAll(dir); err != nil {
			log.Printf("remove dir %s failed: %s", dir, err)
		}
	}()

	arg := path.Join(pkgName, "...")
	cmd := exec.Command("go", "get", arg)
	cmd.Dir = path.Join(".", dir)
	if err := cmd.Run(); err != nil {
		fmt.Printf("go get %s failed: %s", pkgName, err)
		return err
	}
	cmd = exec.Command("go", "vet", "-vettool=/Users/idoharumare/intern/shortenexpression/shortenexpression/shortenexpression", arg)
	cmd.Dir = path.Join(".", dir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("vet output: %s", string(out))
	}

	return nil
}
