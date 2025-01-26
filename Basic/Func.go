package main

func plus(a,b int) int{
	return a+b
}

func values()(int,int,int,int){
	return 10,11,12,13
}

func sum(nums ...int) int{
	ttl :=0
	println("nums:",nums[0])
	for _,num := range nums{
		ttl += num
		println(ttl,num)
	}
	return ttl
}

func main(){

	c := plus(0,10)
	println(c)

	_,_,a,b := values()
	println(a,b)
	println("sum:" ,sum([]int{1,2,3,4,5}...))


}
