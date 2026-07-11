package predictions

import (
	"hash/fnv"
	"math"

	"worldcup-injective/internal/matches"
)

type Prediction struct {
	MatchID    string  `json:"match_id"`
	Winner     string  `json:"winner"`
	ScoreA     int     `json:"score_a"`
	ScoreB     int     `json:"score_b"`
	Confidence float64 `json:"confidence"`
	Rationale  string  `json:"rationale"`
	ModelType  string  `json:"model_type"`
}

// Predict is the core deterministic heuristic (FIFA rank gap + market odds).
// Swap this body for a real AI call (OpenAI/Claude) when ready.
func Predict(m matches.Match) Prediction {
	// 对阵未定（半决赛2/决赛尚未产生对手）时不做无意义预测。
	if m.TeamARank == 0 && m.TeamBRank == 0 {
		return Prediction{
			MatchID:    m.ID,
			Winner:     "TBD",
			Confidence: 0,
			Rationale:  "对阵尚未产生，等待 1/4 决赛与半决赛结果后再预测。",
			ModelType:  "heuristic-v1",
		}
	}

	h := fnv.New32a()
	h.Write([]byte(m.ID))
	seed := h.Sum32()

	rankDiff := float64(m.TeamBRank - m.TeamARank)
	oddsSignal := m.TeamBOdds - m.TeamAOdds
	strength := rankDiff*0.05 + oddsSignal*0.3

	confidence := 0.5 + math.Tanh(strength)/2.0
	if confidence < 0.5 {
		confidence = 1 - confidence
	}

	var winner string
	var scoreA, scoreB int
	if strength >= 0 {
		winner = m.TeamA
		scoreA, scoreB = 2+int(seed%2), 1
	} else {
		winner = m.TeamB
		scoreA, scoreB = 1, 2+int(seed%2)
	}

	return Prediction{
		MatchID:    m.ID,
		Winner:     winner,
		ScoreA:     scoreA,
		ScoreB:     scoreB,
		Confidence: confidence,
		Rationale:  "Based on FIFA ranking gap and market odds.",
		ModelType:  "heuristic-v1",
	}
}

// Free returns the basic public prediction (free tier, no payment).
func Free(m matches.Match) Prediction {
	p := Predict(m)
	p.ModelType = "heuristic-v1-free"
	p.Rationale = "Free tier: winner + scoreline only. Unlock Premium for full analysis."
	return p
}

// Premium returns the full analysis gated behind x402 payment.
// Uses a real LLM when configured; falls back to the heuristic otherwise.
func Premium(m matches.Match) Prediction {
	if p, ok := PredictAI(m); ok {
		return p
	}
	// AI 未配置或调用失败 → 回退规则模型。
	p := Predict(m)
	p.ModelType = "heuristic-v1-premium"
	if p.Winner == "TBD" {
		p.Rationale = "Premium: 对阵确定后将提供完整概率模型与投注价值分析。"
		return p
	}
	p.Rationale = buildPremiumRationale(m, p) + "（AI 未启用或调用失败，当前为规则模型回退）"
	return p
}

func buildPremiumRationale(m matches.Match, p Prediction) string {
	rankGap := m.TeamARank - m.TeamBRank
	oddsGap := m.TeamAOdds - m.TeamBOdds
	return "Premium analysis: " + p.Winner + " 被模型看好（置信度 " +
		formatPct(p.Confidence) + "）。FIFA 排名差 " + itoa(rankGap) +
		"，赔率差 " + ftoa(oddsGap) + "。建议结合实时赔率与阵容新闻再下注。"
}

func formatPct(v float64) string {
	return ftoa(math.Round(v*1000)/10) + "%"
}

func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	neg := v < 0
	if neg {
		v = -v
	}
	buf := []byte{}
	for v > 0 {
		buf = append([]byte{byte('0' + v%10)}, buf...)
		v /= 10
	}
	if neg {
		buf = append([]byte{'-'}, buf...)
	}
	return string(buf)
}

func ftoa(v float64) string {
	// simple 1-decimal formatter without fmt import in hot path
	sign := ""
	if v < 0 {
		sign = "-"
		v = -v
	}
	iv := int(v + 0.05)
	dec := int((v*10)+0.5) % 10
	return sign + itoa(iv) + "." + itoa(dec)
}
