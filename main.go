package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	apiKey    = flag.String("k", "", "Have I Been Pwned API key")
	hashMode  = flag.Bool("hash", false, "Check password hash (k-anon)")
	quiet     = flag.Bool("q", false, "Quiet mode")
	checkFile = flag.String("f", "", "Check passwords from file")
)

func main() {
	flag.Parse()

	args := flag.Args()

	if *checkFile != "" {
		checkFilePasswords(*checkFile)
		return
	}

	if len(args) == 0 {
		fmt.Println("PwnCheck - Password Breach Checker")
		fmt.Println("Usage: pwncheck [options] password")
		flag.Usage()
		return
	}

	for _, pwd := range args {
		checkPassword(pwd)
	}
}

func checkPassword(pwd string) {
	if *hashMode {
		checkHash(pwd)
		return
	}

	hash := sha256.Sum256([]byte(pwd))
	hashHex := hex.EncodeToString(hash[:])

	prefix := hashHex[:5]
	suffix := hashHex[5:]

	url := "https://api.pwnedpasswords.com/range/" + prefix
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	lines := strings.Split(string(body), "\r\n")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		if strings.EqualFold(parts[0], suffix) {
			count := parts[1]
			if !*quiet {
				fmt.Printf("⚠️  FOUND: '%s' appears %s times in breaches\n", pwd, count)
			} else {
				fmt.Printf("BREACHED:%s:%s\n", pwd, count)
			}
			return
		}
	}

	if !*quiet {
		fmt.Printf("✅ '%s' NOT found in breaches\n", pwd)
	}
}

func checkHash(pwd string) {
	hash := sha256.Sum256([]byte(pwd))
	fmt.Printf("SHA256: %x\n", hash[:])
}

func checkFilePasswords(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ Error reading file: %v\n", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		pwd := strings.TrimSpace(line)
		if len(pwd) > 0 {
			checkPassword(pwd)
		}
	}
}
