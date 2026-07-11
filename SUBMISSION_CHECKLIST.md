# Submission Checklist — The Injective Global Cup

**Deadline:** 2026-07-19 (remaining: ~8 days from 2026-07-11)  
**Project location:** `D:\WorkBuddy\InjectiveGlobalCup\`  
**Typeform submission link:** https://xsxo494365r.typeform.com/to/TMaGb1du  
**X account:** [@ljnhu73965779](https://x.com/ljnhu73965779)

---

## Phase 1: Account & repo setup (do this first)

- [ ] **X account verified accessible** — confirmed @ljnhu73965779 can log in and post.
- [ ] **Test post on X** — publish a simple tweet to confirm no phone-number lock or restriction.
- [ ] **Create public GitHub repository** — e.g., `https://github.com/YOUR_USERNAME/worldcup-injective`.
- [ ] **Push project code** to the repository (do NOT commit `.env`).
- [ ] **Update README** with real GitHub link and author info (already written).
- [ ] **Update X post draft** (`X_POST_DRAFT.md`) with the real repo link.

## Phase 2: Make it demoable (before 2026-07-17)

- [ ] **Fill `.env` with API key** (optional but recommended): `OPENAI_API_KEY`, `OPENAI_BASE_URL`, `OPENAI_MODEL`.
- [ ] **Run backend locally**: `cd backend && go run main.go`.
- [ ] **Open frontend**: `frontend/index.html` in browser.
- [ ] **Verify free prediction** works for at least one match.
- [ ] **Verify premium prediction** returns AI-style rationale (if API key set).
- [ ] **Record 3-minute demo video** showing:
  - Match list
  - Free prediction
  - Premium prediction flow (or x402 flow if you have a wallet with USDC)
  - Brief mention of Injective tech used (x402 / MCP / CCTP)
- [ ] **Upload demo video** to a public link (YouTube unlisted, Loom, Google Drive public, etc.).

## Phase 3: Optional on-chain polish (if time permits)

- [ ] Set up **Keplr or Leap wallet** on Injective.
- [ ] Get testnet INJ/USDC from Injective testnet faucet.
- [ ] Start the x402 gateway: `cd x402-gateway && npm install && npm start`.
- [ ] Test a real 0.01 USDC payment flow on Injective testnet.
- [ ] Capture the payment receipt / wallet screenshot for the demo.

> **Note:** On-chain interactions are *not* mandatory. The hackathon explicitly says "on-chain features are optional." Skip this phase if it risks missing the deadline.

## Phase 4: Final submission (2026-07-18 to 2026-07-19)

- [ ] **Post the X tweet** using the draft in `X_POST_DRAFT.md` (with real repo link).
- [ ] **Copy the tweet URL** after posting.
- [ ] **Open Typeform**: https://xsxo494365r.typeform.com/to/TMaGb1du.
- [ ] **Fill the form** with:
  - Project name: `Injective World Cup Predictor`
  - GitHub repo link
  - Demo video link
  - X tweet link
  - Description: mention x402, MCP server, CCTP, real LLM predictions
- [ ] **Submit before deadline** (preferably 2026-07-18 to avoid last-minute issues).
- [ ] **Screenshot the submission confirmation** for records.

## Phase 5: Post-submission (optional but valuable)

- [ ] Engage with the Injective / Ninja Labs community on X.
- [ ] Reply to your own tweet with the demo video if not included in the original post.
- [ ] Track HackQuest announcements for winner notification.

## Emergency contacts / links

- HackQuest event page: https://www.hackquest.io/zh-cn/hackathons/The-Injective-Global-Cup
- Typeform submission: https://xsxo494365r.typeform.com/to/TMaGb1du
- Injective docs: https://docs.injective.network/
- x402 docs: https://x402.org
- CCTP tutorial: https://docs.injective.network/developers-defi/usdc-cctp-tutorial

## Common blockers and fixes

| Problem | Fix |
|---|---|
| `go` not found in terminal | `D:\Software\Go\bin` is already in PATH; restart terminal if needed. |
| Port 8080 already in use | `netstat -ano \| findstr :8080` then `taskkill /F /PID <PID>`. |
| No real AI output | `OPENAI_API_KEY` missing or invalid; app falls back to heuristic model. |
| X tweet fails to post | Add/verify phone number on X; post a test tweet first. |
| `.env` pushed to GitHub | Remove it, add to `.gitignore`, rotate any exposed API key. |
