package prototype

//原型抽象
type Inode interface {
	Print(string)
	Clone() Inode
}
