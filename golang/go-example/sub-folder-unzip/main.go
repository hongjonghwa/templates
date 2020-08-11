package main
import (
    "archive/zip"
    "path/filepath"
    "io"
    "fmt"
    "os"
    "strings"
    "os/exec"
)

func printBytes(s string) {  
    fmt.Printf("Bytes: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func ExternalUnzip(src, dest string) error {
    
    // commandString := fmt.Sprintf(`busybox.exe unzip %s -d %s`, src, dest) // busybox
    // fmt.Println(commandString)   
    // commandSlice := strings.Fields(commandString)
    // c := exec.Command(commandSlice[0], commandSlice[1:]...)

    // commandString := fmt.Sprintf(`"C:\Program Files\7-Zip\7za.exe" e %s %s`, src, dest) // 7-zip
    c := exec.Command("cmd", "/C", "start", `C:\Program Files\7-Zip\7za.exe`, "e", src, dest)
    e := c.Run()
    if e != nil{
        return e
    }
    return nil
}

func Unzip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer func() {
        if err := r.Close(); err != nil {
            panic(err)
        }
    }()

    os.MkdirAll(dest, 0755)

    // Closure to address file descriptors issue with all the deferred .Close() methods
    extractAndWriteFile := func(f *zip.File) error {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer func() {
            if err := rc.Close(); err != nil {
                panic(err)
            }
        }()

        path := filepath.Join(dest, f.Name)
        fmt.Println(path)
        // printBytes(path)
        fmt.Printf("%+q\n", path)

        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            os.MkdirAll(filepath.Dir(path), f.Mode())
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer func() {
                if err := f.Close(); err != nil {
                    panic(err)
                }
            }()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
        return nil
    }

    for _, f := range r.File {
        err := extractAndWriteFile(f)
        if err != nil {
            return err
        }
    }

    return nil
}

func getAllSubDirs(root string) []string {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        // fmt.Println(info.Name())
        if info.IsDir() && info.Name() != "." && info.Name() != ".." {
            files = append(files, path)
        }
        return nil
    })
    if err != nil {
        panic(err)
    }
    return files
}


func getDirs(root string) []string {
    files, err := filepath.Glob("*")
    if err != nil {
        panic(err)
    }
    var dirs []string
    for _, f := range files {
        fi, err := os.Stat(f)
        if err != nil {
            fmt.Println(err)
        }
        if fi.Mode().IsDir(){
            dirs = append(dirs, f)
        }
    }

    return dirs
}


func getZips(root string) []string {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        // fmt.Println(info.Name())
        if filepath.Ext(path) == ".zip" {
            files = append(files, path)
        }
        return nil
    })
    if err != nil {
        panic(err)
    }
    return files
}

func main() {

    fmt.Println()
    fmt.Println("### 첫 번째 서브 폴더에 존재하는 모든 ZIP 파일의 압축을 풀어줍니다.###")
    fmt.Println()

    fmt.Print("(엔터) 키를 누르면 시작합니다! ")
    var input string
    fmt.Scanln(&input)


    fmt.Println()
    fmt.Println()

    dirs := getDirs(".");
    for _, d := range dirs {

        zips := getZips(d)
        fmt.Printf("▷ \"%v\" (zip 파일 수 %v개)\n", d, len(zips))

        for _, z := range zips {

            // fmt.Println(z)
            // fmt.Println(filepath.Base(z))
            extract_folder_name := strings.TrimSuffix(filepath.Base(z), ".zip")
            // fmt.Println(extract_folder_name)
            dest_folder := filepath.Join(d,extract_folder_name)
            // fmt.Println(dest_folder)
            if _, err := os.Stat(dest_folder); os.IsNotExist(err) {
                
                // err := Unzip(z, dest_folder);
                err := ExternalUnzip(z, dest_folder);
                if err != nil {
                    fmt.Printf(" └  Error 압축 해제에 실패하였습니다.")
                    panic(err)
                } 
                fmt.Printf(" └  OK \"%v\"\n", filepath.Base(z))
            }else{
                fmt.Printf(" └  Error \"%v\" 폴더가 존재합니다.\n", extract_folder_name)
            }
        }
        fmt.Println()
    }

    fmt.Printf("Good Bye !\n")
    fmt.Scanf("h")
	
}