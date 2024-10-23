package models

import (
	"time"
)

// User는 사용자 정보를 나타내는 모델이다.
// 틸드 문자(`) 옆에 있는 부분은 태그(tag)라고 불린다.
// Go에서 구조체의 필드에 추가할 수 있는 메타데이터 정보로 이는 자바의 어노테이션과 비슷한 역할을 한다.
// Go의 태그는 주로 데이터베이스 ORM, JSON 마샬링/언마샬링, 검증 라이브러리 등에서 사용되는데,
// 각 태그는 Go의 리플렉션(reflection) 기능을 사용하여 런타임에 읽히고 처리된다.
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    CreatedAt time.Time
}
