package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"form-survey-cs-service/internal/config"
	"form-survey-cs-service/internal/domain"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

var (
	messageFormat = ""
	accountInfo   map[string]string
)

type Message struct {
	To     string
	Body   string
	Member domain.Member
}

type Notifier interface {
	SendMessage(message *Message)
}

type SMSService struct {
	Log chan string
}

func NewSMSService() *SMSService {
	log := make(chan string)
	return &SMSService{Log: log}

}

// JSON 파일을 읽어서 전역 변수에 맵을 저장하는 함수
func loadAccountInfo() error {
	// 파일을 읽음
	file, err := os.Open(config.BankAccountSecretPath)
	if err != nil {
		return fmt.Errorf("파일을 여는데 실패했습니다: %w", err)
	}
	defer file.Close()

	// 파일 내용을 읽음
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("파일 내용을 읽는데 실패했습니다: %w", err)
	}

	// JSON 데이터를 맵으로 변환
	err = json.Unmarshal(bytes, &accountInfo)
	if err != nil {
		return fmt.Errorf("JSON을 맵으로 변환하는데 실패했습니다: %w", err)
	}

	return nil
}

func loadMessageFormat(filename string) (string, error) {
	// load message format
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("파일을 여는데 실패했습니다: %w", err)
	}
	defer file.Close()

	// 파일 내용을 읽음
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("파일 내용을 읽는데 실패했습니다: %w", err)
	}
	return string(bytes), nil
}

func (s *SMSService) buildMessage(message *Message) (string, error) {
	// 해당 지역의 정보가 존재하는지 확인
	if accountInfo == nil {
		loadAccountInfo()
	} else if messageFormat == "" {
		messageFormat, _ = loadMessageFormat(config.MessageFormatPath)
	}

	info, exists := accountInfo[message.Member.Region]
	if !exists {
		return "", fmt.Errorf("해당 지역(%s)의 정보가 없습니다", message.Member.Region)
	}

	// 문자열을 공백으로 분리하여 은행명, 계좌번호, 예금주를 추출
	parts := strings.Fields(info)
	if len(parts) < 3 {
		return "", fmt.Errorf("잘못된 형식의 데이터: %s", info)
	}

	bank := parts[0]          // 은행명
	accountNumber := parts[1] // 계좌번호
	accountHolder := parts[2] // 예금주

	// 문자 메시지 포맷 생성
	msg := fmt.Sprintf(messageFormat, message.Member.Region, bank, accountNumber, accountHolder)
	message.Body = msg
	return msg, nil
}

func (s *SMSService) SendMessage(message *Message) {
	// send message
	_, err := s.buildMessage(message)
	if err != nil {
		return
	}
	err = s.sendSMS("", message)

	log := fmt.Sprintf("sms sent to %s(%s) content: %s", message.Member.Name, message.Member.Phone, message.Body)
	if err != nil {
		log = fmt.Sprintf("sms sent failed to %s content: %s", message.Member.Phone, message.Body)
	}
	s.Log <- log
}

func (s *SMSService) sendSMS(title string, message *Message) error {
	// API URL
	url := "https://apis.aligo.in/send/"
	conf := config.SmsAPIConfig()
	// 요청을 위한 데이터를 설정 (URL 인코딩된 폼 데이터)
	data := map[string]string{
		"key":         conf.ApiKey,          // API 키
		"user_id":     conf.UserId,          // 사용자 ID
		"sender":      conf.Sender,          // 발신자 번호
		"receiver":    message.Member.Phone, // 수신자 번호
		"msg":         message.Body,         // 메시지 내용
		"title":       title,                // 제목
		"testmode_yn": "Y",                  // 테스트 모드
	}

	log.Info().Msgf("API Key %s", conf.ApiKey)

	// 새로운 멀티파트 폼 데이터 생성
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 문자열 데이터를 멀티파트에 추가
	for key, value := range data {
		_ = writer.WriteField(key, value)
	}

	// HTTP 요청 생성
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("HTTP 요청 생성 실패: %w", err)
	}

	// Content-Type을 멀티파트로 설정
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 클라이언트를 사용해 요청 보내기
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	// 응답 상태 코드 확인
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("요청 실패, 상태 코드: %d", resp.StatusCode)
	}
	return nil
}
