# PwnCheck | Password Breach Checker

![Security Tool](https://img.shields.io/badge/Purpose-Breach%20Checker-green?style=for-the-badge)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

---

## Check If Passwords Have Been Pwned

Check passwords and email addresses against the Have I Been Pwned database to see if they've appeared in data breaches.

**Why PwnCheck?**
- Privacy-first: Uses k-Anonymity model
- Never sends full password to any server
- Fast batch checking
- CLI-first design

---

## How It Works

PwnCheck uses the Have I Been Pwned Passwords API with k-Anonymity:
1. Hash your password (SHA-1)
2. Send only the first 5 characters to the API
3. API returns all hashes starting with those characters
4. Check locally if your full hash is in the response

**Your password never leaves your machine in plain text.**

---

## Features

- 🔐 **Privacy-first** - k-Anonymity implementation
- 📧 **Email checking** - Check if emails were in breaches
- 📊 **Password auditing** - Check company password policies
- 🔄 **Batch mode** - Check multiple passwords at once
- ⚡ **Fast** - Concurrent checking

---

## Installation

```bash
git clone https://github.com/simplestar-992/pwncheck.git
cd pwncheck
go build -o pwncheck -ldflags="-s -w"
```

---

## Usage

```bash
# Check a password
./pwncheck password "mypassword123"

# Check an email
./pwncheck email "user@example.com"

# Check password list
./pwncheck passwords -f passwords.txt

# Check with quiet mode (exit 0 if safe, 1 if pwned)
./pwncheck password "mypassword" -q
```

---

## Examples

```bash
# In scripts
if ./pwncheck password "$USER_PASS" -q; then
    echo "Password is safe"
else
    echo "⚠️ Password found in breach database!"
fi

# Password policy enforcement
./pwncheck passwords -f /etc/shadow -v
```

---

## License

MIT © 2024 [simplestar-992](https://github.com/simplestar-992)
