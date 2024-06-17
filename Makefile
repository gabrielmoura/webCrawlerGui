# Opções de compilação
LDFLAGS=-ldflags="-s -w"
WINLDFLAGS=-ldflags="-s -w -H=windowsgui"
BUILD_CMD=wails build
UPX_CMD=-upx

# Alvos principais
.PHONY: all build upx win winupx

all:
	@echo "Comandos disponíveis:"
	@echo "  make build   - Compila o binário"
	@echo "  make upx     - Compila e compacta o binário com UPX"
	@echo "  make win     - Compila o binário para Windows"
	@echo "  make winupx  - Compila e compacta o binário para Windows com UPX"

# Regra padrão: build
build:
	$(BUILD_CMD) $(LDFLAGS)

# Regra para build com UPX
upx:
	$(BUILD_CMD) $(LDFLAGS) $(UPX_CMD)

# Regra para build no Windows
win:
	GOOS=windows GOARCH=amd64 $(BUILD_CMD) $(WINLDFLAGS)

# Regra para build no Windows com UPX
winupx:
	GOOS=windows GOARCH=amd64 $(BUILD_CMD) $(WINLDFLAGS) $(UPX_CMD)
