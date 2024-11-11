# GO Server Setup

## Download GO(if you don't have GO installed in your computer)
- Visit the official Go website and download the installer: https://go.dev/dl/


## Install GO
- macOS: Download the .pkg file and open it. Follow the on-screen instructions
- Windows: Run the .msi installer and follow the instructions

## Update PATH Environment Variable
- Windows:
1.Go to Control Panel > System and Security > System > Advanced System Settings > Environment Variables.
2.Find the Path variable in System Variables and click Edit.
3.Add C:\Go\bin to the list.

- macOS & Linux: Add this line to your shell profile (~/.bashrc, ~/.zshrc, or ~/.profile):
  ```bash
  export PATH=$PATH:/usr/local/go/bin
  ```
  Then reload the profile:
  ```bash
  source ~/.bashrc  # or source ~/.zshrc, depending on your shell
  ```
## Verify the Installation
To check if Go was installed correctly, open a terminal and run:
```bash
go version
```

## Setup Instructions
### 1. Create a Virtual Environment
Do this in the directory src/servers/go-server

#### On macOS/Linux/Windows
```sh
cd src/servers/go-server
```

### 2. Install Dependencies
The server requires the `Gin` package to run. You can install the dependencies using the following command:
```bash
go get -u github.com/gin-gonic/gin
```

### 3. Initialize Go Modules
```bash
go mod init <your-module-name>
```

### 4. Running the Server
Once the dependencies are installed and Go module is initialized, you can run the code using the following command:
```bash
go run serverimplementation.go
```

#### Additional Notes
The server will run on port 5004. 

