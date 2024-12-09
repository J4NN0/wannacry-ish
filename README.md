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
- The **AES Key**, a random symmetric key used for file encryption.
- Yhe **RSA Key Pair**, where the RSA public key is used to encrypt the AES key and the private one is kept secret (only used if the victim pays the ransom).

      make run-crypto-key

### Encrypt or Decrypt Files

Once the keys are generated, we can encrypt or decrypt all files within a specified directory.

- Encryption
  
      make run-wannacry-ish-encrypt d=DIR_PATH

- Decryption

      make run-wannacry-ish-decrypt d=DIR_PATH
