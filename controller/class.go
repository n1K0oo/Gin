package controller

type Username_table struct {
	Username  string `json:"username" `
	Password  string `json:"password"`
	Authority int    `json:"authority"`
	Name      string `json:"name"`
	Tel       string `json:"tel"`
	ID        string `json:"id"`
	Area      string `json:"area"`
	Street    string `json:"street"`
}
type Estatemess_table struct {
	Estateno      int     `json:"estateno"`
	Estatename    string  `json:"estatename"`
	Type1         string  `json:"type1"`
	Type2         string  `json:"type2"`
	Btime         string  `json:"btime"`
	Etime         string  `json:"etime"`
	Estateagent   string  `json:"estateagent"`
	Area          string  `json:"area"`
	Street        string  `json:"street"`
	Address       string  `json:"address"`
	Communityname string  `json:"communityname"`
	Roomno        string  `json:"roomno"`
	Toward        string  `json:"toward"`
	Layer         string  `json:"layer"`
	Mianji        float64 `json:"mianji"`
	Price         float64 `json:"price"`
	Allprice      float64 `json:"allprice"`
	Describes     string  `json:"describes"`
	Username      string  `json:"username"`
	Image1        string  `json:"image1"`
	Buyer         string  `json:"buyer"`
	Flag2         int     `json:"flag2"`
	Flag          int     `json:"flag"`
}
type Indoor_table struct {
	Estateno int    `json:"estateno"`
	Image2   string `json:"image2"`
}
