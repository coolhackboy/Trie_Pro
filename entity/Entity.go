package entity

/**
{
	"COIN_NUM":5,     清算的币的总数
	"GAS_ALL":5,      币的GAS消耗
  "COUNT":{               统计
		"SPACEID":77,     空间ID=77
		"GAS":5,          GAS=5
		"TXS":            所有交易
			[
				"TX_A_HASH":               交易的hash:9988asa...
				{
					"COIN":"111,116,117",  交易的币
					"GAS":3,               GAS消耗
				},
				"TX_A_HASH":
				{
					"COIN":"299,999",
					"GAS":2,
				},
			],
			"RECOINAGE":                   重铸
			{
				"NUM":2                    数量
				"COIN":"111,117",          编号
			},
		},
	},
}
 */
type Recoingage struct {
	NUM int `json:"num"`
	COIN string `json:"coin"`
}

type Txs struct {
	COIN string `json:"coin"`
	GAS int `json:"gas"`
}
type Count struct {
	SPACEID int `json:"spaceid"`
	GAS int `json:"gas"`
	Txs []Txs `json:txs`
	Recoingage Recoingage `json:"recoinage"`
}
type DataEntity struct {
	COIN_NUM int `json:"coin_num"`
	GAS_ALL int `json:"gas_all"`
	Count Count `json:"count"`
}

type DataMap struct {
	K string `json:"k"`
	V string `json:"v"`
}
