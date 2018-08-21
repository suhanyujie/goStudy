package file

import (
	"os"
	"errors"
)

/**
我们通常会用到以下标志：

os.O_RDONLY：只读
os.O_WRONLY：只写
os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为0。
在读文件的时候，文件的权限是被忽略的，所以在使用 OpenFile 时传入的第三个参数可以用0。
而在写文件时，不管是 Unix 还是 Windows，都需要使用 0666。
http://wiki.jikexueyuan.com/project/the-way-to-go/12.2.html
 */
func SsWriteFile(filename string, data []byte) error {
	filePointer, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return errors.New(err.Error()+"---[30022]")
	}
	_, err = filePointer.Write(data)
	if err != nil {
		return err
	}
	return nil
}
