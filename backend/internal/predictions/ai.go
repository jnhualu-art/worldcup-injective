package predictions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"worldcup-injective/internal/matches"
)

var (
	aiAPIKey  string
	aiBaseURL string
	aiModel   string
	aiEnabled bool
)

func init() {
	aiAPIKey = os.Getenv("OPENAI_API_KEY")
	aiBaseURL = os.Getenv("OPENAI_BASE_URL")
	if aiBaseURL == "" {
		aiBaseURL = "https://api.openai.com/v1"
	}
	aiModel = os.Getenv("OPENAI_MODEL")
	if aiModel == "" {
		aiModel = "gpt-4o-mini"
	}
	aiEnabled = aiAPIKey != ""
}

// aiResponse is the structured prediction we ask the model to return as JSON.
type aiResponse struct {
	Winner     string  `json:"winner"`
	ScoreA     int     `json:"score_a"`
	ScoreB     int     `json:"score_b"`
	Confidence float64 `json:"confidence"`
	Rationale  string  `json:"rationale"`
}

// PredictAI calls a real LLM (OpenAI-compatible) for a grounded prediction.
// Returns (prediction, true) on success, or (zero, false) if AI is not
// configured or the call fails — callers should fall back to the heuristic.
func PredictAI(m matches.Match) (Prediction, bool) {
	if !aiEnabled {
		return Prediction{}, false
	}
	// 对阵未定则无法预测。
	if m.TeamARank == 0 && m.TeamBRank == 0 {
		return Prediction{}, false
	}

	system := "You are an expert football analyst. Given World Cup match data, " +
		"return a realistic JSON prediction with exactly these fields: " +
		"\"winner\" (must be one of team_a or team_b), \"score_a\" (int), " +
		"\"score_b\" (int), \"confidence\" (float 0-1 for the winner), " +
		"\"rationale\" (one concise sentence in English). Be realistic."

	user := fmt.Sprintf(
		"Match: %s vs %s (%s). FIFA rank: %d vs %d. "+
			"Market odds: %.2f vs %.2f (draw %.2f). Kickoff: %s. Return only JSON.",
		m.TeamA, m.TeamB, m.Group,
		m.TeamARank, m.TeamBRank,
		m.TeamAOdds, m.TeamBOdds, m.DrawOdds,
		m.Kickoff.Format(time.RFC3339),
	)

	payload := map[string]any{
		"model": aiModel,
		"messages": []map[string]string{
			{"role": "system", "content": system},
			{"role": "user", "content": user},
		},
		"response_format": map[string]string{"type": "json_object"},
		"temperature":     0.7,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return Prediction{}, false
	}

	req, err := http.NewRequest(http.MethodPost, aiBaseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return Prediction{}, false
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+aiAPIKey)

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return Prediction{}, false
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return Prediction{}, false
	}

	var or struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(raw, &or); err != nil || len(or.Choices) == 0 {
		return Prediction{}, false
	}

	var ar aiResponse
	if err := json.Unmarshal([]byte(or.Choices[0].Message.Content), &ar); err != nil {
		return Prediction{}, false
	}

	// 校验 winner 合法性，非法时按比分推导。
	winner := ar.Winner
	if winner != m.TeamA && winner != m.TeamB {
		if ar.ScoreA >= ar.ScoreB {
			winner = m.TeamA
		} else {
			winner = m.TeamB
		}
	}

	return Prediction{
		MatchID:    m.ID,
		Winner:     winner,
		ScoreA:     ar.ScoreA,
		ScoreB:     ar.ScoreB,
		Confidence: ar.Confidence,
		Rationale:  ar.Rationale,
		ModelType:  "ai-" + aiModel,
	}, true
}
