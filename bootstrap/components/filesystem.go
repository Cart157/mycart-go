package components

import (
    "os"
    "io/ioutil"
)

// ConfigFacade有点小用，可已缓存所有配置，不过也可以不缓存
type FileSystemFacade string

var FileSystem *FileSystemFacade

func init() {
    FileSystem = new(FileSystemFacade)
}

func (fs *FileSystemFacade) PathExists(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }

    if os.IsNotExist(err) {
        return false
    }

    panic(err)
}

func (fs *FileSystemFacade) ReadAll(filePath string) ([]byte, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }


    return ioutil.ReadAll(f)
}
