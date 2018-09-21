package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io"
	"github.com/davecgh/go-spew/spew"
	"log"
	"fmt"
)

//定义一个区块结构体
type Block struct {
	Index     int    //区块高度
	Timestamp string //区块生产时间
	BPM       int    //区块数据(心跳数)
	Hash      string //区块的HashId
	PreHash   string //上一区块的HashId
}

//区块链 slice
var BlockChain []Block


//计算区块的HashId
func caclulateHash(block Block) string{
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PreHash
	h := sha256.New()
	h.Write([]byte(record))
	hashId := h.Sum(nil)
	return hex.EncodeToString(hashId)
}

//生成区块
func generateBlock(oldBlock Block , BMP int)(Block , error)  {
	var newBlock Block
	t := time.Now()
	newBlock.Timestamp = t.String()
	newBlock.BPM = BMP
	newBlock.PreHash = oldBlock.Hash
	newBlock.Index = oldBlock.Index + 1
	newBlock.Hash = caclulateHash(newBlock)

	return newBlock , nil
}

//校验区块
func isBlockValid(newBlock Block , oldBlock Block) bool{
	//校验区块高度
	if oldBlock.Index+1 != newBlock.Index{return false}
	//检验新区块的PreHash是否等于上一区块的Hash
	if oldBlock.Hash != newBlock.PreHash{return false}
	//检验新区块的Hash是否已经改变
	if caclulateHash(newBlock) != newBlock.Hash{return false}
	return true
}

//选择区块链最长的链
func replaceChani(newBlock []Block){
	if len(newBlock) > len(BlockChain){
		BlockChain = newBlock
	}
}

//Web服务,借助Gorilla/mux包可以做初始化web的服务
func run() error  {
	//创建http请求响应
	mux := makeMuxRouter()
	//httpAddr := os.Getenv("ADDR")
	//log.Println("Listening on ",os.Getenv("ADDR"))
	s := &http.Server{
		//Addr: 			":"+httpAddr,
		Addr: 			":8080",
		Handler:		mux,
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
		MaxHeaderBytes:	1 << 20,
	}
	if err := s.ListenAndServe(); err != nil{
		return err
	}
	return nil
}
//创建http请求响应
func makeMuxRouter() http.Handler{
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/",handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/",handleWriteBlock).Methods("POST")
	return muxRouter
}
//Get请求的handler
func handleGetBlockchain(w http.ResponseWriter , r *http.Request)  {
	bytes, err := json.MarshalIndent(BlockChain, "","")
	if err != nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}
//Post请求的handler稍微有些复杂,先定义post请求的数据
type Message struct {
	BPM int
}
//Post请求响应的Handler
func handleWriteBlock(w http.ResponseWriter, r *http.Request)  {
	var m Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	fmt.Println("请求体数据:",m)

	newBlock, err := generateBlock(BlockChain[len(BlockChain) - 1],m.BPM)
	if err != nil{
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	if isBlockValid(newBlock , BlockChain[len(BlockChain) - 1]){
		newBlockChain := append(BlockChain,newBlock)
		replaceChani(newBlockChain)
		spew.Dump(BlockChain)
	}
	respondWithJSON(w, r, http.StatusCreated, newBlock)
}
//http请求无论成功与否都需要做出响应
func respondWithJSON(w http.ResponseWriter , r *http.Request , code int , payload interface{})  {
	response, err := json.MarshalIndent(payload , "" , "  ")
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

//main函数入口
func main()  {
	//err := godotenv.Load()
	//if err != nil{
	//	log.Fatal(err)
	//}

	go func() {
		t := time.Now()
		//创世区块
		genesisBlock := Block{0,t.String() ,0 ,"",""}
		spew.Dump(genesisBlock)
		BlockChain = append(BlockChain,genesisBlock)
	}()
	//开始运行wen程序
	log.Fatal(run())
}