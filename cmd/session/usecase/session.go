package usecase

import (
	"backend/cmd/session/repository"
	"backend/pkg/models"
)

type SessionUseCase struct {
	sessRepo repository.SessionRepository
}

func CreateSessionUseCase(sessionRepository repository.SessionRepository) SessionUseCase {
	return SessionUseCase{sessRepo: sessionRepository}
}

func (sessUseCase *SessionUseCase) SetSession(session models.Session, cookieDuration int) error {
	return sessUseCase.sessRepo.SetSession(session, cookieDuration)
}

func (sessUseCase *SessionUseCase) GetSession(cookieVal string) (uint64, error) {
	return sessUseCase.sessRepo.GetSession(cookieVal)
}

func (sessUseCase *SessionUseCase) DeleteSession(cookieVal string) error {
	return sessUseCase.sessRepo.DeleteSession(cookieVal)
}

func (sessUseCase *SessionUseCase) SetCode(code models.Code) error {
	return sessUseCase.sessRepo.SetCode(code)
}

func (sessUseCase *SessionUseCase) GetCode(userId uint64) (string, error) {
	return sessUseCase.sessRepo.GetCode(userId)
}

func (sessUseCase *SessionUseCase) DeleteCode(userId uint64) error {
	return sessUseCase.sessRepo.DeleteCode(userId)
}
