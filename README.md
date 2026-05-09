# Skool CLI

Discovered API spec for skool

Learn more at [Skool](https://r.stripe.com).

## Install

The recommended path installs both the `skool-pp-cli` binary and the `pp-skool` agent skill in one shot:

```bash
npx -y @mvanhorn/printing-press install skool
```

For CLI only (no skill):

```bash
npx -y @mvanhorn/printing-press install skool --cli-only
```


### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/skool-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-skool --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-skool --force
```

## Install for OpenClaw

Tell your OpenClaw agent (copy this):

```
Install the pp-skool skill from https://github.com/mvanhorn/printing-press-library/tree/main/cli-skills/pp-skool. The skill defines how its required CLI can be installed.
```

## Quick Start

### 1. Install

See [Install](#install) above.

### 2. Set Up Credentials

Get your API key from your API provider's developer portal. The key typically looks like a long alphanumeric string.

```bash
export SKOOL_API_KEY="<paste-your-key>"
```

You can also persist this in your config file at `~/.config/skool-pp-cli/config.toml`.

### 3. Verify Setup

```bash
skool-pp-cli doctor
```

This checks your configuration and credentials.

### 4. Try Your First Command

```bash
skool-pp-cli .deploy_status_henson.json --canary-percentage 42 --deployed-revisions example-value
```

## Usage

Run `skool-pp-cli --help` for the full command reference and flag list.

## Commands

### .deploy_status_henson.json

Operations on .deploy_status_henson.json

- **`skool-pp-cli .deploy_status_henson.json list_.deploy_status_henson.json`** - GET /v3/.deploy_status_henson.json

### 017ae153ccc5

Operations on inputs

- **`skool-pp-cli 017ae153ccc5 create_mp_verify`** - POST /017ae153ccc5/4aa4380fa03e/mp_verify
- **`skool-pp-cli 017ae153ccc5 create_telemetry`** - POST /017ae153ccc5/4aa4380fa03e/telemetry
- **`skool-pp-cli 017ae153ccc5 list_inputs`** - GET /017ae153ccc5/4aa4380fa03e/inputs

### _next

Operations on members.json

- **`skool-pp-cli _next get_about.json`** - GET /_next/data/{id}/ai-academy-with-robby-6849/about.json
- **`skool-pp-cli _next get_calendar.json`** - GET /_next/data/{id}/ai-academy-with-robby-6849/calendar.json
- **`skool-pp-cli _next get_classroom.json`** - GET /_next/data/{id}/ai-academy-with-robby-6849/classroom.json
- **`skool-pp-cli _next get_leaderboards.json`** - GET /_next/data/{id}/ai-academy-with-robby-6849/-/leaderboards.json
- **`skool-pp-cli _next get_map.json`** - GET /_next/data/{id}/ai-academy-with-robby-6849/-/map.json
- **`skool-pp-cli _next get_members.json`** - GET /_next/data/{id}/ai-academy-with-robby-6849/-/members.json

### affiliates

Operations on compensations

- **`skool-pp-cli affiliates list_compensations`** - GET /affiliates/v2/compensations
- **`skool-pp-cli affiliates list_payout`** - GET /affiliates/payout
- **`skool-pp-cli affiliates options_compensations`** - OPTIONS /affiliates/v2/compensations

### b

Operations on b

- **`skool-pp-cli b create_b`** - POST /b

### default

Operations on default

- **`skool-pp-cli default create_endpoint`** - POST /{id}

### envelope

Operations on envelope

- **`skool-pp-cli envelope create_envelope`** - POST /api/{id}/envelope/

### f

Operations on f

- **`skool-pp-cli f get_f`** - GET /f/{hash}/{hash}

### groups

Operations on analytics-overview-v2

- **`skool-pp-cli groups create_count_pending_invites`** - POST /groups/{hash}/count-pending-invites
- **`skool-pp-cli groups get_admin_metrics`** - GET /groups/{hash}/admin-metrics
- **`skool-pp-cli groups get_analytics_growth_overview_v2`** - GET /groups/{hash}/analytics-growth-overview-v2
- **`skool-pp-cli groups get_analytics_overview_v2`** - GET /groups/{hash}/analytics-overview-v2
- **`skool-pp-cli groups get_billing_payout_data`** - GET /groups/{hash}/billing-payout-data
- **`skool-pp-cli groups get_discovery`** - GET /groups/{hash}/discovery
- **`skool-pp-cli groups get_member_course_permissions`** - GET /groups/{hash}/member-course-permissions

### maps

Operations on style.json

- **`skool-pp-cli maps list_sprite.json`** - GET /maps/streets-v2/sprite.json
- **`skool-pp-cli maps list_style.json`** - GET /maps/streets-v2/style.json

### tiles

Operations on tiles.json

- **`skool-pp-cli tiles get_1.pbf`** - GET /tiles/v3/{id}/{id}/1.pbf
- **`skool-pp-cli tiles get_2.pbf`** - GET /tiles/v3/{id}/{id}/2.pbf
- **`skool-pp-cli tiles list_tiles.json`** - GET /tiles/v3/tiles.json

### wait

Operations on wait

- **`skool-pp-cli wait list_wait`** - GET /wait


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
skool-pp-cli .deploy_status_henson.json --canary-percentage 42 --deployed-revisions example-value

# JSON for scripting and agents
skool-pp-cli .deploy_status_henson.json --canary-percentage 42 --deployed-revisions example-value --json

# Filter to specific fields
skool-pp-cli .deploy_status_henson.json --canary-percentage 42 --deployed-revisions example-value --json --select id,name,status

# Dry run — show the request without sending
skool-pp-cli .deploy_status_henson.json --canary-percentage 42 --deployed-revisions example-value --dry-run

# Agent mode — JSON + compact + no prompts in one flag
skool-pp-cli .deploy_status_henson.json --canary-percentage 42 --deployed-revisions example-value --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

## Use with Claude Code

Install the focused skill — it auto-installs the CLI on first invocation:

```bash
npx skills add mvanhorn/printing-press-library/cli-skills/pp-skool -g
```

Then invoke `/pp-skool <query>` in Claude Code. The skill is the most efficient path — Claude Code drives the CLI directly without an MCP server in the middle.

<details>
<summary>Use as an MCP server in Claude Code (advanced)</summary>

If you'd rather register this CLI as an MCP server in Claude Code, install the MCP binary first:


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Then register it:

```bash
claude mcp add skool skool-pp-mcp -e SKOOL_API_KEY=<your-key>
```

</details>

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/skool-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.
3. Fill in `SKOOL_API_KEY` when Claude Desktop prompts you.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "skool": {
      "command": "skool-pp-mcp",
      "env": {
        "SKOOL_API_KEY": "<your-key>"
      }
    }
  }
}
```

</details>

## Health Check

```bash
skool-pp-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Config file: `~/.config/skool-pp-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `SKOOL_API_KEY` | per_call | Yes | Set to your API credential. |

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `skool-pp-cli doctor` to check credentials
- Verify the environment variable is set: `echo $SKOOL_API_KEY`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

## HTTP Transport

This CLI uses standard HTTP transport with HTTP/2 disabled for browser-facing endpoints. It does not require a resident browser process for normal API calls.

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)
