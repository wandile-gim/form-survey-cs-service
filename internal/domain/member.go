package domain

import (
	"bytes"
	"encoding/json"
	"form-survey-cs-service/internal/config"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
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

// rateLimiter is a package-level rate limiter for QR API requests
var (
	rateLimiter     = make(chan struct{}, 10) // Allow 10 concurrent requests
	rateLimiterOnce sync.Once
)

// initRateLimiter initializes the rate limiter if it hasn't been initialized yet
func initRateLimiter() {
	rateLimiterOnce.Do(func() {
		// Initialize rate limiter with tokens from environment variable or default to 10
		concurrentLimit, err := strconv.Atoi(config.GetEnv("QR_API_CONCURRENT_LIMIT", "30"))
		if err != nil || concurrentLimit <= 0 {
			concurrentLimit = 30
		}
		rateLimiter = make(chan struct{}, concurrentLimit)
		log.Info().Msgf("QR API rate limiter initialized with concurrent limit: %d", concurrentLimit)
	})
}

func (i *Member) logFailedToSendQR() {
	// Log the failure to send QR code
	log.Info().Msgf("Failed to send QR code to 이름:%s(%s)", i.Name, i.Phone)
}

func (i *Member) ReadyQrTask() {
	// Initialize rate limiter
	initRateLimiter()

	// Acquire a token from the rate limiter
	rateLimiter <- struct{}{}
	defer func() {
		// Release the token back to the rate limiter when done
		<-rateLimiter
	}()

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
		"pay_amount":    i.PayAmount,
		"food":          i.Food,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Error().Msgf("Failed to marshal JSON data: %v", err)
		return
	}

	// Add exponential backoff retry logic
	backoff := 1 * time.Second
	maxRetries := 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		// Convert to io.Reader
		reader := bytes.NewReader(jsonData)
		post, err := http.Post(config.GetEnv("QR_API_HOST", "http://localhost")+"/api/v0/apply", "application/json", reader)
		if err != nil {
			log.Info().Msgf("Failed to post data (attempt %d/%d): %v", attempt+1, maxRetries, err)
			if attempt < maxRetries-1 {
				time.Sleep(backoff)
				backoff *= 2 // Exponential backoff
			}
			continue
		}

		defer post.Body.Close()
		if post.StatusCode != http.StatusCreated {
			log.Info().Msgf("Failed to post data (attempt %d/%d): status code %v", attempt+1, maxRetries, post.StatusCode)
			respBody, _ := io.ReadAll(post.Body)
			// Print the response body
			if (post.StatusCode == http.StatusInternalServerError || post.StatusCode == http.StatusBadRequest) && len(respBody) > 0 {
				// 누구에게 전송이 안됐는지 확인하기 위해서 실패 로그 남기기
				i.logFailedToSendQR()
			}

			// If server is overloaded (429 or 503), retry with backoff
			if (post.StatusCode == http.StatusTooManyRequests || post.StatusCode == http.StatusServiceUnavailable) && attempt < maxRetries-1 {
				time.Sleep(backoff)
				backoff *= 2 // Exponential backoff
				continue
			}
			return
		}
		return
	}
}
