# Bài thi: Adapter Pattern

```text
Họ và tên: Nguyễn Nhật Huy
Mã sinh viên:
Lớp: K24
Bài thi: Adapter Pattern
```

## Việc cần làm

Đọc `assignment.md` trong bộ đề. Hoàn thành các TODO trong `main.go`: tạo `SMSAdapter` chứa `*LegacySMS`, triển khai `Notifier.Send()` và dùng adapter qua biến kiểu `Notifier` trong `main()`. Không sửa thư viện `LegacySMS`.

Starter hiện biên dịch được nhưng chưa gửi thông báo. Không yêu cầu unit test, database hay HTTP API.

## Kiểm tra và chạy

Từ thư mục chứa `go.mod` (Go 1.22 trở lên):

```bash
go build ./...
go run .
```

Sau khi hoàn thành, kết quả cần là:

```text
SMS: Xin chao, don hang cua ban da duoc xac nhan!
```

## Câu trả lời của sinh viên

Mỗi câu trả lời khoảng 2–3 dòng.

1. Adapter Pattern giải quyết vấn đề gì?

   Adapter Pattern giải quyết vấn đề bất tương thích về giao diện (interface) giữa mã nguồn ứng dụng hiện tại và thư viện có sẵn (hoặc bên thứ ba). Nó hoạt động như một lớp chuyển đổi trung gian, giúp các thành phần không cùng giao diện có thể làm việc trơn tru với nhau mà không cần thay đổi cấu trúc của chúng.

2. Tại sao không cần chỉnh sửa `LegacySMS`?

   Không cần chỉnh sửa `LegacySMS` để tuân thủ nguyên lý Open/Closed Principle (OCP trong SOLID) – mở rộng chức năng mà không làm thay đổi mã nguồn đã ổn định. Điều này tránh gây lỗi hồi quy (regression bugs) cho các hệ thống khác đang dùng thư viện này, đồng thời giải quyết trường hợp mã nguồn cũ là thư viện của bên thứ ba không thể can thiệp trực tiếp.

## Nộp bài

Nộp `main.go`, `go.mod`, `README.md` và link GitHub repository. Điền đủ thông tin cá nhân phía trên và bảo đảm người chấm truy cập được repository.
