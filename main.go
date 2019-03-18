package tools

import "fmt"
import "io/ioutil"
import "io"
import "path"
import "time"
import "os"
import "strings"

func FindImgs(){
	i:=0
	fmt.Println("整理图片开始")
	tstr:=time.Now().Format("2006_01_02_15_04_05")
	folerName:="整理图片"+tstr
	os.Mkdir(folerName,os.ModePerm)
	files,_:=ioutil.ReadDir("./")
	for _,f := range files{
		fname:=f.Name()
		fmt.Println("扫描文件",fname)
		lowname := strings.ToLower(fname)
		if strings.HasSuffix(lowname,".jpg") || 
			strings.HasSuffix(lowname,".gif") || 
				strings.HasSuffix(lowname,".png") || 
					strings.HasSuffix(lowname,".jpeg") {
				fmt.Println("正在整理：",fname)
				copyToFolder(folerName,fname)
				i=i+1
			}
		
	}

	time.Sleep(3*1000*1000*1000)
	fmt.Println("正在整理完毕共",i,"个")
}
func copyToFolder(folderName,file string)(int64, error){
	sFile,err:= os.Open(file)
	defer sFile.Close()
	if err!=nil {
		return 999,err
	}
	dFile,err:=os.OpenFile(path.Join(folderName,file),os.O_WRONLY|os.O_CREATE,0)
	defer dFile.Close()
	if err != nil {
		return 998,err
	}
	return io.Copy(dFile,sFile)
}
