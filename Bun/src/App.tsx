import { useState } from "react"; //เอาไว้แสดง state
import "./App.css";
import Myprofile from "./components/Myprofile.tsx";

function App() {
  const [num, setNum] = useState(0);

  // const user = {
  //     name : "bun",
  //     age :  20,
  //     major : "CPE",
  //     img1: "https://images.unsplash.com/photo-1773332585754-f1436987743b?q=80&w=1170&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
  // }

  const student = {
    // คือ object
    name: "Bunkham yolai", //ใน object จะมี key และ value
    age: 20,
    major: "Computer Science",
    img1: "https://images.unsplash.com/photo-1773332585754-f1436987743b?q=80&w=1170&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  };
    

  const products = [
    { title: "Cabbage", id: 1 },
    { title: "Garlic", id: 2 },
    { title: "Apple", id: 3 },
  ];

  function inc() {
    setNum(num + 1);
  }
  function dec() {
    setNum(num - 1);
  }


  return (
    <>
      <Myprofile data = {student} />
      <p>Counter: {num}</p>
      <button onClick={inc}>เพิ่ม</button>
      <button onClick={dec}>ลด</button>
    </>
  );
}
export default App;
