package main

import (
	"fmt"

	"gorm.io/driver/postgres" // ใช้ driver ของ PostgreSQL
	"gorm.io/gorm"            //ดึงมาเพื่อเชื่อมกับ sql
)

// โครงสร้างa
type Student struct {
	gorm.Model
	StudentID string
	Name      string
	Team      string
	Age       int
}

func main() {
	// เชื่อมต่อเข้ากับ sql ที่รันอยู่ใน Dockers
	dsn := "host=localhost user=postgres password=postgres dbname=golangdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("เชื่อมไม่ได้") //เเพนิคเตือนว่ามันเชื่อมไม่ได้
	}

	db.AutoMigrate(&Student{}) //สร้างตารางอัตโนมัติถ้าไม่มีอยู่แล้ว
	myInfo := Student{
		StudentID: "B6741990",
		Name:      "นายบุญค้ำ โยลัย",
		Team:      "G09",
		Age:       20,
	}

	db.Create(&myInfo)
	fmt.Println("บันทึกแล้ว!")
}

// var existingStudent Student

// err = db.Where("student_id = ?", myInfo.StudentID).First(&existingStudent).Error //เอา column เช็คว่ามีมั้ย
// if err != nil {

// 	 //insert ข้อมูลลง sql

// } else {

// 	fmt.Println("มีข้อมูลแล้ว")
// }
// --------------------------------------------------------------------------------
// --------------------------------------------------------------------------------
// --------------------------------------------------------------------------------
// --------------------------------------------------------------------------------
// --------------------------------------------------------------------------------

// บันทึกลง sql
// var existingStudent Student //มีข้อมูลในฐานข้อมูลไหม
// if err := db.Where("student_id = ?", myInfo.StudentID).First(&existingStudent).Error; err != nil {
// 	db.Create(&myInfo)
// 	fmt.Println("บันทึกแล้ว!")
// } else {
// 	fmt.Println("มีข้อมูลแล้ว")
// }

// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // 1. กำหนดโครงสร้างตารางข้อมูลนักศึกษา [cite: 797, 798]
// type Student struct {
// 	gorm.Model
// 	StudentID string `gorm:"unique;not null"`
// 	Name      string
// 	Team      string
// 	Age       int
// }

// func main() {
// 	// 2. เชื่อมต่อฐานข้อมูล PostgreSQL (ตั้งค่าตาม docker-compose.yml) [cite: 808, 809]
// 	dsn := "host=localhost user=postgres password=postgres dbname=golangdb port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("เชื่อมต่อฐานข้อมูลล้มเหลว!")
// 	}

// 	// 3. สั่งให้ GORM สร้างตารางให้อัตโนมัติ (ตารางจะชื่อ students) [cite: 813]
// 	db.AutoMigrate(&Student{})

// 	// 4. เตรียมข้อมูลตัวตนของคุณ [cite: 807]
// 	myInfo := Student{
// 		StudentID: "B6741990",
// 		Name:      "นายบุญค้ำ โยลัย",
// 		Team:      "G01",
// 		Age:       21,
// 	}

// 	// 5. บันทึกข้อมูลแบบสั้น (ถ้าไม่มีให้สร้างใหม่ทันที)
// 	db.FirstOrCreate(&myInfo, Student{StudentID: myInfo.StudentID})

// 	fmt.Println("➔ [Done] ตรวจสอบและจัดการข้อมูลนักศึกษาใน Database สำเร็จแล้ว!")
// }

// // package main

// // import (
// // 	"fmt"

// // 	"gorm.io/driver/postgres"
// // 	"gorm.io/gorm"
// // )

// // type Student struct {
// // 	gorm.Model
// // 	StudentID string `gorm:"unique;not null"`
// // 	Name      string
// // 	Team      string
// // 	Age       int
// // }

// // func main() {
// // 	// 1. เชื่อมต่อฐานข้อมูลปกติ
// // 	dsn := "host=localhost user=postgres password=postgres dbname=golangdb port=5432 sslmode=disable"
// // 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// // 	if err != nil {
// // 		panic("เชื่อมต่อฐานข้อมูลล้มเหลว!")
// // 	}
// // 	db.AutoMigrate(&Student{})

// // 	// 2. ประกาศตัวแปรเปล่าเพื่อมารองรับ Input จากคีย์บอร์ด
// // 	var inputID string
// // 	var inputName string
// // 	var inputTeam string
// // 	var inputAge int

// // 	// 3. พิมพ์ข้อความบอกผู้ใช้ และรับค่าเข้ามาทีละตัว
// // 	fmt.Print("กรุณากรอกรหัสนักศึกษา (เช่น B67XXXXX): ")
// // 	fmt.Scanln(&inputID) // รับค่าตัวอักษรจนกว่าจะกด Enter

// // 	fmt.Print("กรุณากรอกชื่อ-นามสกุล: ")
// // 	// หมายเหตุ: fmt.Scanln จะมองว่าการเคาะเว้นวรรคคือการจบตัวแปร
// // 	// หากต้องการกรอกชื่อแบบมีเว้นวรรค แนะนำให้พิมพ์ติดกัน หรือใช้ fmt.Scanf("%s", &inputName)
// // 	fmt.Scanln(&inputName)

// // 	fmt.Print("กรุณากรอกหมายเลขทีม (เช่น G01): ")
// // 	fmt.Scanln(&inputTeam)

// // 	fmt.Print("กรุณากรอกอายุ: ")
// // 	fmt.Scanln(&inputAge)

// // 	// 4. นำตัวแปรที่ได้จากคีย์บอร์ดมาจัดใส่ Struct
// // 	myInfo := Student{
// // 		StudentID: inputID,
// // 		Name:      inputName,
// // 		Team:      inputTeam,
// // 		Age:       inputAge,
// // 	}

// // 	// 5. บันทึกลงฐานข้อมูล
// // 	db.FirstOrCreate(&myInfo, Student{StudentID: myInfo.StudentID})

// // 	fmt.Println("\n➔ [Done] บันทึกข้อมูลที่คุณกรอกลง Database สำเร็จแล้ว!")
// // }

// --------------------------------------------------------------------------------
