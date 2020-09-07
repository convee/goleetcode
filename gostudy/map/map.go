package map

//注意点
//map 键需要可比较，不能为 slice、map、function
//map 值都有默认值，可以直接操作默认值，如：m[age]++ 值由 0 变为 1
//比较两个 map 需要遍历，其中的 kv 是否相同，因为有默认值关系，所以需要检查 val 和 ok 两个值
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