package dict;
import
(
     "io/ioutil"
     "strconv"
     "strings"
     //"fmt"
)
var dict []byte
var entries [][]byte
var dictPath string
var hmap map[string]int
func check(e error) {
     if e != nil {
          panic(e)
     }
}

func Initi(dictpath string){
     dictPath=dictpath
     d,err:=ioutil.ReadFile(dictPath)
     dict=d
     check(err)
     CalcEntries()
}
func MapLength() int{
     return len(hmap)
}
func MapSet(k string, v int){
     hmap[k]=v
}
func MapGet(k string) int {
     return hmap[k]
}
func ToMap(){
     hmap=make(map[string]int)
     for e:=0;e<len(entries);e++ {
          val:=parseValueFromLine(e)
          vstr:=strconv.Itoa(val)
          nstr:=strings.Trim(string(entries[e]),";"+vstr+"\r\x00\n")
          hmap[string(nstr)]=val
     }
}
func UpdateFile() {
     dict=make([]byte,0)
     for d1:=0;d1<len(entries);d1++ {
          str:=string(entries[d1])
          str=strings.Trim(str,"\x00")
          dict=append(dict,[]byte(str+"\n")...)
     }
     dict=append(dict,'\n')
     ioutil.WriteFile(dictPath,dict,0644)
}
func SetOfKeys() []string{
     fin :=make([]string,len(entries))
     for i:=0;i<len(entries);i++{
          //fmt.Println("ENTRIES",entries[i])
          for b:=0;b<len(entries) && entries[i][b]!=0 && entries[i][b]!=';';b++ {
               fin[i]+=string(entries[i][b])
          }
     }
     return fin
}
func Set(keyValue string,mapValue int) {
     tl:=len(entries)
     l:=tl/2
     p:=0
     up:=false
     for{
          val:=parseValueFromLine(l)
          vstr:=strconv.Itoa(val)
          nstr:=strings.Trim(string(entries[l]),";"+vstr+"\r\x00\n")
          if nstr==keyValue{
               entries[l]=[]byte(nstr+";"+strconv.Itoa(mapValue))
               CalcEntries()
               return
          } else {
               up=compareAlphabetically(keyValue,nstr)
               if up{
                    p=l
                    i:=(tl-l)/2
                    if i==0{
                         i=1
                    }
                    l+=i
               } else {
                    tl=l
                    i2:=(l-p)/2
                    if i2==0{
                         i2=1
                    }
                    l-=i2
               }
          }
     }
}
func Get(keyValue string) int{
     tl:=len(entries)
     l:=tl/2
     p:=0
     up:=false
     for{
          val:=parseValueFromLine(l)
          vstr:=strconv.Itoa(val)
          nstr:=strings.Trim(string(entries[l]),vstr+";\r\x00\n")
          if nstr==keyValue{
               ret:=val
               return ret
          } else {
               up=compareAlphabetically(keyValue,nstr)
               if up{
                    p=l
                    i3:=(tl-l)/2
                    if i3==0{
                         i3=1
                    }
                    l+=i3
               } else {
                    tl=l
                    i4:=(l-p)/2
                    if i4==0{
                         i4=1
                    }
                    l-=i4
               }
          }
     }
}
func parseValueFromLine(i int) int{
     s:=entries[i]
     startRecording:=false
     ret:=make([]byte,0)
     for _,b:=range s{
          if startRecording{
               ret=append(ret,b)
          }
          if b==59||b==';'{
               startRecording=true
          }
     }
     if !startRecording{
          return -1
     }
     v,err:=strconv.Atoi(strings.Trim(string(ret),"\x00\n\r"))
     check(err)
     return v
}
func compareAlphabetically(new string, old string) bool{//is new ahead of or behind? (Only works on lowercase because I'm lazy)
     old=strings.Trim(old,"\n")
     b1:=[]byte(new)
     lb1:=len(b1)
     b2:=[]byte(old)
     lb2:=len(b2)
     least:=0
     larger:=false

     if (lb1>lb2){
          least=lb2
          larger=true
     } else {
          least=lb1
          larger=false
     }
     for check:=0;check<least;check++{
          if int(b1[check])>int(b2[check]) {
               return true
          }
          if int(b1[check])<int(b2[check]) {
               return false
          }
     }
     return larger
}
func CalcEntries(){
     fnumberOfEntries:=0
     for _,b:=range dict{
          if b=='\n'{
               fnumberOfEntries++
          }
     }
     entries=make([][]byte,fnumberOfEntries)
     entries[0]=make([]byte,1024)
     numberOfEntries:=0
     subIterator:=0

     for _,b:=range dict{
          if b=='\n'{
               subIterator=0
               numberOfEntries++
               if numberOfEntries==fnumberOfEntries{
                    return
               }
               entries[numberOfEntries]=make([]byte,1024)
          } else {
               entries[numberOfEntries][subIterator]=b
               subIterator++
          }

     }
     
}
