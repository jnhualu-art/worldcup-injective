# 🎬 Demo 录制口播脚本（3 分钟）

> 目标：录一段能说服评委"这项目真的用上了 Injective 技术"的视频。
> 评分核心 = 你**口播**点名用了 **x402 / MCP / CCTP** 这三项。

---

## 录制前准备

开两个终端：

```powershell
# 终端 1：Go 后端（必须）
cd D:\WorkBuddy\InjectiveGlobalCup\backend
go run main.go

# 终端 2：x402 网关（想展示付费链路才需要）
cd D:\WorkBuddy\InjectiveGlobalCup\x402-gateway
npm install
npm start
```

浏览器打开 `frontend/index.html`（建议 1080p 录屏，Win+G 游戏栏或 OBS）。

---

## 分镜（总时长 ~3:00）

### 镜头 1 · 开场（0:00–0:25）
**画面**：首页标题 "Injective World Cup Predictor" 清晰可见。
**中文旁白**：
> 大家好，我是陆俊华。我为 Injective Global Cup 黑客松做了一个世界杯 AI 预测应用——Injective World Cup Predictor。它把 Injective 最新的 x402 支付、MCP Server 和 CCTP 跨链技术，用在一个人人都懂的场景里：世界杯预测。
**英文（可选）**：
> Hi, I'm Junhua. I built the Injective World Cup Predictor for the Injective Global Cup — showing how x402, MCP and CCTP power a real consumer app.

### 镜头 2 · 真实比赛数据（0:25–0:55）
**画面**：滚动比赛列表，展示「法国 2-0 摩洛哥」「西班牙 2-1 比利时」等已完赛比分，以及半决赛「法国 vs 西班牙（7/15）」。
**中文旁白**：
> 这是真实的 2026 世界杯淘汰赛数据。已完赛的显示比分，待踢的显示对阵和日期，全部由后端 API 实时提供。
**英文**：
> Real 2026 World Cup knockout data — finished matches show scores, upcoming ones show fixtures, served live by the backend API.

### 镜头 3 · 免费预测（0:55–1:30）
**画面**：点一场「法国 vs 西班牙」→ 展示预测结果（winner / 比分 / 置信度 / 简短理由）。
**中文旁白**：
> 每场都有免费预测：基于球队实力和赛程的启发式模型，零成本。比如这场半决赛，模型给出法国胜、置信度 54%。

### 镜头 4 · 付费 AI 预测 + x402（1:30–2:15）
**画面 A（没配钱包也能演示）**：打开开发者工具 Network，点付费预测 → 请求返回 `402 Payment Required`，响应体含 `x402` 支付要求（Network: Injective EVM，0.01 USDC）。
**画面 B（配了钱包）**：展示钱包完成 0.01 USDC 支付后，返回 LLM 详细分析。
**中文旁白**：
> 想要更深入的分析，可以付 0.01 USDC 解锁 AI 预测——这里用的是 Injective 的 **x402 协议**，支付直接在链上结算、亚秒级确认。付费后调用真实 LLM，给出带自然语言推理的完整分析。
**英文**：
> Unlock a premium AI prediction by paying 0.01 USDC via x402, settled on Injective EVM with sub-second finality.

### 镜头 5 · MCP Server（2:15–2:45）
**画面**：展示 `claude_desktop_config.json` 里接入 MCP server，或终端列出 `list_worldcup_matches` / `predict_match` 工具。
**中文旁白**：
> 除了网页，我还用 Go 写了一个 MCP Server，暴露 `list_worldcup_matches` 和 `predict_match` 两个工具。这样 Claude、Cursor 这类 AI 代理能直接查询比赛和预测，不用自己接 API。

### 镜头 6 · 收尾（2:45–3:00）
**画面**：回到首页，展示 GitHub 链接。
**中文旁白**：
> 代码已开源在 GitHub。这就是 Injective World Cup Predictor——一个用 x402、MCP、CCTP 构建的世界杯预测 dApp。谢谢大家！
**英文**：
> Open source on GitHub. Thanks for watching!

---

## ✅ 必须口播到的技术关键词（评分点，缺一不可）

| 技术 | 口播要点 |
|---|---|
| **x402** | 0.01 USDC 付费预测，Injective EVM 结算，亚秒级确认 |
| **MCP Server** | Go 写的 MCP server，AI 代理可调 `list_worldcup_matches` / `predict_match` |
| **CCTP** | 作为跨链 USDC 奖励分发路径（路线图，未来预测池奖金跨链发放） |

---

## 录制小贴士

- 语速放慢，每镜留 1–2 秒空场，后期好剪。
- **付费镜头**：没配 Injective 钱包也没关系，直接展示 `402 Payment Required` 响应（Network 面板）就足以证明 x402 集成是真的。
- 导出 MP4（建议 < 200MB）。GitHub 放不了视频，可放：YouTube 不公开链接 / 或直接挂到 X 推文。
- 录完把视频链接填回 `README.md` 第 25 行 `Demo video:` 处。
