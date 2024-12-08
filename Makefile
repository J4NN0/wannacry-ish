# === CONFIG =======================================================
WANNACRY_ISH_PROJECT_NAME="wannacry-ish"
CRYPTO_KEY_PROJECT_NAME="crypto-key"

# === BUILD =======================================================
build-wannacry-ish:
	@echo "---> Building $(WANNACRY_ISH_PROJECT_NAME)"
	go build -o build/wannacry_ish cmd/$(WANNACRY_ISH_PROJECT_NAME)/main.go
.PHONY: build-wannacry-ish

build-crypto-key:
	@echo "---> Building $(CRYPTO_KEY_PROJECT_NAME)"
	go build -o build/crypto_key cmd/$(CRYPTO_KEY_PROJECT_NAME)/main.go
.PHONY: build-crypto-key

build: build-wannacry-ish build-crypto-key

# === RUN =======================================================
# Generate AES and RSA key pair
run-crypto-key:
	go run ./cmd/$(CRYPTO_KEY_PROJECT_NAME)/main.go
.PHONY: run-crypto-key

# Encrypt files
run-wannacry-ish-encrypt:
	go run ./cmd/$(WANNACRY_ISH_PROJECT_NAME)/main.go encrypt -d $(d)
.PHONY: run-wannacry-ish

# Decrypt files
run-wannacry-ish-decrypt:
	go run ./cmd/$(WANNACRY_ISH_PROJECT_NAME)/main.go decrypt -d $(d)
.PHONY: run-wannacry-ish
