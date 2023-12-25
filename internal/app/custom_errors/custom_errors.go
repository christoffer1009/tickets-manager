package custom_errors

type ErroValidacao struct {
	Campo    string `json:"campo"`
	Mensagem string `json:"mensagem"`
}
