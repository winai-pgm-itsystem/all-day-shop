package alldaylogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/winai-pgm-itsystem/all-day-shop/pkg/utils"
)

type IAlldayLogger interface {
	Print() IAlldayLogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
}

type alldayLogger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitAlldayLogger(c *fiber.Ctx, code int, res any) IAlldayLogger {
	log := &alldayLogger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		Path:       c.Path(),
		StatusCode: code,
	}
	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(res)

	return log
}

func (l *alldayLogger) Print() IAlldayLogger {
	utils.Debug(l)
	return l
}

func (l *alldayLogger) Save() {
	data := utils.Output(l)

	filename := fmt.Sprintf("./assets/logs/alldaylogger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	file.WriteString(string(data) + "\n")
}

func (l *alldayLogger) SetQuery(c *fiber.Ctx) {
	var body any
	if err := c.QueryParser(&body); err != nil {
		log.Printf("query parser error: %v", err)
	}
	l.Query = body
}

func (l *alldayLogger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("body parser error: %v", err)
	}

	switch l.Path {
	case "v1/users/signup":
		l.Body = "never gonna give you up"
	default:
		l.Body = body
	}
}

func (l *alldayLogger) SetResponse(res any) {
	l.Response = res
}
