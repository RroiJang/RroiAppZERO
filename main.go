package main

func main() {
	
}






















// func sum (a, b int) int {
// 	return a + b
// }

// func main(){
// 	fmt.Println("RroiJang")
// 	fmt.Println(sum(1,2))
// }

















// //(Generics) การประยุกต์ใช้กับ object หรือ struct และต้องการเขียนเป็น 1 Function

// //ใช้งานสำหรับตัว struct
// type GameOrMovie interface {
// 	getPrice() int
// }


// type Game struct {
// 	Title string
// 	Platform string
// 	Price int
// }

// type Movie struct {
// 	Title string
// 	Price int
// }

// func(g Game)getPrice() int{
// 	return g.Price
// }

// func(g Movie)getPrice() int{
// 	return g.Price
// }

// func sum[V GameOrMovie](objs []V) int {
// 	var result int
// 	for _, obj := range objs {
// 		result += obj.getPrice()
// 	}
// 	return result
// }

// func main () {
// 	games := []Game {
// 		{
// 			Title:		"Minecraft",
// 			Platform:	"PC",
// 			Price:		30,

// 		},
// 		{
// 			Title:		"The Sims 4",
// 			Platform:	"PC",
// 			Price:		20,
// 		},
// 	}
// 	movies := [] Movie {
// 		{
// 			Title: "Movie1",
// 			Price: 10,
// 		},
// 		{
// 			Title: "Movie2",
// 			Price: 20,
// 		},
// 	}

// 	fmt.Println(sum(games))
// 	fmt.Println(sum(movies))
// }












//Generics
// type Number interface {
// 	int | float64
// }

// func main() {
// 	numsInt := []int{1, 2, 3, 4, 5}
// 	numsFloat64 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

// 	fmt.Println(sum(numsInt))
// 	fmt.Println(sum(numsFloat64))
// }

// func sum [V Number](nums []V) V{ //ผลบวกเราต้องการรับ type ได้ ก็คือ int จำนวนเต็ม | float64 โดยจะรับเป็น int หรือ float64 ก็ได้
// 	var sum V
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }

// func sumInt(nums []int) int {
// 	var sum int
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }

// func sumFloat64(nums []float64) float64 {
// 	var sum float64
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }










// func Debug(obj any){
// 	raw, _ := json.MarshalIndent(&obj, "", "\t")
// 	fmt.Println(string(raw))
// }

// type IBook interface {
// 	GetBook() *book
// 	SetTitle(title string)
// }

// type book struct {
// 	*models.Book
// }

// func NewBook(title, author string) IBook {
// 	return &book{
// 		&models.Book{
// 			Id: 1,
// 			Title: title,
// 			Author: author,
// 		},
// 	}
// }

// func (b *book) GetBook() *book {
// 	return b
// }

// func (b *book) SetTitle(title string) {
// 	b.Title = title
// }

// func main() {
// 	b := NewBook("Golang Basic", "RroiJang")
// 	b.SetTitle("ABC")
// 	Debug(b.GetBook())
// }




	// // a := [3]int{1,2,3}
	// // fmt.Println(a[2])

	// //slice
	// a := make([]int, 0) // a := []int{1,2,3}
	// a = append(a, 4)
	// // a = append(a, 10)

	// fmt.Println(a)

	// arrayWathcer(a)

	// fmt.Println(a)




// func sum(name string, a, b int) (int, string){
// 	fmt.Println(name,"!")
// 	return a + b, "Hello : " + name
// }

// func changer(x int){
// 	x = 20
// }

// func main() {
	// fmt.Println(sum("RroiJang",30, 44))
	// x := "success"
	// // changer(x)
	// fmt.Println(x)

	//while loop
	// for i := 10; i > 5; i-- {
	// 	fmt.Println(i)
	// }
// }

