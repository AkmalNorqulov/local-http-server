
# 🌐 LAN Web Server with DNS Redirect

This project is a simple Go server that:  
✅ Serves static websites from a `dist/` folder to all devices on your local network.  
✅ Redirects requests for `ddd.uz` to your PC’s IP address.  
✅ Forwards all other DNS queries to Google DNS (`8.8.8.8`).  

It allows devices connected to your router (phones, laptops, TVs, etc.) to access your site by visiting `http://ddd.uz`.

---

## ⚙️ How It Works
- Starts a web server on your PC (port `80`) for files in the `dist/` folder.  
- Starts a DNS server (port `53`) that:  
  - Redirects `ddd.uz` → your PC’s IP (e.g., `192.168.100.191`).  
  - Forwards all other domains to `8.8.8.8`.  

✅ Works for all devices on your LAN.  
✅ Stops and restores normal DNS resolution when you exit the program.

---

## 🖥 Setup

### 1️⃣ Set your PC’s LAN IP
Make sure your PC has a **static LAN IP** (e.g., `192.168.100.191`).  

### 2️⃣ Set your router’s DNS server
- Log in to your router’s admin panel (often at `192.168.0.1` or `192.168.100.1`).  
- Set:  
  - **Primary DNS** = your PC’s IP (e.g., `192.168.100.191`).  
  - **Secondary DNS** = `8.8.8.8` (fallback).  
- Save and reboot the router.  

### 3️⃣ Prepare your website
- Place all your static website files (HTML, CSS, JS, etc.) inside a folder named `dist`.  
- Your project folder should look like this:  

```
project/
├── dist/
│   ├── index.html
│   ├── assets/
├── main.go
```

---

## 🚀 Build and Run

### 🛠 Build the server

#### 🔵 Windows:
Open PowerShell or CMD:
```bash
go build -o server.exe main.go
```

#### 🟢 Linux/macOS:
Open Terminal:
```bash
go build -o server main.go
```

---

### ▶ Run the server

#### 🔵 Windows:
Run as Administrator (right-click → *Run as administrator*):  
```bash
server.exe
```

#### 🟢 Linux/macOS:
Run with `sudo` (required for ports 80 and 53):  
```bash
sudo ./server
```

---

When running, the console will show:  
```
[AIO] Starting up...
[DNS] Intercepting ddd.uz on port 53
[HTTP] Serving files from ./dist at http://192.168.100.191
```

---

## 🌐 Access the Web App
On any device connected to your router:  
```
http://ddd.uz
```

✅ Works in all browsers.  

---

## 🛑 Stopping
- Press `Ctrl+C` in the console window or close it.  
- The server and DNS will stop.  
- Your router will automatically use the secondary DNS (e.g., `8.8.8.8`) for internet access.  

---

## ⚠️ Notes
- This app requires **administrator/root permissions** to bind to ports `80` (HTTP) and `53` (DNS).  
- Works only within your local network (LAN).  
- Make sure no other application is using port 80 or 53 on your PC.  
