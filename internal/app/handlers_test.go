package app_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndreyKosinskiy/go-blog/internal/app"
	"github.com/AndreyKosinskiy/go-blog/internal/user/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func testUser() *model.User {
	u := &model.User{
		Id:       uuid.UUID{},
		UserName: "TestUser",
		Email:    "testuser@test.test",
		Password: ";j9124lkn0a9ul2k3n4",
	}
	return u
}

func testUserJSON() []byte {
	u := testUser()
	ujson, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return ujson
}

func testUserXML() []byte {
	u := testUser()
	uxml, err := xml.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return uxml
}

func TestCreateUser(t *testing.T) {
	e := echo.New()
	testcase := []struct {
		name        string
		body        []byte
		contentType string
		want        string
	}{
		{
			name:        "user xml",
			body:        testUserXML(),
			contentType: echo.MIMEApplicationXML,
			want:        string(testUserXML()),
		},
		{
			name:        "user json",
			body:        testUserJSON(),
			contentType: echo.MIMEApplicationJSON,
			want:        string(testUserJSON()),
		},
	}
	for _, v := range testcase {
		t.Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(v.body))
			r.Header.Set(echo.HeaderContentType, v.contentType)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)
			a := &app.App{}
			fmt.Printf("\n%s\n", v.body)
			assert.NoError(t, a.CreateUserHandle(c))
		})
	}
}
