---
name: pp-skool
description: "Printing Press CLI for Skool. Discovered API spec for skool"
author: "robj1925"
license: "Apache-2.0"
argument-hint: "<command> [args] | install cli|mcp"
allowed-tools: "Read Bash"
metadata:
  openclaw:
    requires:
      bins:
        - skool-pp-cli
---

# Skool — Printing Press CLI

## Prerequisites: Install the CLI

This skill drives the `skool-pp-cli` binary. **You must verify the CLI is installed before invoking any command from this skill.** If it is missing, install it first:

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
3. **Add to Path (Optional):**
   ```bash
   sudo mv skool-pp-cli /usr/local/bin/
   ```
4. **Verify:** `skool-pp-cli version`

Discovered API spec for skool

## HTTP Transport

This CLI uses standard HTTP transport with HTTP/2 disabled for browser-facing endpoints. It does not require a resident browser process for normal API calls.

## Command Reference

**.deploy_status_henson.json** — Operations on .deploy_status_henson.json

- `skool-pp-cli .deploy_status_henson.json` — GET /v3/.deploy_status_henson.json

**017ae153ccc5** — Operations on inputs

- `skool-pp-cli 017ae153ccc5 create_mp_verify` — POST /017ae153ccc5/4aa4380fa03e/mp_verify
- `skool-pp-cli 017ae153ccc5 create_telemetry` — POST /017ae153ccc5/4aa4380fa03e/telemetry
- `skool-pp-cli 017ae153ccc5 list_inputs` — GET /017ae153ccc5/4aa4380fa03e/inputs

**_next** — Operations on members.json

- `skool-pp-cli _next get_about.json` — GET /_next/data/{id}/ai-academy-with-robby-6849/about.json
- `skool-pp-cli _next get_calendar.json` — GET /_next/data/{id}/ai-academy-with-robby-6849/calendar.json
- `skool-pp-cli _next get_classroom.json` — GET /_next/data/{id}/ai-academy-with-robby-6849/classroom.json
- `skool-pp-cli _next get_leaderboards.json` — GET /_next/data/{id}/ai-academy-with-robby-6849/-/leaderboards.json
- `skool-pp-cli _next get_map.json` — GET /_next/data/{id}/ai-academy-with-robby-6849/-/map.json
- `skool-pp-cli _next get_members.json` — GET /_next/data/{id}/ai-academy-with-robby-6849/-/members.json

**affiliates** — Operations on compensations

- `skool-pp-cli affiliates list_compensations` — GET /affiliates/v2/compensations
- `skool-pp-cli affiliates list_payout` — GET /affiliates/payout
- `skool-pp-cli affiliates options_compensations` — OPTIONS /affiliates/v2/compensations

**b** — Operations on b

- `skool-pp-cli b` — POST /b

**default** — Operations on default

- `skool-pp-cli default <id>` — POST /{id}

**envelope** — Operations on envelope

- `skool-pp-cli envelope <id>` — POST /api/{id}/envelope/

**f** — Operations on f

- `skool-pp-cli f <hash>` — GET /f/{hash}/{hash}

**groups** — Operations on analytics-overview-v2

- `skool-pp-cli groups create_count_pending_invites` — POST /groups/{hash}/count-pending-invites
- `skool-pp-cli groups get_admin_metrics` — GET /groups/{hash}/admin-metrics
- `skool-pp-cli groups get_analytics_growth_overview_v2` — GET /groups/{hash}/analytics-growth-overview-v2
- `skool-pp-cli groups get_analytics_overview_v2` — GET /groups/{hash}/analytics-overview-v2
- `skool-pp-cli groups get_billing_payout_data` — GET /groups/{hash}/billing-payout-data
- `skool-pp-cli groups get_discovery` — GET /groups/{hash}/discovery
- `skool-pp-cli groups get_member_course_permissions` — GET /groups/{hash}/member-course-permissions

**maps** — Operations on style.json

- `skool-pp-cli maps list_sprite.json` — GET /maps/streets-v2/sprite.json
- `skool-pp-cli maps list_style.json` — GET /maps/streets-v2/style.json

**tiles** — Operations on tiles.json

- `skool-pp-cli tiles get_1.pbf` — GET /tiles/v3/{id}/{id}/1.pbf
- `skool-pp-cli tiles get_2.pbf` — GET /tiles/v3/{id}/{id}/2.pbf
- `skool-pp-cli tiles list_tiles.json` — GET /tiles/v3/tiles.json

**wait** — Operations on wait

- `skool-pp-cli wait` — GET /wait


### Finding the right command

When you know what you want to do but not which command does it, ask the CLI directly:

```bash
skool-pp-cli which "<capability in your own words>"
```

`which` resolves a natural-language capability query to the best matching command from this CLI's curated feature index. Exit code `0` means at least one match; exit code `2` means no confident match — fall back to `--help` or use a narrower query.

## Auth Setup
Skool uses session-based authentication. The CLI provides an automated login flow.

1. Run the login command:
   ```bash
   skool-pp-cli auth login
   ```
   **CRITICAL FOR AGENTS:** This command opens a GUI browser on the user's machine. **Do not attempt to automate the browser login yourself.** Run the command and immediately instruct the human user to complete the login in the window that opens.
   - **WARNING:** Do NOT run `node scripts/refresh_auth.js` directly. The Node script only extracts the token; the Go binary (`./skool-pp-cli auth login`) is required to actually save it to the config file.
   - If running the command in the background, use `command_status` and wait patiently (up to 120 seconds) for the human to finish. Do NOT attempt to run other commands in the meantime.
2. Run `skool-pp-cli doctor --agent` to verify setup.

## Agent Mode

Add `--agent` to any command. Expands to: `--json --compact --no-input --no-color --yes`.

- **Pipeable** — JSON on stdout, errors on stderr
- **Filterable** — `--select` keeps a subset of fields. Dotted paths descend into nested structures; arrays traverse element-wise. Critical for keeping context small on verbose APIs:

  ```bash
  skool-pp-cli .deploy_status_henson.json --canary-percentage 42 --deployed-revisions example-value --agent --select id,name,status
  ```
- **Previewable** — `--dry-run` shows the request without sending
- **Offline-friendly** — sync/search commands can use the local SQLite store when available
- **Non-interactive** — never prompts, every input is a flag
- **Explicit retries** — use `--idempotent` only when an already-existing create should count as success

### Response envelope

Commands that read from the local store or the API wrap output in a provenance envelope:

```json
{
  "meta": {"source": "live" | "local", "synced_at": "...", "reason": "..."},
  "results": <data>
}
```

Parse `.results` for data and `.meta.source` to know whether it's live or local. A human-readable `N results (live)` summary is printed to stderr only when stdout is a terminal — piped/agent consumers get pure JSON on stdout.

## Agent Feedback

When you (or the agent) notice something off about this CLI, record it:

```
skool-pp-cli feedback "the --since flag is inclusive but docs say exclusive"
skool-pp-cli feedback --stdin < notes.txt
skool-pp-cli feedback list --json --limit 10
```

Entries are stored locally at `~/.skool-pp-cli/feedback.jsonl`. They are never POSTed unless `SKOOL_FEEDBACK_ENDPOINT` is set AND either `--send` is passed or `SKOOL_FEEDBACK_AUTO_SEND=true`. Default behavior is local-only.

Write what *surprised* you, not a bug report. Short, specific, one line: that is the part that compounds.

## Output Delivery

Every command accepts `--deliver <sink>`. The output goes to the named sink in addition to (or instead of) stdout, so agents can route command results without hand-piping. Three sinks are supported:

| Sink | Effect |
|------|--------|
| `stdout` | Default; write to stdout only |
| `file:<path>` | Atomically write output to `<path>` (tmp + rename) |
| `webhook:<url>` | POST the output body to the URL (`application/json` or `application/x-ndjson` when `--compact`) |

Unknown schemes are refused with a structured error naming the supported set. Webhook failures return non-zero and log the URL + HTTP status on stderr.

## Named Profiles

A profile is a saved set of flag values, reused across invocations. Use it when a scheduled agent calls the same command every run with the same configuration - HeyGen's "Beacon" pattern.

```
skool-pp-cli profile save briefing --json
skool-pp-cli --profile briefing .deploy_status_henson.json --canary-percentage 42 --deployed-revisions example-value
skool-pp-cli profile list --json
skool-pp-cli profile show briefing
skool-pp-cli profile delete briefing --yes
```

Explicit flags always win over profile values; profile values win over defaults. `agent-context` lists all available profiles under `available_profiles` so introspecting agents discover them at runtime.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 2 | Usage error (wrong arguments) |
| 3 | Resource not found |
| 4 | Authentication required |
| 5 | API error (upstream issue) |
| 7 | Rate limited (wait and retry) |
| 10 | Config error |

## Argument Parsing

Parse `$ARGUMENTS`:

1. **Empty, `help`, or `--help`** → show `skool-pp-cli --help` output
2. **Starts with `install`** → ends with `mcp` → MCP installation; otherwise → see Prerequisites above
3. **Anything else** → Direct Use (execute as CLI command with `--agent`)

## MCP Server Installation

Install the MCP binary from this CLI's published public-library entry or pre-built release, then register it:

```bash
claude mcp add skool-pp-mcp -- skool-pp-mcp
```

Verify: `claude mcp list`

## Direct Use

1. Check if installed: `which skool-pp-cli`
   If not found, offer to install (see Prerequisites at the top of this skill).
2. Match the user query to the best command from the Unique Capabilities and Command Reference above.
3. Execute with the `--agent` flag:
   ```bash
   skool-pp-cli <command> [subcommand] [args] --agent
   ```
4. If ambiguous, drill into subcommand help: `skool-pp-cli <command> --help`.
