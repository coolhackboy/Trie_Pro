package trie

import (
	"fmt"
	"github.com/coolhackboy/trie/entity"
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)
// Used for testing
func newEmpty() *trie.Trie {
	diskdb, _ := ethdb.NewMemDatabase()
	trie, _ := trie.New(common.Hash{}, trie.NewDatabase(diskdb))
	return trie
}

func updateString(trie *trie.Trie, k string, v []byte) {
	trie.Update([]byte(k),v)
}

func getString(trie *trie.Trie, k string) []byte {
	return trie.Get([]byte(k))
}

func TestInsert(t *testing.T) {

	txs := []entity.Txs{
		entity.Txs{"1",1},
		entity.Txs{"2",1},
	}

	recoingage := entity.Recoingage{1,"2"}

	count := entity.Count{1,2,txs,recoingage}

	a := entity.DataEntity{1,2,count}

	bytes, err := json.Marshal(a)

	if err != nil {
		t.Errorf("json convert error")
	}

	trie := newEmpty()
	updateString(trie,"all",bytes);

	res := getString(trie, "all")
	fmt.Printf("%s\n",res)

}

func TestGetTxs(t *testing.T)  {
	txs := []entity.Txs{
		entity.Txs{"1",1},
		entity.Txs{"2",1},
	}

	//recoingage := entity.Recoingage{1,"2"}
	//
	//count := entity.Count{1,2,txs,recoingage}
	//
	//a := entity.DataEntity{1,2,count}

	bytes, err := json.Marshal(txs)
	if err != nil {
		t.Errorf("json convert error")
	}
	trie := newEmpty()
	updateString(trie,"txs",bytes);

	res := getString(trie, "txs")
	fmt.Printf("%s\n",res)
}

func TestGetCount(t *testing.T)  {
	txs := []entity.Txs{
		entity.Txs{"1",1},
		entity.Txs{"2",1},
	}

	recoingage := entity.Recoingage{1,"2"}

	count := entity.Count{1,2,txs,recoingage}

	//a := entity.DataEntity{1,2,count}


	bytes, err := json.Marshal(count)
	if err != nil {
		t.Errorf("json convert error")
	}
	trie := newEmpty()
	updateString(trie,"count",bytes);

	res := getString(trie, "count")
	fmt.Printf("%s\n",res)
}

func TestGetRecoinage(t *testing.T){
	//txs := []entity.Txs{
	//	entity.Txs{"1",1},
	//	entity.Txs{"2",1},
	//}

	recoingage := entity.Recoingage{1,"2"}

	//count := entity.Count{1,2,txs,recoingage}

	//a := entity.DataEntity{1,2,count}


	bytes, err := json.Marshal(recoingage)
	if err != nil {
		t.Errorf("json convert error")
	}
	trie := newEmpty()
	updateString(trie,"recoinage",bytes);

	res := getString(trie, "recoinage")
	fmt.Printf("%s\n",res)
}

func TestFormatEntity(t *testing.T)  {
	txs := []entity.Txs{
		entity.Txs{"1",1},
		entity.Txs{"2",1},
	}

	recoingage := entity.Recoingage{1,"2"}

	count := entity.Count{1,2,txs,recoingage}

	a := entity.DataEntity{1,2,count}


	bytes, err := json.Marshal(a)
	if err != nil {
		t.Errorf("json convert error")
	}
	fmt.Printf("%s\n",bytes)
}