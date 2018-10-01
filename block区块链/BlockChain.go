package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

/*
定义区块结构
*/
type Block struct {
	Timestamp     int64  //时间戳,当前区块的生成时间
	Data          []byte //当前区块存放的行为信息(如账单记录)
	Hash          []byte //当前区块的 加密哈希值
	PrevBlockHash []byte //上一个区块的 加密哈希值
}

/*
为区块结构体绑定一个方法
将当前区块的Timestamp + Date + PrevBlockHash ----> 绑定并加密为一个Hash(赋予唯一性)
*/
func (this *Block) SetHash() {
	//将时间戳转换为[]byte
	timestamp := []byte(strconv.FormatInt(this.Timestamp, 10))
	//将三个二进制切片进行拼接
	headers := bytes.Join([][]byte{this.Data, this.PrevBlockHash, timestamp}, []byte{})
	//将拼接之后的 headers 进行 sha256加密
	hash := sha256.Sum256(headers)
	//将加密后的hash值 赋值给 区块结构
	this.Hash = hash[:]
}

/*
NewBlock 新建一个区块的 API,date 当前区块所保存的数据,prevBlockHash 当前区块的上一块Hash,return 返回新的区块结构体
*/
func NewBlock(date, prevBlockHash []byte) *Block {
	// 生成一个区块
	block := Block{}
	// 并赋值-->(创建时间,date,前驱Hash)
	block.Timestamp = time.Now().Unix()
	block.Data = date
	block.PrevBlockHash = prevBlockHash
	// 进行加密
	block.SetHash() //使用方法进行Hash加密,将加密后的Hash值赋值给block.Hash
	// 将当前区块的加密Hash值返回给上级
	return &block
}

/*
 定义区块链结构(创世块+>下一块+>下一块...)
*/
type BlockChain struct {
	Blocks []*Block //有序的区块集合
}

//定义创世块
func NewGenesisBlock() *Block {
	genesisBlock := Block{}
	genesisBlock.Data = []byte("Genesis Block")
	genesisBlock.PrevBlockHash = []byte{}
	genesisBlock.Timestamp = time.Now().Unix()
	return &genesisBlock
}

//新建一个区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

//添加新区块到区块链中
func (this *BlockChain) AddBlock(data []byte) {
	// 1 得到当前新区块的前置区块
	prevBlock := this.Blocks[len(this.Blocks)-1]
	// 2 根据date创建一个新的区块
	newBlock := NewBlock(data, prevBlock.Hash)
	// 3 依照新区块和前置区块 ,添加到区块链Blocks中
	this.Blocks = append(this.Blocks, newBlock)
}
func main() {
	//创建一个区块链
	bc := NewBlockChain()
	//获取用户指令,并作出相应操作
	var cmd string
	for {
		fmt.Println("按`1`添加信息并创建新区块\n",
			"按`0`查看当前区块链信息\n",
			"按`exit`退出")
		fmt.Scanf("%v\n", &cmd) //fmt.Scanf格式内要加上\n
		switch cmd {
		case "1":
			//添加一个区块
			input := make([]byte, 1024)
			fmt.Println("请输入新区块信息 :")
			fmt.Scanf("%v\n", &input) //fmt.Scanf格式内要加上\n
			bc.AddBlock(input)
		case "0":
			for i, block := range bc.Blocks {
				fmt.Println("----------------------")
				fmt.Println("第", i, "个区块信息为 :")
				fmt.Printf("timeStamp : %v\n", time.Unix(block.Timestamp, 0))
				fmt.Printf("data : %s\n", block.Data)
				fmt.Printf("hash : %x\n", block.Hash)
				fmt.Printf("PrevBlockHash : %x\n", block.PrevBlockHash)
			}
		case "exit":
			//退出程序
			fmt.Println("您已经退出")
			return
		}
	}
}
