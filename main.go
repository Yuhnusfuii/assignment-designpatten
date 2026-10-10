package main

import "fmt"

// Notifier là interface mà ứng dụng sử dụng để gửi thông báo.
type Notifier interface {
	Send(message string) error
}

// LegacySMS là thư viện cũ. Không chỉnh sửa struct và phương thức này.
type LegacySMS struct{}

func (s *LegacySMS) SendSMS(content string) {
	fmt.Println("SMS:", content)
}

// SMSAdapter bọc LegacySMS để tương thích với interface Notifier.
type SMSAdapter struct {
	legacySMS *LegacySMS
}

// Send triển khai interface Notifier và gọi SendSMS của LegacySMS.
func (a *SMSAdapter) Send(message string) error {
	a.legacySMS.SendSMS(message)
	return nil
}

func main() {
	// Khởi tạo LegacySMS và SMSAdapter
	legacy := &LegacySMS{}
	adapter := &SMSAdapter{
		legacySMS: legacy,
	}

	// Gán adapter cho biến kiểu Notifier
	var notifier Notifier = adapter

	// Gọi Send với thông báo yêu cầu
	_ = notifier.Send("Xin chao, don hang cua ban da duoc xac nhan!")
}

