import React from "react";
// import reactLogo from './assets/react.svg'
// import viteLogo from './assets/vite.svg'
// import heroImg from './assets/hero.png'
// import './App.css'

function App() {

  const studentId = "B6741990";
  const Name = "นายบุญค้ำ โยลัย";
  const team = "G08";
  const age = "21";
  
  return (
    <>
      <h1>สวัสดีจ้า</h1>
      <div className="App">
      <header className="App-header">
        <h1>{studentId}, {Name}, {team}, อายุ {age}</h1>
        </header>
      </div>

    </>
  )
}

export default App













// import React from "react";

// // 1. สร้างคอมโพเนนต์ลูกมารอรับของ (สืบทอดค่าผ่าน props)
// function StudentDisplay(props: { id: string; name: string; team: string }) {
//   return (
//     <div style={{ border: "2px solid blue", padding: "10px", margin: "10px" }}>
//       <h3>ข้อมูลที่รับทอดมา:</h3>
//       <p>รหัส: {props.id}</p>
//       <p>ชื่อ: {props.name}</p>
//       <p>ทีม: {props.team}</p>
//     </div>
//   );
// }

// // 2. คอมโพเนนต์แม่ (App) เป็นคนเก็บข้อมูลหลัก
// function App() {
//   const studentId = "B6741990";
//   const studentName = "นายบุญค้ำ โยลัย";
//   const teamNumber = "G08";

//   return (
//     <div className="App" style={{ textAlign: "center", marginTop: "50px" }}>
//       <h1>Welcome to My App</h1>
//       {/* ➔ ส่งทอดข้อมูลลงไปให้คอมโพเนนต์ลูกตรงนี้ */}
//       <StudentDisplay id={studentId} name={studentName} team={teamNumber} />
//     </div>
//   );
// }

// export default App;