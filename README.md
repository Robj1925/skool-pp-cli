# Skool Community CLI (`skool-pp-cli`)

A high-performance command-line interface for Skool community owners and power users. This tool allows you to bypass the browser for community management, perform deep domain analytics, and gain behavioral insights into your members.

<p align="center">
  <img src="assets/logo.png" alt="Skool Community CLI Logo" width="200">
</p>

## 🚀 Key Features

- **Full Data Sync:** Synchronize your entire community (members, posts, etc.) to a local SQLite database.
- **Domain Analytics (Rung 4):** Identify "orphaned" members (joined but never posted), stale members, and top contributors.
- **Behavioral Insights (Rung 5):** Predictive modeling for churn risk and member growth velocity.
- **Cloudflare Optimized:** Built-in support for HTTP/2 and modern browser headers to ensure reliable API connectivity.

---

## 🛠 Installation

### From Source (Recommended)

Ensure you have [Go](https://go.dev/dl/) installed (v1.21+) and [Node.js](https://nodejs.org/).

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

## 🤖 Instructions for AI Agents
If you are an AI Agent (e.g., Claude Code, Codex, Antigravity) instructed to "install this CLI tool", follow these exact steps:

1. **Install Dependencies:** If `go` or `npm` are missing, install them first.
   - On macOS: 
     - If `brew` is not installed, install it first: `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
     - Then run: `brew install go node`
   - On Linux (Debian/Ubuntu): `sudo apt-get update && sudo apt-get install -y golang nodejs npm`
   - If you cannot install them, explicitly ask the human to install Go (v1.21+) and Node.js.
2. **Clone and Build:** Run the following one-liner to clone, install dependencies, and build the binary in the current directory:
   ```bash
   git clone https://github.com/Robj1925/skool-pp-cli.git && cd skool-pp-cli && npm install && npx playwright install chromium && go build -o skool-pp-cli ./cmd/skool-pp-cli
   ```
3. **Handle Authentication (CRITICAL):** The automated browser login (`skool-pp-cli auth login`) can be unstable. **Always instruct the human user to authenticate manually:**
   - Ask the human to open their browser, log in to Skool, and open Developer Tools (Network tab).
   - Tell them to find a request (like `members.json`) and copy the entire value of the `cookie:` header.
   - Instruct the human to provide you with the cookie string, or tell them to run `./skool-pp-cli auth set-token "YOUR_COOKIE_STRING"` themselves.
4. **Verify:** Once the session token is saved, run `./skool-pp-cli doctor --agent` to verify connectivity.

---

## 🔑 Setup & Authentication

Skool uses session-based authentication. The CLI provides a "One-Click" automated login flow to capture your credentials securely.

### 1. Manual Login (Recommended)
Type the following command to get step-by-step instructions for extracting your session cookie:

```bash
skool-pp-cli auth refresh
```

This will guide you on how to find your `cookie:` string from your browser's Developer Tools. Once you have copied the cookie string, save it to the CLI config by running:

```bash
skool-pp-cli auth set-token "YOUR_COOKIE_STRING"
```

*(Note: The automated `skool-pp-cli auth login` command is still available but relies on Playwright and can be unstable on some environments. The manual method is guaranteed to work 100% of the time.)*

### 2. Verify Connection
Run the "doctor" command to ensure everything is working:
```bash
skool-pp-cli doctor
```

---

## 📖 Basic Workflow

### 💬 Direct Messaging
You can now list your conversations and send messages directly from the CLI:

1. **List your chat history:**
   ```bash
   skool-pp-cli channels list
   ```
   This will output a list of your recent conversations with their unique `channel_id`.

2. **Send a message:**
   ```bash
   skool-pp-cli channels send <CHANNEL_ID> --content "Hello from the CLI!"
   ```

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
| `posts create` | Create a new post in a group. |
| `sync` | Fetches all data from Skool and stores it locally. |
| `analytics-domain` | Advanced reports on member status and contributions. |
| `insights` | Predictive analytics (Churn, Velocity, Health). |
| `auth set-token` | Configures your session cookie. |
| `doctor` | Connectivity and auth health check. |
| `search` | Instant full-text search across your local community data. |

### ✍️ Writing Content

You can create posts in any group where you have permission. You need the group's unique hash (found in the URL or via `skool-pp-cli groups list`).

```bash
# Create a post (defaults to 'General discussion' category)
skool-pp-cli posts create <group-hash> \
  --title "My New Post" \
  --content "This is the body of the post. Supports markdown!"

# Create a post in a specific category
skool-pp-cli posts create <group-hash> \
  --title "Q&A" \
  --content "Ask me anything!" \
  --category-id <category-id>
```

**Note:** The Skool API requires a category ID for every post. If `--category-id` is omitted, the CLI defaults to the "General discussion" category.

---

## ❓ Troubleshooting

**"No such table: members"**
This usually means you haven't run `skool-pp-cli sync` yet, or the sync failed. Run `skool-pp-cli sync` and check for errors.

**"Cloudflare / HTTP 403"**
If you are getting blocked, ensure you are using the latest version of the CLI. We've implemented HTTP/2 support to align with Skool's production environment.

**"Missing config file"**
The CLI expects a config file at `~/.config/skool-pp-cli/config.toml`. Running `skool-pp-cli auth login` will create this for you automatically.

---

## 📜 License

Licensed under Apache-2.0. See [LICENSE](LICENSE) for details.

Developed with ❤️ for the Skool Community.
