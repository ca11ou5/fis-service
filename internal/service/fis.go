package service

import (
	"fis/internal/entity"
	"fis/internal/pb"
	"slices"
	"sync"
	"time"
)

const (
	// Application statuses
	pendingStatus  = "pending"
	repeatedStatus = "repeated"
	approveStatus  = "approve"
	rejectStatus   = "reject"
	errorStatus    = "error"

	// Scoring statuses
	scoringErrorStatus     = "ERROR"
	scoringCompletedStatus = "COMPLETED"

	// Decision statuses
	positive = "POSITIVE"
	negative = "NEGATIVE"

	// Scopes
	fullScope      = "full"
	softScoreScope = "soft-score"
)

func (s *Service) FullMappingAndSave(app *pb.Application, requestId string) error {
	var wg sync.WaitGroup
	var locality, work, registration, birthplace, realEstate *entity.Suggestion

	wg.Add(5)
	go func() {
		// TODO: Узнать если нам не приходит поле, стоит ли его подменять пустым suggestion'ом с дадаты, или же стоит оставлять как есть
		locality = s.repo.GetDadataAddress(app.LifeAddress)
		defer wg.Done()
	}()
	go func() {
		work = s.repo.GetDadataAddress(app.WorkAddress)
		defer wg.Done()
	}()
	go func() {
		registration = s.repo.GetDadataAddress(app.RegistrationAddress)
		defer wg.Done()
	}()
	go func() {
		birthplace = s.repo.GetDadataAddress(app.BirthPlace)
		defer wg.Done()
	}()
	go func() {
		realEstate = s.repo.GetDadataAddress(app.RealEstateAddress)
		defer wg.Done()
	}()
	wg.Wait()

	var application *entity.Application
	var err error
	if requestId != "" {
		application, err = s.repo.GetMongoApplication(requestId)
		if err != nil {
			return err
		}
	} else {
		// TODO: fis.service.ts 837 line / do logic from new this.appModel()
		// TODO: 100% Сделать application не nil
	}

	// TODO: Узнать нужна ли проверка в fis.service.ts 839 line / Она скорее всего никогда не будет выполнена

	// TODO: Убедиться что у поля в монге у заявки OpenApiSocialStatus тип string
	if application.OpenApiSocialStatus == "" {
		if app.SocialStatus != "" {
			application.OpenApiSocialStatus = app.SocialStatus
		}
	}

	// TODO: fis.service.ts 849 line / Узнать нужно ли это?
	app.WorkSector = nil

	// TODO: From cleanEmpty
}

func (s *Service) UpdateStatus(application *entity.Application) error {
	if application.OnlineDecision == pendingStatus {
		status := application.OnlineDecision
		var scoring *entity.ScoringStatus

		defer func() {
			if scoring != nil {
				application.LastScoringStatus.UpdatedAt = time.Now()
				application.LastScoringStatus.RequestID = scoring.RequestID
				application.LastScoringStatus.Status = scoring.Status
				application.LastScoringStatus.Decision = scoring.Decision
			}
			err := s.repo.SaveApplication(application)
			if err != nil {
				s.log.WithError(err).Error("failed to save application")
				return
			}

			if status != pendingStatus {
				application.OnlineDecision = status
				application.CreditScore = scoring.Comment
				err = s.repo.SaveApplication(application)
				if err != nil {
					s.log.WithError(err).Error("failed to save application")
					return
				}
			}

			if scoring.Status == scoringCompletedStatus && status != errorStatus {
				// TODO: fis.service.ts 963 line / Узнать нужна ли проверка на typeof string для citizenship
				if application.PartnerName != "" {
					err = s.repo.SaveApplication(application)
					if err != nil {
						s.log.WithError(err).Error("failed to save application")
						return
					}
					// TODO: Сохранить в мускул
					return
				}

				if application.PartnerName != "" {
					user, err := s.repo.GetUserByUsername(application.PartnerName)
					if err != nil {
						s.log.WithError(err).Error("failed to get user by username")
						return
					}

					switch scoring.Decision {
					case negative:
						if !slices.Contains(user.Scope, softScoreScope) {
							return
						}
					case positive:
						if slices.Contains(user.Scope, fullScope) {
							// TODO: logic from fis.service.ts 1003 line / await this.full(instance);
							return
						}
					}
				}

				// TODO: Узнать как интерпретировать fis.service.ts 1012 line / Везде поменять application.Citizenship / Нужна ли проверка вообще?
			}
		}()

		err := func() error {
			scoring, err := s.repo.GetScoringStatus(application.OrderID)
			if err != nil {
				return err
			}

			err = s.repo.UpdateApplicationScoring(application.ID, scoring)
			if err != nil {
				return err
			}

			switch scoring.Status {
			case scoringErrorStatus:
				status = errorStatus
			case scoringCompletedStatus:
				switch scoring.Decision {
				case negative:
					status = rejectStatus
				case positive:
					status = approveStatus
				}
			}
			return nil
		}()
		if err != nil {
			err = s.repo.UpdateApplicationScoring(application.ID, err.Error())
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (s *Service) SendToFis(requestId string) error {
	application, err := s.repo.GetMongoApplication(requestId)
	if err != nil {
		return err
	}

	// True if not nil, false if nil
	isLongApplication := application.Citizenship != nil

	switch isLongApplication {
	case true:
		s.longApplicationMapping(application)
	default:
		s.fullApplicationToFisMapping(application)
	}

}
