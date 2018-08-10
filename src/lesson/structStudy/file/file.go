package file

type File struct {
	fd int//文件描述符
	name string//文件名
}

func NewFile(fd int,name string)*File {
	if fd<0 {
		return nil
	}

	return &File{fd, name}
}
