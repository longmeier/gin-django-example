package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// SliceStrFindEle 切片字符串找元素  "222" in ["222"]
func SliceStrFindEle(sli *[]string, val string) (int, bool) {
	for i, item := range *sli {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// SliceIntFindEle 切片字符串找元素   222 in [222]
func SliceIntFindEle(sli *[]int, val int) (int, bool) {
	for i, item := range *sli {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// StringAppendString 字符串拼接  +,
func StringAppendString(msg, apd string) string {
	if msg == "" {
		msg = apd
	} else {
		msg += "," + apd
	}
	return msg
}

// StringSliceJoinStr slice->字符串  [1,2,3]->"1","2","3"
func StringSliceJoinStr(arr []string) (ret string) {
	arrLen := len(arr)
	for i, col := range arr {
		if i+1 == arrLen { // 如果是最后一个不加,
			ret += fmt.Sprintf(`"%v"`, col)
		} else {
			ret += fmt.Sprintf(`"%v",`, col)
		}
	}
	return
}

// PyArrStrTsfSlice python数组字符串转slice  ['111','222'] -> ["111" "222"]
func PyArrStrTsfSlice(str string) (sli []string) {
	str1 := strings.Replace(str, "[", "", -1)
	str2 := strings.Replace(str1, "]", "", -1)
	str3 := strings.Replace(str2, "'", "\"", -1)
	sli = strings.Split(str3, ",")
	return
}

// SliceParamsToStr 数据库切片参数转字符串 {"111" "222"} ->  ?,?
func SliceParamsToStr(ids []string) ([]interface{}, string) {
	inIds := ""
	params := make([]interface{}, 0)
	for i := 0; i < len(ids); i++ {
		if i == 0 {
			inIds += "?"
		} else {
			inIds += ",?"
		}
		params = append(params, ids[i])
	}
	return params, inIds
}

// SliceStringToStrInSQL 切片参数转字符串 {"111" "222"} ->  111,222
func SliceStringToStrInSQL(str []string) (ret string) {
	ret = ""
	for i, s := range str {
		if i == 0 {
			ret += fmt.Sprintf(`"%v"`, s)
		} else {
			ret += fmt.Sprintf(`,"%v"`, s)
		}
	}
	return ret
}

// SliceStringToStr 切片参数转字符串 {"111" "222"} ->  111,222
func SliceStringToStr(str []string) (ret string) {
	ret = ""
	for i, s := range str {
		if i == 0 {
			ret += s
		} else {
			ret += "," + s
		}
	}
	return ret
}

// SliceIntToStr 切片参数转字符串 {111, 222} ->  111,222
func SliceIntToStr(str []int) (ret string) {
	ret = ""
	for i, s := range str {
		if i == 0 {
			ret += strconv.Itoa(s)
		} else {
			ret += "," + strconv.Itoa(s)
		}
	}
	return ret
}

// PyArrJsonStrTsfSlice python数组json字符串转slice [{"id":1}, {"id":2}] -> [{"id":1}, {"id":2}]
func PyArrJsonStrTsfSlice(jsonStr string) (arrSlice []map[string]string, err error) {
	jStr := strings.Replace(jsonStr, "\\", "", -1)
	jStr = strings.Replace(jStr, "'", "\"", -1)
	err = json.Unmarshal([]byte(jStr), &arrSlice)
	return
}

// PyArrJsonStrTsfString python数组json字符串转slice [{"id":1}, {"id":2}] -> [{"id":1}, {"id":2}]
func PyArrJsonStrTsfString(jsonStr string) (res string, err error) {
	jStr := strings.Replace(jsonStr, "\\", "", -1)
	res1 := strings.Replace(jStr, "'", "\"", -1)
	res2 := strings.Replace(res1, "\"None\"", "\"\"", -1)
	res = strings.Replace(res2, "None", "\"\"", -1)
	return
}

// PyArrJsonStrTsfInterface python数组json字符串转interface [{"id":1}, {"id":2}] -> [{"id":1}, {"id":2}]
func PyArrJsonStrTsfInterface(jsonStr string) (arrInter []interface{}) {
	jStr := strings.Replace(jsonStr, "'", "\"", -1)
	err := json.Unmarshal([]byte(jStr), &arrInter)
	if err != nil {
		fmt.Printf("PyArrJsonStrTsfSlice failed, err:%v\n", err)
	}
	return
}

// FloatTsfDecimal 四舍五入保留两位小数
func FloatTsfDecimal(f float64) float64 {
	d, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
	return d
}

// FloatTsfString 四舍五入保留两位小数
func FloatTsfString(f float64) string {
	d := fmt.Sprintf("%.2f", f)
	return d
}

// StringTsfFloat  字符串转float
func StringTsfFloat(f string) float64 {
	d, _ := strconv.ParseFloat(f, 64)
	return d
}

// StringIsExist  判断元素是否存在字符串中  bool
func StringIsExist(strTxt, ele string) bool {

	if ele != "" && strTxt != "" && strings.Contains(strTxt, ele) {
		return true
	}
	return false
}

// EqualAmt 均算 amtR / tAmtR * amt   num 当前第几次   tNum 总数量
func EqualAmt(amtR, tAmtR, amt, oddAmt float64, currentNum, tNum int) (rAmt, odAmt float64) {
	rt := float64(0)
	if tAmtR > 0 {
		if amtR != 0 {
			if currentNum == tNum { // 如果是最后一次
				if amt != 0 {
					rt = FloatTsfDecimal(amt - oddAmt)
				} else {
					rt = 0
				}
			} else {
				rt = FloatTsfDecimal(amtR / tAmtR * amt)
			}
		}
		oddAmt += rt
	} else {
		oddAmt = 0
	}
	rAmt = FloatTsfDecimal(rt)
	odAmt = FloatTsfDecimal(oddAmt)
	return
}

// EqualAmtRod 均算取整     (amtR / tAmtR) * amt   num 当前第几次   tNum 总数量
func EqualAmtRod(amtR, tAmtR, amt, oddAmt float64, currentNum, tNum int) (rAmt, odAmt int) {
	rt := float64(0)
	if tAmtR > 0 {
		if currentNum == tNum { // 如果是最后一次
			if amt != 0 {
				rt = amt - oddAmt
			} else {
				rt = 0
			}
		} else {
			rt = FloatTsfDecimal(amtR / tAmtR * amt)
		}
		oddAmt += rt
	} else {
		oddAmt = 0
	}
	rAmt = int(rt)
	odAmt = int(oddAmt)
	return
}

// DBStructParamsTsfStr 数据库struct 转字符串
func DBStructParamsTsfStr(par interface{}) (str string) {
	params := reflect.ValueOf(par)
	str = ""
	for i := 0; i < params.NumField(); i++ {
		field := params.Field(i)
		value := fmt.Sprintf("%v", field)
		if value == "" {
			value = "null"
		} else {
			fmt.Printf("%d不为空:%v\n", i, value)
		}
		if value == "null" {
			if i == 0 {
				str += fmt.Sprintf("%v", value)
			} else {
				str += fmt.Sprintf(",%v", value)
			}
		} else {
			if i == 0 {
				str += fmt.Sprintf("'%v'", value)
			} else {
				str += fmt.Sprintf(",'%v'", value)
			}
		}

	}
	return str
}

// GetNextMonthTime 获取下几个月时间
func GetNextMonthTime(n int) (ret time.Time) {
	ret = GetNowTime().AddDate(0, n, 0)
	return
}

func GetNextDayTime(n int) (ret time.Time) {
	ret = GetNowTime().AddDate(0, 0, n)
	return
}

// DBCommaGenerateQ 根据sql语句逗号生成带参数问号语句
func DBCommaGenerateQ(sqlStr string) (ret string) {
	sqlSlice := strings.Split(sqlStr, ",")
	num := len(sqlSlice)
	qStr := ""
	for i := 0; i < num; i++ {
		if i == 0 {
			qStr += "?"
		} else {
			qStr += ",?"
		}
	}
	ret = fmt.Sprintf(sqlStr, qStr)
	return
}

// DBCommaGenerateQOfIn 根据sql语句逗号生成带参数问号语句  in "xxx","xxx"
func DBCommaGenerateQOfIn(arr []string) (sql string) {
	sql = ""
	for i, ar := range arr {
		if i == 0 {
			sql += fmt.Sprintf(`"%s"`, ar)
		} else {
			sql += fmt.Sprintf(`,"%s"`, ar)
		}
	}
	return
}

// DBCommaGenerateQIntOfIn 根据sql语句逗号生成带参数问号语句  in "xxx","xxx"
func DBCommaGenerateQIntOfIn(arr []int) (sql string) {
	sql = ""
	for i, ar := range arr {
		if i == 0 {
			sql += fmt.Sprintf(`%v`, ar)
		} else {
			sql += fmt.Sprintf(`,%v`, ar)
		}
	}
	return
}

// DBCommaGenerateQByLen 根据sql语句逗号生成带参数问号语句
func DBCommaGenerateQByLen(num int) (ret string) {
	qStr := ""
	for i := 0; i < num; i++ {
		if i == 0 {
			qStr += "?"
		} else {
			qStr += ",?"
		}
	}
	ret = qStr
	return
}

// DBEmptyStrTsfNullStr 字段空字符串转 "null"
func DBEmptyStrTsfNullStr(str string) (ret sql.NullString) {
	if str == "" || str == "0" || len(str) == 0 {
		ret = sql.NullString{}
		return
	}
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

// DBEmptyStrTsfNullInt 字段空字符串转 "0"
func DBEmptyStrTsfNullInt(str string) (ret int) {
	if str == "" || str == "0" || len(str) == 0 {
		ret = 0
		return
	}
	ret, _ = strconv.Atoi(str)
	return
}

// MapStringTsfStr map接口转字符串
func MapStringTsfStr(param []map[string]string) string {
	dataType, _ := json.Marshal(param)
	dataStr := string(dataType)
	return dataStr
}

// TPFunDateTimeStringFormat 时间格式化  字符串格式化----结果字符串
func TPFunDateTimeStringFormat(timeValue string, fmt string) string {
	timeLayout := "2006-01-02T15:04:05Z"   //所需模板
	loc, err := time.LoadLocation("Local") //***获取时区***
	if err != nil {
		loc = time.FixedZone("CST-8", 8*3600)
	}
	theTime, _ := time.ParseInLocation(timeLayout, timeValue, loc) //使用模板在对应时区转化为time.time类型

	// 0001-01-01T00:00:00Z这里也表示时间为null
	if theTime.IsZero() {
		return ""
	} else {
		//时间戳转日期
		//dataTimeStr := theTime.Format("2006-01-02 15:04:05") //使用模板格式化为日期字符串
		dataTimeStr := theTime.Format(fmt) //使用模板格式化为日期字符串
		return dataTimeStr
	}
}

// TPFuncStringDateTime 字符串转时间  2006-01-02T15:04:05Z
func TPFuncStringDateTime(timeValue string, fmt string) time.Time {
	if fmt == "" {
		fmt = "2006-01-02T15:04:05Z" //所需模板
	}

	loc, err := time.LoadLocation("Asia/Shanghai") //***获取时区***
	if err != nil {
		loc = time.FixedZone("CST-8", 8*3600)
	}
	theTime, _ := time.ParseInLocation(fmt, timeValue, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

// DiffTimeObject 时间相差
func DiffTimeObject(nowTime, lastTime time.Time) time.Duration {
	return nowTime.Sub(lastTime)
}

// GetNowTimeFormat 获取时区当前时间 2006-01-02 15:04:05
func GetNowTimeFormat(format string) (ret string) {
	//go语言的time.Now()返回的是当地时区时间
	//但是部署之后，有的服务器会默认使用世界标准时间（UTC），所以需要主动设置一下时区
	var cstSh, err = time.LoadLocation("Asia/Shanghai") //上海
	// 因为docker 没有Asia/Shanghai会提示；time: missing Location in call to Date
	//  centos、ubuntu 都存放了/usr/share/zoneinfo/下的时区，alpine镜像没有
	if err != nil {
		cstSh = time.FixedZone("CST-8", 8*3600)
	}
	ret = time.Now().In(cstSh).Format(format)
	return
}

// GetNowTime 获取时区当前时间
func GetNowTime() (ret time.Time) {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		timeLocal = time.FixedZone("CST-8", 8*3600)
	}
	time.Local = timeLocal
	ret = time.Now().Local()
	return
}

// map 转 string
func MapToString(mapObj map[string]string) string {
	data, _ := json.Marshal(mapObj)
	dataString := string(data)
	return dataString
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// 获取执行路径
func GetCurrentAbPath() string {
	//dir := getCurrentAbPathByExecutable()
	//tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	//if strings.Contains(dir, tmpDir) {
	//	return getCurrentAbPathByCaller()
	//}
	//return dir
	return fmt.Sprintf("%v/src/paycenter", os.Getenv("GOPATH"))
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 获取上级路径
func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func ProjectPath() (path string) {
	var (
		sp = "/"
		ss []string
	)
	if runtime.GOOS == "windows" {
		sp = "\\"
	}

	stdout, _ := exec.Command("go", "env", "GOMOD").Output()
	fmt.Println(stdout)
	path = string(bytes.TrimSpace(stdout))
	if path != "" {
		ss = strings.Split(path, sp)
		ss = ss[:len(ss)-1]
		path = strings.Join(ss, sp)
		return
	}
	fileDir, _ := os.Getwd()
	path = os.Getenv("GOPATH")
	ss = strings.Split(fileDir, path)
	log.Print(fmt.Sprintf("%v-%v", fileDir, path))
	if path != "" {
		ss2 := strings.Split(ss[1], sp)
		path += sp
		for i := 1; i < len(ss2); i++ {
			path += ss2[i]
			if Exists(path) {
				return path
			}
		}
	}
	return
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 删除字符串前的某字符
func DeleteFrontSpace(str string, deleteStr string) string {
	var end_str string
	var number int = 0
	temp := strings.Split(str, "")
	for index, value := range temp {
		if value != deleteStr {
			number = index
			break
		}
	}
	for i := number; i < len(str); i++ {
		end_str += temp[i]
	}

	return end_str
}

// Map -> struct
func Map2Struct(map_data interface{}, struct_data interface{}) error {
	tempData, err := json.Marshal(map_data)
	if err != nil {
		return err
	}
	err2 := json.Unmarshal([]byte(tempData), &struct_data)
	if err2 != nil {
		return err2
	}
	return nil
}

func TypeToFlag(proType string) (proFlag int) {
	if proType != "combo" { // 非套餐
		// 更新项目收费状态
		switch proType {
		case "exam":
			proFlag = 1
		case "goods":
			proFlag = 5
		case "treat":
			proFlag = 6
		case "addfee":
			proFlag = 7
		case "medicine":
			proFlag = 4
		case "coupon_pack": // 券包
			proFlag = 9
		case "member_card": // 会员卡
			proFlag = 10
		case "swap_card": // 会员卡
			proFlag = 11
		default:
			proFlag = 2
		}
	} else {
		proFlag = 3
	}
	return
}

func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			if strings.HasSuffix(tagValue, ",omitempty") {
				tagValue = tagValue[0 : len(tagValue)-10]
			}
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

// CategroyValueToCluster 商品分类-归类
func CategroyValueToCluster(values string) int {
	switch values {
	case "1":
		return 0
	case "2":
		return 1
	case "3":
		return 1
	case "4":
		return 4
	case "5":
		return 2
	case "6":
		return 1
	case "7":
		return 3
	case "8":
		return 3
	case "9":
		return 3
	case "10":
		return 3
	case "11":
		return 0
	case "12":
		return 0
	default:
		return 0
	}
}

func GetZeroPayMethodList() []map[string]string {
	return []map[string]string{
		{"name": "现金", "english_name": "xj", "amt": "0.00"},
	}
}

func GetAvailablePort() (int, error) {
	// 使用 net 包中的 Listen 函数监听一个随机端口
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	// 获取监听的端口
	address := listener.Addr().(*net.TCPAddr)
	return address.Port, nil
}

// AppendRemoveDuplicates 追加去重
func AppendRemoveDuplicates(arr *[]string, item string) {
	if item == "" {
		return
	}
	itemArry := strings.Split(item, ",")
	for _, col := range itemArry {
		_, isOk := SliceStrFindEle(arr, col)
		if !isOk {
			*arr = append(*arr, col)
		}
	}
	return
}

// AppendRemoveDuplicatesOfSlice 追加去重
func AppendRemoveDuplicatesOfSlice(arr, item *[]string) {
	if len(*item) == 0 {
		return
	}
	for _, it := range *item {
		_, isOk := SliceStrFindEle(arr, it)
		if !isOk {
			*arr = append(*arr, it)
		}
	}
	return
}

// RemoveSliceDuplicates 切片去重
func RemoveSliceDuplicates(arr []string) []string {
	uniqueMap := make(map[string]bool)
	result := []string{}
	for _, str := range arr {
		// 如果映射中没有这个字符串，将其添加到结果数组，并在映射中标记为true
		if !uniqueMap[str] && str != "" {
			uniqueMap[str] = true
			result = append(result, str)
		}
	}
	return result
}

func RemoveEmptyStrings(input []string) []string {
	var result []string
	for _, str := range input {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// PayMethodAppendMerge 支付方式追加合并
func PayMethodAppendMerge(arr, brr string) (res string, err error) {
	arrObj, err := PyArrJsonStrTsfSlice(arr)
	if err != nil {
		return
	}
	brrObj, err := PyArrJsonStrTsfSlice(brr)
	if err != nil {
		return
	}
	resObj := make([]map[string]string, 0)
	resObj = append(resObj, brrObj...)
	for _, ao := range arrObj {
		flag := false
		aoAmt := StringTsfFloat(ao["amt"])
		for _, bo := range resObj {
			boAmt := StringTsfFloat(bo["amt"])
			if ao["english_name"] == bo["english_name"] {
				boAmt += aoAmt
				boAmt = FloatTsfDecimal(boAmt)
				bo["amt"] = FloatTsfString(boAmt)
				flag = true
				break
			}
		}
		if !flag {
			resObj = append(resObj, ao)
		}
	}
	resByte, _ := json.Marshal(resObj)
	res = string(resByte)
	return
}

// PayMethodDeduct 支付方式扣除
func PayMethodDeduct(originArr, deductArr string) (res string, err error) {
	originObj, err := PyArrJsonStrTsfSlice(originArr)
	if err != nil {
		return
	}
	deductObj, err := PyArrJsonStrTsfSlice(deductArr)
	if err != nil {
		return
	}
	resObj := make([]map[string]string, 0)
	for _, ar := range originObj {
		aoAmt := StringTsfFloat(ar["amt"])
		for _, dr := range deductObj {
			doAmt := StringTsfFloat(dr["amt"])
			if dr["english_name"] == ar["english_name"] {
				aoAmt -= doAmt
				aoAmt = FloatTsfDecimal(aoAmt)
				ar["amt"] = FloatTsfString(aoAmt)
				break
			}
		}
		if aoAmt > 0 {
			resObj = append(resObj, ar)
		}
	}
	if len(resObj) == 0 {
		resObj = GetZeroPayMethodList()
	}
	resByte, _ := json.Marshal(resObj)
	res = string(resByte)
	return
}

// VerifyNowDateIsValidity 验证时间是否有效
func VerifyNowDateIsValidity(startDate, endDate string) (res bool) {
	nowDate := GetNowTimeFormat("2006-01-02")
	res = false
	startFlag, endFlag := true, true
	if startDate != "" {
		startFlag = false
		if startDate <= nowDate {
			startFlag = true
		}
	}
	if endDate != "" {
		endFlag = false
		if nowDate <= endDate {
			endFlag = true
		}
	}
	if startFlag && endFlag {
		res = true
	}
	return res
}
