package validators

import (
	"github.com/christoffer1009/tickets-manager/internal/app/custom_errors"
	"github.com/christoffer1009/tickets-manager/internal/app/models"
	"github.com/go-playground/validator/v10"
)

func ValidarCriarTecnico(tecnicoDTO models.TecnicoDTO) []custom_errors.ErroValidacao {
	validate := validator.New()

	if err := validate.Struct(tecnicoDTO); err != nil {
		var errosValidacao []custom_errors.ErroValidacao

		for _, fieldError := range err.(validator.ValidationErrors) {
			erroValidacao := custom_errors.ErroValidacao{
				Campo:    fieldError.Field(),
				Mensagem: fieldError.Tag(),
			}
			errosValidacao = append(errosValidacao, erroValidacao)
		}

		return errosValidacao

	}

	return nil
}

func ValidarAtualizarTecnico(tecnicoDTO models.AtualizarTecnicoDTO) []custom_errors.ErroValidacao {
	validate := validator.New()

	if err := validate.Struct(tecnicoDTO); err != nil {
		var errosValidacao []custom_errors.ErroValidacao

		for _, fieldError := range err.(validator.ValidationErrors) {
			erroValidacao := custom_errors.ErroValidacao{
				Campo:    fieldError.Field(),
				Mensagem: fieldError.Tag(),
			}
			errosValidacao = append(errosValidacao, erroValidacao)
		}

		return errosValidacao

	}

	return nil
}
