package main

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"sort"
	"log"
)

type DPOS_Block struct {
	Index		int
	Timestamp	string
	BPM			int
	Hash 		string
	PreHash 	string
	Delegate	string
}

//创建区块函数
func generate_DPOS_Block(oldBlock DPOS_Block, _BMP int , address string)(DPOS_Block, error)  {
	var newBlock DPOS_Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = _BMP
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = createDPOSBlockHash(newBlock)
	newBlock.Delegate = address

	return newBlock,nil
}
//创建区块的Hash
func createDPOSBlockHash(block DPOS_Block)string{
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PreHash
	sha3 := sha256.New()
	sha3.Write([]byte(record))
	hash := sha3.Sum(nil)
	return hex.EncodeToString(hash)
}
//检验区块
func isDPOSBlockValid(newBlock DPOS_Block, oldBlock DPOS_Block)bool  {
	if oldBlock.Index+1 != newBlock.Index{
		return false
	}
	if oldBlock.Hash != newBlock.PreHash{
		return false
	}
	return true
}

//区块集合
var blockChain []DPOS_Block

//超级节点结构体
type Trustee struct {
	name	string
	votes	int
}
//超级节点集合
type trusteeList []Trustee

//三个函数进行超级节点的排序
func (_trusteeList trusteeList) Len() int  {
	return len(_trusteeList)
}
func (_trusteeList trusteeList) Swap(i int , j int)  {
	_trusteeList[i],_trusteeList[j] = _trusteeList[j] , _trusteeList[i]
}
func (_truesteeList trusteeList) Less(i int , j int)bool  {
	return _truesteeList[j].votes < _truesteeList[i].votes
}

//选举得票数最高的5个节点作为前缀节点,并打乱其顺序
func selectTruestee() ([]Trustee) {
	//随机产生12个节点
	_trusteeList := []Trustee{
		{"node1", rand.Intn(100)},
		{"node2", rand.Intn(100)},
		{"node3", rand.Intn(100)},
		{"node4", rand.Intn(100)},
		{"node5", rand.Intn(100)},
		{"node6", rand.Intn(100)},
		{"node7", rand.Intn(100)},
		{"node8", rand.Intn(100)},
		{"node9", rand.Intn(100)},
		{"node10", rand.Intn(100)},
		{"node11", rand.Intn(100)},
		{"node12", rand.Intn(100)},
	}
	sort.Sort(trusteeList(_trusteeList))
	//取前5个
	result := _trusteeList[:5]
	_trusteeList = result[1:]
	_trusteeList = append(_trusteeList,result[0])
	log.Println("当前超级节点列表是", _trusteeList)
	return _trusteeList

}
func main()  {
	t :=time.Now()
	//创建创世区块
	genesisBlock := DPOS_Block{0, t.String(), 0, createDPOSBlockHash(DPOS_Block{}),"",""}
	blockChain = append(blockChain , genesisBlock)
	//超级节点轮流出块
	for _,trustee := range selectTruestee(){
		_BPM := rand.Intn(100)
		blockHeight := len(blockChain)
		oldBlock := blockChain[blockHeight - 1]
		newBlock,err := generate_DPOS_Block(oldBlock , _BPM , trustee.name)
		if err != nil{
			log.Println(err)
			continue
		}
		if isDPOSBlockValid(newBlock,oldBlock){
			blockChain = append(blockChain,newBlock)
			log.Println("当前出块节点:",trustee.name)
			log.Println("当前出块数量:",len(blockChain) -1)
			log.Println("当前区块信息:",blockChain[len(blockChain) - 1])
		}
	}
}