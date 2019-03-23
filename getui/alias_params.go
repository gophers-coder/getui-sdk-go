package getui

type AliasParams struct {
	AliasList []OneCidAlias `json:"alias_list"`
}

type OneCidAlias struct {
	Cid   string `json:"cid"`
	Alias string `json:"alias"`
}
