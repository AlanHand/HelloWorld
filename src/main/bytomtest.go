package main

import (
	"io"
	"crypto/rand"
	"crypto/hmac"
	"crypto/sha512"
	"github.com/bytom/crypto/ed25519/ecmath"
	"encoding/hex"
	"fmt"
	"github.com/bytom/crypto"
	"github.com/bytom/common/bech32"
	"bytes"
)

type (
	XPrv [64]byte
	XPub [64]byte
)
type PublicKey []byte
func main()  {
	//创建秘钥对
	privateKey, publicKey,_ := newXKeys(nil)
	//字节数据不能直接转换为字符串,可以将字节数组转换为slice
	prk := privateKey[0:64]
	pubk := publicKey[0:64]
	fmt.Println(hex.EncodeToString(prk))
	fmt.Println(hex.EncodeToString(pubk))

	//根据公钥哈希得到地址
	publicHash := crypto.Ripemd160(publicKey.PublicKey())
	address ,err := getAddressByPublicKeyHash("bm" , 0x00 , publicHash)
	if err != nil{

	}
	fmt.Println("地址:"+address)
}

//公钥哈希得到地址
func getAddressByPublicKeyHash(hrp string, witnessVersion byte, witnessProgram []byte) (string,error) {

	//将公钥Hash160进行位移变化
	converted, err := ConvertBits(witnessProgram,8,5,true)
	if err != nil {
		return "", err
	}
	combined := make([]byte, len(converted)+1)
	//第一个字节为版本号
	combined[0] = witnessVersion
	copy(combined[1:], converted)
	//编码
	bech , err := bech32.Bech32Encode(hrp,combined)

	if err != nil {
		return "", err
	}

	//解码校验一次
	version, program, err := decodeSegWitAddress(bech)
	if err != nil {
		return "", fmt.Errorf("invalid segwit address: %v", err)
	}

	if version != witnessVersion || !bytes.Equal(program, witnessProgram) {
		return "", fmt.Errorf("invalid segwit address")
	}

	return bech, nil
}
func decodeSegWitAddress(address string) (byte, []byte, error) {
	// Decode the bech32 encoded address.
	_, data, err := bech32.Bech32Decode(address)
	if err != nil {
		return 0, nil, err
	}

	// The first byte of the decoded address is the witness version, it must
	// exist.
	if len(data) < 1 {
		return 0, nil, fmt.Errorf("no witness version")
	}

	// ...and be <= 16.
	version := data[0]
	if version > 16 {
		return 0, nil, fmt.Errorf("invalid witness version: %v", version)
	}

	// The remaining characters of the address returned are grouped into
	// words of 5 bits. In order to restore the original witness program
	// bytes, we'll need to regroup into 8 bit words.
	regrouped, err := bech32.ConvertBits(data[1:], 5, 8, false)
	if err != nil {
		return 0, nil, err
	}

	// The regrouped data must be between 2 and 40 bytes.
	if len(regrouped) < 2 || len(regrouped) > 40 {
		return 0, nil, fmt.Errorf("invalid data length")
	}

	// For witness version 0, address MUST be exactly 20 or 32 bytes.
	if version == 0 && len(regrouped) != 20 && len(regrouped) != 32 {
		return 0, nil, fmt.Errorf("invalid data length for witness "+
			"version 0: %v", len(regrouped))
	}

	return version, regrouped, nil
}
func ConvertBits(data []byte, fromBits, toBits uint8, pad bool) ([]byte, error) {
	if fromBits < 1 || fromBits > 8 || toBits < 1 || toBits > 8 {
	return nil, fmt.Errorf("only bit groups between 1 and 8 allowed")
	}

	// The final bytes, each byte encoding toBits bits.
	var regrouped []byte

	// Keep track of the next byte we create and how many bits we have
	// added to it out of the toBits goal.
	nextByte := byte(0)
	filledBits := uint8(0)

	for _, b := range data {

		// Discard unused bits.
		b = b << (8 - fromBits)

		// How many bits remaining to extract from the input data.
		remFromBits := fromBits
		for remFromBits > 0 {
			// How many bits remaining to be added to the next byte.
			remToBits := toBits - filledBits

			// The number of bytes to next extract is the minimum of
			// remFromBits and remToBits.
			toExtract := remFromBits
			if remToBits < toExtract {
				toExtract = remToBits
			}

			// Add the next bits to nextByte, shifting the already
			// added bits to the left.
			nextByte = (nextByte << toExtract) | (b >> (8 - toExtract))

			// Discard the bits we just extracted and get ready for
			// next iteration.
			b = b << toExtract
			remFromBits -= toExtract
			filledBits += toExtract

			// If the nextByte is completely filled, we add it to
			// our regrouped bytes and start on the next byte.
			if filledBits == toBits {
				regrouped = append(regrouped, nextByte)
				filledBits = 0
				nextByte = 0
			}
		}
	}

	// We pad any unfinished group if specified.
	if pad && filledBits > 0 {
	nextByte = nextByte << (toBits - filledBits)
	regrouped = append(regrouped, nextByte)
	filledBits = 0
	nextByte = 0
	}

	// Any incomplete group must be <= 4 bits, and all zeroes.
	if filledBits > 0 && (filledBits > 4 || nextByte != 0) {
	return nil, fmt.Errorf("invalid incomplete group")
	}

	return regrouped, nil
}

func newXKeys(r io.Reader)(xprv XPrv , xpub XPub , err error)  {
	xprv,err = NewXPrv(r)
	if err != nil{
		return
	}
	return xprv,xprv.XPub(),nil
}

func NewXPrv(r io.Reader) (xprv XPrv , err error){
	if r == nil{
		r = rand.Reader
	}
	var entropy [64]byte
	_,err = io.ReadFull(r,entropy[:])

	if err != nil{
		return xprv,err
	}
	return RootXPrv(entropy[:]),nil
}

func RootXPrv(seed []byte)(xprv XPrv){
	h := hmac.New(sha512.New , []byte{'R','o','o','t'})
	h.Write(seed)
	h.Sum(xprv[:0])
	pruneRootScalar(xprv[:32])
	return
}

func pruneRootScalar(s []byte)  {
	s[0] &= 248
	s[31] &= 31 // clear top 3 bits
	s[31] |= 64 // set second highest bit
}

//根据私钥得到公钥,方法中的返回变量指定了名称的话那么返回的就是这个xpub变量
func (xprv XPrv)XPub()(xpub XPub)  {
	var scalar ecmath.Scalar
	//将私钥的前32个字节数据拷贝到scalar数组中
	copy(scalar[:] , xprv[:32])
	//由私钥得到32个字节的公钥
	var P ecmath.Point
	P.ScMulBase(&scalar)
	buf := P.Encode()

	//将32个字节的公钥拷贝到64个字节变量xpub的前32个字节中
	copy(xpub[:32],buf[:])
	//将32个字节的私钥拷贝到64个字节变量xpub的后32个字节中
	copy(xpub[32:],xprv[32:])
	return
}

func (xpub XPub)PublicKey() PublicKey  {
	return PublicKey(xpub[:32])
}

