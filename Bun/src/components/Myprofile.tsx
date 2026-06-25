import React from "react";

// 1. สร้าง Interface เพื่อกำหนดโครงสร้างและชนิดข้อมูลของ student
interface Student {
  name: string;
  age: number;
  major: string;
  img1: string;
}

// 2. สร้าง Interface สำหรับ Props ของคอมโพเนนต์ Myprofile
interface MyProfileProps {
  data: Student; 
}
// 3. กำหนด Type ให้กับ props และใช้ destructuring เพื่อดึง data ออกมา
function Myprofile({ data }: MyProfileProps) {
  return (
    <div>
      <p>{data.name}</p>
      <p>{data.major}</p>
      
       
     
    </div>
  );
}

export default Myprofile;