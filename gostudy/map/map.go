package map

func main() {
	//创建
	m := make(map[string]int)
	//设置 kv
	m["hello"] = 1
	//删除k
	delete(m,"hello")
	//遍历
	for k,v:=range m{
		fmt.Println(k,v)
	}
}