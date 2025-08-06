package main

import "fmt"

func main ()  {
	

// name := map  [type] type
//                key  value  
var studentID = map[string]int{
"Hung":1,    //need "," for new key:value
"Nam":2,
"Hai":3,
"Cuong":4,
"Vu":5,
"Long":6,
}

// lấy giá trị của key "Hung"
idHung := studentID["Hung"]
   //%s is string placeholder, %d is decimal integer
  // %s = "Hung"
 //  %d = idHung
fmt.Printf("The ID of %s is %d \n", "Hung", idHung)    //\n xuống dòng

/*Go hiểu chuỗi "Hung" điền vào %s vì thứ tự đối số và quy tắc khớp placeholder của hàm fmt.Printf.
Chuỗi định dạng đầu tiên: fmt.Printf luôn mong đợi đối số đầu tiên là chuỗi định dạng ("The ID of %s is %d").
Đối số tiếp theo theo thứ tự: Sau chuỗi định dạng, 
các đối số còn lại được coi là các giá trị sẽ được điền vào các placeholder theo đúng thứ tự chúng xuất hiện.
%s là placeholder đầu tiên trong chuỗi định dạng.
"Hung" là đối số đầu tiên sau chuỗi định dạng.
Go sẽ khớp %s với "Hung".
Nói cách khác, Go không "hiểu" dựa trên tên biến hay bất cứ ngữ cảnh đặc biệt nào 
ngoài việc khớp tuần tự các placeholder trong chuỗi định dạng với các đối số được cung cấp.*/


// à ra vậy truyền key vào biến rồi lấy giá trị của biến




studentName := "Nam"
idNam:= studentID[studentName]
fmt.Printf("The ID of %s is %d", studentName, idNam)   //dùng printf để định dạng

}