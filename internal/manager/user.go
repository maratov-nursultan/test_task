package user

import (
	"context"
	"fmt"
	"github.com/maratov-nursultan/profile/internal/model"
	"github.com/maratov-nursultan/profile/internal/repository"
	"strconv"
	"strings"
	"unicode"
)

var listNumberV1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
var listNumberV2 = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}

type ManagerSDK interface {
	CheckIin(iin string) (*model.IinCheckResponse, error)
	CreateUser(ctx context.Context, request *model.InfoRequest) error
	ListUserByName(ctx context.Context, name string) ([]*model.User, error)
	GetUserByIin(ctx context.Context, iin string) (*model.User, error)
}

type Manager struct {
	userRepo repository.UserSDK
}

func NewManager(userRepo repository.UserSDK) ManagerSDK {
	return &Manager{
		userRepo: userRepo,
	}
}

func (m *Manager) CheckIin(iin string) (*model.IinCheckResponse, error) {
	listIinNumber := make([]int, 12)
	for i := 0; i < 12; i++ {
		number, _ := strconv.Atoi(string(iin[i]))
		listIinNumber[i] = number
	}

	sum := 0
	for i := 0; i < 11; i++ {
		sum += listIinNumber[i] * listNumberV1[i]
	}
	controlNumber := sum % 11

	if controlNumber == 10 {
		sum = 0
		for i := 0; i < 11; i++ {
			sum += listIinNumber[i] * listNumberV2[i]
		}
		controlNumber = sum % 11
	}

	if controlNumber != listIinNumber[11] {
		return nil, model.ErrIinInvalid
	}

	intIndexSex, err := strconv.Atoi(string(iin[6]))
	if err != nil {
		return nil, err
	}

	sex := "male"
	if intIndexSex%2 == 0 {
		sex = "female"
	}

	var century string
	switch intIndexSex {
	case 1, 2:
		century = "18"
	case 3, 4:
		century = "19"
	case 5, 6:
		century = "20"
	default:
		return nil, model.ErrIinInvalid
	}

	day := iin[4:6]
	month := iin[2:4]
	year := century + iin[0:2]

	dateOfBirth := fmt.Sprintf("%s.%s.%s", day, month, year)

	return &model.IinCheckResponse{
		Correct:     true,
		Sex:         sex,
		DateOfBirth: dateOfBirth,
	}, nil
}

func (m *Manager) CreateUser(ctx context.Context, request *model.InfoRequest) error {
	fio := make([]string, 3)

	request.Name += " "
	name := ""
	count := 0
	for _, letter := range request.Name {
		if string(letter) == " " {
			fio[count] = name
			name = ""
			count++
		} else {
			name += string(letter)
		}
	}

	err := m.userRepo.Create(ctx, &repository.User{
		Iin:        request.Iin,
		Firstname:  fio[1],
		Lastname:   fio[0],
		Middlename: fio[2],
		Phone:      request.Phone,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) ListUserByName(ctx context.Context, name string) ([]*model.User, error) {
	listUserResp, err := m.userRepo.ListUserByName(ctx, name)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.User, len(listUserResp))
	for i, user := range listUserResp {
		fullName := upFirstLetterWord(user.Lastname) + " " + upFirstLetterWord(user.Firstname) + " " + upFirstLetterWord(user.Middlename)
		resp[i] = &model.User{
			Name:  fullName,
			Iin:   user.Iin,
			Phone: user.Phone,
		}
	}

	return resp, nil
}

func (m *Manager) GetUserByIin(ctx context.Context, iin string) (*model.User, error) {
	user, err := m.userRepo.GetUserByIin(ctx, iin)
	if err != nil {
		return nil, err
	}

	fullName := upFirstLetterWord(user.Lastname) + " " + upFirstLetterWord(user.Firstname) + " " + upFirstLetterWord(user.Middlename)
	resp := &model.User{
		Name:  fullName,
		Iin:   user.Iin,
		Phone: user.Phone,
	}

	return resp, nil
}

func upFirstLetterWord(word string) string {
	if word == "" {
		return ""
	}

	listWord := []rune(strings.ToLower(word))
	listWord[0] = unicode.ToUpper(listWord[0])
	return string(listWord)
}
