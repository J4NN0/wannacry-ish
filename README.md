# WannaCry-ish

This repository provides a program designed as an educational resource to understand how ransomware works under the hood. It does not exploit any vulnerabilities, but executing ransomware typically requires one in real-world scenarios. A famous ransomware attack is [WannaCry](https://en.wikipedia.org/wiki/WannaCry_ransomware_attack), which used a Microsoft vulnerability to encrypt data and demand ransom payments in Bitcoin. 

### Disclaimer 

This program is for educational purposes only. It is intended to demonstrate the cryptographic principles and workflows commonly employed by ransomware. It should never be used maliciously or to harm others. Misusing this tool for malicious purposes is illegal and unethical. Unauthorized use of this tool may violate laws and result in severe legal consequences. Always act responsibly and ethically. The authors do not endorse or condone any unauthorized or illegal use of this software.

# How does ransomware work?

Three types of keys are involved, each playing a critical role in the cryptographic workflow:
1. AES Key
   - A symmetric key used for encrypting and decrypting files within the target directory.
   - Fast and efficient for large-scale encryption of data.
2. RSA Key Pair
   - Public Key: It is used to encrypt the AES key after its usage.
   - Private Key (kept safe by the attacker): Used to decrypt the AES key in case the ransom is paid.

Key workflow:
- Encryption
  1. A random AES key is generated for file encryption.
  2. The AES key is used to encrypt the contents of each file in the target directory.
  3. The original files are replaced with their encrypted versions.
  4. The AES key is encrypted with the RSA public key preventing the victim from decrypting files.
  5. Note that the RSA private key is stored locally for demonstration purposes.
- Decryption
  1. The AES key is decrypted with the RSA private key.
  2. The decrypted AES key is used to restore the original file contents.

> Why Keep the Private Key Secret?

In a real-world scenario, the private RSA key would never be stored with the encrypted files. The attacker would keep the private key securely to prevent the victim from decrypting the AES key, and thus their files, without paying a ransom. In this educational tool, the private key is stored locally (alongside the encrypted AES key) for demonstration purposes only.

# Usage

### Generate Keys

Before running the encryption or decryption functionality, we need to generate the required keys:

    make run-crypto-key

Two files will be created locally in the working directory
1.	The AES key `encrypted_aes_key.bin` (encrypted using the RSA public key to protect it), which will be used to encrypt files.
2.	The RSA private key `rsa_private_key.pem`, which will be used to decrypt the encrypted AES key.
3. Note that the RSA public is not stored locally as it can be easily derived from the private key.

### Encrypt or Decrypt Files

Once the keys are generated, we can encrypt or decrypt all files within a specified directory.

- Encryption
  
      make run-wannacry-ish-encrypt d=DIR_PATH

   When files are encrypted, the `.cry` extension is appended to each encrypted file. If encryption is run in the sample directory `data`, then you should see

      omfg.txt.cry
      sample.txt.cry

   And content should look like this

      OQ<C6><96>c/(<CC><E6><F5>6^R<)<A7><D7><D1><C0>&<BA><9D><9B>^MQ<CD>+^M<94>Co^TS#^L<95>|iS<C8>i\m^S<A6><FB>i<B8>

- Decryption

      make run-wannacry-ish-decrypt d=DIR_PATH

   When files are decrypted, the `.cry` extension is removed from each file and their content should be readable again.
