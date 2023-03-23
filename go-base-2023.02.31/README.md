## I. Các thành phần chính trong project

1. **Routes**: quản lý luồng đi giữa các middleware
2. **Controller**: điều hướng để xử lý request và và trả về kết quả, tương tác nhiều với các thành phần trong project như Service, Dto,...
3. **Service**: Xử lý business logic, tương tác với repository để CRUD data trong Database
4. **Repository**: tương tác với database để query data
5. **Validate request**: Controller gọi validate request để validate các thông tin gửi lên
6. **Dto**: là thành phần chứa thông tin trả về cho client nên mỗi Dto cần có 1 MapFrom function
7. **CustomError**: nếu có lỗi thì gọi thành phần này, định dạng và xử lý lỗi trả về
8. **CustomResponse**: định dạng và xử lý data trước khi trả về cho client

---

## II. Ý nghĩa Cấu trúc Project

1. **auth** chứa các vấn đề liên quan đến xử lý xác thực trước khi Request tới **Controller**
2. **helpers** giống **utils** nhưng được sài ít hơn trong dự án thường chỉ gọi 1 lần, thường dùng để viết các function xử lý tạo instance kêt nối tới các bên thứ 3
3. **models** định nghĩa các model để mapping tới database
4. **modules** chứa các module chính trong project. Mỗi module gồm 4 thành phần chính là:
   - **Controller**: điều hướng
   - **Service**: xử lý business logic
   - **Repository**: tương tác với Database
   - **Dto**: đinh nghĩa những nên trả về những gì cho client
5. **routes** quản lý route chỉ định logic luồng đi chính của dự án như khi request tói thì đi qua những middleware nào
6. **utils** chứa các function xử lý được dùng nhiều nơi và nhiều lần được gọi trong project
7. **middlewares** định nghĩa các luồng để routes quản lý điều hướng

---

## III. Lưu ý

- Để tránh import cycle hay cycle dependency thì package utils không được dùng các function ở trong các model. Mà hãy viết trực tiếp. Cần data từ database thì query trực tiếp.
  - [resolve import cycle](https://jogendra.dev/import-cycles-in-golang-and-how-to-deal-with-them)
