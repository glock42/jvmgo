package classpath

import "io/ioutil"
import "path/filepath"
import "archive/zip"
import "errors"


type ZipEntry struct {
    absPath string
}

func newZipEntry(path string) *ZipEntry {
    absPath, err := filepath.Abs(path)
    if err!= nil {
        panic(err)
    }
    return &ZipEntry{absPath}
}

func (self *ZipEntry) String() string {
    return self.absPath
}

func (self *ZipEntry) readClass(classname string) ([]byte, Entry, error){
    r, err := zip.OpenReader(self.absPath)
    if err != nil{
        return nil, nil, err
    }

    defer r.Close()

    for _, f := range r.File {
        if f.Name == classname {
            rc, err := f.Open()
            if err != nil {
                return nil, nil, err
            }

            defer rc.Close()

            data, err := ioutil.ReadAll(rc)
            if err != nil {
                return nil, nil, err
            }

            return data, self, nil
        }
    }

    return nil, nil, errors.New("class not found: " + classname)
}