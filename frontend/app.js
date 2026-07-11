const API = 'http://localhost:8080'; // 直连 Go 后端（本地预览无需 x402 网关）

async function loadMatches() {
  const res = await fetch(`${API}/api/matches`);
  const data = await res.json();
  const container = document.getElementById('matches');
  container.innerHTML = '';

  for (const m of data) {
    const el = document.createElement('div');
    el.className = 'match';
    const resultBadge = m.result
      ? `<span class="badge result">${m.team_a} ${m.result} ${m.team_b}</span>`
      : '';
    el.innerHTML = `
      <div class="teams">${m.team_a} vs ${m.team_b} <span class="badge">${m.status}</span> ${resultBadge}</div>
      <div class="meta">${m.group} · ${m.stadium} · ${new Date(m.kickoff).toLocaleString()}</div>
      <div class="actions">
        <button class="free" onclick="getFree('${m.id}', this)">免费预测</button>
        <button class="paid" onclick="getPremium('${m.id}', this)">付费 AI 预测 (0.01 USDC)</button>
      </div>
      <div class="prediction" id="pred-${m.id}" style="display:none"></div>
    `;
    container.appendChild(el);
  }
}

async function getFree(id, btn) {
  btn.disabled = true;
  const res = await fetch(`${API}/api/predict/${id}`);
  const pred = await res.json();
  showPred(id, pred);
  btn.disabled = false;
}

async function getPremium(id, btn) {
  btn.disabled = true;
  // 本地预览直连后端返回完整预测；生产环境经 x402 网关(:3000) 会先返回 402 要求 0.01 USDC 支付。
  const res = await fetch(`${API}/api/premium-predict/${id}`);
  const pred = await res.json();
  showPred(id, pred);
  btn.disabled = false;
}

function showPred(id, pred) {
  const box = document.getElementById(`pred-${id}`);
  box.style.display = 'block';
  const scoreline = (pred.winner === 'TBD' || (pred.score_a === 0 && pred.score_b === 0))
    ? '待定'
    : `${pred.score_a} - ${pred.score_b}`;
  const confidence = pred.confidence ? (pred.confidence * 100).toFixed(1) + '%' : '—';
  box.innerHTML = `
    <strong>预测结果：${pred.winner} ${scoreline}</strong>
    <div>置信度：${confidence}</div>
    <div>模型：${pred.model_type || '—'}</div>
    <div style="margin-top:0.4rem;color:#555">${pred.rationale || ''}</div>
  `;
}

loadMatches();
