package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"main/block"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// ipfs in çalışması gerekiyor ki ona bağlansın.
	sh := shell.NewShell("localhost:5001")
	//cid, err := sh.Add(strings.NewReader("hello world selam")) // yazı için

	cid, err := sh.AddDir("./Qr.png") //dosya için

	//Güvenliği arttırmak için anahtarı rastgele döngü halinde değiştirmek.

	bytes := make([]byte, 32) //AES-256 için rastgele bir 32 bayt anahtar oluşturun.
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	fmt.Printf("\n")
	key := hex.EncodeToString(bytes) //anahtarı bayt cinsinden kodlayın ve gizli olarak saklayın, bir kasaya koyun
	fmt.Printf("key(anahtar) : %s\n", key)

	encrypted := block.Encrypt(cid, key)
	fmt.Printf("encrypted(şifreli) : %s\n", encrypted)

	decrypted := block.Decrypt(encrypted, key)
	fmt.Printf("decrypted(şifre çözüm) : %s\n", decrypted)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("https://ipfs.io/ipfs/%s", decrypted)
}
