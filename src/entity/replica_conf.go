package entity

type ReplicaConf struct {
	Gid           string   `json:"gid"`
	Pda_code      string   `json:"pda_code"`
	Group_members []string `json:"group_members"`
}
