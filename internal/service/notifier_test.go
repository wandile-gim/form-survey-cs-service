package service

import (
	"bytes"
	"fmt"
	"form-survey-cs-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestSMSService_buildMessage(t *testing.T) {
	info := map[string]string{
		"서울": "기업 111-123455-01-011 홍길동",
	}
	messageFormat := `[총동문회_%s지역]
2024 하반기 총동문회 회비 입금 계좌는 %s %s (예금주:%s) 입니다.

입금자명은 나라/기수/성함 (ex 미국11소피아)으로 기재하시고 5,000원입금 부탁드립니다.

입금이 완료되면 접수용 큐알코드를 보내드리겠습니다. 감사합니다.`

	message := &Message{
		Member: domain.Member{
			Id:           0,
			Name:         "Test",
			Phone:        "01012345678",
			Region:       "서울",
			RegisteredAt: time.Now(),
			Record: &domain.TaskRecord{
				Id:    0,
				Retry: 0,
			},
		},
	}
	infos, exists := info[message.Member.Region]
	if !exists {
		t.Errorf("지역 정보가 없습니다: %s", message.Member.Region)
	}

	// 문자열을 공백으로 분리하여 은행명, 계좌번호, 예금주를 추출
	parts := strings.Fields(infos)
	if len(parts) < 3 {
		t.Errorf("잘못된 형식의 데이터: %s", infos)
	}

	bank := parts[0]          // 은행명
	accountNumber := parts[1] // 계좌번호
	accountHolder := parts[2] // 예금주

	// 문자 메시지 포맷 생성
	msg := fmt.Sprintf(messageFormat, message.Member.Region, bank, accountNumber, accountHolder)
	message.Body = msg
	if message.Body == "" {
		t.Errorf("메시지 생성에 실패했습니다")
	}
	assert.Contains(t, message.Body, "서울")
	log.Printf("message: %s", message.Body)
}

func TestSMSService_sendSMS(t *testing.T) {
	url := "https://apis.aligo.in/send/"
	messageFormat := `[총동문회_%s지역]
2024 하반기 총동문회 회비 입금 계좌는 %s %s (예금주:%s) 입니다.

입금자명은 나라/기수/성함 (ex 미국11소피아)으로 기재하시고 5,000원입금 부탁드립니다.

입금이 완료되면 접수용 큐알코드를 보내드리겠습니다. 감사합니다.`
	msg := fmt.Sprintf(messageFormat, "서울", "기업", "111-123455-01-011", "홍길동")
	// 요청을 위한 데이터를 설정 (URL 인코딩된 폼 데이터)
	data := map[string]string{
		"key":         "",             // API 키
		"user_id":     "",             // 유저 ID
		"sender":      "",             // 발신자 번호
		"receiver":    "01012345678",  // 수신자 번호
		"msg":         msg,            // 메시지 내용
		"title":       "API TEST 입니다", // 제목
		"testmode_yn": "Y",            // 테스트 모드
	}

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
		t.Errorf("HTTP 요청 생성 실패: %v", err)
	}

	// Content-Type을 멀티파트로 설정
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 클라이언트를 사용해 요청 보내기
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("요청 실패: %v", err)
	}
	defer resp.Body.Close()

	// 응답 상태 코드 확인
	if resp.StatusCode != http.StatusOK {
		t.Errorf("요청 실패, 상태 코드: %d", resp.StatusCode)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
