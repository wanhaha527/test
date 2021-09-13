package main

import (
	"fmt"
)

//实现函数 int sqrt(int x).
//计算并返回x的平方根（向下取整）
var a,b,c=2,2,2
func main() {
	p:=(a+b+c)/2
	s:=sqrt(p*(p-a)*(p-b)*(p-c))
	fmt.Print(s)
	fmt.Print("\n")
	se:=sqrt(26)
	fmt.Print(se)
}
//9=1+3+5
func sqrt( x int ) int {
	// write code here

	i:= 1
	var res = 0
	for x >= 0 {//9 8 5 0 -7
		x -= i  //8 5 0 -7
		res++   //1 2 3 4
		i += 2  //3 5 7 9
	}
	return res - 1
}
/*
1.
9=1+3+5
public class Solution {
    public int sqrt(int x) {
        int i = 1;
        int res = 0;
        while (x >= 0) {
            x -= i;
            res++;
            i += 2;
        }
        return res - 1;
    }
}
2.
public class Solution {
    public int sqrt(int x) {//9
		if (x== 0)
			return 0;
		int left = 1, right = x;//9
		while (true) {
			int mid = left + (right - left) / 2;//5
			//这里判断不用if (mid * mid > x)，因为使用mid > x / mid一定会有结果
			if (mid > x / mid)//5>1
				right = mid - 1;//4 3
			else {
				if(mid+1>x/(mid+1))
					return mid;
				left=mid+1;
			}
		}
	}
}

*/