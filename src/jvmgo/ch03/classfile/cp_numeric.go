package classfile

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}
type ConstantLongInfo struct {
	val int64
}
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader)  {
	bytes :=reader.readUint32()
	self.val =int32(bytes)
}
func (self *ConstantFloatInfo) readInfo(reader *ClassReader)  {
	bytes :=reader.readUint32()
	self.val =float32(bytes)
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader)  {
	bytes :=reader.readUint64()
	self.val =int64(bytes)
}
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader)  {
	bytes :=reader.readUint64()
	self.val =float64(bytes)
}


