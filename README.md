# Ipfs_Encrypted_Distribution



https://user-images.githubusercontent.com/92402372/178027155-832d4687-45c1-4281-a963-9c218ae53c55.mov


![1](https://user-images.githubusercontent.com/92402372/179562775-b8d39439-2bb6-4c65-a7a5-d94a82eb9545.png)


![2](https://user-images.githubusercontent.com/92402372/179562790-1bf24833-8fb5-460e-8a6b-71e5f4a8bfe5.png)


# Gereksinimler / Requirements
1- Golang

2- VsCode

3- IPFS Dowloand url to Desktop

# main.go
Ayağa kaldırmış olduğunuz ipfs'e bağlanınız.(Ipfs standart Portudur.)

Your ipfs connection that you have raised. (Ipfs is the standard port.)

sh := shell.NewShell("localhost:5001")

- Işleme / Processing

Text işlemleri için / For text operations;



sh.Add(strings.NewReader("hello world selam"))

- Dosya işlemleri için(png,jpeg,GIF,File vb.) / For file operations(png,jpeg,GIF,File etc.);

sh.AddDir("./Qr.png")

# BLOCK
# Golang-AES-GCM
Golang ile kendi şifreleme algoritmamı yazmak istedik biraz araştırma yaptım ve aes kütüphanesini buldum https://pkg.go.dev/crypto/aes#NewCipher ve anahtarları kayıt etmemiz gerekiyor bu sıkıcı ve güvensiz bunun yerine 32 bitlik bir anahtar değişkeni yazdım sonra rastgele(rand) işlemine gönderdim sonrasında ise hexadecimale dönüştürdüm sonuç sürekli değişen bir key ve anahtarıda bende ;)

We wanted to write my own encryption algorithm with Golang I did some research and found the aes library https://pkg.go.dev/crypto/aes#NewCipher and we need to register the keys it's boring and insecure instead I wrote a 32 bit key variable then randomly(rand ) process, then I converted it to hexadecimal, the result is a constantly changing key and I have the key ;)

# şifreleme adımları / encryption steps
- Anahtar (32-bytes için AES-256 şifreleme) / key

- nonce (rastgele sayı)
 
- şifrelenecek veri / data to be encrypted

- aesGCM ile şifrele / encrypt with aesGCM

- şifre çözme / decoding

- Anahtar (32-bytes AES-256) / key

- nonce

- şifrelenmiş veriden nonce çıkartma / removing nonce from encrypted data

- aesGCM ile çözümle şifre çözmede nonce(anahtar tekrarlama riski için kullanıldı) değişkenine dikkat edin doğru nonce olmaz ise şifre çözülemez. / Pay attention to the nonce variable (used for key duplication risk) in decrypt with aes GCM resolve, if not the correct nonce cannot be decrypted.

# ++ Ornek / Example
# Dosya Şifreleme / File Encryption
projemiz de encryptFile fonksiyon ile dosya oluşturup şifreleyebiliriz.

In our project, we can create and encrypt files with the encryptFile function.

func encryptFile(filename string, data []byte, keyString string) { f, _ := os.Create(filename) defer f.Close() f.Write(encrypt(data, passphrase)) }

# Dosya Şifre Çözme / File Decryption
Projemiz ile Şifrelediğimiz dosyamızı decryptFile fonksiyon ile işlediğimiz key ile çözüyor.

It decrypts the file that we encrypted with our project with the key that we processed with the decrypt File function.

func decryptFile(filename string, keyString string) []byte { data, _ := ioutil.ReadFile(filename) return decrypt(data, passphrase) }
