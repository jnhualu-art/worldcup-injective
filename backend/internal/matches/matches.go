package matches

import "time"

type Match struct {
	ID        string    `json:"id"`
	TeamA     string    `json:"team_a"`
	TeamB     string    `json:"team_b"`
	Group     string    `json:"group"`
	Kickoff   time.Time `json:"kickoff"`
	Stadium   string    `json:"stadium"`
	TeamARank int       `json:"team_a_rank"`
	TeamBRank int       `json:"team_b_rank"`
	TeamAOdds float64   `json:"team_a_odds"`
	TeamBOdds float64   `json:"team_b_odds"`
	DrawOdds  float64   `json:"draw_odds"`
	Status    string    `json:"status"` // upcoming | live | finished
	Result    string    `json:"result"` // e.g. "2-0", empty if not played
}

// 真实 2026 美加墨世界杯淘汰赛数据（截至 2026-07-11）。
// 时间为 UTC。Rank 参考 FIFA 7 月排名近似值，Odds 为示意赔率。
var all = []Match{
	{
		ID: "qf-fra-mar", TeamA: "France", TeamB: "Morocco",
		Group: "Quarter-final", Stadium: "Boston Stadium",
		Kickoff:   time.Date(2026, 7, 9, 20, 0, 0, 0, time.UTC),
		TeamARank: 2, TeamBRank: 12, TeamAOdds: 2.0, TeamBOdds: 5.0, DrawOdds: 3.4,
		Status: "finished", Result: "2-0",
	},
	{
		ID: "qf-esp-bel", TeamA: "Spain", TeamB: "Belgium",
		Group: "Quarter-final", Stadium: "Los Angeles Stadium",
		Kickoff:   time.Date(2026, 7, 10, 19, 0, 0, 0, time.UTC),
		TeamARank: 3, TeamBRank: 8, TeamAOdds: 1.9, TeamBOdds: 4.5, DrawOdds: 3.5,
		Status: "finished", Result: "2-1",
	},
	{
		ID: "qf-nor-eng", TeamA: "Norway", TeamB: "England",
		Group: "Quarter-final", Stadium: "Miami Stadium",
		Kickoff:   time.Date(2026, 7, 11, 21, 0, 0, 0, time.UTC),
		TeamARank: 11, TeamBRank: 4, TeamAOdds: 3.0, TeamBOdds: 2.4, DrawOdds: 3.3,
		Status: "upcoming",
	},
	{
		ID: "qf-arg-sui", TeamA: "Argentina", TeamB: "Switzerland",
		Group: "Quarter-final", Stadium: "Kansas City Stadium",
		Kickoff:   time.Date(2026, 7, 12, 1, 0, 0, 0, time.UTC),
		TeamARank: 1, TeamBRank: 14, TeamAOdds: 1.8, TeamBOdds: 5.5, DrawOdds: 3.6,
		Status: "upcoming",
	},
	{
		ID: "sf-fra-esp", TeamA: "France", TeamB: "Spain",
		Group: "Semi-final", Stadium: "Dallas Stadium",
		Kickoff:   time.Date(2026, 7, 14, 19, 0, 0, 0, time.UTC),
		TeamARank: 2, TeamBRank: 3, TeamAOdds: 2.1, TeamBOdds: 2.2, DrawOdds: 3.3,
		Status: "upcoming",
	},
	{
		ID: "sf-tbd", TeamA: "TBD", TeamB: "TBD",
		Group: "Semi-final", Stadium: "Atlanta Stadium",
		Kickoff:   time.Date(2026, 7, 15, 19, 0, 0, 0, time.UTC),
		TeamARank: 0, TeamBRank: 0, TeamAOdds: 0, TeamBOdds: 0, DrawOdds: 0,
		Status: "upcoming",
	},
}

func All() []Match { return all }

func ByID(id string) (Match, bool) {
	for _, m := range all {
		if m.ID == id {
			return m, true
		}
	}
	return Match{}, false
}
