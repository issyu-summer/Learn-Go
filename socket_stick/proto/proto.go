package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(msg string) ([]byte, error) {
	//四个字节的消息长度
	length := int32(len(msg))
	pkg := new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err!=nil{
		return nil,err
	}
	return pkg.Bytes(),nil
}

func Decode(reader *bufio.Reader) (string, error) {
	msgHeader, _ := reader.Peek(4)
	buf := bytes.NewBuffer(msgHeader)
	var len int32
	err := binary.Read(buf, binary.LittleEndian, &len)
	if err!=nil{
		return "", err
	}
	//如果没有可以读取的字节
	if int32(reader.Buffered())<len+4{
		return "", err
	}
	pack:= make([]byte, int(4+len))
	_,err = reader.Read(pack)
	if err!=nil{
		return "", err
	}
	return string(pack[4:]),nil
}