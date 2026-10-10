# Tích hợp hệ thống gửi thông báo bằng Adapter Pattern

**Ngôn ngữ:** Golang · **Pattern:** Adapter · **Trình độ:** Beginner  
**Hình thức:** Cá nhân · **Thời gian:** 90 phút · **Thang điểm:** 100

## Bối cảnh

Ứng dụng cần gửi thông báo cho người dùng. Mọi dịch vụ gửi thông báo phải triển khai interface `Notifier`, nhưng thư viện SMS cũ dùng phương thức `SendSMS` nên không trực tiếp triển khai interface này:

```go
type Notifier interface {
    Send(message string) error
}

type LegacySMS struct{}

func (s *LegacySMS) SendSMS(content string) {
    fmt.Println("SMS:", content)
}
```

Hãy dùng Adapter Pattern để kết nối thư viện cũ với ứng dụng. Không sửa `LegacySMS` hoặc `SendSMS`.

## Yêu cầu

1. **Xây dựng SMSAdapter (60 điểm):** Tạo struct `SMSAdapter` chứa `*LegacySMS`; triển khai `Send(message string) error`; trong phương thức này gọi `LegacySMS.SendSMS(message)`. Thư viện cũ không trả lỗi nên trả `nil` sau khi gửi.
2. **Sử dụng Adapter (30 điểm):** Trong `main()`, khởi tạo `LegacySMS`, tạo `SMSAdapter`, gán adapter cho biến kiểu `Notifier` và gọi `Send()` để gửi thông báo.
3. **Giải thích (10 điểm):** Trong `README.md`, trả lời mỗi câu khoảng 2–3 dòng: Adapter Pattern giải quyết vấn đề gì? Tại sao không cần chỉnh sửa `LegacySMS`?

Output mong đợi khi chạy `go run .`:

```text
SMS: Xin chao, don hang cua ban da duoc xac nhan!
```

## Các bước thực hiện

1. Sao chép các file trong `student/starter/` vào repository nộp bài của bạn.
2. Đọc interface và thư viện cũ trong `main.go`.
3. Hoàn thành TODO tạo adapter, sau đó hoàn thành TODO trong `main()`.
4. Chạy `go build ./...` và `go run .`; đối chiếu output phía trên.
5. Điền thông tin cá nhân và hai câu trả lời vào `README.md`.

Chỉ cần code ngắn gọn, dễ đọc. Không yêu cầu unit test, database hoặc HTTP API. Không sử dụng generic, reflection, framework Dependency Injection, Factory Pattern hoặc abstraction không cần thiết.

## Thang điểm

| Nội dung | Điểm |
|---|---:|
| Tạo đúng SMSAdapter | 20 |
| Triển khai đúng Notifier.Send() | 25 |
| Gọi đúng LegacySMS.SendSMS() | 15 |
| Sử dụng Adapter trong main() | 20 |
| Chương trình chạy đúng output | 10 |
| Giải thích trong README | 10 |
| **Tổng** | **100** |

## Nộp bài

Nộp `main.go`, `go.mod`, `README.md` và link GitHub repository. README cần ghi:

```text
Họ và tên:
Mã sinh viên:
Lớp:
Bài thi: Adapter Pattern
```

Ghi hướng dẫn chạy `go run .` trong README. Đảm bảo repository chứa đủ ba file và người chấm có quyền truy cập.
