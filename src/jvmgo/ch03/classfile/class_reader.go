package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}
//Java 虚拟机规范定义了u1，u2，u4三种数据类型，分别表示1,2,4字节无符号整数
//u1
func (self *ClassReader) readUint8() uint8 {
	val :=self.data[0]
	self.data=self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16  {
	val :=binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val

}
func (self *ClassReader) readUint32() uint32 {
	val :=binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}
func (self *ClassReader) readUint64() uint64 {
	val :=binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}
//读取uint16表，表的大小由开头的uint16数据指出
func (self *ClassReader) readUint16s() []uint16  {
	n :=self.readUint16()
	s :=make([]uint16, n)
	for i :=range s{
		s[i]=self.readUint16()
	}
	return s
}
func (self *ClassReader) readBytes(n uint32) []byte  {
	bytes :=self.data[:n]
	self.data =self.data[n:]
	return bytes
}
