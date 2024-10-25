package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"form-survey-cs-service/internal/config"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const defaultDues = "5000원"

type Member struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Group        string    `json:"group"`
	Corps        string    `json:"corps"`
	Gender       string    `json:"gender"`
	Generation   string    `json:"generation"`
	Region       string    `json:"region"`
	RegisteredAt time.Time `json:"registered_at"`
	PayAmount    float64   `json:"pay_amount"`
	PaidAt       string    `json:"paid_at"`
	Food         string    `json:"food"`

	Record *TaskRecord `json:"record"`
}

var regionMap map[string]string

// JSON 파일 읽고 매칭된 값으로 변환하는 함수
func loadJSONData(filePath string) {
	// JSON 파일 읽기
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal().Msgf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	// JSON 데이터 읽기
	byteValue, _ := ioutil.ReadAll(file)

	// JSON을 map[string]string 형식으로 파싱
	if err := json.Unmarshal(byteValue, &regionMap); err != nil {
		log.Fatal().Msgf("Failed to unmarshal JSON data: %v", err)
	}
}

// 매칭된 값을 반환하는 함수
func matchRegion(region string) string {
	if val, exists := regionMap[region]; exists {
		return val
	}
	return ""
}

func (i *Member) encode() string {
	// encode to json
	loadJSONData(config.EncodeJsonPath)

	// 테스트할 지역 이름
	return matchRegion(i.Region)
}

func (i *Member) DefineTransitCode() {
	// define transit code
	i.Group = i.encode()
}

func (i *Member) RecordTask(state string) *Member {
	i.Record = &TaskRecord{
		state: state,
	}
	return i
}

func (i *Member) CalcDues() {
	if i.Food == "" {
		i.Food = defaultDues
	}
}

func (i *Member) ReadyQrTask() {
	// ready qr task
	if i.Region == "" {
		i.Region = "s"
	}
	body := map[string]interface{}{
		"cms":           "y",
		"greetings":     "y",
		"agree":         "Y",
		"created_at":    "",
		"privacy_agree": "y",
		"time":          i.RegisteredAt.Format("2006-01-02 15:04:05"),
		"name":          i.Name,
		"gender":        i.Gender,
		"period":        "굿뉴스코 가족",
		"job":           "s",
		"region":        i.Region,
		"phone_number":  i.Phone,
		"transport":     "s",
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Error().Msgf("Failed to marshal JSON data: %v", err)
		return
	}

	// Convert to io.Reader
	reader := bytes.NewReader(jsonData)
	post, err := http.Post(config.GetEnv("QR_API_HOST", "http://localhost:8000")+"/api/v0/apply", "application/json", reader)
	if err != nil {
		log.Info().Msgf("Failed to post data: %v", err)
		return
	}
	defer post.Body.Close()
	if post.StatusCode != http.StatusCreated {
		log.Info().Msgf("Failed to post data: %v", post.StatusCode)
		respBody, _ := io.ReadAll(post.Body)
		// Print the response body
		if len(respBody) > 0 {
			fmt.Println(string(respBody))
		}
		return
	}
	log.Info().Msgf("Success to post data: %v", post.StatusCode)
}
