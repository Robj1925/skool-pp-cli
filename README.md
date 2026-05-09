# Skool Community CLI (`skool-pp-cli`)

A high-performance command-line interface for Skool community owners and power users. This tool allows you to bypass the browser for community management, perform deep domain analytics, and gain behavioral insights into your members.

## 🚀 Key Features

- **Full Data Sync:** Synchronize your entire community (members, posts, etc.) to a local SQLite database.
- **Domain Analytics (Rung 4):** Identify "orphaned" members (joined but never posted), stale members, and top contributors.
- **Behavioral Insights (Rung 5):** Predictive modeling for churn risk and member growth velocity.
- **Cloudflare Optimized:** Built-in support for HTTP/2 and modern browser headers to ensure reliable API connectivity.

---

## 🛠 Installation

### From Source (Recommended)

Ensure you have [Go](https://go.dev/dl/) installed (v1.21+).

1. Clone the repository:
   ```bash
   git clone https://github.com/Robj1925/skool-pp-cli.git
   cd skool-pp-cli
   ```

2. Build the binary:
   ```bash
   go build -o skool-pp-cli ./cmd/skool-pp-cli
   ```

3. (Optional) Move it to your path:
   ```bash
   sudo mv skool-pp-cli /usr/local/bin/
   ```

---

## 🔑 Setup & Authentication

Skool uses a session-based cookie for authentication. You will need to "sniff" your session cookie from your browser.

### 1. Capture your Session Cookie
1. Log in to [Skool.com](https://www.skool.com) in your browser (Chrome/Edge/Brave).
2. Right-click anywhere and select **Inspect**.
3. Go to the **Application** tab (or **Storage** in Firefox).
4. In the left sidebar, expand **Cookies** and select `https://www.skool.com`.
5. Find the row where the Name is **NOT** `auth_token` but rather the entire string containing your session info. 
   > **Tip:** You can also just go to the **Network** tab, refresh the page, click on any request to `www.skool.com`, and copy the entire value of the `Cookie` header from the **Request Headers** section.

### 2. Set the Token in CLI
Copy the entire cookie string and run:
```bash
skool-pp-cli auth set-token "client_id=...; auth_token=...; etc..."
```

### 3. Verify Connection
Run the "doctor" command to ensure everything is working:
```bash
skool-pp-cli doctor
```

---

## 📖 Basic Workflow

### 1. Synchronize Data
Before running analytics, you need a local copy of your data:
```bash
skool-pp-cli sync
```
This will create a SQLite database at `~/.config/skool-pp-cli/skool.db`.

### 2. Run Domain Analytics
Find "orphaned" members who joined but haven't contributed yet:
```bash
skool-pp-cli analytics-domain orphans
```

Find "stale" members who haven't been active in 30 days:
```bash
skool-pp-cli analytics-domain stale --days 30
```

### 3. Get Behavioral Insights
Predict which members are at risk of churning:
```bash
skool-pp-cli insights churn-risk
```

Identify members who are leveling up at high velocity:
```bash
skool-pp-cli insights level-velocity
```

---

## 🖥️ Command Reference

| Command | Description |
| --- | --- |
| `sync` | Fetches all data from Skool and stores it locally. |
| `analytics-domain` | Advanced reports on member status and contributions. |
| `insights` | Predictive analytics (Churn, Velocity, Health). |
| `auth set-token` | Configures your session cookie. |
| `doctor` | Connectivity and auth health check. |
| `search` | Instant full-text search across your local community data. |

---

## ❓ Troubleshooting

**"No such table: members"**
This usually means you haven't run `skool-pp-cli sync` yet, or the sync failed. Run `skool-pp-cli sync` and check for errors.

**"Cloudflare / HTTP 403"**
If you are getting blocked, ensure you are using the latest version of the CLI. We've implemented HTTP/2 support to align with Skool's production environment.

**"Missing config file"**
The CLI expects a config file at `~/.config/skool-pp-cli/config.toml`. Running `auth set-token` will create this for you automatically.

---

## 📜 License

Licensed under Apache-2.0. See [LICENSE](LICENSE) for details.

Developed with ❤️ for the Skool Community.
