package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"sync"
	"testing"
	"time"
)

func TestEmailSend(t *testing.T) {
	tos := []string{"1533181183@qq.com"}
	subject := "Awesome Subject"
	filename := "email_test.go"
	sendEmail(tos, subject, filename)

}

func sendEmail(tos []string, subject string, filename string) {
	e := NewEmail()
	e.From = "gowb <15123784108@163.com>"
	e.To = tos
	// 抄送
	e.Cc = []string{"cqu20141693@gmail.com"}
	// 秘密抄送
	//e.Bcc = []string{""cqu20141693@gmail.com""}

	e.Subject = subject
	text := "Text Body is, of course, supported!"
	e.Text = []byte(text)
	if filename != "" {
		_, err := e.AttachFile(filename)
		if err != nil {
			log.Fatal("attachFile not found")
			return
		}
	}
	// 465 ssl
	// 22
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "15123784108@163.com", "wq114655", "smtp.163.com"))
	if err != nil {
		log.Fatal("email send failed. ", err.Error())
		return
	}
	log.Print("email send success")
}

func TestEmailPool(t *testing.T) {
	counter := 2
	ch := make(chan *Email, counter)
	pool, err := NewPool("smtp.163.com:25",
		4,
		smtp.PlainAuth("", "15123784108@163.com", "wq114655", "smtp.163.com"))
	if err != nil {
		log.Fatal("failed to create pool:", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(counter)
	for i := 0; i < counter; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				err := pool.Send(e, 10*time.Second)
				if err != nil {
					_, err := fmt.Fprintf(os.Stderr, "email:%v sent error:%v\n", e, err)
					if err != nil {
					}
				}
			}
		}()
	}
	for i := 0; i < counter; i++ {
		e := NewEmail()
		e.From = "gowb <15123784108@163.com>"
		e.To = []string{"1533181183@qq.com"}
		e.Subject = "email pool test"
		e.Text = []byte(fmt.Sprintf("Awesome Web %d", i+1))
		ch <- e
	}
	time.Sleep(1 * time.Second)
	close(ch)
	wg.Wait()

}
