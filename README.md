# ⚽ Injective World Cup Predictor

A lightweight World Cup prediction dApp built for **The Injective Global Cup** hackathon. It demonstrates how Injective's newest primitives can power everyday consumer apps — even for users who don't know what a smart contract is.

## What it does

Fans can browse the live 2026 World Cup knockout bracket, get free heuristic predictions, and unlock a **premium AI prediction** for any match by paying **0.01 USDC** via the x402 protocol. The project also ships a **Go MCP server** so Claude/Cursor/other AI agents can query fixtures and predictions through native tools.

## Injective technologies demonstrated

| Technology | How it is used in this project | Link |
|---|---|---|
| **x402** | Pay-per-prediction gateway. Premium AI insights are gated behind a 0.01 USDC x402 payment, settled on Injective EVM with sub-second finality. | [x402](https://x402.org) |
| **MCP Server / Agent Skills** | The Go backend exposes `list_worldcup_matches` and `predict_match` as MCP tools, so AI agents can consume World Cup data without custom API integration. | [mcp-go](https://github.com/mark3labs/mcp-go) |
| **CCTP** | Documented as the cross-chain USDC reward rail for future prediction-pool prize distribution. | [Injective CCTP tutorial](https://docs.injective.network/developers-defi/usdc-cctp-tutorial) |
| **Injective EVM** | x402 payments settle on Injective EVM, which inherits the speed and low fees of the Injective chain. | [Injective EVM](https://docs.injective.network/) |

> **Note:** On-chain interactions are encouraged but not mandatory for this hackathon. The app works end-to-end as a Web2+AI app and can settle on Injective testnet/mainnet when the user turns on x402.

## Live demo

- **Frontend:** open `frontend/index.html` directly in a browser, or serve it with any static server.
- **Backend API:** `http://localhost:8080`
- **x402 Gateway:** `http://localhost:3000`
- **Demo video:** *(record and replace this link)*

## Project layout

```
.
├── backend/              # Go backend + MCP server
│   ├── internal/api      # HTTP handlers
│   ├── internal/matches  # Real World Cup match data
│   ├── internal/predictions  # AI prediction engine
│   └── internal/mcp      # Injective MCP server
├── x402-gateway/         # Node/Express x402 paywall gateway
└── frontend/             # Plain HTML/JS UI
```

## Quick start

### 1. Backend (Go)

```bash
cd backend
cp ../.env.example ../.env
# Edit ../.env and optionally fill OPENAI_API_KEY for real AI predictions
go mod tidy
go run main.go
```

Backend runs on `http://localhost:8080`.

### 2. x402 Gateway (Node)

Requires Node.js 20+.

```bash
cd x402-gateway
npm install
npm start
```

Gateway runs on `http://localhost:3000`.

### 3. Frontend

Open `frontend/index.html` in a browser, or serve it:

```bash
cd frontend
python -m http.server 5500
# visit http://localhost:5500
```

## Real AI predictions

Premium predictions are powered by a **real LLM** (OpenAI-compatible). The free tier stays on a fast heuristic model so it costs nothing.

Configure in `.env`:

```bash
OPENAI_API_KEY=sk-...          # required for real AI
OPENAI_BASE_URL=https://api.openai.com/v1   # or DeepSeek/Claude-compatible endpoint
OPENAI_MODEL=gpt-4o-mini       # any chat model
```

- If `OPENAI_API_KEY` is set, `/api/premium-predict/:id` (and the MCP `predict_match` tool) calls the LLM and returns a grounded prediction with a natural-language rationale.
- If unset or the call fails, it **falls back** to the heuristic model. The app never crashes.

The model is asked for strict JSON: `winner`, `score_a`, `score_b`, `confidence` (0–1), `rationale`.

## API

| Endpoint | Description | x402 gated? |
|---|---|---|
| `GET /api/health` | Health check | No |
| `GET /api/matches` | List all matches | No |
| `GET /api/matches/:id` | Match details | No |
| `GET /api/predict/:id` | Free heuristic prediction | No |
| `GET /api/premium-predict/:id` | **Real LLM** prediction (falls back to heuristic if no key) | **Yes** (0.01 USDC) |

## MCP Server

The backend also exposes an MCP server over stdio. Tools:

- `list_worldcup_matches` – returns all matches.
- `predict_match(match_id)` – returns a premium prediction.

To use with Claude Desktop, add to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "worldcup-injective": {
      "command": "D:\\WorkBuddy\\InjectiveGlobalCup\\backend\\worldcup-injective.exe",
      "args": []
    }
  }
}
```

(Or run `go run main.go` and point Claude to the built binary.)

## CCTP roadmap

Future iterations will distribute prediction-pool winnings to top predictors using **Circle CCTP** to send native USDC cross-chain (e.g., from Ethereum Sepolia to Injective testnet). Contract references and testnet faucets are documented in the [Injective CCTP tutorial](https://docs.injective.network/developers-defi/usdc-cctp-tutorial).

## Technologies used

- Go + chi router
- `@injectivelabs/x402` (Node)
- `mcp-go` (Go MCP SDK)
- Plain HTML/JS frontend

## Submission

Built for **The Injective Global Cup** hackathon on HackQuest.  
See [`SUBMISSION_CHECKLIST.md`](SUBMISSION_CHECKLIST.md) for the full submission checklist and timeline.

## Author

Built by **陆俊华** ([@ljnhu73965779](https://x.com/ljnhu73965779)) for The Injective Global Cup.
