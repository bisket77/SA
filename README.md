# 📌 สรุปปัญหา Git และวิธีแก้

## 1. Git ถูกสร้างผิดที่

### อาการ

```bash
git status
```

แสดงไฟล์ทั้งเครื่อง เช่น

```text
Documents/
Downloads/
Pictures/
AppData/
```

### วิธีแก้

```powershell
cd "C:\Users\Asus TUF"
Remove-Item -Recurse -Force .git
```

จากนั้นเข้าโปรเจกต์จริงแล้วสร้าง Git ใหม่

```bash
git init
```

---

## 2. Push ไม่ได้ (fetch first)

### อาการ

```text
! [rejected] main -> main (fetch first)
```

### สาเหตุ

GitHub มี Commit อยู่แล้ว แต่ Git ในเครื่องเป็นคนละประวัติ

### ตรวจสอบ

```bash
git remote -v
git log --oneline -n 5
```

### ถ้า GitHub มีแค่ README

```bash
git push -u origin main --force
```

> ⚠️ คำสั่งนี้จะเขียนทับข้อมูลบน GitHub

### ถ้าต้องการเก็บข้อมูลเดิมบน GitHub

```bash
git pull origin main --allow-unrelated-histories
git push origin main
```

---

## 3. อัปโหลดโปรเจกต์ขึ้น GitHub

เข้าโฟลเดอร์โปรเจกต์

```bash
cd "C:\Users\Asus TUF\Desktop\STUDENT 3\TERM 1\SA"
```

ตรวจสอบ Git Repository

```bash
git rev-parse --show-toplevel
```

ถ้าขึ้น

```text
C:/Users/Asus TUF/Desktop/STUDENT 3/TERM 1/SA
```

แสดงว่าใช้งานได้

### อัปโหลดขึ้น GitHub

```bash
git add .
git commit -m "Update SA project"
git push
```

---

## 4. ไม่มีไฟล์ให้ Commit

### อาการ

```text
nothing to commit, working tree clean
```

### ความหมาย

- ไม่มีไฟล์ใหม่
- ไม่มีไฟล์ที่แก้ไข
- GitHub กับเครื่องตรงกันแล้ว

---

## 5. มี Git ซ้อนกัน (Nested Repository)

### อาการ

```text
modified: week1 (modified content)
```

หรือ

```text
modified: my-project (modified content)
```

### ตัวอย่างโครงสร้าง

```text
SA
└── week1
    └── .git
```

หรือ

```text
SA
└── week1
    └── my-project
        └── .git
```

### วิธีแก้

เข้าไปในโปรเจกต์ย่อยก่อน

```bash
cd week1
```

หรือ

```bash
cd my-project
```

จากนั้น

```bash
git add .
git commit -m "Update project"
git push
```

---

## 6. อัปเดตโปรเจกต์หลังแก้ไขโค้ด

```bash
git add .
git commit -m "Update project"
git push
```

### ตัวอย่าง

```bash
git commit -m "Add login page"
```

```bash
git commit -m "Fix database connection"
```

```bash
git commit -m "Update frontend"
```

---

## 7. คำสั่งตรวจสอบที่ใช้บ่อย

### ตรวจสอบสถานะไฟล์

```bash
git status
```

### ดูประวัติ Commit

```bash
git log --oneline
```

### ดู Repository ที่เชื่อมอยู่

```bash
git remote -v
```
=========================================================================
# 🔄 วิธีเปลี่ยน Repository ปลายทางที่จะอัปโหลดขึ้น GitHub

## 1️⃣ ตรวจสอบ Repository ปัจจุบัน

รันคำสั่ง

```bash
git remote -v
