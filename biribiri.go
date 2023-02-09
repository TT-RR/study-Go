package main

import "fmt"

type SYOHIN struct {
	price int
	name  string
}

func main() {
	//変数の宣言
	v := SYOHIN{120, "リンゴジュース"}
	var money, no int
	nextBuyFlag := "yes"

	//所持金を入力・表示
	fmt.Println("所持金を入力してください:")
	fmt.Scan(&money)
	fmt.Printf("所持金[%d円]\n", money)

	//買い物をやめるまで続ける
	if nextBuyFlag != "no" {
		fmt.Println("\n[No][値段][商品名]")
		fmt.Println("-------------------------")

		//商品の表示
		for i := 0; i < 5; i++ {
			fmt.Println(i+1, v.price, v.name)
		}

		//商品の購入・残金の計算
		fmt.Println("\n購入したい商品のNoを入力してください:")
		fmt.Scan(&no)
		money = money - v.price
		fmt.Printf("\n[%s]のお買い上げありがとうございます。", v.name)
		fmt.Printf("\n残金[%d円]", money)

		//買い物を続けるか聞く
		fmt.Println("\n買い物を続けますか？:")
		fmt.Scan(&nextBuyFlag)
	}
	fmt.Println("お買い上げありがとうございました。")
}
