###########################################
# build Go binary
###########################################
FROM golang:1.23 AS builder

WORKDIR /app

# Go 의존성 다운로드 및 검증
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Go 코드 복사 및 빌드
COPY . ./
ENV CGO_ENABLED=0
RUN go build -o /app/form_survey_cs ./cmd
RUN ls -l /app

###########################################
# final image
############################################
FROM alpine:latest

ARG BUILD_TIME=unknown
ARG GIT_SHORT_COMMIT_ID=unknown

ENV BUILD_TIME=${BUILD_TIME} \
    GIT_SHORT_COMMIT_ID=${GIT_SHORT_COMMIT_ID}

# 필수 패키지 설치
RUN apk add --no-cache ca-certificates

# Go 바이너리 복사
COPY --from=builder /app/form_survey_cs /app/form_survey_cs

# 실행 권한 부여
RUN chmod +x /app/form_survey_cs

# 컨테이너 시작 시 서비스 실행
ENTRYPOINT ["/app/form_survey_cs"]
